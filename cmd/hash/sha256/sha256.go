package sha256

import (
	hashService "github.com/Ice-nebula/nixl/internal/hash"
	"github.com/spf13/cobra"
)

func NewSha256Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sha256 [text]",
		Short: "Compute sha256 hash  for the provided text",
		Long:  "The 'sha256' command allows you to generate hash values for input text using sha256 algorithm.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			hasher := hashService.NewHashService()
			textToHash := args[0]
			hashText, err := hasher.Sha256(textToHash)
			if err != nil {
				return err
			}
			cmd.Printf("result of %s string : %s", textToHash, hashText)
			return nil
		},
	} //end cmd
	return cmd
} //end method
