package model

import (
	"fmt"
	"path"
	"strings"

	"github.com/vmkteam/mfd-generator/mfd"

	"github.com/dizzyfool/genna/generators/base"
	"github.com/spf13/cobra"
)

const (
	mfdFlag = "mfd"
	pkgFlag = "package"
	nsFlag  = "ns"
)

// CreateCommand creates generator command
func CreateCommand() *cobra.Command {
	return base.CreateCommand("model", "Create model from xml", New())
}

// Generator represents mfd generator
type Generator struct {
	options Options

	printNamespaces bool
}

// New creates basic generator
func New() *Generator {
	return &Generator{}
}

// AddFlags adds flags to command
func (g *Generator) AddFlags(command *cobra.Command) {
	flags := command.Flags()
	flags.SortFlags = false

	flags.StringP(base.Output, "o", "", "output dir path")
	if err := command.MarkFlagRequired(base.Output); err != nil {
		panic(err)
	}

	flags.StringP(mfdFlag, "m", "", "mfd file")
	if err := command.MarkFlagRequired(mfdFlag); err != nil {
		panic(err)
	}

	flags.StringP(pkgFlag, "p", "", "package name")

	flags.BoolP(nsFlag, "n", false, "print package struct")
}

// ReadFlags read flags from command
func (g *Generator) ReadFlags(command *cobra.Command) error {
	var err error

	flags := command.Flags()

	if g.options.Output, err = flags.GetString(base.Output); err != nil {
		return err
	}

	if g.options.MFDPath, err = flags.GetString(mfdFlag); err != nil {
		return err
	}

	if g.options.Package, err = flags.GetString(pkgFlag); err != nil {
		return err
	}

	if g.printNamespaces, err = flags.GetBool(nsFlag); err != nil {
		return err
	}

	g.options.Def()

	return nil
}

// Generate runs generator
func (g *Generator) Generate() error {
	// loading project from file
	project, err := mfd.LoadProject(g.options.MFDPath, false)
	if err != nil {
		return err
	}

	// printing namespaces string
	if g.printNamespaces {
		fmt.Print(PrintNamespaces(project))
		return nil
	}

	g.options.GoPGVer = project.GoPGVer

	// basic generator
	output := path.Join(g.options.Output, "model.go")
	modelData := PackNamespace(project.Namespaces, g.options)
	if _, err := mfd.FormatAndSave(modelData, output, modelTemplate, true); err != nil {
		return fmt.Errorf("generate project model error: %w", err)
	}

	// generating search
	output = path.Join(g.options.Output, "model_search.go")
	searchData := PackSearchNamespace(project.Namespaces, g.options)
	if _, err := mfd.FormatAndSave(searchData, output, searchTemplate, true); err != nil {
		return fmt.Errorf("generate project search error: %w", err)
	}

	// generating validate
	output = path.Join(g.options.Output, "model_validate.go")
	validateDate := PackValidateNamespace(project.Namespaces, g.options)
	if _, err := mfd.FormatAndSave(validateDate, output, validateTemplate, true); err != nil {
		return fmt.Errorf("generate project validate error: %w", err)
	}

	// generating params
	output = path.Join(g.options.Output, "model_params.go")
	if _, err := GenerateParams(project.Namespaces, output, g.options); err != nil {
		return fmt.Errorf("generate project params error: %w", err)
	}

	return nil
}

func PrintNamespaces(project *mfd.Project) string {
	var formats []string

	for _, namespace := range project.Namespaces {
		var format []string
		for _, entity := range namespace.Entities {
			format = append(format, entity.Table)
		}
		formats = append(formats, fmt.Sprintf("%s:%s", namespace, strings.Join(format, ",")))
	}

	return strings.Join(formats, ";")
}
