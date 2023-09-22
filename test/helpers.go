package test

import (
	"bytes"
	"strings"

	"github.com/spf13/cobra"
)

func RunCobraCommand(c cobra.Command, args ...string) string {

	buffer := new(bytes.Buffer)
	c.SetOut(buffer)
	c.SetErr(buffer)
	c.SetArgs(args)
	c.Execute()
	return strings.TrimSpace(buffer.String())
} //end method
