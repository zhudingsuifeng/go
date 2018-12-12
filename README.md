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

$go env     #查看go相关的环境变量
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/fly/.cache/go-build"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/fly/go"
GORACE=""
GOROOT="/usr/lib/go-1.10"
GOTMPDIR=""
GOTOOLDIR="/usr/lib/go-1.10/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build162333287=/tmp/go-build -gno-record-gcc-switches"
```

可以看到，GOPATH="/home/fly/go"。在没有弄清楚一些东西之前，建议使用默认的工作目录。下面修改GOPATH的方法虽然正确，但是还会存在其他问题。留待后面解决。

在linux下，go工作空间在$HOME/go目录，可以通过[设置环境变量](https://github.com/zhudingsuifeng/go/blob/master/docs/environment.md)GOPATH来指定Go的工作空间为其他目录:

```
$vim ~/.profile

在文件最后添加：

export GOPATH=/home/fly/github/go
```

使生效：

```
source ~/.profile
```

如果使用默认的工作空间，却在其他目录下创建项目，在vscode中debug的时候会报错：

```
Failed to continue :"Cannot find Delve debugger.Install from https://github.com/derekparker/delve&ensure it is in your "GOPATH/bin" or "PATH"."
```

在vscode中debug go 程序，报错：

```
cd $GOPATH
go get github.com/derekparker/delve/cmd/dlv
```

如果没有报错，就可以在vscode中直接debug go 程序了。

在文件上右键open in terminal，就可以在vscode中执行shell命令了。