package main

import "fmt"

/* 这里是接口的定义，定义部分包含3部分
 * type关键字， 名称  interface关键字。 之后
 * 是接口定义的主体，里面包含该接口具有的
 * 方法定义 */
type Student interface {
    get_name() string
    get_age() uint8
    
}

/* 这里定义的是结构体的数据类型
 * */
type StudentS struct {
    name string
    age uint32
}

/* 下面是接口的具体实现，对接口的get_name方法
 * 进行了实现，并且绑定到了StudentS类型上。
 * 这种定义方式与C++和Java的非常不同，这里没有
 * this或者self，而是有开发者自己定义。*/
func (student StudentS) get_name () string{
    return student.name
}

func main() {
    var student = StudentS{"sunnyzhang", 18}

    /* 获取成员的值 */
    fmt.Println( student.get_name() )
}


