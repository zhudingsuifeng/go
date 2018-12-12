## GO

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
