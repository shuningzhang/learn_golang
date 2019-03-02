package main 

import "fmt"

func main() { 

    /* 下面一行定义一个int8类型的变量，也称为声明
     * 变量。 可以看出声明方式由3部分构成，关键字
     * var，变量名var_name和变量类型int8。Go语言的
     * 变量声明方式与C或者Java等语言的方式不太一
     * 样，刚开始学可能觉得有些怪，但并非Go语言独创。*/
    var var_name int8

    /* 变量的赋值，也就是让变量的内容变为某个值。 */
    var_name = 5

    /* 定义一个16位的整型数，第二行进行2个变量运
     * 算。 由于变量的类型不一样，因此需要进行强制
     * 类型转换。 */
    var var_name_x int16
    var_name_x = 11 + int16(var_name)
	
    /* 上面介绍的都是数字型的数据类型，除了上面的
     * 类型外，还有int32，int64，uint8，uint16，
     * uint32，uint64等类型，这些表示的是整型数。
     * 而以uint头的表示无符号整型数。 */

    fmt.Println(var_name_x)
	
    /* 除了整型数外，还有浮点型数，浮点型数就是我们
     * 日常说的小数 */
    var float_name float32
    float_name = 1.2
	
    fmt.Println(float_name)
    /* 浮点型数包括32位和64位两种，也就是float32和float64。 */
	
    /* 下面介绍的类型是字符串类型。 */
    var str_name string
    str_name = "test"
    fmt.Println(str_name)
	
    /* 变量可以在声明的时候直接初始化。 */
    var int_name int16 = 12
    fmt.Println(int_name)
	
    /* 可以同时声明多个变量，如下是定义3个都为
     * int32类型的变量，并进行初始化，初始化并非
     * 强制要求的 */
    var int_name_1, int_name_2, int_name_3 int32 = 1, 2, 3 
    fmt.Println(int_name_1, int_name_2, int_name_3)
	
    /* 可以不声明变量类型，这样编译器会自动选择类型 */
    var int_name_4 = 160
	
    /* 下面是最简单的变量声明方式，可以将var关键字省略 */
    int_name_5 := 170
    fmt.Println(int_name_4, int_name_5)
	
    /* 最后说明一下，变量名可以字母、下划线和数字，但只能以非数字开头。*/
	
}
