//go:build linux && cgo
// +build linux,cgo

package plctag

// acgo pkg-config: libplctag

// #cgo CFLAGS: -I${SRCDIR}/third_party/libplctag/src/lib -g -Wall
// #cgo LDFLAGS: -L${SRCDIR}/libs -lplctag -lutil
// #include <stdlib.h>
// #include "libplctag.h"
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// TagID ...
type TagID int32

//TagDecodeError bind #const char *plc_tag_decode_error(int err);
func TagDecodeError(errCode int) string {
	errStr := C.plc_tag_decode_error(C.int(errCode))
	//defer C.free(unsafe.Pointer(errStr))
	return C.GoString(errStr)
}

// TagCreate bind #int32_t plc_tag_create(const char *attrib_str, int timeout);
func TagCreate(tagAttrs string, timeout int) (TagID, error) {
	fmt.Printf("recieved attrs=%s", tagAttrs)
	attribStr := C.CString(tagAttrs)
	defer C.free(unsafe.Pointer(attribStr))
	id := int32(C.plc_tag_create(attribStr, C.int(timeout)))
	if id < 0 {
		return TagID(0), errors.New(TagDecodeError(int(id)))
	}
	return TagID(id), nil
}

// TagDestroy bind #int plc_tag_destroy(int32_t tag);
func TagDestroy(tag TagID) error {
	res := int(C.plc_tag_destroy(C.int32_t(tag)))

	if res < 0 {
		return errors.New(TagDecodeError(res))
	}
	return nil
}

//TagStatus bind #int plc_tag_status(int32_t tag)
func TagStatus(tag TagID) error {
	res := int(C.plc_tag_status(C.int32_t(tag)))
	if res < 0 {
		return errors.New(TagDecodeError(res))
	}
	return nil
}

//TagRead bind #int plc_tag_read(int32_t tag, int timeout);
func TagRead(tag TagID, timeout int) error {
	res := int(C.plc_tag_read(C.int32_t(tag), C.int(timeout)))
	if res < 0 {
		es := TagDecodeError(res)
		return errors.New(es)
	}
	return nil
}

// TagGetInt32 bind #int32_t plc_tag_get_int32(int32_t tag, int offset);
func TagGetInt32(tag TagID, offset int) (int32, error) {
	res := int32(C.plc_tag_get_int32(C.int32_t(tag), C.int(offset)))
	if res < 0 {
		return 0, errors.New(TagDecodeError(int(res)))
	}
	return res, nil
}

// TagGetUint8 bind #uint8_t plc_tag_get_uint8(int32_t tag, int offset);
func TagGetUint8(tag TagID, offset int) (uint8, error) {
	res := uint8(C.plc_tag_get_uint8(C.int32_t(tag), C.int(offset)))
	if res < 0 {
		return 0, errors.New(TagDecodeError(int(res)))
	}
	return res, nil
}
