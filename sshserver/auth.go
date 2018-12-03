package sshserver

import (
	"log"
)

func PasswordCallback(ctx Context, password string) bool{
	log.Printf("get user: %v password is: %v",ctx.User(),string(password))
	return true
	//return &sshserver.Permissions{
	//	CriticalOptions: map[string]string{
	//		"name": conn.User(),
	//		"localports":  conn.LocalAddr().String(),
	//		"remoteports": conn.RemoteAddr().String(),
	//	},
	//}, nil
	//return nil,fmt.Errorf("Password auth failed\n")
}

func PublicKeyCallback(ctx Context, key PublicKey) bool {
	authmutex.Lock()
	defer authmutex.Unlock()
	return false
	//if clientinfo, found := authorisedKeys[string(key.Marshal())]; found {
	//	return &sshserver.Permissions{
	//		CriticalOptions: map[string]string{"name": clientinfo.Comment,
	//			"localports":  clientinfo.LocalPorts,
	//			"remoteports": clientinfo.RemotePorts},
	//	}, nil
	//}
	//
	//return nil, fmt.Errorf("Unknown public key\n")
}
