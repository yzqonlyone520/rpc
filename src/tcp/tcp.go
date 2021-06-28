package tcpRpc

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type RpcHandle int

func Server() {
	rpcHandle := new(RpcHandle)

	listen, _ := net.Listen("tcp", ":8070")
	go func() {
		for {
			conn, _ := listen.Accept()
			serve := rpc.NewServer()
			serve.Register(rpcHandle)
			go serve.ServeConn(conn)
		}
	}()
}

type Args struct {
	In string
}

type Reply struct {
	Out string
}

func (r RpcHandle) RpcFunc(args *Args, reply *Reply) error {
	reply.Out = "tcp rpc reply : " + args.In
	return nil
}

func Client() {
	var args = Args{In:"tcp test"}
	var reply Reply

	var begin, end time.Time
	begin = time.Now()
	client, _ := rpc.Dial("tcp", "127.0.0.1:8070")
	client.Call("RpcHandle.RpcFunc", &args, &reply)
	end = time.Now()
	fmt.Printf("返回结果[%s] 耗时[%v]\n", reply.Out, end.Sub(begin))
}