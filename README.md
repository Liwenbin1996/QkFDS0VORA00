# QkFDS0VORA00

1.函数注释应以“函数名”为开头 如：  

```go
// FuncA ...
func FuncA(a int) {
    ...
}
```

2.每个目录 需要有独立的README.md
README.md 应解释本目录的代码和组织结构
Unix系统可使用tree指令获取目录的树形结构
形如：

```
.
├── README.md
├── api
│   ├── README.md
│   ├── api.go
│   └── eg.go
├── cache
│   └── README.md
├── config
├── go.mod
├── go.sum
├── job
│   └── README.md
├── main.go
├── middleware
│   ├── README.md
│   ├── jwt.go
│   └── middleware.go
├── model
│   └── README.md
├── route
│   ├── README.md
│   └── routes.go
└── util
    └── README.md
```

3.关于Go的一些代码规范：

https://github.com/golang/go/wiki/CodeReviewComments

https://colobu.com/2017/02/07/write-idiomatic-golang-codes/

https://zhuanlan.zhihu.com/p/69445822

