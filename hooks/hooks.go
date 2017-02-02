package hooks

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/vdemeester/tiger/config"
)

// Run runs the specified hook
func Run(hook string) error {
	hookFolder, err := config.GetAsString("hooks.dir", filepath.Join(os.Getenv("HOME"), ".config/tiger/hooks"))
	if err != nil {
		return errors.Wrapf(err, "error running hook %q", hook)
	}
	hookPath := filepath.Join(hookFolder, hook)
	files, err := ioutil.ReadDir(hookPath)
	if err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "error running hook %q", hook)
	}
	// Execute each hook
	for _, f := range files {
		fmt.Fprintf(os.Stderr, "🐅 Running hook %s/%s\n", hook, f.Name())
		cmd := exec.Command(filepath.Join(hookPath, f.Name()))
		if err := cmd.Run(); err != nil {
			return errors.Wrapf(err, "error running hook %s/%s", hook, f.Name())
		}
	}
	return nil
}
