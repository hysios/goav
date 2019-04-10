package avformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avio.h>
import "C"

import (
	"errors"
	"unsafe"

	"github.com/giorgisio/goav/avutil"
)

// int avio_open_dyn_buf(AVIOContext **s)
func AvioOpenDynBuf(ctxt **AvIOContext) error {
	return avutil.ErrorFromCode(int(C.avio_open_dyn_buf((**C.struct_AVIOContext)(unsafe.Pointer(ctxt)))))
}

func AvioGetBuffer(ctxt *AvIOContext) ([]byte, error) {
	var pbuf uintptr
	size := C.avio_get_dyn_buf((*C.struct_AVIOContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(&pbuf)))
	if size <= 0 {
		return nil, errors.New("invalid buffer")
	}
	defer C.avio_close_dyn_buf((*C.struct_AVIOContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(&pbuf)))
	buf := C.GoBytes(unsafe.Pointer(pbuf), size)
	return buf, nil
}

func (ctxt *AvIOContext) GetBuffer() ([]byte, error) {
	return AvioGetBuffer(ctxt)
}

func (ctxt *AvIOContext) Pause(on bool) error {
	var yes int
	if on {
		yes = 1
	}
	return avutil.ErrorFromCode(int(C.avio_pause((*C.struct_AVIOContext)(ctxt), (C.int)(yes))))
}
