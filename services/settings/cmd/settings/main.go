package main

import (
	"os"

	_ "github.com/owncloud/ocis/v2/ocis-pkg/log/gomicro"
	"github.com/owncloud/ocis/v2/services/settings/pkg/command"
	"github.com/owncloud/ocis/v2/services/settings/pkg/config/defaults"
)

func main() {
	if err := command.Execute(defaults.DefaultConfig()); err != nil {
		os.Exit(1)
	}
}
