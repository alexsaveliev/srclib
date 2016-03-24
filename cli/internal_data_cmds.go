package cli

import (
	"encoding/json"
	"log"
	"os"

	"sourcegraph.com/sourcegraph/go-flags"

	"sourcegraph.com/sourcegraph/srclib/graph"
	"sourcegraph.com/sourcegraph/srclib/grapher"
	"sourcegraph.com/sourcegraph/srclib/unit"
)

func init() {
	cliInit = append(cliInit, func(cli *flags.Command) {
		c, err := cli.AddCommand("internal", "(internal subcommands - do not use)", "Internal subcommands. Do not use.", &struct{}{})
		if err != nil {
			log.Fatal(err)
		}

		_, err = c.AddCommand("emit-unit-data", "", "", &emitUnitDataCmd)
		if err != nil {
			log.Fatal(err)
		}

		_, err = c.AddCommand("normalize-graph-data", "", "", &normalizeGraphDataCmd)
		if err != nil {
			log.Fatal(err)
		}
	})
}

type EmitUnitDataCmd struct {
	Args struct {
		Units []string `name:"units" description:"Paths to source units."`
	} `positional-args:"yes"`
}

var emitUnitDataCmd EmitUnitDataCmd

func (c *EmitUnitDataCmd) Execute(args []string) error {
	var units unit.SourceUnits

	for _, path := range c.Args.Units {
		unitFile, err := os.Open(path)
		if err != nil {
			return err
		}
		var u *unit.SourceUnit
		if err := json.NewDecoder(unitFile).Decode(&u); err != nil {
			return err
		}
		units = append(units, u)
	}

	if err := json.NewEncoder(os.Stdout).Encode(units); err != nil {
		return err
	}

	return nil
}

type NormalizeGraphDataCmd struct {
	UnitType string `long:"unit-type" description:"source unit type (e.g., GoPackage)"`
	Dir      string `long:"dir" description:"directory of source unit (SourceUnit.Dir field)"`
}

var normalizeGraphDataCmd NormalizeGraphDataCmd

func (c *NormalizeGraphDataCmd) Execute(args []string) error {
	in := os.Stdin

	var o *graph.Output
	if err := json.NewDecoder(in).Decode(&o); err != nil {
		return err
	}

	if err := grapher.NormalizeData(c.UnitType, c.Dir, o); err != nil {
		return err
	}

	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}

	if _, err := os.Stdout.Write(data); err != nil {
		return err
	}

	return nil
}
