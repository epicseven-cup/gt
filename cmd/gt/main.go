package gt

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/epicseven-cup/gt/internal/ui/commit"
)

const VERSION = "0.0.1"

func gt(mode string, args []string) {
	if err := tea.NewProgram(commit.NewModel()); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	flag.Parsed()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("usage: gt <command>, gt -h")
	}

	mode := args[0]
	switch mode {
	case "help":
		flag.Usage()
	case "version":
		fmt.Printf("gt version: %s\n", VERSION)
	case "commit":
		gt(mode, args[1:])
	default:
		fmt.Printf("unknown command: %s\n", mode)
	}
}
