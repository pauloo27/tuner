package mpv

//From go-gl

/*
#cgo windows CFLAGS: -DTAG_WINDOWS
#cgo windows LDFLAGS: -lopengl32
#cgo darwin CFLAGS: -DTAG_DARWIN
#cgo darwin LDFLAGS: -framework OpenGL
#cgo linux CFLAGS: -DTAG_LINUX
#cgo linux LDFLAGS: -lGL
#cgo egl CFLAGS: -DTAG_EGL
#cgo egl LDFLAGS: -lEGL
// Check the EGL tag first as it takes priority over the platform's default
// configuration of WGL/GLX/CGL.
#if defined(TAG_EGL)
	#include <stdlib.h>
	#include <EGL/egl.h>
	void* glGetProcAddress(const char* name) {
		return eglGetProcAddress(name);
	}
#elif defined(TAG_WINDOWS)
	#define WIN32_LEAN_AND_MEAN 1
	#include <windows.h>
	#include <stdlib.h>
	static HMODULE ogl32dll = NULL;
	void* glGetProcAddress(const char* name) {
		void* pf = wglGetProcAddress((LPCSTR) name);
		if (pf) {
			return pf;
		}
		if (ogl32dll == NULL) {
			ogl32dll = LoadLibraryA("opengl32.dll");
		}
		return GetProcAddress(ogl32dll, (LPCSTR) name);
	}
#elif defined(TAG_DARWIN)
	#include <stdlib.h>
	#include <dlfcn.h>
	void* glGetProcAddress(const char* name) {
		return dlsym(RTLD_DEFAULT, name);
	}
#elif defined(TAG_LINUX)
	#include <stdlib.h>
	#include <GL/glx.h>
	void* glGetProcAddress(const char* name) {
		return glXGetProcAddress(name);
	}
#endif
*/
import "C"
import "unsafe"

func getProcAddress(namea string) unsafe.Pointer {
	cname := C.CString(namea)
	defer C.free(unsafe.Pointer(cname))
	return C.glGetProcAddress(cname)
}
