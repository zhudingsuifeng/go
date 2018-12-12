### Environment variable fle

/etc/profile 

在用户登录时，操作系统定制用户环境时使用的第一个文件，此文件为系统的每个用户设置环境信息，当用户第一次登录时，该文件被执行。

/etc/environment

在用户登录时，操作系统使用的第二个文件，系统在读取用户个人的prifile前，设置环境文件的环境变量。

～/.profile

在用户登录时，用到的第三个文件.profile文件，每个用户都可使用该文件输入专用于自己使用的shell信息，当用户登录时，该文件仅仅执行一次！默认情况下，会设置一些环境变量，执行用户的.bashrc文件。

/etc/bashrc

为每一个运行bash shell的用户执行此文件，当bash shell被打开时，该文件被读取。

~/.bashrc

该文件包含专用于用户的bash shell 的 bash 信息，当登录时以及每一打开新的shell时，该文件被读取。

### How to set environment variables into ubuntu18

1. 临时设置

```
export PATH=/home/fly/go:$PATH
```

2. 当前用户的全局设置

打开～/.bashrc (如果使用的是bash[, 或者~/.zshrc使用的是zsh]),添加行：

```
export PATH=/home/fly/go:$PATH
```

使生效：

```
source .bashrc
```

3. 所有用户的全局设置

```
$vim /etc/profile
```

里面加入：

```
export PATH=/home/fly/go:$PATH
```

使生效

```
source profile
```

### Test current environment variables

```
echo $PATH
or
env
```

### How to set GOPATH and GOROOT

```
vi ~/.profile

可能会用到的命令：

go env   # 查看当前go的一些环境变量
which go # go程序在当前系统中的位置
```

在文件最后添加：

```
export GOROOT="/usr/lib/go-1.10"
export GOBIN=$GOROOT/bin
export GOPATH="/home/fly/github/go"
export PATH=$PATH:$GOPATH/bin
```

使文件立即生效：

```
$source ~/.profile
```

### Profile loading order

用户登录后加载profile和bashr的顺序：

```
1. /etc/profile
    ->/etc/profile.d/*.sh

2. $HOME/.profile
    ->$HOME/.bashrc
        ->/etc/bashrc
```

说明：

bash首先执行/etc/profile脚本,/etc/profile脚本先依次执行/etc/profile.d/*.sh

随后bash会执行用户主目录下的.profile脚本，.profile脚本会执行用户主目录下的.bashrc脚本

而.bashrc脚本会执行/etc/bashrc脚本

至此，所有的环境变量和初始化设定都已经加载完成。