package ttf

import (
    "syscall"
    "unsafe"
    "sdl"
)

var (
    sdl_ttf, _ = syscall.LoadLibrary("SDL_ttf.dll")

    // SDL_ttf
    TTF_Init, _ = syscall.GetProcAddress(sdl_ttf, "TTF_Init")
    TTF_OpenFont, _ = syscall.GetProcAddress(sdl_ttf, "TTF_OpenFont")
    TTF_CloseFont, _ = syscall.GetProcAddress(sdl_ttf, "TTF_CloseFont")
    TTF_RenderUTF8_Solid, _ = syscall.GetProcAddress(sdl_ttf, "TTF_RenderUTF8_Solid")
    TTF_RenderUTF8_Shaded, _ = syscall.GetProcAddress(sdl_ttf, "TTF_RenderUTF8_Shaded")
    TTF_RenderUTF8_Blended, _ = syscall.GetProcAddress(sdl_ttf, "TTF_RenderUTF8_Blended")
)

type cast unsafe.Pointer

// Call Proc
func call(proc uint32, arg ... uintptr) uintptr {
    var args [12]uintptr
    var zero uint = 0
    for i := 0; i < 12; i++ {
        if i < len(arg) {
            args[i] = arg[i]
        } else {
            args[i] = uintptr(unsafe.Pointer(&zero))
        }
    }
    ret, _, _ := syscall.Syscall12(uintptr(proc),
        args[0], args[1], args[2], args[3], 
        args[4], args[5], args[6], args[7], 
        args[8], args[9], args[10], args[11])
    return uintptr(ret)
}

// TTF

// Initializes SDL_ttf.
func Init() int { return int(call(TTF_Init)) }

func OpenFont(file string, ptsize int) *Font {
    var font uintptr
    font = call(TTF_OpenFont, uintptr(cast(syscall.StringBytePtr(file))), uintptr(ptsize))
    return (*Font)(cast(font))
}

func (f *Font) Close() { call(TTF_CloseFont, uintptr(cast(f))) }

func RenderUTF8_Solid(font *Font, text string, color sdl.Color) *sdl.Surface {
    var surface uintptr
    var c uint32 = (uint32(color.B) << 16) | (uint32(color.G) << 8) | uint32(color.R)
    surface = call(TTF_RenderUTF8_Solid, uintptr(cast(font)), uintptr(cast(syscall.StringBytePtr(text))), uintptr(c))
    return (*sdl.Surface)(cast(surface))
}

func RenderUTF8_Shaded(font *Font, text string, color sdl.Color) *sdl.Surface {
    var surface uintptr
    var c uint32 = (uint32(color.B) << 16) | (uint32(color.G) << 8) | uint32(color.R)
    surface = call(TTF_RenderUTF8_Shaded, uintptr(cast(font)), uintptr(cast(syscall.StringBytePtr(text))), uintptr(c))
    return (*sdl.Surface)(cast(surface))
}

func RenderUTF8_Blended(font *Font, text string, color sdl.Color) *sdl.Surface {
    var surface uintptr
    var c uint32 = (uint32(color.B) << 16) | (uint32(color.G) << 8) | uint32(color.R)
    surface = call(TTF_RenderUTF8_Blended, uintptr(cast(font)), uintptr(cast(syscall.StringBytePtr(text))), uintptr(c))
    return (*sdl.Surface)(cast(surface))
}
