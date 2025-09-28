package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
#include <stdlib.h>

#include "callback.h"
*/
import "C"
import (
	"fmt"
	"io"
	"log/slog"
	"sync"
	"unsafe"
)

type nativeCallback[T any] struct {
	dic   map[uintptr]T
	mut   sync.RWMutex
	count uintptr
}

func newNativeCallback[T any]() *nativeCallback[T] {
	return &nativeCallback[T]{
		dic:   make(map[uintptr]T),
		count: 1,
	}
}

func (c *nativeCallback[T]) register(item T) uintptr {
	c.mut.Lock()
	defer c.mut.Unlock()

	id := c.count
	c.count++
	c.dic[id] = item
	return id
}

func (c *nativeCallback[T]) unregister(id uintptr) {
	c.mut.Lock()
	defer c.mut.Unlock()

	delete(c.dic, id)
}

func (c *nativeCallback[T]) get(id uintptr) T {
	c.mut.RLock()
	defer c.mut.RUnlock()

	return c.dic[id]
}

// StreamWriter wraps a io.Writer for C callbacks
type StreamWriter struct {
	writer io.Writer
	err    error
	mutex  sync.Mutex
}

var (
	streamWriters   *nativeCallback[*StreamWriter]       = newNativeCallback[*StreamWriter]()
	resourceLoaders *nativeCallback[*resourceLoaderData] = newNativeCallback[*resourceLoaderData]()
)

//export goStreamWriteCallback
func goStreamWriteCallback(userData unsafe.Pointer, data unsafe.Pointer, length C.int) C.int {
	id := *(*uintptr)(userData)
	// bin := unsafe.Slice((*byte)(data), length)
	bin := C.GoBytes(data, length)

	sw := streamWriters.get(id)
	if sw == nil {
		return C.PLUTOBOOK_STREAM_STATUS_WRITE_ERROR
	}
	if sw.err != nil {
		return C.PLUTOBOOK_STREAM_STATUS_WRITE_ERROR
	}

	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	_, err := sw.writer.Write(bin)
	if err != nil {
		sw.err = err
		return C.PLUTOBOOK_STREAM_STATUS_WRITE_ERROR
	}

	return C.PLUTOBOOK_STREAM_STATUS_SUCCESS
}

//export goResourceLoaderCallback
func goResourceLoaderCallback(closure unsafe.Pointer, cUrl unsafe.Pointer) unsafe.Pointer {
	id := *(*uintptr)(closure)
	url := C.GoString((*C.char)(cUrl))

	rl := resourceLoaders.get(id)
	if rl == nil {
		return nil
	}
	if rl.err != nil {
		return nil
	}

	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	data, err := rl.loader(url)
	if err != nil {
		return nil
	}

	binPtr := (*C.char)(unsafe.Pointer((*byte)(&data.Bin[0])))
	binLen := C.uint(len(data.Bin))
	cMime := C.CString(data.Mime)
	cTextEncoding := C.CString(data.TextEncoding)
	res := C.plutobook_resource_data_create(binPtr, binLen, cMime, cTextEncoding)

	refCount := C.plutobook_resource_data_get_reference_count(res)
	slog.Debug(fmt.Sprintf("go: resource refCount=%d", refCount))

	return unsafe.Pointer(res)
}
