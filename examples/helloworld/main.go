package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"gitlab.com/gobl/gobl/pkg/cmd"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}

func main() {
	cmd.SetRootCmd(rootCmd)
	cmd.Execute(context.Background(), nil, nil)
}
