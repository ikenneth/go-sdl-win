// SDL binding from SDL.dll
package sdl

import (
    "syscall"
    "unsafe"
)

var (
    // Load DLL
    sdl, _ = syscall.LoadLibrary("SDL.dll")
    sdl_image, _ = syscall.LoadLibrary("SDL_image.dll")
    
    // SDL
    SDL_Init, _ = syscall.GetProcAddress(sdl, "SDL_Init")
    SDL_Quit, _ = syscall.GetProcAddress(sdl, "SDL_Quit")
    SDL_GetError, _ = syscall.GetProcAddress(sdl, "SDL_GetError")
    SDL_SetVideoMode, _ = syscall.GetProcAddress(sdl, "SDL_SetVideoMode")
    SDL_WM_SetCaption, _ = syscall.GetProcAddress(sdl, "SDL_WM_SetCaption")
    SDL_FillRect, _ = syscall.GetProcAddress(sdl, "SDL_FillRect")
    SDL_SetAlpha, _ = syscall.GetProcAddress(sdl, "SDL_SetAlpha")
    SDL_InitSubSystem, _ = syscall.GetProcAddress(sdl, "SDL_InitSubSystem")
    SDL_PollEvent, _ = syscall.GetProcAddress(sdl, "SDL_PollEvent")
    SDL_Flip, _ = syscall.GetProcAddress(sdl, "SDL_Flip")
    SDL_GetTicks, _ = syscall.GetProcAddress(sdl, "SDL_GetTicks")
    SDL_Delay, _ = syscall.GetProcAddress(sdl, "SDL_Delay")
    SDL_UpperBlit, _ = syscall.GetProcAddress(sdl, "SDL_UpperBlit")
    SDL_GetMouseState, _ = syscall.GetProcAddress(sdl, "SDL_GetMouseState")
    SDL_FreeSurface, _ = syscall.GetProcAddress(sdl, "SDL_FreeSurface")
    
    // SDL_image
    IMG_Load, _ = syscall.GetProcAddress(sdl_image, "IMG_Load")
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

// General

// Initializes SDL.
func Init(flags uint32) int { return int(call(SDL_Init, uintptr(flags))) }

// Shuts down SDL
func Quit() { call(SDL_Quit) }

// Initializes subsystems.
// func InitSubSystem(flags uint32) int { return int(C.SDL_InitSubSystem(C.Uint32(flags))) }

// Shuts down a subsystem.
// func QuitSubSystem(flags uint32) { C.SDL_QuitSubSystem(C.Uint32(flags)) }

// Checks which subsystems are initialized.
// func WasInit(flags uint32) int { return int(C.SDL_WasInit(C.Uint32(flags))) }

// Error Handling

// Gets SDL error string
func GetError() string { return string(call(SDL_GetError)) }

// Set a string describing an error to be submitted to the SDL Error system.
// func SetError(description string) {
	// cdescription := C.CString(description)
	// C.SetError(cdescription)
	// C.free(unsafe.Pointer(cdescription))
// }

// TODO SDL_Error

// Clear the current SDL error
// func ClearError() { C.SDL_ClearError() }

// Video

// Sets up a video mode with the specified width, height, bits-per-pixel and
// returns a corresponding surface.  You don't need to call the Free method
// of the returned surface, as it will be done automatically by sdl.Quit.
func SetVideoMode(w int, h int, bpp int, flags uint32) *Surface {
	var screen = call(SDL_SetVideoMode, uintptr(w), uintptr(h), uintptr(bpp), uintptr(flags))
	return (*Surface)(cast(screen))
}

// Returns a pointer to the current display surface.
// func GetVideoSurface() *Surface { return (*Surface)(cast(C.SDL_GetVideoSurface())) }

// Checks to see if a particular video mode is supported.  Returns 0 if not
// supported, or the bits-per-pixel of the closest available mode.
// func VideoModeOK(width int, height int, bpp int, flags uint32) int {
	// return int(C.SDL_VideoModeOK(C.int(width), C.int(height), C.int(bpp), C.Uint32(flags)))
// }

// func ListModes(format *PixelFormat, flags uint32) []Rect {
	// modes := C.SDL_ListModes((*C.SDL_PixelFormat)(cast(format)), C.Uint32(flags))
	// if modes == nil { //no modes available
		// return make([]Rect, 0)
	// }
	// var any int
	// *((***C.SDL_Rect)(unsafe.Pointer(&any))) = modes
	// if any == -1 { //any dimension is ok
		// return nil
	// }

	// var count int
	// ptr := *modes //first element in the list
	// for count = 0; ptr != nil; count++ {
		// ptr = *(**C.SDL_Rect)(unsafe.Pointer( uintptr(unsafe.Pointer(modes)) + uintptr(count * unsafe.Sizeof(ptr)) ))
	// }
	// var ret = make([]Rect, count - 1)

	// *((***C.SDL_Rect)(unsafe.Pointer(&ret))) = modes // TODO
	// return ret
// }

// type VideoInfo struct {
	// HW_available bool         "Flag: Can you create hardware surfaces?"
	// WM_available bool         "Flag: Can you talk to a window manager?"
	// Blit_hw      bool         "Flag: Accelerated blits HW --> HW"
	// Blit_hw_CC   bool         "Flag: Accelerated blits with Colorkey"
	// Blit_hw_A    bool         "Flag: Accelerated blits with Alpha"
	// Blit_sw      bool         "Flag: Accelerated blits SW --> HW"
	// Blit_sw_CC   bool         "Flag: Accelerated blits with Colorkey"
	// Blit_sw_A    bool         "Flag: Accelerated blits with Alpha"
	// Blit_fill    bool         "Flag: Accelerated color fill"
	// Video_mem    uint32       "The total amount of video memory (in K)"
	// Vfmt         *PixelFormat "Value: The format of the video surface"
	// Current_w    int32        "Value: The current video mode width"
	// Current_h    int32        "Value: The current video mode height"
// }

// func GetVideoInfo() *VideoInfo {
	// vinfo := (*internalVideoInfo)(cast(C.SDL_GetVideoInfo()))

	// flags := vinfo.Flags

	// return &VideoInfo{
		// HW_available: flags&(1<<0) != 0,
		// WM_available: flags&(1<<1) != 0,
		// Blit_hw:      flags&(1<<9) != 0,
		// Blit_hw_CC:   flags&(1<<10) != 0,
		// Blit_hw_A:    flags&(1<<11) != 0,
		// Blit_sw:      flags&(1<<12) != 0,
		// Blit_sw_CC:   flags&(1<<13) != 0,
		// Blit_sw_A:    flags&(1<<14) != 0,
		// Blit_fill:    flags&(1<<15) != 0,
		// Video_mem:    vinfo.Video_mem,
		// Vfmt:         vinfo.Vfmt,
		// Current_w:    vinfo.Current_w,
		// Current_h:    vinfo.Current_h,
	// }
// }

// Makes sure the given area is updated on the given screen.  If x, y, w, and
// h are all 0, the whole screen will be updated.
// func (screen *Surface) UpdateRect(x int32, y int32, w uint32, h uint32) {
	// C.SDL_UpdateRect((*C.SDL_Surface)(cast(screen)), C.Sint32(x), C.Sint32(y), C.Uint32(w), C.Uint32(h))
// }

// func (screen *Surface) UpdateRects(rects []Rect) {
	// if len(rects) > 0 {
		// C.SDL_UpdateRects((*C.SDL_Surface)(cast(screen)), C.int(len(rects)), (*C.SDL_Rect)(cast(&rects[0])))
	// }
// }

// Gets the window title and icon name.
// func WM_GetCaption() (title, icon string) {
	// SDL seems to free these strings.  TODO: Check to see if that's the case
	// var ctitle, cicon *C.char
	// C.SDL_WM_GetCaption(&ctitle, &cicon)
	// title = C.GoString(ctitle)
	// icon = C.GoString(cicon)
	// return
// }

// Sets the window title and icon name.
func WM_SetCaption(title, icon string) {
    call(SDL_WM_SetCaption, uintptr(cast(syscall.StringBytePtr(title))), uintptr(cast(syscall.StringBytePtr(icon))))
}

// Sets the icon for the display window.
// func WM_SetIcon(icon *Surface, mask *uint8) {
	// C.SDL_WM_SetIcon((*C.SDL_Surface)(cast(icon)), (*C.Uint8)(mask))
// }

// Minimizes the window
// func WM_IconifyWindow() int { return int(C.SDL_WM_IconifyWindow()) }

// Toggles fullscreen mode
// func WM_ToggleFullScreen(surface *Surface) int {
	// return int(C.SDL_WM_ToggleFullScreen((*C.SDL_Surface)(cast(surface))))
// }

// Swaps OpenGL framebuffers/Update Display.
// func GL_SwapBuffers() { C.SDL_GL_SwapBuffers() }

// func GL_SetAttribute(attr int, value int) int {
  // return int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
// }

// Swaps screen buffers.
func (screen *Surface) Flip() int { return int(call(SDL_Flip, uintptr(cast(screen)))) }

// Frees (deletes) a Surface
func (screen *Surface) Free() { call(SDL_FreeSurface, (uintptr(cast(screen)))) }

// Locks a surface for direct access.
// func (screen *Surface) Lock() int {
	// return int(C.SDL_LockSurface((*C.SDL_Surface)(cast(screen))))
// }

// Unlocks a previously locked surface.
// func (screen *Surface) Unlock() int { return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

func (dst *Surface) Blit(dstrect *Rect, src *Surface, srcrect *Rect) int {
	var ret = call(SDL_UpperBlit,
        uintptr(cast(src)), uintptr(cast(srcrect)),
        uintptr(cast(dst)), uintptr(cast(dstrect)))
	return int(ret)
}

// This function performs a fast fill of the given rectangle with some color.
func (dst *Surface) FillRect(dstrect *Rect, color uint32) int {
	var ret = call(SDL_FillRect, uintptr(cast(dst)), uintptr(cast(dstrect)), uintptr(color))
	return int(ret)
}

// Adjusts the alpha properties of a Surface.
func (s *Surface) SetAlpha(flags uint32, alpha uint8) int {
	return int(call(SDL_SetAlpha, uintptr(cast(s)), uintptr(flags), uintptr(alpha)))
}

// Gets the clipping rectangle for a surface.
// func (s *Surface) GetClipRect(r *Rect) {
	// C.SDL_GetClipRect((*C.SDL_Surface)(cast(s)), (*C.SDL_Rect)(cast(r)))
	// return
// }

// Sets the clipping rectangle for a surface.
// func (s *Surface) SetClipRect(r *Rect) {
	// C.SDL_SetClipRect((*C.SDL_Surface)(cast(s)), (*C.SDL_Rect)(cast(r)))
	// return
// }

// Map a RGBA color value to a pixel format.
// func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	// return (uint32)(C.SDL_MapRGBA((*C.SDL_PixelFormat)(cast(format)), (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b), (C.Uint8)(a)))
// }

// Gets RGBA values from a pixel in the specified pixel format.
// func GetRGBA(color uint32, format *PixelFormat, r, g, b, a *uint8) {
	// C.SDL_GetRGBA(C.Uint32(color), (*C.SDL_PixelFormat)(cast(format)), (*C.Uint8)(r), (*C.Uint8)(g), (*C.Uint8)(b), (*C.Uint8)(a))
// }

// Loads Surface from file (using IMG_Load).
func Load(file string) *Surface {
	var screen = call(IMG_Load, uintptr(cast(syscall.StringBytePtr(file))))
	return (*Surface)(cast(screen))
}

// Creates an empty Surface.
// func CreateRGBSurface(flags uint32, width int, height int, bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) *Surface {
	// p := C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), C.int(bpp),
		// C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))
	// return (*Surface)(cast(p))
// }

// Events

// Enables UNICODE translation.
// func EnableUNICODE(enable int) int { return int(C.SDL_EnableUNICODE(C.int(enable))) }

// Sets keyboard repeat rate.
// func EnableKeyRepeat(delay, interval int) int {
	// return int(C.SDL_EnableKeyRepeat(C.int(delay), C.int(interval)))
// }

// Gets keyboard repeat rate.
// func GetKeyRepeat() (int, int) {

	// var delay int
	// var interval int

	// C.SDL_GetKeyRepeat((*C.int)(cast(&delay)), (*C.int)(cast(&interval)))

	// return delay, interval
// }

// Gets a snapshot of the current keyboard state
// func GetKeyState() []uint8 {
	// var numkeys C.int
	// array := C.SDL_GetKeyState(&numkeys)

	// var ptr = make([]uint8, numkeys)

	// *((**C.Uint8)(unsafe.Pointer(&ptr))) = array // TODO

	// return ptr

// }

// Modifier
type Mod int

// Key
type Key int

// Retrieves the current state of the mouse.
func GetMouseState(x, y *int) uint8 {
	return uint8(call(SDL_GetMouseState, uintptr((cast(x))), uintptr(cast(y))))
}

// Retrieves the current state of the mouse relative to the last time this
// function was called.
// func GetRelativeMouseState(x, y *int) uint8 {
	// return uint8(C.SDL_GetRelativeMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
// }

// Gets the state of modifier keys
// func GetModState() Mod { return Mod(C.SDL_GetModState()) }

// Sets the state of modifier keys
// func SetModState(modstate Mod) { C.SDL_SetModState(C.SDLMod(modstate)) }

// Gets the name of an SDL virtual keysym
// func GetKeyName(key Key) string { return C.GoString(C.SDL_GetKeyName(C.SDLKey(key))) }

// Events

// Waits indefinitely for the next available event
// func (event *Event) Wait() bool {
	// var ret = C.SDL_WaitEvent((*C.SDL_Event)(cast(event)))
	// return ret != 0
// }

// Polls for currently pending events
func (event *Event) Poll() bool {
	var ret = call(SDL_PollEvent,uintptr(cast(event)))
	return ret != 0
}

// Returns KeyboardEvent or nil if event has other type
func (event *Event) Keyboard() *KeyboardEvent {
	if event.Type == KEYUP || event.Type == KEYDOWN {
		return (*KeyboardEvent)(cast(event))
	}

	return nil
}

// Returns MouseButtonEvent or nil if event has other type
// func (event *Event) MouseButton() *MouseButtonEvent {
	// if event.Type == MOUSEBUTTONDOWN || event.Type == MOUSEBUTTONUP {
		// return (*MouseButtonEvent)(cast(event))
	// }

	// return nil
// }

// Returns MouseMotion or nil if event has other type
// func (event *Event) MouseMotion() *MouseMotionEvent {
	// if event.Type == MOUSEMOTION {
		// return (*MouseMotionEvent)(cast(event))
	// }

	// return nil
// }

// Returns ActiveEvent or nil if event has other type
// func (event *Event) Active() *ActiveEvent {
  // if event.Type == ACTIVEEVENT {
    // return (*ActiveEvent)(cast(event))
  // }

  // return nil
// }

// Returns ResizeEvent or nil if event has other type
// func (event *Event) Resize() *ResizeEvent {
  // if event.Type == VIDEORESIZE {
    // return (*ResizeEvent)(cast(event))
  // }

  // return nil
// }

// Time

// Gets the number of milliseconds since the SDL library initialization.
func GetTicks() uint32 { return uint32(call(SDL_GetTicks)) }

// Waits a specified number of milliseconds before returning.
func Delay(ms uint32) { call(SDL_Delay, uintptr(ms)) }
