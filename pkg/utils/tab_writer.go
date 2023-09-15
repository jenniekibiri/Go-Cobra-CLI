package utils

import (
	"text/tabwriter"

	"github.com/spf13/cobra"
)


// NewTabWriter returns a new tab writer with the given padding, minwidth, tabwidth, and padding characters.
func NewTabWriter(cmd *cobra.Command) *tabwriter.Writer {
	return tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 1, ' ', tabwriter.TabIndent)
}
