package sshserver

import (
	"fmt"
	"log"
	"connect/conf"
	"connect/lib"
	"github.com/kr/pty"
	"golang.org/x/crypto/ssh"
)

func InitSsh(){
	Handle(StartSession)

	bind := fmt.Sprintf("[%s]:%d", conf.Setting.Sshd.ListenAddr, conf.Setting.Sshd.ListenPort)
	log.Printf("starting sshserver server on %v",bind)
	log.Fatal(ListenAndServe(bind, nil))
}

type EnvMsg struct{
	Key string
	Value string
}

func StartSession(s Session) {
	//cmd := exec.Command("/bin/bash")
	//ptyReq, winCh, isPty := s.Pty()
	//
	//if isPty {
	//	userHome := fmt.Sprintf("/Users/%s",s.User())
	//
	//	cmd.Dir = userHome
	//
	//	pwd := fmt.Sprintf("PWD=%s",userHome)
	//	home := fmt.Sprintf("HOME=%s",userHome)
	//	cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
	//	cmd.Env = append(cmd.Env,conf.Setting.Ps1)
	//	cmd.Env = append(cmd.Env,"LANG=zh_CN.utf8")
	//	cmd.Env = append(cmd.Env,"LC_ALL=zh_CN.utf8")
	//	cmd.Env = append(cmd.Env,pwd)
	//	cmd.Env = append(cmd.Env,home)
	//
	//	f, err := pty.Start(cmd)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//	go func() {
	//		for win := range winCh {
	//			setWinsize(f, win.Width, win.Height)
	//		}
	//	}()
	//	go func() {
	//		io.Copy(f, s) // stdin
	//	}()
	//	io.Copy(s, f) // stdout
	//
	//	fmt.Fprintf(f,"nihao %s color",color.Red("red"))
	//	//s.Write([]byte(d))
	//
	//} else {
	//	io.WriteString(s, "No PTY requested.\n")
	//	s.Exit(1)
	//}


	s.Write([]byte(lib.Title()))
	s.Write([]byte(lib.Usage()))
	s.Write([]byte(lib.Prompt("")))
	ptyReq, winCh, isPty := s.Pty()
	if isPty{
		msg := &EnvMsg{
			Key: "TERM",
			Value: ptyReq.Term,
		}
		s.SendRequest("env",true,ssh.Marshal(msg))

		pty1,tty,err := pty.Open()

		if err != nil{
			log.Fatalf("open pty err: %v",err)
		}
		defer tty.Close()


		go func() {
			for win := range winCh {
				setWinsize(pty1, win.Width, win.Height)
			}
		}()

		go func(){
			//io.Copy(pty1,s) //stdin
			copyBuffer(pty1,s,nil,writeCallBack,s)
		}()

		go func(){
			//io.Copy(s,pty1) //stdout
			copyBuffer(s,pty1,nil,readCallBack,s)
		}()

		select {

		}

	} else {
		log.Fatalf("is not pty")
	}

}

func writeCallBack(s Session, writedStr []byte)(error){
	log.Printf("user: %v Write data: %v",s.User(),string(writedStr))
	return nil
}

func readCallBack(s Session, writedStr []byte)(error){
	log.Printf("user: %v Recv data: %v",s.User(),string(writedStr))
	return nil
}