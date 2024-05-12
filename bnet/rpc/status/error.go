package status

import (
	"fmt"

	"github.com/Gophercraft/core/bnet/rpc/codes"
)

type CodedError struct {
	Code codes.Code
	Err  error
}

func (coded_error *CodedError) Error() string {
	return coded_error.Err.Error()
}

func Errorf(code codes.Code, fm string, args ...any) error {
	coded_error := new(CodedError)
	coded_error.Code = code
	coded_error.Err = fmt.Errorf(fm, args...)
	return coded_error
}
