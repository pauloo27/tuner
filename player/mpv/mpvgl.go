package mpv

/*
#include <mpv/opengl_cb.h>
extern void* getProcAddr(void*fn_ctx,char* text);

*/
import "C"

import (
	"unsafe"
)

type MpvGL struct {
	ctx *C.mpv_opengl_cb_context
}

type get_proc_addr func(string) unsafe.Pointer

var callback_get_proc_address_fn get_proc_addr

//export getProcAddr
func getProcAddr(fn_ctx unsafe.Pointer, name *C.char) unsafe.Pointer {
	if callback_get_proc_address_fn != nil {
		return callback_get_proc_address_fn(C.GoString(name))
	}
	return nil
}

func (mgl *MpvGL) InitGL() error {
	callback_get_proc_address_fn = getProcAddress
	return NewError(C.mpv_opengl_cb_init_gl(mgl.ctx,
		nil,
		(*[0]byte)(C.getProcAddr),
		nil))
}

func (mgl *MpvGL) Draw(fbo, width, height int) int {
	return int(C.mpv_opengl_cb_draw(mgl.ctx, C.int(fbo), C.int(width), C.int(-height)))
}

//@deprecated
func (mgl *MpvGL) Render(fbo int, vp []int) int {
	return int(C.mpv_opengl_cb_render(mgl.ctx, C.int(fbo), (*C.int)(unsafe.Pointer(&vp[0]))))
}

func (mgl *MpvGL) ReportFlip(time int64) error {
	return NewError(C.mpv_opengl_cb_report_flip(mgl.ctx, C.int64_t(time)))
}

func (mgl *MpvGL) UninitGL() error {
	return NewError(C.mpv_opengl_cb_uninit_gl(mgl.ctx))
}
