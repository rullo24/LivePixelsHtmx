package cli

import (
	"fmt"
	"livepixelshtmx/internal/game"

	"github.com/spf13/cobra"
)

type CliConfig struct {
	NumCols int
	NumRows int
}

// DESCRIPTION
// Captures arguments from user/CLI into a CliConfig obj
func (c *CliConfig) ParseArgs() error {
	var rootCmd *cobra.Command = &cobra.Command{
		Use:     "Live Pixel Canvas",
		Short:   "A pixel canvas server that updates in real-time.",
		Example: `./livepixelshtmx --num-cols=3 --num-rows=3`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if c.NumCols <= 0 {
				return fmt.Errorf("num-cols must be +ve (num-cols=%d)", c.NumCols)
			}

			if c.NumRows <= 0 {
				return fmt.Errorf("num-rows must be +ve (num-rows=%d)", c.NumRows)
			}

			return nil
		},
	}

	rootCmd.Flags().IntVar(&c.NumCols, "num-cols", game.DEFAULT_NUM_COLUMNS, "number of grid columns")
	rootCmd.Flags().IntVar(&c.NumRows, "num-rows", game.DEFAULT_NUM_ROWS, "number of grid rows")

	return rootCmd.Execute()
}
