## Go type

type是go语法里的重要且常用的关键字。

type 有几种用法：定义结构体， 定义接口， 类型别名， 类型定义， 类型开关

### 定义结构体

```
// type 定义结构体
type persion struct{
    name string   // 属性后面不能加逗号
    age int 
}

// 结构体初始化
p := persion{
    name: "fly",   // 后面要加逗号
    age: 24,
}
```

### 定义接口

```
package main
import "fmt"
// type 定义接口类型
type Person interface{
    Name()
}
// 定义Pearson 接口，接口包含Name() 方法

// 定义一个student结构体, 包含name 和 age 属性
type student struct{
    name string
    age int
}

// 接口实现, 是在结构体上实现的Name()方法，只有student类型的变量才能使用Name()方法
// (类似class的对象，才能调用对象方法)
func (s student)Name(){
    fmt.Println("pearson name is ", s.name)
}

func main(){
    s := student{"fly",30}
    s.Name()
}
```

### 类型别名

```
// 类型等价定义，相当于类型重命名
type name string
// name 类型与string 等价， 就是将一个string 类型起一个别名叫做name

// 区别var声明方式
var name string
// 定义一个string变量
```

类型别名，只能对包内的类型产生作用，对包外的类型采用类型别名，将会报错：

```
cannot define new methods on non-local type string
```

### 类型定义

除了给已知的类型其别名，还可以针对新类型(自定义类型)函数进行定义

```
type handle func(str string)
// 自定义一个函数func, 别名叫做handle,传入一个string参数
```

### 类型开关

在Go语言中存在interface{}类型，可以用来保存任何类型的值，如果我们需要知道具体保存了哪些类型，就需要使用类型开关来判断，具体代码如下：

```
func classifier(items ...interface{}){
    for i,x := range items{
        switch x.(type){
            case bool:
                fmt.Printf("type #%d is bool",i)
            case string:
                fmt.Printf("type #%d is string", i)
            case int:
                fmt.Printf("type #%d is int", i)
        }
    }
}
```