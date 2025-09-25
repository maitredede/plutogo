package plutogo

/*
#cgo pkg-config: cairo freetype2 harfbuzz fontconfig expat icu-uc
#cgo LDFLAGS: -lplutobook

#include <plutobook.h>
#include <stdlib.h>
*/
import "C"

func BuildInfo() string {
	cVersion := C.plutobook_build_info()
	return C.GoString(cVersion)
}
