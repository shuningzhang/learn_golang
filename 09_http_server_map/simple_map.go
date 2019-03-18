package main
  
import ("fmt")

var content_t  map[string]string

func main() {
		/* map变量定义 */
		        var tt map[string]string
			        tt = make( map[string]string )

				content_t = make( map[string]string)
					/* 变量内容初始化 */
					        tt["Beijing"] = "北京"
							tt["Hebei"] = "河北"
								tt["Tianjin"] = "天津"

									/* 变量内容访问 */
									        fmt.Println(tt["Hebei"])

										        delete(tt, "Hebei")

											        fmt.Println(tt["Hebei"])

			content_t["t"] = "tt"
											}




