package root

import (
	"fmt"
	cfg "gfetch/internal/config"
	sys "gfetch/internal/system"
	bugCmd "gfetch/pkg/cmd/bug"
	"gfetch/pkg/user"
	"os"
	s "strings"

	"github.com/logrusorgru/aurora/v4"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var fetch = `
    .--.	%s@%s
   |o_o |	%s
   |:_/ |	%s: %s
  //   \ \	%s: %s
 (|     | )     %s:  %s 	
/'\_   _/'\	%s: %s	
\___)=(___/ 	%s: %s

`

func init() {
	/*if err := cfg.Configs.Store("config.toml"); err != nil {
		os.Exit(2)
	}*/

	if err := cfg.Configs.Load("config.toml"); err != nil {
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
		Run:     newRunRoot,
	}

	cmd.AddCommand(bugCmd.NewCmdBug())

	return cmd
}

func newRunRoot(cmd *cobra.Command, args []string) {
	var (
		// titles
		titleDistro string = "distro"
		titleKernel string = "kernel"
		titleShell  string = "shell"
		titleMemory string = "memory"

		// values
		username string = user.GetUsername()
		hostname string = user.GetHostname()
		distro   string = sys.GetDistro()
		kernel   string = sys.GetKernel()
		shell    string = sys.GetShell()
		memory   string = sys.GetMemoryInfo()
	)

	if cfg.Configs.Title.BoldTitle {
		titleDistro = fmt.Sprint(aurora.Bold(titleDistro))
		titleKernel = fmt.Sprint(aurora.Bold(titleKernel))
		titleShell = fmt.Sprint(aurora.Bold(titleShell))
		titleMemory = fmt.Sprint(aurora.Bold(titleMemory))
	}
	if cfg.Configs.Title.ItalicTitle {
		titleDistro = fmt.Sprint(aurora.Italic(titleDistro))
		titleKernel = fmt.Sprint(aurora.Italic(titleKernel))
		titleShell = fmt.Sprint(aurora.Italic(titleShell))
		titleMemory = fmt.Sprint(aurora.Italic(titleMemory))
	}

	fmt.Printf(
		fetch,
		username, hostname,
		s.Repeat("-", len(user.GetUsername()+user.GetHostname())+1),
		titleDistro, distro,
		titleKernel, kernel,
		titleShell, shell,
		titleMemory, memory,
		"desktop", sys.GetDE(),
	)
}
