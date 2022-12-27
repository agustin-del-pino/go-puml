package cmd

import "github.com/spf13/cobra"

var cli = &cobra.Command{
	Use:   "puml",
	Short: "CLI for Plant UML",
}

func Run() error {
	return cli.Execute()
}
