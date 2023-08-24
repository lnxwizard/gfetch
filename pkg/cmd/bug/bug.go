package bug

import (
	"github.com/cli/browser"
	"github.com/spf13/cobra"
)

func NewCmdBug() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bug",
		Short:   "Create a issue on GitHub",
		Long:    "Create a issue on GitHub with your default browser. If you see an bug/issue, please report the bug/issue with details on GitHub.",
		Example: "$ gfetch bug",
		RunE:    openOnBrowser,
	}

	return cmd
}

func openOnBrowser(cmd *cobra.Command, args []string) error {
	if err := browser.OpenURL("https://github.com/lnxwizard/gfetch/issues/new"); err != nil {
		return err
	}

	return nil
}
