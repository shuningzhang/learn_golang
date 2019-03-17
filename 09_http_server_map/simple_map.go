package main
  
import ("fmt")

func main() {
        var tt map[string]string
        tt = make( map[string]string )

        tt["Beijing"] = "北京"
	tt["Hebei"] = "河北"
	tt["Tianjin"] = "天津"

        fmt.Println(tt["Hebei"])
}

