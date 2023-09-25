package md5

import (
	hashService "github.com/Ice-nebula/nixl/internal/hash"
	"github.com/spf13/cobra"
)

func NewMd5Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "md5 [text]",
		Short: "Compute md5 hash  for the provided text",
		Long:  "The 'md5' command allows you to generate hash values for input text using md5 algorithm.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			hasher := hashService.NewHashService()
			textToHash := args[0]
			hashText, err := hasher.Md5(textToHash)
			if err != nil {
				return err
			}
			cmd.Printf("result of %s string : %s", textToHash, hashText)
			return nil
		},
	} //end cmd
	return cmd
} //end method
