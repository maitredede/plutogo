package plutogo

/*
#cgo pkg-config: cairo freetype2 harfbuzz fontconfig expat icu-uc
#cgo LDFLAGS: -lplutobook

#include <plutobook.h>
#include <stdlib.h>
*/
import "C"

func Version() string {
	cVersion := C.plutobook_version_string()
	return C.GoString(cVersion)
}

func VersionNumber() int {
	cVersionNum := C.plutobook_version()
	return int(cVersionNum)
}
