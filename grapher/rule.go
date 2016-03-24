package grapher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sourcegraph.com/sourcegraph/makex"
	"sourcegraph.com/sourcegraph/srclib"
	"sourcegraph.com/sourcegraph/srclib/buildstore"
	"sourcegraph.com/sourcegraph/srclib/config"
	"sourcegraph.com/sourcegraph/srclib/graph"
	"sourcegraph.com/sourcegraph/srclib/plan"
	"sourcegraph.com/sourcegraph/srclib/toolchain"
	"sourcegraph.com/sourcegraph/srclib/unit"
	"sourcegraph.com/sourcegraph/srclib/util"
)

const graphOp = "graph"
const goUnitType = "GoPackage"

func init() {
	plan.RegisterRuleMaker(graphOp, makeGraphRules)
	buildstore.RegisterDataType("graph", &graph.Output{})
}

func makeGraphRules(c *config.Tree, dataDir string, existing []makex.Rule) ([]makex.Rule, error) {
	const op = graphOp
	var rules []makex.Rule
	var goUnits unit.SourceUnits
	for _, u := range c.SourceUnits {
		if u.Type == goUnitType {
			goUnits = append(goUnits, u)
			continue
		}

		toolRef := u.Ops[op]
		if toolRef == nil {
			choice, err := toolchain.ChooseTool(graphOp, u.Type)
			if err != nil {
				return nil, err
			}
			toolRef = choice
		}

		rules = append(rules, &GraphUnitRule{dataDir, u, toolRef})
	}

	// Graph all GoPackage units together
	if len(goUnits) > 0 {
		toolRef, err := toolchain.ChooseTool(graphOp, goUnitType)
		if err != nil {
			return nil, err
		}
		rules = append(rules, &GraphMultiUnitsRule{dataDir, goUnits, goUnitType, toolRef})
	}
	return rules, nil
}

type GraphUnitRule struct {
	dataDir string
	Unit    *unit.SourceUnit
	Tool    *srclib.ToolRef
}

func (r *GraphUnitRule) Target() string {
	return filepath.ToSlash(filepath.Join(r.dataDir, plan.SourceUnitDataFilename(&graph.Output{}, r.Unit)))
}

func (r *GraphUnitRule) Prereqs() []string {
	ps := []string{filepath.ToSlash(filepath.Join(r.dataDir, plan.SourceUnitDataFilename(unit.SourceUnit{}, r.Unit)))}
	for _, file := range r.Unit.Files {
		if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
			// skip not-existent files listed in source unit
			continue
		}
		ps = append(ps, file)
	}
	return ps
}

func (r *GraphUnitRule) Recipes() []string {
	if r.Tool == nil {
		return nil
	}
	safeCommand := util.SafeCommandName(srclib.CommandName)
	return []string{
		fmt.Sprintf("%s tool %q %q < $< | %s internal normalize-graph-data --unit-type %q --dir . 1> $@", safeCommand, r.Tool.Toolchain, r.Tool.Subcmd, safeCommand, r.Unit.Type),
	}
}

func (r *GraphUnitRule) SourceUnit() *unit.SourceUnit { return r.Unit }

type GraphMultiUnitsRule struct {
	dataDir   string
	Units     unit.SourceUnits
	UnitsType string
	Tool      *srclib.ToolRef
}

func (r *GraphMultiUnitsRule) Target() string {
	return filepath.ToSlash(filepath.Join(r.dataDir, plan.SourceUnitDataFilename(&graph.Output{}, &unit.SourceUnit{Type: r.UnitsType})))
}

func (r *GraphMultiUnitsRule) Prereqs() []string {
	ps := []string{}
	for _, u := range r.Units {
		ps = append(ps, filepath.ToSlash(filepath.Join(r.dataDir, plan.SourceUnitDataFilename(unit.SourceUnit{}, u))))
		for _, file := range u.Files {
			if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
				// skip not-existent files listed in source unit
				continue
			}
			ps = append(ps, file)
		}
	}
	return ps
}

func (r *GraphMultiUnitsRule) Recipes() []string {
	if r.Tool == nil {
		return nil
	}
	safeCommand := util.SafeCommandName(srclib.CommandName)
	unitFiles := []string{}
	for _, u := range r.Units {
		unitFiles = append(unitFiles, filepath.ToSlash(filepath.Join(r.dataDir, plan.SourceUnitDataFilename(unit.SourceUnit{}, u))))
	}
	return []string{
		fmt.Sprintf("%s internal emit-unit-data %s | %s tool %q %q | %s internal normalize-graph-data --unit-type %q --dir . 1> $@", safeCommand, strings.Join(unitFiles, " "), safeCommand, r.Tool.Toolchain, r.Tool.Subcmd, safeCommand, r.UnitsType),
	}
}
