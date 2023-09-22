package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var stringForHash string = "testHash"

func Test_Shar256_with_string_for_hash_should_be_expected(t *testing.T) {
	service := NewHashService()
	expected := "fdd8157ddd7d2ade12a3799aa9998a8de76d291c1f3ddce3b3bb7edb2f42c7a8"
	actual, err := service.Sha256(stringForHash)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

} //end method

func Test_shar256_with_empty_string_should_be_error(t *testing.T) {
	service := NewHashService()
	actual, err := service.Sha256("")
	if assert.Error(t, err) {
		assert.Equal(t, "text can not be empty", err.Error())
		assert.Equal(t, "", actual)
	} //end if
} //end method

func Test_HashMd5_with_valid_string_should_get_correct_md5_hash(t *testing.T) {
	service := NewHashService()
	expected := "32269ae63a25306bb46a03d6f38bd2b7"
	actual, err := service.Md5("testmd5")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_HashMd5_with_empty_argument_should_be_error(t *testing.T) {
	service := NewHashService()
	expected := "text can not be empty"
	actual, err := service.Md5("")
	assert.Error(t, err)
	assert.Equal(t, expected, err.Error())
	assert.Equal(t, "", actual)
}
