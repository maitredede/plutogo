package plutopure

import (
	"github.com/jupiterrider/ffi"
)

var (
	ffiPageMarginsType = ffi.NewType(&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	// ffiPageMarginsType = ffi.NewType(&ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble)
)

type PageMargins struct {
	// _ structs.HostLayout

	Top    float32
	Right  float32
	Bottom float32
	Left   float32
}
