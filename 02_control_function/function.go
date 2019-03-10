package main

import "fmt"

/* 这里是函数的定义，其中a,  b     是函数的参数
 * int32是参数类型，）和{中间的int32是返
 * 回值的类型。
 * 本函数实现2个32位整数的加法，并返回结果*/
func int_add (a, b int32) int32 {
    var result int32
    
    result = a + b

    return result
}

func main() {
    /* 这里是函数的调用，通过函数的名字进行
     * 调用， 而后面跟着括弧，括弧中是参数。*/
    fmt.Println( int_add(2, 3) )
}

