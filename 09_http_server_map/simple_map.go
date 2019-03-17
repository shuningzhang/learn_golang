package main
  
import ("fmt")

func main() {
        var tt map[string]string
        tt = make( map[string]string )

        tt["China"] = "中国"

        fmt.Println(tt["China"])
}

