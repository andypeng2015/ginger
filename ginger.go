package main

import (
	"fmt"
	"os"

	"github.com/ysugimoto/ginger/command"
	"github.com/ysugimoto/go-args"
)

func main() {
	ctx := args.New().
		Alias("help", "h", nil).
		Parse(os.Args[1:])

	var cmd command.Command
	switch ctx.At(0) {
	case command.INIT:
		cmd = command.NewInit()
	case command.CONFIG:
		cmd = command.NewConfig()
	case command.CREATE:
		cmd = command.NewCreate()
	case command.BUILD:
		cmd = command.NewBuild()

	default:
		if ctx.Len() == 0 {
			cmd = command.NewHelp()
		}
	}

	if cmd == nil {
		fmt.Printf("Command :%s not found. Abort.", ctx.At(0))
		os.Exit(1)
	}

	if ctx.Has("help") {
		fmt.Println(cmd.Help())
	} else if err := cmd.Run(ctx); err != nil {
		fmt.Printf("Command %s failed: %s\n", ctx.At(0), err.Error())
	}
}
