package md5

import (
	"testing"

	"github.com/Ice-nebula/nixl/test"
	"github.com/stretchr/testify/assert"
)

func Test_md5_with_provide_text_should_be_get_md5_hash(t *testing.T) {
	cmd := NewMd5Command()
	actual := test.RunCobraCommand(*cmd, []string{"md5"}...)
	expected := "result of md5 string : 1bc29b36f623ba82aaf6724fd3b16718"
	assert.Equal(t, expected, actual, "expected: %s, but got: %s", expected, actual)
}

func Test_md5_with_empty_argument_should_be_get_error(t *testing.T) {
	cmd := NewMd5Command()
	actual := test.RunCobraCommand(*cmd, []string{}...)
	expected := "Error: accepts 1 arg(s), received 0"
	assert.Contains(t, actual, expected, "expected: %s, but got: %s", expected, actual)
} //end method
