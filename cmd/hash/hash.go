package hash

import (
	sha256Cmd "github.com/Ice-nebula/nixl/cmd/hash/sha256"
	"github.com/spf13/cobra"
)

func NewHashCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash <Algorithm> <text>",
		Short: "Compute various hash algorithms for the provided text",
		Long: `The 'hash' command allows you to generate hash values for input text using various algorithms.

Supported algorithms:
   - md5: Calculate an MD5 hash (e.g., 'nixl hash md5 "hello world"')
   - sha1: Calculate a SHA-1 hash (e.g., 'nixl hash sha1 [text]')
- sha256: Calculate a SHA-256 hash (e.g., 'nixl hash [text]'
   - sha512: Calculate a SHA-512 hash (e.g., 'nixl hash sha512 [text]')
`,
	} //end cmd
	cmd.AddCommand(sha256Cmd.NewSha256Command())
	return cmd
} //end method
