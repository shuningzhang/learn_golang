package main

import (
    "fmt"
    "net"
    "os"
    "io/ioutil"
    "strconv"
    "strings"
)

func read_html_file(name string) ([]byte, error) {
	var ret []byte
	file_obj, err := os.Open(name)
	if err == nil {
		contents, err := ioutil.ReadAll(file_obj)
		if err == nil {
			return contents, err
		}
	}

	return ret, err
}


func http_resp_header(file_len int, content_type string) ([]byte) {
	header :=         "HTTP/1.1 200 OK\r\n" +
                          "Date: Mon, 27 Jul 2009 12:28:53 GMT\r\n" +
                          "Server: Nginx\r\n" +
                          "Content-Length: " +
			  strconv.Itoa(file_len) +
			  "\r\n" +
                          "Content-Type: " +
			  content_type +
			  "\r\n" +
                          "\r\n" 

	return []byte(header)
}

func parse_request(req_content string) (method, file_name, req_type string){
	position := strings.Index(req_content, "\r\n")
	req_header := req_content[0 : position]
	header_item := strings.Split(req_header, " ")
	req_t := "text/html"

	if header_item[1] != "/" {
		position = strings.Index(header_item[1], ".")
		str_len := len(header_item[1])
		req_t = header_item[1][position+1 : str_len]
	}

	return header_item[0], header_item[1], req_t
}

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
	m, path, t := parse_request(string(buf))
	fmt.Println(m, path, t)
	if path == "/" {
		path = "index.html"
	} else {
		path = "." + path
	}

	body, err := read_html_file(path)
	
	file_len := len(body)

	header := http_resp_header(file_len, t)

	conn.Write([]byte(header))
	
	conn.Write(body)
        
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


