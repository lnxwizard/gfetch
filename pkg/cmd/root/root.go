package root

import (
	"fmt"
	cfg "gfetch/internal/config"
	sys "gfetch/internal/system"
	"gfetch/pkg/user"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"os"
	s "strings"
)

var fetch = `
    .--.	%s@%s
   |o_o |	%s
   |:_/ |	distro: %s
  //   \ \	kernel: %s
 (|     | )     shell:  %s 	
/'\_   _/'\	memory: %s	
\___)=(___/     %s

`

func init() {
	if err := cfg.Configs.Load("/.gfetch/.config.toml"); err != nil {
		os.Exit(2)
	}
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gfetch",
		Short: "A minimal system fetch written in Go.",
		Long:  "A minimal system fetch for Linux written in Go. You can customize `gfetch` by editing `.gfetch.toml` file in home directory.",
		Example: heredoc.Doc(`
		$ gfetch
		$ gfetch --help`),
		Version: "0.1.0",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Printf(
				fetch,
				user.GetUsername(),
				user.GetHostname(),
				s.Repeat("-", len(user.GetUsername()+user.GetHostname())+1),
				sys.GetDistro(),
				sys.GetKernel(),
				sys.GetShell(),
				sys.GetMemoryInfo(),
				sys.GetColors(),
			)

			return nil
		},
	}

	return cmd
}
