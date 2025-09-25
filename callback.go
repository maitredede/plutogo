package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
#include <stdlib.h>

#include "callback.h"
*/
import "C"
import (
	"io"
	"sync"
	"unsafe"
)

// StreamWriter wraps a io.Writer for C callbacks
type StreamWriter struct {
	writer io.Writer
	err    error
	mutex  sync.Mutex
}

// global map for stream writers
var streamWriters = make(map[uintptr]*StreamWriter)
var streamWritersMutex sync.RWMutex
var streamWriterCounter uintptr = 1

// registerStreamWriter register a StreamWriter and returns its ID
func registerStreamWriter(sw *StreamWriter) uintptr {
	streamWritersMutex.Lock()
	defer streamWritersMutex.Unlock()

	id := streamWriterCounter
	streamWriterCounter++
	streamWriters[id] = sw
	return id
}

// unregisterStreamWriter remove a StreamWriter from the global map
func unregisterStreamWriter(id uintptr) {
	streamWritersMutex.Lock()
	defer streamWritersMutex.Unlock()

	delete(streamWriters, id)
}

// getStreamWriter get a StreamWriter with its ID
func getStreamWriter(id uintptr) *StreamWriter {
	streamWritersMutex.RLock()
	defer streamWritersMutex.RUnlock()

	return streamWriters[id]
}

//export goStreamWriteCallback
func goStreamWriteCallback(userData unsafe.Pointer, data unsafe.Pointer, length C.int) C.int {
	id := *(*uintptr)(userData)
	// bin := unsafe.Slice((*byte)(data), length)
	bin := C.GoBytes(data, length)

	sw := getStreamWriter(id)
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
