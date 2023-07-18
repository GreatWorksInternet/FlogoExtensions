package string

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestSuccessfulReplaceAllMultiple_Eval(t *testing.T) {
	var in = &ReplaceAllMultiple{}

	// Test with a string containing multiple occurrences of old values
	final, err := in.Eval("Frank & Danny", []interface{}{"&", "and", "Danny", "Joe"})
	assert.Nil(t, err)
	assert.Equal(t, "Frank and Joe", final)

	// Test with a string containing no old values
	final, err = in.Eval("Frank and Danny", []interface{}{"&", "and"})
	assert.Nil(t, err)
	assert.Equal(t, "Frank and Danny", final)
}

func TestFailingReplaceAllMultiple_Eval(t *testing.T) {
	var in = &ReplaceAllMultiple{}

	// Test with an odd number of elements in the replaceList
	final, err := in.Eval("Frank and Danny", []interface{}{"&"})
	assert.Empty(t, final)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "string.replaceAllMultiple replaceList must not have an odd number of elements")
}
