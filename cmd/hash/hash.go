package hash

import (
	md5Cmd "github.com/Ice-nebula/nixl/cmd/hash/md5"
	sha256Cmd "github.com/Ice-nebula/nixl/cmd/hash/sha256"
	sha512Cmd "github.com/Ice-nebula/nixl/cmd/hash/sha512"
	"github.com/spf13/cobra"
)

func NewHashCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash <Algorithm> <text>",
		Short: "Compute various hash algorithms for the provided text",
		Long: `The 'hash' command allows you to generate hash values for input text using various algorithms.

Supported algorithms:
   - md5: Calculate an MD5 hash (e.g., 'nixl hash md5 "hello world"')
- sha256: Calculate a SHA-256 hash (e.g., 'nixl hash [text]'
   - sha512: Calculate a SHA-512 hash (e.g., 'nixl hash sha512 [text]')
`,
	} //end cmd
	cmd.AddCommand(sha512Cmd.NewSha512Command())
	cmd.AddCommand(sha256Cmd.NewSha256Command())
	cmd.AddCommand(md5Cmd.NewMd5Command())
	return cmd
} //end method
