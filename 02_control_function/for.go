package main

import "fmt"

/* 这个函数用于计算从1到end的累加和，参数
 * end是数据的最后一个值。 */
func calc_sum (end uint32) (uint32) {
    var begin, result uint32

    result = 0

	/* 这里是循环，其中{}分别是循环体的开始和
	 * 结束，与函数和条件判断类似。
	 * for是关键字，表示循环；关键字和循环体中间
	 * 分为3部分：
	 *   begin = 0 在循环开始执行，只执行一次
	 *   begin < end 是条件判断，满足的情况下
	 *     才会执行循环体中的内容。
	 *   begin ++每次执行完循环体中的内容后执行*/
    for begin = 0; begin <= end; begin ++ {
		result += begin  //这里是循环体
    }

    return result
}

/* 这个函数的功能与上面一个一致 */
func calc_sum_ex (end uint32) (uint32) {
    var begin, result uint32

    result = 0

	/* 上一个函数中所说的3部分并不是强制的，可以
	 * 只有一部分，也就是条件判断。这个时候其实
	 * 退化成了C或者Java中的while语句。 */
    for begin <= end {
		result += begin  //这里是循环体
		begin ++
    }

    return result
}

func main() {
    /* 这里是函数的调用 */
    fmt.Println( calc_sum_ex(5) )
}


