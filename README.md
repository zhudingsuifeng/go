## How to work with golang
### Install golang
```javascript
$sudo su
#apt update
#apt upgrade
#apt install golang
```
### Hello world
```javascript
$vi hello.go

package main
import "fmt"
func main(){
    fmt.Println("hello, world!")
}

$go run hello.go    # run the hello.go golang file 
$go build hello.go  # generate the executable file from hello.go
```
