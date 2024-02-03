package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/luanmateuz/go-impl-ls-wc/cmd"
)

func main() {
	msg := "Please select a command, ls or wc."

	wcCmd := flag.NewFlagSet("wc", flag.ExitOnError)
	wcLines := wcCmd.Bool("l", false, "count lines")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println(msg)
		return
	}

	switch os.Args[1] {
	case "wc":
		wcCmd.Parse(os.Args[2:])
		fmt.Printf("%d\n", cmd.Wc(os.Stdin, *wcLines))
	default:
		fmt.Println(msg)
	}

}
