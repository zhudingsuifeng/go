## 1 - GO

下一节： [第二篇 函数](https://github.com/zhudingsuifeng/go/blob/master/docs/goFunction.md)

### hello world

package main:每个Go文件都必须以package name 语句开头。包(package)提供了代码封装和重用。这里包的名字为main。

func main():main函数是一个特殊的函数，它是Go程序的入口点。main函数必须包含在main package中。

### 变量

```
var name type           # 声明单一变量，name 变量名，type 变量类型，默认初始化为0
var name type = value   # 声明带初始值的变量，value变量值
var name = value        # 带初始值的情况下，可以省略类型
var name1, name2 type = value1, value2    # 一次声明多个变量
var name1, name2 = value1, value2         # 省略类型
```

类型推导： 如果声明一个变量时提供了初始值，Go可以根据初始值来自动推导变量的类型。

```
var (
    name1 = initialvalue1,
    name2 = initialvalue2
)
# 一次声明多个，不同类型的变量
```

简短声明：

Go也支持一种声明变量的简洁形式，称为简短声明，使用:=操作符。

声明语法 name := initialvalue.

```
package main

import main

func main(){
    name, age := "fly", 24
    fmt.Println("my name is ", name , "and age is ", age)
}
```

简短声明要求:

1. := 操作符左边的所有变量都有初始值。

2. := 操作符左边至少有一个变量时尚未声明的。

3. 简短声明只能用在函数中。

变量声明后，不能被赋予与其类型不同的值。

### 类型

布尔类型： bool

有符号整型： int8, int16, int32, int64

无符号整型： uint8, uint16, uint32, uint64

浮点型： float32, float64

复数类型： complex64, complex128

byte是uint8 的别名， rune 是 int32 的别名。

string类型： 字符串，字节的集合。

### 类型转换

Go 有着非常严格的强类型特征。Go没有自动类型提升或类型转换。类型转换时，需要显式类型转换。

### 常量

常量(constant)表示固定的值。

关键字 const 修饰的名字为常量，不能被重新赋予任何值。

```
package main

func main(){
    const a = 5
}
```

常量的值必须在编译期间确定。因此不能将函数的返回值复制给常量，因为函数调用发生在运行期。

```
package main

import (
    "fmt"
    "math"
)

func main(){
    fmt.Println("Hello, playground")
    var a = math.Sqrt(4)   // allowed
    const b = math.Sqrt(4) // not allowed
}
```

Go 是一种强类型语言。 所有变量都需要显式类型。但是常量可以无类型。 

无类型常量有一个默认的类型，当且仅当代码中需要无类型常量提供类型时，他才会提供该默认类型。

### 字符串常量

```
const helle = "hello world"
```

字符串没有任何类型。

```
type mystring string  //创建一个新的类型mystring,他是string 的别名。
```

（可以通过 type newtype type 的语法来创建一个新的类型。类似C语言的typedef。）

### 布尔常量

```
const trueConst = true
```

### 数值常量

```
const a = 5
var intv int = a
var int32v int32 = a
var float64v float64 =a
```

上面程序中，const a 是无类型的，值为 5. a可以赋值给任何与之类型兼容的变量。

数值常量可以在表达式中自由的混合和匹配，尽当将他们赋值给变量或者代码中明确需要类型的时候，才需要他们的类型。

```
package main

import (
    "fmt"
)

func main(){
    var a = 5.9/8
    fmt.Pringf("a's type %T value %v", a, a)
}
```

上面的程序中，语法上5.9是一个浮点数，8是一个整数。因为他们都是数值常量，因此5.9/8是合法的。相处的结果为0.7375， 是一个浮点数。因此变量a 的类型为浮点数。

下一节： [第二篇 函数](https://github.com/zhudingsuifeng/go/blob/master/docs/goFunction.md)