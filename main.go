package main

import (
	"errors"
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/epicseven-cup/gt/internal/ui/commit"
	"os"
	"os/exec"
	"strings"
)

const VERSION = "0.0.1"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("usage: gt <command>, gt -h")
		os.Exit(1)
	}
	fmt.Println(args)
	mode := args[0]
	switch mode {
	case "help":
		flag.Usage()
	case "version":
		fmt.Printf("gt version: %s\n", VERSION)
	case "commit":
		err := gtMessage(mode, args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		out, err := exec.Command("git", args...).Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}
}
