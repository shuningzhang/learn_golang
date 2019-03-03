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

/* 对请求体进行解析，解析出请求的类型及请求路径等内容*/
func parse_request(req_content string) (method, file_name, req_type string){
	position := strings.Index(req_content, "\r\n")  //取请求的第一行
	req_header := req_content[0 : position]
	header_item := strings.Split(req_header, " ")
	req_content_type := "text/html"

	if header_item[1] != "/" {
		dot_position := strings.Index(header_item[1], ".")
		qm_position := strings.Index(header_item[1], "?")
		/* 默认情况下通过字符串结尾截取扩展名 */
		str_len := len(header_item[1]) 
		fmt.Println(dot_position, qm_position, str_len, "\r\n")
		/* 判断是否有参数，如果有则采用?位置作为结尾
		 * 也就是截取.和?中间的部分作为请求内容的类型*/
		if qm_position >= 0 && qm_position >= dot_position {
			str_len = qm_position 
		}
		req_content_type = header_item[1][dot_position+1 : str_len]
	}

	return header_item[0], header_item[1], req_content_type
}

/* 根据请求路径，解析出其中的参数 */
func parse_path_params(path string) (string){
 	var content string
 	var file_path string
 	var params string

	qm_position := strings.Index(path, "?")
	if qm_position >= 0 {
		file_path = path[1 : qm_position-1]
		str_len := len(path)
		params = path[qm_position+1 : str_len]
	}

	fmt.Println(file_path, params, "\r\n")

	content = "<html><body><p style=\"color:#FF0000\";>" +
		  params +
		  "</p></body></html>"

	
 	if ioutil.WriteFile(file_path,[]byte(content),0644) == nil {	
		fmt.Println("ff")
	}
	return "/" + file_path
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

	if path != "/" {
		path = parse_path_params(path)
	}
	
	if path == "/" {
		path = "index.html"
	} else {
		path = "." + path
	}

	
	fmt.Println(path)
	body, err := read_html_file(path)

	fmt.Println(body)
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



