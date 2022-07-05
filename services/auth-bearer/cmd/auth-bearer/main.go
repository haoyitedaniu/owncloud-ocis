package main

import (
	"os"

	_ "github.com/owncloud/ocis/v2/ocis-pkg/log/gomicro"
	"github.com/owncloud/ocis/v2/services/auth-bearer/pkg/command"
	"github.com/owncloud/ocis/v2/services/auth-bearer/pkg/config/defaults"
)

func main() {
	if err := command.Execute(defaults.DefaultConfig()); err != nil {
		os.Exit(1)
	}
}
