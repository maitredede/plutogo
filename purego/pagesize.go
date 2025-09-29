package plutopure

import (
	"github.com/jupiterrider/ffi"
)

var (
	ffiPageSizeType = ffi.NewType(&ffi.TypeFloat, &ffi.TypeFloat)
)

type PageSize struct {
	// _ structs.HostLayout

	Width  float32
	Height float32
}
