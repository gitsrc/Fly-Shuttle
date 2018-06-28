package TcpTunPipe

import (
	"context"
)

type TunData struct {
	Host string
	Port int
	Data []byte
}

type ReplyData struct {
	Data []byte
}

type TcpTunPipe struct {
}

func (*TcpTunPipe) SendTcp(ctx context.Context, args *TunData, reply *ReplyData) error {
	//fmt.Println(string(args.Data))
	reply.Data = []byte("Hello")
	return nil
}

func (*TcpTunPipe) hello(ctx context.Context, args *string, reply *string) error {
	reply = args
	return nil
}
