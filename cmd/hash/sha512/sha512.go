package sha512

import (
	hashService "github.com/Ice-nebula/nixl/internal/hash"
	"github.com/spf13/cobra"
)

func NewSha512Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sha512 [text]",
		Short: "Compute sha512 hash  for the provided text",
		Long:  "The 'sha512' command allows you to generate hash values for input text using sha512 algorithm.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			hasher := hashService.NewHashService()
			stringToHash := args[0]
			hashedString, err := hasher.Sha512(stringToHash)
			if err != nil {
				return err
			}
			cmd.Printf("result of %s string : %s", stringToHash, hashedString)
			return nil
		},
	} //end cmd
	return cmd
} //end method
