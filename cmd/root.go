/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Ice-nebula/nixl/cmd/hash"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	var RootCmd = &cobra.Command{
		Use:   "nixl<command> <sub command> [flags]",
		Short: "versatile command-line utility",
		Long: `
Nixl is a versatile command-line utility crafted in Go, designed to streamline a range of developer tasks. It simplifies common operations, making them accessible through an intuitive command-line interface.
`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	RootCmd.AddCommand(hash.NewHashCommand())
	return RootCmd
} //end method.root command
