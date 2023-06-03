package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	fmt.Println("ls")
	showMessageBox("Hello Universe!", "This is a useless program!")
}

func showMessageBox(title, message string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	// Convert strings to UTF-16 encoding
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	messagePtr, _ := syscall.UTF16PtrFromString(message)

	// Call the MessageBoxW function
	messageBox.Call(0, uintptr(unsafe.Pointer(messagePtr)), uintptr(unsafe.Pointer(titlePtr)), 0)
}
