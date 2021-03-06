package command

import (
	"os"
	"sync"

	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/ysugimoto/go-args"

	"github.com/ysugimoto/ginger/config"
	"github.com/ysugimoto/ginger/logger"
)

var dependencyPackages = []string{
	"github.com/aws/aws-lambda-go",
}

// Install is the struct for install ginger project dependencies.
// This command installs project dependencies.
type Install struct {
	Command
	log *logger.Logger
}

func NewInstall() *Install {
	return &Install{
		log: logger.WithNamespace("ginger.install"),
	}
}

func (i *Install) Help() string {
	return "No Help"
}

func (i *Install) Run(ctx *args.Context) {
	c := config.Load()
	if !c.Exists() {
		i.log.Error("Configuration file could not load. Run `ginger init` before.")
		return
	}

	i.log.Print("Install function dependencies.")

	if _, err := os.Stat(c.LibPath); err != nil {
		i.log.Printf("Create library directory: %s\n", c.LibPath)
		os.Mkdir(c.LibPath, 0755)
	}

	tmpDir, _ := ioutil.TempDir("", "ginger-tmp-packages")
	defer os.RemoveAll(tmpDir)

	var wg sync.WaitGroup
	for _, pkg := range dependencyPackages {
		wg.Add(1)
		i.log.Printf("Installing %s...\n", pkg)
		go i.installDependencies(pkg, tmpDir, &wg)
	}
	wg.Wait()

	// Recursive copy
	if err := i.movePackages(tmpDir, c.LibPath); err != nil {
		i.log.Error(err.Error())
	}
}

// installDependencies installs dependencies via "go get".
//
// >>> doc
//
// ## Install dependencies
//
// Install dependency packages for build lambda function.
//
// ```
// $ ginger install
// ```
//
// This command is run automatically on initialize, but if you checkout project after initialize,
// You can install dependency packages via this command.
//
// <<< doc
func (i *Install) installDependencies(pkg, tmpDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("go", "get", pkg)
	cmd.Env = buildEnv(map[string]string{
		"GOPATH": tmpDir,
	})
	cmd.Run()
}

func (i *Install) movePackages(src, dest string) error {
	items, err := ioutil.ReadDir(src)
	if err != nil {
		return exception("Failed to read directory: %s", src)
	}
	for _, item := range items {
		from := filepath.Join(src, item.Name())
		to := filepath.Join(dest, item.Name())
		if err := os.Rename(from, to); err != nil {
			return exception("Failed to move file: %s => %s, %s", from, to, err.Error())
		}
	}
	return nil
}
