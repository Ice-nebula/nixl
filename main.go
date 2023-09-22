/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Ice-nebula/nixl/cmd"
)

func main() {
	root := cmd.RootCommand()
	root.Execute()
}
