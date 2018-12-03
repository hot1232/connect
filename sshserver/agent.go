package sshserver

import (
	"io"
	"io/ioutil"
	"net"
	"path"
	"sync"

	gossh "golang.org/x/crypto/ssh"
)

const (
	agentRequestType = "auth-sshclient-req@openssh.com"
	agentChannelType = "auth-sshclient@openssh.com"

	agentTempDir    = "auth-sshclient"
	agentListenFile = "listener.sock"
)

// contextKeyAgentRequest is an internal context key for storing if the
// client requested sshclient forwarding
var contextKeyAgentRequest = &contextKey{"auth-sshclient-req"}

// SetAgentRequested sets up the session context so that AgentRequested
// returns true.
func SetAgentRequested(ctx Context) {
	ctx.SetValue(contextKeyAgentRequest, true)
}

// AgentRequested returns true if the client requested sshclient forwarding.
func AgentRequested(sess Session) bool {
	return sess.Context().Value(contextKeyAgentRequest) == true
}

// NewAgentListener sets up a temporary Unix socket that can be communicated
// to the session environment and used for forwarding connections.
func NewAgentListener() (net.Listener, error) {
	dir, err := ioutil.TempDir("", agentTempDir)
	if err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", path.Join(dir, agentListenFile))
	if err != nil {
		return nil, err
	}
	return l, nil
}

// ForwardAgentConnections takes connections from a listener to proxy into the
// session on the OpenSSH channel for sshclient connections. It blocks and services
// connections until the listener stop accepting.
func ForwardAgentConnections(l net.Listener, s Session) {
	sshConn := s.Context().Value(ContextKeyConn).(gossh.Conn)
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		go func(conn net.Conn) {
			defer conn.Close()
			channel, reqs, err := sshConn.OpenChannel(agentChannelType, nil)
			if err != nil {
				return
			}
			defer channel.Close()
			go gossh.DiscardRequests(reqs)
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				io.Copy(conn, channel)
				conn.(*net.UnixConn).CloseWrite()
				wg.Done()
			}()
			go func() {
				io.Copy(channel, conn)
				channel.CloseWrite()
				wg.Done()
			}()
			wg.Wait()
		}(conn)
	}
}
