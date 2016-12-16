package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, param string) string {
	return "ECHO:" + param
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, err := client1.Ca/ll("post", "From client1")
	if err != nil {
		fmt.Println(err)
	}
	resp2, err := client2.Call("post", "From client2")
	if err != nil {
		fmt.Println(err)
	}
	if resp1 != "ECHO:From client1" || resp2 != "ECHO:From client2" {
		t.Error("IpcClient.Call failed resp1：", resp1, "resp2", resp2)
	}

	client1.Close()
	client2.Close()
}
