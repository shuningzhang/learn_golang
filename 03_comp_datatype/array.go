package main

import "fmt"

/* 这个函数用于计算几个学生成绩的均值 */
func calc_average () (float32) {
    var student_1 uint32 = 67
    var student_2 uint32 = 89
    var student_3 uint32 = 96
    var student_4 uint32 = 87
    var student_5 uint32 = 79
    var sum uint32 = 0
    var result float32 = 0;

    sum = student_1 + student_2 +
          student_3 + student_4 +
          student_5

    result = float32(sum) / 5
    
    return result
}

/* 上面函数计算均值数据数量少，可能并看不出有
 * 多麻烦，如果有100个学生，上面一个函数将会非常
 * 复杂。
 * 下面这个函数利用了数组，功能是一样的。*/
func calc_average_ex() (float32) {
    /* 这里定义一个数组，并进行初始化，初始化不
     * 是必须的。
     *  数组的名称      数量 类型       数组初始化的数据     。*/
    var students = [5] uint32 {67, 89, 96, 87, 79}
    var sum uint32 = 0
    var result float32 = 0
    var i uint32 = 0

    /* 通过数组，结合for循环可以非常容易的计算
     * 累加和。
     * 数组索引是从0开始的，0表示第一个元素。 */
    for i = 0; i<5; i++ {
    
        /* 数组的访问通过[]进行，其中i是索引。
         * 读或者写其中的值都通过[]访问。 */
	    sum += students[i] 
    }

    result = float32(sum) / 5

    return result
}

func main() {
    /* 这里是函数的调用 */
    fmt.Println( calc_average_ex() )
}


