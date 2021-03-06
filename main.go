package main

import (
	"fmt"
	"os"

	"github.com/landru29/etopo2stl/convert"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "etopo2stl",
	Short: "convert etopo xyz to stl",
}

func addCommands() {
	rootCmd.AddCommand(convert.Cmd)
}

func main() {
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
