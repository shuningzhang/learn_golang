package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, 
        			"Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }

    /* 这里解析从命令行输入的参数，包括地址和端口 */
    service := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    /* 建立与服务端的连接 */
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    /* 向服务端发送数据 */
    _, err = conn.Write([]byte("Hello World!\r\n"))
    checkError(err)

    /* 接收服务端发送的数据，并在终端打印           */
    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
    os.Exit(0)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, 
                    "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}



