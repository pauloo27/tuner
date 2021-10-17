package libmpv

//#include <mpv/client.h>
import "C"

import (
	"fmt"
)

type Error int

//const char *mpv_error_string(int error);
func NewError(errcode C.int) error {
	if errcode == C.MPV_ERROR_SUCCESS {
		return nil
	}
	return Error(errcode)
}

func (m Error) Error() string {
	return fmt.Sprintln(int(m), C.GoString(C.mpv_error_string(C.int(m))))
}
