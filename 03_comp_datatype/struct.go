package main

import "fmt"

/* 这里是结构体的定义，定义部分包含3部分
 * type关键字， 名称  struct关键字。 之后
 * 是结构体定义的主体，里面具体的成员定义。
 * 成员定义有名称和类型构成。*/
type Student struct {
    name string
    age uint8
    score uint8
}

func main() {
    /* 这里是结构体类型的定义，与普通变量无疑 */
    var zhang Student

    /* 还可以在定义的时候进行初始化 */
    var wang = Student{"ericwang", 17, 99}
    

	/* 成员访问， 通过        变量名.成员       进行访问。 */
    zhang.name = "sunnyzhang"
    zhang.age = 18
    zhang.score = 98

    /* 获取成员的值 */
    fmt.Println( zhang.age, wang.age )
}


