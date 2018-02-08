package command

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"archive/zip"
	"io/ioutil"

	"github.com/ysugimoto/ginger/config"
	"github.com/ysugimoto/ginger/entity"
	"github.com/ysugimoto/ginger/logger"
	"github.com/ysugimoto/ginger/request"
	"github.com/ysugimoto/go-args"
)

const (
	DEPLOY_FUNCTION = "function"
	DEPLOY_FN       = "fn"
	DEPLOY_API      = "api"
)

type Deploy struct {
	Command
	log *logger.Logger
}

func NewDeploy() *Deploy {
	return &Deploy{
		log: logger.WithNamespace("ginger.deploy"),
	}
}

func (d *Deploy) Help() string {
	return `
	ginger deploy [subcommand] [options]

Subcommand:
  function: Deploy functions (default: all, one of function if --name option supplied)
  api:      Deploy apis (default: all, one of path if --name option supplied)

Options:
  --all:   Deploy all functions/apis
  --name:  Target fucntion name
  --stage: Target api stage
`
}

func (d *Deploy) Run(ctx *args.Context) error {
	c := config.Load()
	if !c.Exists() {
		d.log.Error("Configuration file could not load. Run `ginger init` before.")
		return nil
	}
	switch ctx.At(1) {
	case DEPLOY_FUNCTION, DEPLOY_FN:
		if c.Project.LambdaExecutionRole == "" {
			d.log.Warn("Lambda execution role isn't set. Please open Ginger.toml ant put sutable role into 'lambda_execution_role' section.")
		}
		d.log.AddNamespace("function")
		return d.deployFunction(c, ctx)
	case DEPLOY_API:
		d.log.AddNamespace("api")
		return d.deployAPI(c, ctx)
	default:
		if ctx.Has("all") {
			d.log.AddNamespace("all")
			d.deployFunction(c, ctx)
			d.deployAPI(c, ctx)
		} else {
			fmt.Println(d.Help())
		}
		return nil
	}
}

func (d *Deploy) deployFunction(c *config.Config, ctx *args.Context) error {
	targets := c.Functions
	if ctx.Has("name") {
		name := ctx.String("name")
		if !c.Functions.Exists(name) {
			d.log.Errorf("Target function %s doesn't exist", name)
			return nil
		}
		targets = entity.Functions{c.Functions.Find(name)}
	}

	buildDir, err := ioutil.TempDir("", "ginger-builds")
	if err != nil {
		return err
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
		if err := lambda.DeployFunction(fn.Name, buffer); err == nil {
			d.log.Infof("Function %s deployed successfully!\n", fn.Name)
		}
	}
	return nil
}

func (d *Deploy) archive(fn *entity.Function, binPath string) ([]byte, error) {
	buf := new(bytes.Buffer)
	z := zip.NewWriter(buf)
	bin, err := ioutil.ReadFile(binPath)
	if err != nil {
		return nil, err
	}
	header := &zip.FileHeader{
		Name:           fn.Name,
		Method:         zip.Deflate,
		ExternalAttrs:  0777 << 16,
		CreatorVersion: 3 << 8,
	}
	if f, err := z.CreateHeader(header); err != nil {
		return nil, err
	} else if _, err := f.Write(bin); err != nil {
		return nil, err
	} else if err := z.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *Deploy) deployAPI(c *config.Config, ctx *args.Context) (err error) {
	api := request.NewAPIGateway(c)
	restId := c.API.RestId
	if restId == "" {
		restId, err = api.CreateRestAPI(fmt.Sprintf("ginger-%s", c.Project.Name))
		if err != nil {
			return nil
		}
		c.API.RestId = restId
		c.Write()
	}

	var rootId string
	if r := c.API.Find("/"); r == nil {
		rootId, err = api.GetResourceIdByPath(restId, "/")
		if err != nil {
			return nil
		}
		resource := entity.NewResource("/")
		resource.Id = rootId
		c.API.Resources = append(c.API.Resources, resource)
		c.Write()
	} else {
		rootId = r.Id
	}

	for _, r := range c.API.Resources {
		// If "Id" exists, the resource has already been deployed
		if r.Id != "" {
			continue
		}
		if err := d.deployResources(c, restId, r); err != nil {
			return nil
		}
	}
}

func (d *Deploy) deployResources(c *config.Config, restId string, resource *entity.Resource) (err error) {
	var parentId string
	var parts string
	var resource *entity.Resource
	// Split by path part and create recursively
	for _, part := range strings.Split(r.Path) {
		parts += "/" + part
		// Exists in config
		if resource = c.API.Find(parts); resource != nil {
			if resource.Id == "" {
				resource.Id, err = api.CreateResource(restId, parentId, part)
				if err != nil {
					return err
				}
			}
			parentId = resource.Id
			continue
		}
		// Create for sub path
		resourceId, err = api.CreateResource(restId, parentId, part)
		if err != nil {
			return err
		}
		resource = entity.NewResource(resourceId, parts)
		c.API.Resources = append(c.API.Resources, resource)
		c.Write()
		parentId = resourceId
		if resource.IntegratedFunction != "" {
			d.putLambdaInteragation(restId, resource)
		}
	}
}

func (d *Deploy) putLambdaIntegration(restId string, resource *entity.Resource) error {
	return nil
}
