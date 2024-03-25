/*
Copyright © 2024 Nikolay Grinko <ngrink.io@gmail.com>
*/

package main

import (
	"fmt"
	"os"

	"github.com/ngrink/tasker/cmd"
	"github.com/ngrink/tasker/core"
)

func main() {
	if err := core.Tasks.Load(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}
