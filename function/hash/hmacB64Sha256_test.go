package hash

import (
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

var HmacB64Sha256In = &HmacB64Sha256{}

func TestHmacB64Sha256_1(t *testing.T) {
	// Test case 1
	key1 := "secret"
	body1 := "hello"
	expectedOutput1 := "iKqz7ejTrflNJquQ07r9SiCDBww7zOnAFO4EpEOEfAs="
	output1, err := HmacB64Sha256In.Eval(body1, key1)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput1, output1, "Key: %s, Body: %s", key1, body1)
}

func TestHmacB64Sha256_2(t *testing.T) {
	// Test case 2
	key2 := "password"
	body2 := "world"
	expectedOutput2 := "PB0mXvA6XMyX9EaXmGhTi2iXxwNZG32f7T31y2mSFds="
	output2, err := HmacB64Sha256In.Eval(body2, key2)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput2, output2, "Key: %s, Body: %s", key2, body2)
}

func TestHmacB64Sha256_3(t *testing.T) {
	// Test case 3
	key3 := "123456"
	body3 := "abcdef"
	expectedOutput3 := "kTFrn+pNxeJdPWcJSVN+URt4kL0uL00GKDNhP4Ulv48="
	output3, err := HmacB64Sha256In.Eval(body3, key3)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput3, output3, "Key: %s, Body: %s", key3, body3)
}