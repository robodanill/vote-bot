package handlers

import "fmt"

func handleHelp(cmd string) string {
	return fmt.Sprintf("[%s] Available commands:\ncreate\nvote\nresults\nend\ndelete\nhelp", cmd)
}
