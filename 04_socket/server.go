package main

import (
    "fmt"
    "net"  //网络通信的库
    "os"
)

func main() {
	/* 构建服务端的地址，端口 */
    service := ":8080"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    /* 建立对端口的监听 */
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        fmt.Println("begin accept")
        /* 等待客户端的连接 */
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        buf := make([]byte,1024)

        /* 读取客户端发送的数据，并打印 */
        result, err := conn.Read(buf)
        fmt.Println(result, string(buf))

        /* 将数据原封返回给客户端 */
        conn.Write(buf) 
        conn.Close()   
    }
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, 
                    "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}


