package formula

import (
	"fmt"

	"github.com/spf13/cobra"
)

func printFormula() {
	fmt.Println("The more point you get, the less you have to pay..")
}

// CreateCommand - Create new formula command
func CreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "formula",
		Short: "print out fine formula",
		Long:  "explain how this fine work, its formula to calculate the fine fee",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			printFormula()
		},
	}
}
