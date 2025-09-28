package plutogo

/*
#cgo pkg-config: cairo freetype2 harfbuzz fontconfig expat icu-uc plutobook

#include <plutobook.h>
#include <stdlib.h>

#include "callback.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"io"
	"unsafe"
)

// Canvas represents a rendering canvas
type Canvas struct {
	ptr *C.plutobook_canvas_t
}

// NewImageCanvas creates a new image canvas
func NewImageCanvas(width, height int, format ImageFormat) (*Canvas, error) {
	var cFormat C.plutobook_image_format_t

	switch format {
	case ImageFormatARGB32:
		cFormat = C.PLUTOBOOK_IMAGE_FORMAT_ARGB32
	case ImageFormatRGB24:
		cFormat = C.PLUTOBOOK_IMAGE_FORMAT_RGB24
	case ImageFormatA8:
		cFormat = C.PLUTOBOOK_IMAGE_FORMAT_A8
	case ImageFormatA1:
		cFormat = C.PLUTOBOOK_IMAGE_FORMAT_A1
	default:
		return nil, errors.New("invalid image format")
	}

	ptr := C.plutobook_image_canvas_create(C.int(width), C.int(height), cFormat)
	if ptr == nil {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return nil, errors.New(errMsg)
	}

	return &Canvas{ptr: ptr}, nil
}

// Close releases the canvas
func (c *Canvas) Close() {
	if c.ptr != nil {
		C.plutobook_canvas_destroy(c.ptr)
		c.ptr = nil
	}
}

// ClearSurface clears the canvas with specified color
func (c *Canvas) ClearSurface(r, g, b, a float64) error {
	if c.ptr == nil {
		return ErrCanvasIsClosed
	}

	C.plutobook_canvas_clear_surface(c.ptr, C.float(r), C.float(g), C.float(b), C.float(a))
	return nil
}

// WriteToPNG saves the canvas as PNG
func (c *Canvas) WriteToPNG(filename string) error {
	if c.ptr == nil {
		return ErrCanvasIsClosed
	}

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	result := C.plutobook_image_canvas_write_to_png(c.ptr, cFilename)
	if !result {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return errors.New(errMsg)
	}

	return nil
}

// WritePNGToWriter écrit le contenu du canvas en PNG vers un io.Writer
func (c *Canvas) WritePNGToWriter(writer io.Writer) error {
	if c.ptr == nil {
		return ErrCanvasIsClosed
	}

	// Créer le StreamWriter
	sw := &StreamWriter{writer: writer}

	// Enregistrer le StreamWriter
	id := streamWriters.register(sw)
	defer streamWriters.unregister(id)

	// Appeler la fonction C avec notre callback
	result := C.plutobook_image_canvas_write_to_png_stream(
		c.ptr,
		C.plutobook_stream_write_callback_t(C.stream_write_wrapper),
		unsafe.Pointer(&id),
	)

	if !result {
		err := sw.err
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		if err != nil {
			return fmt.Errorf("%s: %w", errMsg, err)
		}
		return errors.New(errMsg)
	}

	return sw.err
}

// Flush Flushes any pending drawing operations on the canvas
func (c *Canvas) Flush() error {
	if c.ptr == nil {
		return ErrCanvasIsClosed
	}

	C.plutobook_canvas_flush(c.ptr)
	return nil
}

// Finish Finishes all drawing operations and performs cleanup on the canvas
func (c *Canvas) Finish() error {
	if c.ptr == nil {
		return ErrCanvasIsClosed
	}

	C.plutobook_canvas_finish(c.ptr)
	return nil
}
