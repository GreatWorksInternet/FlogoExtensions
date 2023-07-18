package string

import (
	"errors"
	"fmt"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type ReplaceAllMultiple struct {
}

func init() {
	function.Register(&ReplaceAllMultiple{})
}

func (s *ReplaceAllMultiple) Name() string {
	return "replaceAllMultiple"
}

func (s *ReplaceAllMultiple) GetCategory() string {
	return "string"
}

func (s *ReplaceAllMultiple) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeArray}, false
}

func (s *ReplaceAllMultiple) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("string.replaceAllMultiple function first parameter [%+v] must be string", params[0])
	}
	replaceList, err := coerce.ToArray(params[1])
	if err != nil {
		return nil, fmt.Errorf("string.replaceAllMultiple function second parameter [%+v] must be array", params[1])
	}
	if len(replaceList)%2 == 1 {
		return nil, errors.New("string.replaceAllMultiple replaceList must not have an odd number of elements")
	}
	return replaceAllMultiple(str, replaceList), nil
}

func replaceAllMultiple(str string, replaceList []any) string {
	stringList := coerceToStringArray(replaceList)
	r := strings.NewReplacer(stringList...)
	return r.Replace(str)
}

func coerceToStringArray(arr []any) []string {
	result := []string{}
	for _, val := range arr {
		str, _ := coerce.ToString(val)
		result = append(result, str)
	}
	return result
}
