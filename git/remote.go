package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// EnsureRemoteIsPresent makes sure the specified remote is present.
// If the specified remote has an URI, it will validate it is the same or error out.
// If the specified remote is not present, it will be added if URI is not empty, otherwise
// it will error out.
func EnsureRemoteIsPresent(remote Remote) error {
	cmd := exec.Command("git", "remote", "-v")
	content, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "error looking at remotes")
	}
	if strings.Contains(string(content), remote.Name) {
		if strings.Contains(string(content), fmt.Sprintf("%s (fetch)", remote.URI)) {
			return nil
		}
		return errors.Wrapf(err, "incoherent remote: expected %q, got %s", remote.URI, string(content))
	}

	cmd = exec.Command("git", "remote", "add", remote.Name, remote.URI)
	return cmd.Run()
}
