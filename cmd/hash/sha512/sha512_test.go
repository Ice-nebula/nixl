package sha512

import (
	"testing"

	"github.com/Ice-nebula/nixl/test"
	"github.com/stretchr/testify/assert"
)

func Test_sha512_with_text_should_be_get_sha512hash(t *testing.T) {
	cmd := NewSha512Command()
	actual := test.RunCobraCommand(*cmd, []string{"sha512"}...)
	expected := "result of sha512 string : 1f9720f871674c18e5fecff61d92c1355cd4bfac25699fb7ddfe7717c9669b4d085193982402156122dfaa706885fd64741704649795c65b2a5bdec40347e28a"
	assert.Equal(t, expected, actual, "expected: %s, but got: %s", expected, actual)
}
