package hash

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"encoding/base64"
	"crypto/sha256"
)

type B64Sha256 struct {
}

func init() {
	function.Register(&B64Sha256{})
}

func (s *B64Sha256) Name() string {
	return "b64Sha256"
}

func (s *B64Sha256) GetCategory() string {
	return "hash"
}

func (s *B64Sha256) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (s *B64Sha256) Eval(in ...interface{}) (interface{}, error) {

	input, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256([]byte(input))
	return base64.StdEncoding.EncodeToString(hash[:]), nil

}
