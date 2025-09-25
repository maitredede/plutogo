package plutogo

/*
#cgo pkg-config: cairo freetype2 harfbuzz fontconfig expat icu-uc
#cgo LDFLAGS: -lplutobook

#include <plutobook.h>
#include <stdlib.h>

#include "callback.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"io"
	"time"
	"unsafe"
)

// Book represents a PlutoBook instance
type Book struct {
	ptr *C.plutobook_t
}

// NewBook creates a new PlutoBook instance
func NewBook(pageSize PageSize, margins PageMargins, mediaType MediaType) (*Book, error) {
	cPageSize := C.plutobook_page_size_t{
		width:  C.float(pageSize.Width),
		height: C.float(pageSize.Height),
	}

	cMediaType := C.plutobook_media_type_t(mediaType)

	cMargins := C.plutobook_page_margins_t{
		left:   C.float(margins.Left),
		top:    C.float(margins.Top),
		right:  C.float(margins.Right),
		bottom: C.float(margins.Bottom),
	}

	ptr := C.plutobook_create(cPageSize, cMargins, cMediaType)
	if ptr == nil {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return nil, errors.New(errMsg)
	}

	return &Book{ptr: ptr}, nil
}

// NewBookWithPresetPageSize creates a new PlutoBook instance with preset page size
func NewBookWithPresetPageSize(preset int, margins PageMargins, mediaType MediaType) (*Book, error) {
	var cPageSize C.plutobook_page_size_t
	var cMargins C.plutobook_page_margins_t
	var cMediaType C.plutobook_media_type_t

	// Convert preset to C enum
	switch preset {
	case 0: // A4
		cPageSize = C.PLUTOBOOK_PAGE_SIZE_A4
	case 1: // A3
		cPageSize = C.PLUTOBOOK_PAGE_SIZE_A3
	case 2: // Letter
		cPageSize = C.PLUTOBOOK_PAGE_SIZE_LETTER
	case 3: // Legal
		cPageSize = C.PLUTOBOOK_PAGE_SIZE_LEGAL
	default:
		return nil, errors.New("invalid page size preset")
	}

	// Convert margins
	switch margins {
	case PageMarginsNone:
		cMargins = C.PLUTOBOOK_PAGE_MARGINS_NONE
	case PageMarginsNarrow:
		cMargins = C.PLUTOBOOK_PAGE_MARGINS_NARROW
	case PageMarginsNormal:
		cMargins = C.PLUTOBOOK_PAGE_MARGINS_NORMAL
	case PageMarginsWide:
		cMargins = C.PLUTOBOOK_PAGE_MARGINS_WIDE
	default:
		cMargins = C.PLUTOBOOK_PAGE_MARGINS_NONE
	}

	// Convert media type
	switch mediaType {
	case MediaTypePrint:
		cMediaType = C.PLUTOBOOK_MEDIA_TYPE_PRINT
	case MediaTypeScreen:
		cMediaType = C.PLUTOBOOK_MEDIA_TYPE_SCREEN
	default:
		cMediaType = C.PLUTOBOOK_MEDIA_TYPE_SCREEN
	}

	ptr := C.plutobook_create(cPageSize, cMargins, cMediaType)
	if ptr == nil {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return nil, errors.New(errMsg)
	}

	return &Book{ptr: ptr}, nil
}

// Close releases the PlutoBook instance
func (b *Book) Close() {
	if b.ptr != nil {
		C.plutobook_destroy(b.ptr)
		b.ptr = nil
	}
}

func (b *Book) ClearContent() error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}
	C.plutobook_clear_content(b.ptr)
	return nil
}

// LoadHTML loads HTML content from a string
func (b *Book) LoadHTML(htmlContent string, userStyle, userScript string, baseURL string) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	cHTML := C.CString(htmlContent)
	defer C.free(unsafe.Pointer(cHTML))

	cUserStyle := C.CString(userStyle)
	defer C.free(unsafe.Pointer(cUserStyle))

	cUserScript := C.CString(userScript)
	defer C.free(unsafe.Pointer(cUserScript))

	cBaseURL := C.CString(baseURL)
	defer C.free(unsafe.Pointer(cBaseURL))

	result := C.plutobook_load_html(b.ptr, cHTML, -1, cUserStyle, cUserScript, cBaseURL)
	if !result {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return errors.New(errMsg)
	}

	return nil
}

// LoadURL loads HTML content from a file or URL
func (b *Book) LoadURL(url, userStyle, userScript string) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))

	cUserStyle := C.CString(userStyle)
	defer C.free(unsafe.Pointer(cUserStyle))

	sUserScript := C.CString(userScript)
	defer C.free(unsafe.Pointer(sUserScript))

	result := C.plutobook_load_url(b.ptr, cURL, cUserStyle, sUserScript)
	if !result {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return errors.New(errMsg)
	}

	return nil
}

// WriteToPDF exports the document as PDF
func (b *Book) WriteToPDF(filename string) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	result := C.plutobook_write_to_pdf(b.ptr, cFilename)
	if !result {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return errors.New(errMsg)
	}

	return nil
}

func (b *Book) WriteToPNG(filename string, width int, height int) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	cWidth := C.int(width)
	cHeight := C.int(height)

	result := C.plutobook_write_to_png(b.ptr, cFilename, cWidth, cHeight)
	if !result {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return errors.New(errMsg)
	}

	return nil
}

func (b *Book) WriteToPNGStream(writer io.Writer, width int, height int) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	// Créer le StreamWriter
	sw := &StreamWriter{writer: writer}

	// Enregistrer le StreamWriter
	id := registerStreamWriter(sw)
	defer unregisterStreamWriter(id)

	// Appeler la fonction C avec notre callback
	cWidth := C.int(width)
	cHeight := C.int(height)
	result := C.plutobook_write_to_png_stream(
		b.ptr,
		C.plutobook_stream_write_callback_t(C.stream_write_wrapper),
		unsafe.Pointer(&id),
		cWidth,
		cHeight,
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

func (b *Book) WriteToPDFStream(writer io.Writer) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	// Créer le StreamWriter
	sw := &StreamWriter{writer: writer}

	// Enregistrer le StreamWriter
	id := registerStreamWriter(sw)
	defer unregisterStreamWriter(id)

	// Appeler la fonction C avec notre callback
	result := C.plutobook_write_to_pdf_stream(
		b.ptr,
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

// WriteToPDFRange exports a range of pages as PDF
func (b *Book) WriteToPDFRange(filename string, startPage, endPage, step int) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	result := C.plutobook_write_to_pdf_range(b.ptr, cFilename, C.uint(startPage), C.uint(endPage), C.int(step))
	if !result {
		errPtr := C.plutobook_get_error_message()
		errMsg := C.GoString(errPtr)
		C.plutobook_clear_error_message()
		return errors.New(errMsg)
	}

	return nil
}

// GetPageCount returns the total number of pages
func (b *Book) GetPageCount() int {
	if b.ptr == nil {
		return 0
	}
	return int(C.plutobook_get_page_count(b.ptr))
}

// GetPageSize returns the page size in points
func (b *Book) GetPageSize() PageSize {
	if b.ptr == nil {
		return PageSize{}
	}

	size := C.plutobook_get_page_size(b.ptr)
	return PageSize{
		Width:  float64(size.width),
		Height: float64(size.height),
	}
}

// GetPageSize returns the page size in points
func (b *Book) GetPageSizeAt(index int) PageSize {
	if b.ptr == nil {
		return PageSize{}
	}

	size := C.plutobook_get_page_size_at(b.ptr, C.uint(index))
	return PageSize{
		Width:  float64(size.width),
		Height: float64(size.height),
	}
}

// GetDocumentWidth returns the document width
func (b *Book) GetDocumentWidth() float64 {
	if b.ptr == nil {
		return 0
	}
	return float64(C.plutobook_get_document_width(b.ptr))
}

// GetDocumentHeight returns the document height
func (b *Book) GetDocumentHeight() float64 {
	if b.ptr == nil {
		return 0
	}
	return float64(C.plutobook_get_document_height(b.ptr))
}

// RenderPage renders a specific page to the canvas
func (b *Book) RenderPage(canvas *Canvas, pageIndex int) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}
	if canvas.ptr == nil {
		return ErrCanvasIsClosed
	}

	/* result := */
	C.plutobook_render_page(b.ptr, canvas.ptr, C.uint(pageIndex))
	// if !result {
	// 	return errors.New("failed to render page")
	// }

	return nil
}

// RenderDocument renders the entire document to the canvas
func (b *Book) RenderDocument(canvas *Canvas) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}
	if canvas.ptr == nil {
		return ErrCanvasIsClosed
	}

	/* result := */
	C.plutobook_render_document(b.ptr, canvas.ptr)
	// if !result {
	// 	return errors.New("failed to render document")
	// }

	return nil
}

// WriteDocumentPNGToWriter génère un PNG du document complet dans un io.Writer
func (b *Book) WriteDocumentPNGToWriter(writer io.Writer, format ImageFormat) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}

	// Obtenir les dimensions réelles du document
	docWidth := int(b.GetDocumentWidth())
	docHeight := int(b.GetDocumentHeight())

	if docWidth <= 0 || docHeight <= 0 {
		return errors.New("invalid document dimensions")
	}

	// Créer un canvas de la taille du document
	canvas, err := NewImageCanvas(docWidth, docHeight, format)
	if err != nil {
		return fmt.Errorf("failed to create document canvas: %w", err)
	}
	defer canvas.Close()

	// Rendre le document complet
	err = b.RenderDocument(canvas)
	if err != nil {
		return fmt.Errorf("failed to render document: %w", err)
	}

	// Écrire dans le writer
	return canvas.WritePNGToWriter(writer)
}

func (b *Book) SetCreationDate(t time.Time) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}
	cValue := C.CString(t.Format(time.RFC3339))
	C.plutobook_set_metadata(b.ptr, C.PLUTOBOOK_PDF_METADATA_CREATION_DATE, cValue)
	return nil
}

func (b *Book) SetModificationDate(t time.Time) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}
	cValue := C.CString(t.Format(time.RFC3339))
	C.plutobook_set_metadata(b.ptr, C.PLUTOBOOK_PDF_METADATA_MODIFICATION_DATE, cValue)
	return nil
}

// GetMetadata Gets the value of the specified metadata
func (b *Book) GetMetadata(meta PdfMetadata) (string, error) {
	if b.ptr == nil {
		return "", ErrBookIsClosed
	}
	cMeta := C.plutobook_pdf_metadata_t(meta)
	cValue := C.plutobook_get_metadata(b.ptr, cMeta)
	value := C.GoString(cValue)
	return value, nil
}

func (b *Book) SetMetadata(meta PdfMetadata, value string) error {
	if b.ptr == nil {
		return ErrBookIsClosed
	}
	cMeta := C.plutobook_pdf_metadata_t(meta)
	cValue := C.CString(value)
	C.plutobook_set_metadata(b.ptr, cMeta, cValue)
	return nil
}
