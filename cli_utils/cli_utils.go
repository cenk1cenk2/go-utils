package cli_utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gitlab.kilic.dev/libraries/go-utils/logger"
)

var Log *logrus.Logger

func CliLoadEnvironment() {
	if env := os.Getenv("ENV_FILE"); env != "" {
		err := godotenv.Load(env)

		if err != nil {
			Log.Fatalln(err)
		}
	}
}

func CliBeforeFunction(c *cli.Context) error {
	CliLoadEnvironment()

	level, err := logrus.ParseLevel(c.String("utils.log"))

	if err != nil {
		fmt.Println(fmt.Sprintf("Log level is not valid with %s, using default.", level))

		level = logrus.InfoLevel
	}

	if c.String("utils.debug") != "" {
		level = logrus.DebugLevel
	}

	Log = logger.InitiateLogger(level)

	return nil
}

func CliGreet(c *cli.Context) {
	name := fmt.Sprintf("%s - %s", c.App.Name, c.App.Version)
	fmt.Println(name)
	fmt.Println(strings.Repeat("-", len(name)))
}

func CliRun(app *cli.App) {
	if err := app.Run(os.Args); err != nil {
		if Log != nil {
			Log.Fatalln(err)
		} else {
			fmt.Println(err)

			os.Exit(1)
		}
	}
}

func CliCreate(app *cli.App) {
	app.Flags = append(CliDefaultFlags, app.Flags...)

	app.Before = CliBeforeFunction

	CliRun(app)
}
