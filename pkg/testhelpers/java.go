package testhelpers

import (
	"github.com/spring-financial-group/jx-helpers/v3/pkg/cmdrunner"
)

// TestShouldDisableMaven should disable maven
func TestShouldDisableMaven() bool {
	cmd := cmdrunner.Command{
		Name: "mvn",
		Args: []string{"-v"},
	}
	_, err := cmd.RunWithoutRetry()
	return err != nil
}
