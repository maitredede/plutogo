package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
*/
import "C"

// ImageFormat represents the image format for canvas
type ImageFormat int

const (
	ImageFormatInvalid ImageFormat = C.PLUTOBOOK_IMAGE_FORMAT_INVALID
	ImageFormatARGB32  ImageFormat = C.PLUTOBOOK_IMAGE_FORMAT_ARGB32
	ImageFormatRGB24   ImageFormat = C.PLUTOBOOK_IMAGE_FORMAT_RGB24
	ImageFormatA8      ImageFormat = C.PLUTOBOOK_IMAGE_FORMAT_A8
	ImageFormatA1      ImageFormat = C.PLUTOBOOK_IMAGE_FORMAT_A1
)
