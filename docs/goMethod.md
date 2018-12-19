## go Method

一个方法只是一个函数，它有一个特殊的接收者(receiver)类型，该接收者放在func 关键字和函数名之间。接收者可以是结构体类型或者非结构体类型。可以在方法内部访问接收者。

通过下面的语法创建一个方法：

```
func (t Type) methodName(parameter list){

}
```

上面的代码片段创建一个名为methodName的方法，该方法有一个类型为Type的接收者。

编写一个简单的程序，创建一个结构体类型的方法并调用它。

```
package main

import (
    "fmt"
)

type Employee struct{
    name string
    salary int
    currency string
}

/*
displaySalary() method has Employee as the receiver type
*/

func (e Employee) displaySalary(){
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main(){
    emp1 := Employee{
        name: "bob",
        salary:5000,
        currency:"$"
    }
    emp1.displaySalary() // Calling displaySalary() method of Employee type
}
```

我们为Employee创建一个名为displaySalary的方法。在displaySalary()方法内部可以访问它的接收者e (类型为Employee).

### 为什么使用方法而不是函数

上面的程序可以使用函数而不是方法重写如下：

```
package main

import (
    "fmt"
)

type Employee struct{
    name string
    salary int
    currency string
}

/*
displaySalary() method converted to function with Employee as parameter
*/
func displaySalary(e Employee){
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main(){
    emp1 := Employee{
        name: "bob",
        salary:5000,
        currency:"$"
    }
    displaySalary(emp1)
}
```

使用displaySalary函数替换方法，并将Employee 结构体作为参数传给它。输出结果一样。

为什么我们可以使用函数完成的工作，还要使用方法？

1. Go 不是一个春面向对象的编程语言，它不支持class 类型。因此通过在一个类型上建立方法来实现与class 相似的行为。

2. 同名方法可以定义在不同的类型上，但是Go 不允许同名函数。

```
package main

import (
    "fmt"
    "math"
)

type Rectangle struct{
    length int
    width int
}

type Circle struct{
    radius float64
}

func (r Rectangle) Area() int {
    return r.length * r.width
}

func (c Circle) Area() float64{
    return math.Pi * c.radius * c.radius
}

func main(){
    r := Rectangle{
        length:10,
        width:5,
    }
    fmt.Printf("Area of rectangle %d\n", r.Area())

    c := Circle{
        radius:12,
    }
    fmt.Printf("Area of circle %f", c.Area())
}
```
