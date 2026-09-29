package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// Execute runs the CLI with the supplied streams and maps command errors to a non-zero exit status.
func Execute(args []string, stdout, stderr io.Writer) int {
	root := NewRootCommand()
	root.SetArgs(args)
	root.SetOut(stdout)
	root.SetErr(stderr)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	return 0
}

// NewRootCommand builds the command tree shared by the executable and future command tests.
func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "nexspace",
		Short: "Manage Nexspace projects",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
}
