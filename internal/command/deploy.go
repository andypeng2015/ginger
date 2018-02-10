package command

import (
	"bytes"
	"fmt"
	"os"

	"archive/zip"
	"io/ioutil"

	"github.com/ysugimoto/go-args"

	"github.com/ysugimoto/ginger/internal/config"
	"github.com/ysugimoto/ginger/internal/entity"
	"github.com/ysugimoto/ginger/internal/logger"
	"github.com/ysugimoto/ginger/internal/request"
)

const (
	DEPLOY_FUNCTION = "function"
	DEPLOY_FN       = "fn"
	DEPLOY_API      = "api"
	DEPLOY_ALL      = "all"
	DEPLOY_HELP     = "help"
)

// Deploy is the struct that manages function and api deployment.
// deploy syncs between local cofiguration and AWS platform.
type Deploy struct {
	Command
	log *logger.Logger
}

func NewDeploy() *Deploy {
	return &Deploy{
		log: logger.WithNamespace("ginger.deploy"),
	}
}

// Show deloy command help
func (d *Deploy) Help() string {
	return `
deploy - Deploy management functions and apis.

Usage:
  $ ginger deploy [subcommand] [options]

Subcommand:
  function : Deploy functions (default: all, one of function if --name option supplied)
  api      : Deploy apis (default: all, one of path if --name option supplied)
  all      : Deploy both of functions and apis
  help     : Show this help

Options:
  --name  : Target fucntion name
  --stage : Target api stage
`
}

// Run the deploy command
func (d *Deploy) Run(ctx *args.Context) {
	c := config.Load()
	if !c.Exists() {
		d.log.Error("Configuration file could not load. Run `ginger init` before.")
		return
	}
	var err error
	defer func() {
		if err != nil {
			d.log.Error(err.Error())
			debugTrace(err)
		}
		c.Write()
	}()

	switch ctx.At(1) {
	case DEPLOY_FUNCTION, DEPLOY_FN:
		err = d.deployFunction(c, ctx)
	case DEPLOY_API:
		err = d.deployAPI(c, ctx)
	case DEPLOY_ALL:
		d.log.AddNamespace("all")
		d.log.Print("========== Function Deployment ==========")
		if err = d.deployFunction(c, ctx); err != nil {
			return
		}
		d.log.Print("========== API Deployment ==========")
		err = d.deployAPI(c, ctx)
	default:
		fmt.Println(d.Help())
	}
}

// deployFunction deploys functions to AWS Lambda.
func (d *Deploy) deployFunction(c *config.Config, ctx *args.Context) error {
	if c.Project.LambdaExecutionRole == "" {
		d.log.Warn("Lambda execution role isn't set. Run the 'ginger config --role [role-name]' to set it.")
		return nil
	}
	d.log.AddNamespace("function")
	targets := c.Functions
	if ctx.Has("name") {
		name := ctx.String("name")
		if !c.Functions.Exists(name) {
			return exception("Target function %s doesn't exist", name)
		}
		targets = entity.Functions{c.Functions.Find(name)}
	}

	buildDir, err := ioutil.TempDir("", "ginger-builds")
	if err != nil {
		return exception(err.Error())
	}

	// Build functions
	defer os.RemoveAll(buildDir)
	builder := newBuilder(c.FunctionPath, buildDir)
	binaries := builder.build(targets)

	// Deploy to AWS
	lambda := request.NewLambda(c)
	for fn, binary := range binaries {
		d.log.Printf("Archiving zip for %s...\n", fn.Name)
		buffer, err := d.archive(fn, binary)
		if err != nil {
			d.log.Errorf("Archive error for %s: %s", fn.Name, err.Error())
			continue
		}
		d.log.Printf("Deploying function %s to AWS Lambda...\n", fn.Name)
		if arn, err := lambda.DeployFunction(fn, buffer); err == nil {
			d.log.Infof("Function %s deployed successfully!\n", fn.Name)
			fn.Arn = arn
		}
	}
	return nil
}

// archive archives built application binary to zip.
func (d *Deploy) archive(fn *entity.Function, binPath string) ([]byte, error) {
	buf := new(bytes.Buffer)
	z := zip.NewWriter(buf)
	bin, err := ioutil.ReadFile(binPath)
	if err != nil {
		return nil, exception("Binary file read error: %s", err.Error())
	}
	header := &zip.FileHeader{
		Name:           fn.Name,
		Method:         zip.Deflate,
		ExternalAttrs:  0777 << 16,
		CreatorVersion: 3 << 8,
	}
	if f, err := z.CreateHeader(header); err != nil {
		return nil, exception("Failed to create zip header: %s", err.Error())
	} else if _, err := f.Write(bin); err != nil {
		return nil, exception("Failed to write binary to zip stream: %s", err.Error())
	} else if err := z.Close(); err != nil {
		return nil, exception("Failed to close zip stream: %s", err.Error())
	}
	return buf.Bytes(), nil
}

// deployAPI deploys resources to AWS APIGateway.
func (d *Deploy) deployAPI(c *config.Config, ctx *args.Context) (err error) {
	d.log.AddNamespace("api")
	api := request.NewAPIGateway(c)
	restId := c.API.RestId

	if restId == "" {
		restId, err = api.CreateRestAPI(fmt.Sprintf("ginger-%s", c.Project.Name))
		if err != nil {
			return
		}
		c.API.RestId = restId
	}

	var rootId string
	if r := c.API.Find("/"); r == nil {
		rootId, err = api.GetResourceIdByPath(restId, "/")
		if err != nil {
			return
		}
		resource := entity.NewResource(rootId, "/")
		c.API.Resources = append(c.API.Resources, resource)
	} else {
		rootId = r.Id
	}

	for _, r := range c.API.Resources {
		// If "Id" exists, the resource has already been deployed
		if r.Id != "" && api.ResourceExists(restId, r.Id) {
			d.log.Infof("Endpoint %s has already deployed.\n", r.Path)
		} else {
			api.CreateResourceRecursive(restId, r.Path)
		}
		if r.Integration != nil {
			api.PutIntegration(restId, r)
		}
	}

	if s := ctx.String("stage"); s != "" {
		api.Deploy(restId, s)
	}
	return
}
