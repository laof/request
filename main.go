package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/atotto/clipboard"
)

type ProxyData struct {
	code   string
	online string
}

func main() {

	chaDirect := make(chan string)
	chaProxy := make(chan ProxyData)

	proxyData := ProxyData{"", ""}

	go Direct(chaDirect)
	go Proxy(chaProxy, proxyData)

	direct := <-chaDirect
	proxyData = <-chaProxy

	if direct != "" {
		copy(direct)
	} else if proxyData.code != "" {
		copy(proxyData.code)
	} else if proxyData.online != "" {
		openChrome(proxyData.online)
	} else {
		fmt.Println("Sorry, connection failed please try again")
		time.Sleep(time.Second)
	}

}

func intPtr(n int) uintptr {
	return uintptr(n)
}
func strPtr(s string) uintptr {
	str, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(str))
}

func copy(str string) {
	fmt.Println(str)
	clipboard.WriteAll(str)
	showMessage("Successfully", "Copy the successfully paste the board")
}

func showMessage(tittle, text string) {
	user32dll, _ := syscall.LoadLibrary("user32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(intPtr(0), strPtr(text), strPtr(tittle), intPtr(0))
	defer syscall.FreeLibrary(user32dll)
}

func openChrome(url string) {

	commands := map[string][]string{
		"windows": {"cmd", "/c", "start"},
		"darwin":  {"open"},
		"linux":   {"xdg-open"},
	}

	run, ok := commands[runtime.GOOS]
	if !ok {
		fmt.Println("please click this link " + url)
	} else {
		run = append(run, url)
		cmd := exec.Command(run[0], run[1:]...)
		cmd.Start()
	}

}
