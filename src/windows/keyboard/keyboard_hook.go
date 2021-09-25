package keyboard

import (
	"fmt"
	"km/src/windows"
	"syscall"
	"unsafe"
)

var hook uintptr

func Hook(c chan<- windows.KeyboardEvent) {
	go func() {

		hook, _, _ = windows.SetWindowsHookExW.Call(uintptr(13), syscall.NewCallback(func(code int32, wParam uintptr, lParam uintptr) uintptr {
			if lParam != 0 {
				c <- windows.KeyboardEvent{
					Message:         wParam,
					KBDLLHOOKSTRUCT: *(*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
				}
			}
			call, _, _ := windows.CallNextHookEx.Call(0, uintptr(code), wParam, lParam)
			return call
		}), uintptr(0), uintptr(0))

		if hook == 0 {
			fmt.Println("keyboard failed hook")
		}
		var msg *windows.MSG
		for {
			r, _, _ := windows.GetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, uintptr(0), uintptr(0))
			if int32(r) != 0 {
				fmt.Printf("%v", msg)
			}
		}

	}()
}

func UnHook() {
	result, _, _ := windows.UnhookWindowsHookEx.Call(hook)
	if 0 != result {
		fmt.Println("keyboard failed unhook")
	}
}
