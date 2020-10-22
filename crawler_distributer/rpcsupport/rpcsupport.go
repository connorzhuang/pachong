package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//将对象target注册为一个rpc服务
func ServerRpc(host string, target interface{}) error {
	rpc.Register(target)
	listener, err := net.Listen("tcp", host) //在host端口监听
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept() //建立一个连接
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn) //启动服务

	}
	return nil
}

//客户端请求服务
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil

}
