package plutogo

/*
#cgo pkg-config: cairo freetype2 harfbuzz fontconfig expat icu-uc
#cgo LDFLAGS: -lplutobook

#include <plutobook.h>
#include <stdlib.h>
*/
import "C"

var (
	UnitsPT float64 = C.PLUTOBOOK_UNITS_PT
	UnitsPC float64 = C.PLUTOBOOK_UNITS_PC
	UnitsIN float64 = C.PLUTOBOOK_UNITS_IN
	UnitsCM float64 = C.PLUTOBOOK_UNITS_CM
	UnitsMM float64 = C.PLUTOBOOK_UNITS_MM
	UnitsPX float64 = C.PLUTOBOOK_UNITS_PX
)
