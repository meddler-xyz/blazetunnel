package client

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// Init function initializes the client command
// commandline functionality and retursn the cli.Command
func Init() *cli.Command {

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "tunnel",
			Usage:    "Remote public tunnel address to connect to",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "local",
			Usage:    "Local TCP server to proxy the connections to",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "token",
			Usage:    "auth-token ",
			FilePath: ".blazetoken",
			Required: true,
		},
		&cli.UintFlag{
			Name:  "i,idle-timeout",
			Usage: "Idle timeout for the quic sessions (in seconds)",
			Value: 1800,
		},
	}

	return &(cli.Command{
		Name: "client",

		Usage:  "Run a client instance",
		Action: createClient,
		Flags:  flags,
	})
}

func createClient(ctx *cli.Context) error {

	tunnel := ctx.String("tunnel")
	if tunnel == "" {
		return errors.New("Tunnel address cannot be empty")
	}

	local := ctx.String("local")
	if local == "" {
		return errors.New("Local address cannot be empty")
	}

	token := ctx.String("token")
	if token == "" {
		return errors.New("Token can not be empty")
	}

	return NewClient(tunnel, local, ctx.Uint("i"), token).Start()
}
