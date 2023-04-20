# GO

## Go基础类

### 1. 与其他语言相比，GO语言有什么好处？
+ golang针对并发做了优化，
+ 单一的标准代码格式，更具可读性
+ 自动垃圾回收机制更加有效

### 2. golang的数据类型
+ Method
+ Bool
+ String
+ Array
+ Slice
+ Struct
+ Pointer
+ Function
+ Interface
+ Map
+ Channel
+ Numeric

### 3.Go程序中的包
`pkg`是go工作区中包含Go源文件或其他包的目录

包声明：`package <packagename>`

使用包：`packagename`

### 4.Go支持显性形式的类型转化
```go
package main
func main() {
	i := 55
	j := 67.5
	
	sum := i + int(j)
	print(sum)
}
```

### 5.什么是Goroutine？如何停止
Goroutine 比标准线程更轻量级，要创建Goroutine，使用关键字`go`
`go f(x,y,z)`
通过一个信号通道停止它，goroutine只能在被告知检查时响应信号，需要在逻辑位置（例如：`for`循环顶检查)
```go
package main
func main() {
	quit := make(chan bool)
	go func() {
		for  {
			select {
			case <-quit:
				return
			default:
				// do something here
            }
		}
    }()
    // do 
	quit <- true
}
```

### 6.如何在运行时检查变量类型
类型开头(Type Switch)在运行时检查变量类型的最佳方案。

### 7.go两个接口之间可以存在什么关系
如果两个接口有相同的方法列表，那么他们是等价的，可以相互赋值

### 8.go当中同步锁有什么特点，作用是什么？
+ 当一个goruntine（协程）获得`Mutex`，其他goroutine需要等待其释放`Mutex`。`RWMutex`在
读锁占用的情况下，会阻止写，但不会阻止读（RWMutex）。在写锁占用的情况下，会阻止任何其他
goroutine（无论读或者写），整个锁相当于由该goroutine独占

+ 同步锁的作用是保证资源在使用时的独有性，不会因为并发而导致数据错乱，保证系统的稳定性

### 9.go语言当中Channel（通道）有什么特点？
+ 给一个nil的channel中发送或者接收数据，都会造成永久阻塞
+ 给一个已经关闭的channel发送数据，会引起panic
+ 从一个已经关闭的channel接收数据，如果缓冲区为空，则返回一个零值

### 10.go语言当中的channel缓冲有什么特点
无缓冲的channel是同步的，而有缓冲的channel是非同步的

