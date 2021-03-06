package network

import (
    "github.com/urfave/cli"

    cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)


// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
    app.Commands = append(app.Commands, cli.Command{
        Name:      name,
        Aliases:   aliases,
        Usage:     "Manage Rocket Pool network parameters",
        Subcommands: []cli.Command{

            cli.Command{
                Name:      "node-fee",
                Aliases:   []string{"f"},
                Usage:     "Get the current network node commission rate",
                UsageText: "rocketpool network node-fee",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 0); err != nil { return err }

                    // Run
                    return getNodeFee(c)

                },
            },

        },
    })
}

