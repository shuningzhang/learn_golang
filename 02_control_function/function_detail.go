package main

import "fmt"

/* 这个函数接收2个参数，返回2个结果，具体在
 * 内部打印接收的参数。这里定义了一个复杂一点
 * 的函数。
 * 1. 2个参数具有不同的类型，分别是string和int32
 * 2. 具有2个返回值，这一点是跟C或者Java不同的地方
 * 3. 返回值的类型也是不同的。*/
func print_info (st_name string, st_age int32) (string, int32) {
      
    fmt.Println(st_name)
    fmt.Println(st_age)

    return st_name, st_age
}

func main() {
    /* 这里是函数的调用 */
    fmt.Println( print_info("sbbdtest", 21) )
}


