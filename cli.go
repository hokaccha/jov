package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/hokaccha/go-prettyjson"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

var commandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
   
{{.Description}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`

func makeCommandDescription(usage, description string) string {
	format := "USAGE:\n   jov %s"

	if description == "" {
		return fmt.Sprintf(format, usage)
	} else {
		description = strings.Trim(description, "\n")
		description = strings.Replace(description, "\t\t", "   ", -1)
		format += "\nDESCRIPTION:\n%s"
		return fmt.Sprintf(format, usage, description)
	}
}

var InputJson interface{}

func NewCliApp() *cli.App {
	cli.CommandHelpTemplate = commandHelpTemplate

	app := cli.NewApp()
	app.Name = "jov"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Input JSON file path",
		},
		cli.IntFlag{
			Name:  "truncate, t",
			Usage: "truncate string length",
		},
	}
	app.Commands = []cli.Command{
		cmdGet,
		cmdSelect,
		cmdReject,
		cmdSlice,
		cmdHead,
		cmdTail,
	}
	app.Before = doBefore
	app.Action = doMain

	return app
}

func doBefore(c *cli.Context) error {
	filepath := c.String("file")

	var in interface{}
	var err error
	var d *json.Decoder

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		d = json.NewDecoder(os.Stdin)
	} else if filepath != "" {
		f, err := os.Open(filepath)
		if err != nil {
			return err
		}
		defer f.Close()
		d = json.NewDecoder(f)
	} else {
		return nil
	}

	err = d.Decode(&in)

	if err != nil {
		return err
	}

	InputJson = in

	return nil
}

func doMain(c *cli.Context) {
	if InputJson == nil {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	outputJson(InputJson, c)
}

var cmdGet = cli.Command{
	Name:        "get",
	Usage:       "Get value of property",
	Description: makeCommandDescription("get <property>", ""),
	Action: func(c *cli.Context) {
		log.Printf("%#v", c)
		out, err := processor.Get(InputJson, c.Args()[0])

		if err != nil {
			log.Fatal(err)
		}

		outputJson(out, c)
	},
}

var cmdSelect = cli.Command{
	Name:        "select",
	Usage:       "Select properties from collection",
	Description: makeCommandDescription("select <property>...", ""),
	Action: func(c *cli.Context) {
		out, err := processor.Select(InputJson, c.Args()...)

		if err != nil {
			log.Fatal(err)
		}

		outputJson(out, c)
	},
}

var cmdReject = cli.Command{
	Name:        "reject",
	Usage:       "Reject properties from collection",
	Description: makeCommandDescription("reject <property>...", ""),
	Action: func(c *cli.Context) {
		out, err := processor.Reject(InputJson, c.Args()...)

		if err != nil {
			log.Fatal(err)
		}

		outputJson(out, c)
	},
}

var cmdSlice = cli.Command{
	Name:        "slice",
	Usage:       "Slice array",
	Description: makeCommandDescription("slice <start> <length>", ""),
	Action: func(c *cli.Context) {
		args := c.Args()

		if len(args) != 2 {
			argumentsErrorAndExit(c, "slice")
		}

		start, err := strconv.Atoi(args[0])
		if err != nil {
			argumentsErrorAndExit(c, "slice")
		}

		length, err := strconv.Atoi(args[1])
		if err != nil {
			argumentsErrorAndExit(c, "slice")
		}

		out, err := processor.Slice(InputJson, start, length)

		if err != nil {
			log.Fatal(err)
		}

		outputJson(out, c)
	},
}

var cmdHead = cli.Command{
	Name:        "head",
	Usage:       "head array",
	Description: makeCommandDescription("head <length>", ""),
	Action: func(c *cli.Context) {
		args := c.Args()

		if len(args) != 1 {
			argumentsErrorAndExit(c, "head")
		}

		length, err := strconv.Atoi(args[0])
		if err != nil {
			argumentsErrorAndExit(c, "head")
		}

		out, err := processor.Slice(InputJson, 0, length)

		if err != nil {
			log.Fatal(err)
		}

		outputJson(out, c)
	},
}

var cmdTail = cli.Command{
	Name:  "tail",
	Usage: "tail array",
	Action: func(c *cli.Context) {
		args := c.Args()

		if len(args) != 1 {
			argumentsErrorAndExit(c, "tail")
		}

		length, err := strconv.Atoi(args[0])
		if err != nil {
			argumentsErrorAndExit(c, "tail")
		}

		out, err := processor.Slice(InputJson, 0, length)

		if err != nil {
			log.Fatal(err)
		}

		outputJson(out, c)
	},
}

func argumentsErrorAndExit(c *cli.Context, cmd string) {
	fmt.Fprintln(os.Stderr, "Arguments Error\n")
	cli.ShowCommandHelp(c, cmd)
	os.Exit(1)
}

func outputJson(o interface{}, c *cli.Context) {
	var s []byte
	var err error

	prettyjson.StringMaxLength = c.Int("truncate")

	if isatty.IsTerminal(os.Stdout.Fd()) {
		s, err = prettyjson.MarshalPretty(o)
	} else {
		s, err = json.MarshalIndent(o, "", "  ")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(colorable.NewColorableStdout(), string(s)+"\n")
}
