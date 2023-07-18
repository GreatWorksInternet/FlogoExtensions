package hash

import (
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestB64Sha256_1(t *testing.T) {
	var B64Sha256In = &B64Sha256{}
	// Test case 1
	input1 := "hello"
	expectedOutput1 := "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ="
	output1, err := B64Sha256In.Eval(input1)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput1, output1, "Input: %s", input1)
}

func TestB64Sha256_2(t *testing.T) {
	var B64Sha256In = &B64Sha256{}
	// Test case 2
	input2 := "world"
	expectedOutput2 := "SG6kYiTRu0+2gPNPfJrZao8k7Ii+c+qOWmxlJg6cuKc="
	output2, err := B64Sha256In.Eval(input2)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput2, output2, "Input: %s", input2)
}

func TestB64Sha256_3(t *testing.T) {
	var B64Sha256In = &B64Sha256{}
	// Test case 3
	input3 := "1234567890"
	expectedOutput3 := "x3Xnt1ft5jDNCqERO9ECZhqziCnKUqZCKreChi8mhkY="
	output3, err := B64Sha256In.Eval(input3)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput3, output3, "Input: %s", input3)
}