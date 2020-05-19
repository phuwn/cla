package fine

import (
	"github.com/spf13/cobra"

	"github.com/phuwn/clia/cmd/fine/calc"
	"github.com/phuwn/clia/cmd/fine/formula"
)

// New - create the new command 'fine'
func New() *cobra.Command {
	rootCmd := &cobra.Command{Use: "fine"}
	rootCmd.AddCommand(
		formula.CreateCommand(),
		calc.CreateCommand(),
	)

	return rootCmd
}
