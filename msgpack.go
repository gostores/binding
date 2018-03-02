package binding

import (
	"net/http"

	"github.com/gostores/codec"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(req *http.Request, obj interface{}) error {
	if err := codec.NewDecoder(req.Body, new(codec.MsgpackHandle)).Decode(&obj); err != nil {
		return err
	}
	return validate(obj)
}
