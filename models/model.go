package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	gossh "golang.org/x/crypto/ssh"
	"golang.org/x/net/websocket"
	"io"
)

type SSH struct {
	User    string
	Pwd     string
	Addr    string
	Client  *gossh.Client
	Session *gossh.Session
}

func (s *SSH) Connect() (*SSH, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = s.User
	config.Auth = []gossh.AuthMethod{gossh.Password(s.Pwd)}
	client, err := gossh.Dial("tcp", s.Addr, config)
	if nil != err {
		return nil, err
	}
	s.Client = client
	session, err := client.NewSession()
	if nil != err {
		return nil, err
	}
	s.Session = session
	return s, nil
}

func (s *SSH) Exec(cmd string) (string, error) {
	var buf bytes.Buffer
	s.Session.Stdout = &buf
	s.Session.Stderr = &buf
	err := s.Session.Run(cmd)
	if err != nil {
		return "", err
	}
	defer s.Session.Close()
	stdout := buf.String()
	fmt.Printf("Stdout:%v\n", stdout)
	return stdout, nil
}

func SSHWebSocketHandler(ws *websocket.Conn) {
	cols := 141
	rows := 83
	sh := &SSH{
		User: "kora",
		Pwd:  "upbjsxt",
		Addr: "127.0.0.1:22",
	}
	sh, err := sh.Connect()
	if nil != err {
		beego.Error(err)
		return
	}

	session := sh.Session
	defer session.Close()
	modes := gossh.TerminalModes{
		gossh.ECHO:          1,
		gossh.TTY_OP_ISPEED: 14400,
		gossh.TTY_OP_OSPEED: 14400,
	}

	if err = session.RequestPty("xterm-256color", rows, cols, modes); err != nil {
		beego.Error(err)
		return
	}

	w, err := session.StdinPipe()
	if nil != err {
		beego.Error(err)
		return
	}
	go func() {
		io.Copy(w, ws)
	}()

	r, err := session.StdoutPipe()
	if nil != err {
		beego.Error(err)
		return
	}
	go func() {
		io.Copy(ws, r)
	}()

	er, err := session.StderrPipe()
	if nil != err {
		beego.Error(err)
		return
	}
	go func() {
		io.Copy(ws, er)
	}()

	if err := session.Shell(); nil != err {
		beego.Error(err)
		return
	}

	if err := session.Wait(); nil != err {
		beego.Error(err)
		return
	}

}
