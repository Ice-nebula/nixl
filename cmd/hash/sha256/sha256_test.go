package sha256

import (
	"testing"

	"github.com/Ice-nebula/nixl/test"
	"github.com/stretchr/testify/assert"
)

func Test_sha256_with_empty_argument_should_be_get_error(t *testing.T) {
	cmd := NewSha256Command()
	actual := test.RunCobraCommand(*cmd, []string{}...)
	expected := "Error: accepts 1 arg(s), received 0"
	assert.Contains(t, actual, expected, "expected: %s,but got: %s", expected, actual)
}

func Test_sha256_command_with_non_empty_string_should_be_return_sha256hash(t *testing.T) {
	cmd := NewSha256Command()
	actual := test.RunCobraCommand(*cmd, "sha256")
	expected := "result of sha256 string : 5d5b09f6dcb2d53a5fffc60c4ac0d55fabdf556069d6631545f42aa6e3500f2e"
	assert.Equal(t, expected, actual, "expected : %s, but got: %s", expected, actual)
} //end method
