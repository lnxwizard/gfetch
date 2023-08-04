package system

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
)

func GetColors() string {
	red := aurora.Red("⬤")
	green := aurora.Green("⬤")
	yellow := aurora.Yellow("⬤")
	blue := aurora.Blue("⬤")
	magenta := aurora.Magenta("⬤")
	cyan := aurora.Cyan("⬤")
	s := ""

	return fmt.Sprintln(red, s, green, s, yellow, s, blue, s, magenta, s, cyan)
}
