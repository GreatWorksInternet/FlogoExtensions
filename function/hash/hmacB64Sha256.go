package hash

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"encoding/base64"
	"crypto/sha256"
	"crypto/hmac"
)

type HmacB64Sha256 struct {
}

func init() {
	function.Register(&HmacB64Sha256{})
}

func (s *HmacB64Sha256) Name() string {
	return "hmacB64Sha256"
}

func (s *HmacB64Sha256) GetCategory() string {
	return "hash"
}

func (s *HmacB64Sha256) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *HmacB64Sha256) Eval(in ...interface{}) (interface{}, error) {

	body, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	key, err := coerce.ToString(in[1])
	if err != nil {
		return nil, err
	}
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(body))
	hash := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash), nil

}
