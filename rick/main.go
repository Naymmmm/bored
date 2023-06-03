package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	showMessageBox("Error", "Your computer has encountered a critical error.")

	time.Sleep(2 * time.Second)

	showMessageBox("rick.exe", "You've been ricked!")

	time.Sleep(2 * time.Second)

	openBrowserWithVideo("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}

func showMessageBox(title, message string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	titlePtr, _ := syscall.UTF16PtrFromString(title)
	messagePtr, _ := syscall.UTF16PtrFromString(message)

	messageBox.Call(0, uintptr(unsafe.Pointer(messagePtr)), uintptr(unsafe.Pointer(titlePtr)), 0)
}

func openBrowserWithVideo(url string) {
	var err error

	switch runtime.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		fmt.Println("Unable to open the default web browser.")
		return
	}

	if err != nil {
		fmt.Println("An error occurred while opening the web browser:", err)
	}
}
