package windows

import "syscall"

var (
	user32DLL, _ = syscall.LoadDLL("user32.dll")

	CallNextHookEx, _      = user32DLL.FindProc("CallNextHookEx")
	SetWindowsHookExW, _   = user32DLL.FindProc("SetWindowsHookExW")
	GetMessageW, _         = user32DLL.FindProc("GetMessageW")
	UnhookWindowsHookEx, _ = user32DLL.FindProc("UnhookWindowsHookEx")
)

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-kbdllhookstruct
type KBDLLHOOKSTRUCT struct {
	VKCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
}

type KeyboardEvent struct {
	Message uintptr
	KBDLLHOOKSTRUCT
}

const (
	WH_KEYBOARD_LL uintptr = 13
	WH_MOUSE_LL    uintptr = 14
)

type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uint32
	LParam  uint32
	Time    uint32
	POINT
}

type POINT struct {
	X int32
	Y int32
}

type MouseEvent struct {
	Message uintptr
	MSLLHOOKSTRUCT
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msllhookstruct
type MSLLHOOKSTRUCT struct {
	POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
}
