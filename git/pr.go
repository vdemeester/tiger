package git

import (
	"fmt"
	"github.com/pkg/errors"
)

type PRConfig struct {
	Number string
	Remote Remote
}

func CheckoutPR(config PRConfig) error {
	fmt.Println(config)
	// Handle the remote (add if not present)
	if err := EnsureRemoteIsPresent(config.Remote); err != nil {
		return errors.Wrapf(err, "couldn't checkout PR %s", config.Number)
	}
	// Detect the remote and act on it (or bail if it's not supported)
	// Fetch the PR in a branch so that it can be checked-out
	// check it out (if asked)
	return nil
}

func RebasePR(config PRConfig) error {
	return nil
}