package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
*/
import "C"

// PageSize represents a page size configuration
type PageSize struct {
	Width  float64
	Height float64
}

// Pre-defined page sizes
var (
	PageSizeA3     = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_A3, Height: C.PLUTOBOOK_PAGE_HEIGHT_A3}
	PageSizeA4     = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_A4, Height: C.PLUTOBOOK_PAGE_HEIGHT_A4}
	PageSizeA5     = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_A5, Height: C.PLUTOBOOK_PAGE_HEIGHT_A5}
	PageSizeB4     = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_B4, Height: C.PLUTOBOOK_PAGE_HEIGHT_B4}
	PageSizeB5     = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_B5, Height: C.PLUTOBOOK_PAGE_HEIGHT_B5}
	PageSizeLetter = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_LETTER, Height: C.PLUTOBOOK_PAGE_HEIGHT_LETTER}
	PageSizeLegal  = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_LEGAL, Height: C.PLUTOBOOK_PAGE_HEIGHT_LEGAL}
	PageSizeLedger = PageSize{Width: C.PLUTOBOOK_PAGE_WIDTH_LEDGER, Height: C.PLUTOBOOK_PAGE_HEIGHT_LEDGER}
)

func (s PageSize) Portrait() PageSize {
	if s.Width > s.Height {
		return PageSize{
			Width:  s.Height,
			Height: s.Width,
		}
	}
	return s
}

func (s PageSize) Landscape() PageSize {
	if s.Width < s.Height {
		return PageSize{
			Width:  s.Height,
			Height: s.Width,
		}
	}
	return s
}
