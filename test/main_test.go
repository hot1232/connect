package test

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/crypto/ssh"
	"time"
	"net"
	"io"
	"testing"
)

func Test_Main(t *testing.T) {
	config := &ssh.ClientConfig{
		User: "zhangliang",
		Auth: []ssh.AuthMethod{
			ssh.Password("qwe123!"),
		},
		//Config: sshserver.Config{
		//	Ciphers: []string{"aes128-cbc"},
		//},
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// config.Config.Ciphers = append(config.Config.Ciphers, "aes128-cbc")
	clinet, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	checkError(err, "连接交换机")

	session, err := clinet.NewSession()
	defer session.Close()
	checkError(err, "创建shell")

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm-256color", 80, 40, modes); err != nil {
		log.Fatal(err)
	}

	w, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	o, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}
	_, err = session.StderrPipe()
	if err != nil {
		panic(err)
	}


	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}
	go func(){
		var idx int
		for{
			idx ++
			w.Write([]byte("ls\r\n"))
			time.Sleep(time.Second*2)
			if idx>5{
				w.Write([]byte("exit\n"))
			}
		}
	}()

	data := make([]byte,4096)

	go func(){

		for{

			c,err := o.Read(data)
			if err == io.EOF{
				fmt.Printf("read data from stdout :EOF")
				break
			}
			if err != nil {
				fmt.Printf("read data from stdout err: %v",err)
			}
			if c > 0 {
				fmt.Printf("return: %v", string(data))
			}

		}

	}()


	session.Wait()
}

func checkError(err error, info string) {
	if err != nil {
		fmt.Printf("%s. error: %s\n", info, err)
		os.Exit(1)
	}
}
