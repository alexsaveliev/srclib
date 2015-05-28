package src

import (
	"log"
	"os"

	"sourcegraph.com/sourcegraph/makex"
)

func init() {
	_, err := CLI.AddCommand("makefile",
		"prints the Makefile that `src make` executes",
		"The makefile command prints the Makefile that `src make` will execute.",
		&makefileCmd,
	)
	if err != nil {
		log.Fatal(err)
	}
}

type MakefileCmd struct {
	ToolchainExecOpt `group:"execution"`
	BuildCacheOpt    `group:"build cache"`
	Verbose bool     `short:"v" long:"verbose" description:"show more verbose output"`
}

var makefileCmd MakefileCmd

func (c *MakefileCmd) Execute(args []string) error {
	mf, err := CreateMakefile(c.ToolchainExecOpt, c.BuildCacheOpt, c.Verbose)
	if err != nil {
		return err
	}

	mfData, err := makex.Marshal(mf)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(mfData)
	return err
}
