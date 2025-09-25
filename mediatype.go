package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
*/
import "C"

// MediaType represents the media type for rendering
type MediaType int

const (
	MediaTypePrint  MediaType = C.PLUTOBOOK_MEDIA_TYPE_PRINT
	MediaTypeScreen MediaType = C.PLUTOBOOK_MEDIA_TYPE_SCREEN
)
