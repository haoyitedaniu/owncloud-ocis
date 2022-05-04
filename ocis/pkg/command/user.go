package command

import (
	"fmt"

	"github.com/owncloud/ocis/v2/extensions/user/pkg/command"
	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/urfave/cli/v2"
)

// UserCommand is the entrypoint for the User command.
func UserCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     cfg.User.Service.Name,
		Usage:    subcommandDescription(cfg.User.Service.Name),
		Category: "extensions",
		Before: func(c *cli.Context) error {
			if err := parser.ParseConfig(cfg); err != nil {
				fmt.Printf("%v", err)
				return err
			}
			cfg.User.Commons = cfg.Commons
			return nil
		},
		Subcommands: command.GetCommands(cfg.User),
	}
}

func init() {
	register.AddCommand(UserCommand)
}
