package main

import (
	"os"
	"syscall"
	"unsafe"

	"flag"
	"connect/conf"
	"github.com/golang/glog"
	"connect/sshserver"
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

var cf = flag.String("config","config.yaml","config path")

func main() {
	var err error
	flag.Parse()
	config := conf.NewSettings()
	conf.Setting,err = config.Load(*cf)
	if err != nil {
		glog.Fatalf("load config: %v err: %v",*cf,err)
	}

	sshserver.InitSsh()
}
