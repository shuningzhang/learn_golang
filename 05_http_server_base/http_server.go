package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    /* 这里代码与普通TCP服务端并没有差异 */
    service := ":8080"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        fmt.Println("begin accept")
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        buf := make([]byte,1024)
        result, err := conn.Read(buf)
        fmt.Println(result, string(buf))
        /* 只有下面这个发送数据是有差异的，下面按照HTTP
         * 协议发送数据。 */
        conn.Write([]byte("HTTP/1.1 200 OK\r\n" +
                          "Date: Mon, 27 Jul 2009 12:28:53 GMT\r\n" +
                          "Server: Nginx\r\n" +
                          "Content-Length: 12\r\n" +
                          "Content-Type: text/plain\r\n" +
                          "\r\n" +
                          "Hello World!"))
        
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


