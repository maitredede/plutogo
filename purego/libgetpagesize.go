package plutopure

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
)

var (
	libGetPageSize func(book uintptr) PageSize
)

func registerFFIGetPageSize() {
	libGetPageSizeSym := mustGetSymbol("plutobook_get_page_size")

	var libGetPageSizeCIF ffi.Cif
	if ok := ffi.PrepCif(&libGetPageSizeCIF, ffi.DefaultAbi, 1, &ffiPageSizeType, &ffi.TypePointer); ok != ffi.OK {
		panic("plutobook_get_page_size cif prep is not OK")
	}

	libGetPageSize = func(book uintptr) PageSize {
		var ret PageSize
		args := []unsafe.Pointer{
			unsafe.Pointer(book),
		}
		ffi.Call(&libGetPageSizeCIF, libGetPageSizeSym, unsafe.Pointer(&ret), args...)
		return ret
	}
}
