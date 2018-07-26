package jpush_im

import (
	"fmt"
	"encoding/json"
)


var _ error = (*Errors)(nil)
type Errors struct {
	Err struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (e *Errors) Error() string {
	bs, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("code: %q, message: %q", e.Err.Code, e.Err.Message)
	}
	return string(bs)
}


