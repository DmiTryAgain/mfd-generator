package cmd

import (
	"flag"
	xmllang "github.com/vmkteam/mfd-generator/generators/xml-lang"
	"os"

	"github.com/vmkteam/mfd-generator/generators/model"
	"github.com/vmkteam/mfd-generator/generators/repo"
	"github.com/vmkteam/mfd-generator/generators/vt"
	"github.com/vmkteam/mfd-generator/generators/vt-template"
	"github.com/vmkteam/mfd-generator/generators/xml"
	"github.com/vmkteam/mfd-generator/generators/xml-vt"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "mfd",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.HasSubCommands() {
			if err := cmd.Help(); err != nil {
				panic("help not found")
			}
			os.Exit(0)
		}
	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
}

var debug = flag.Bool("debug", false, "enable debug output")

func init() {
	root.AddCommand(
		xml.CreateCommand(),
		xmlvt.CreateCommand(),
		xmllang.CreateCommand(),
		model.CreateCommand(),
		repo.CreateCommand(),
		vt.CreateCommand(),
		vttmpl.CreateCommand(),
	)
}

// Execute runs root cmd
func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
