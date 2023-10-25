
# d:\code\go\ui\main.go
	cd d:\code\go\ui
	go mod init ui
	
# go get github.com/manifoldco/promptui失败原因:
	go env -w GOPROXY=https://goproxy.cn
	go get github.com/manifoldco/promptui
	
	go run main.go
	go build main.go  、  go build ui

# 调用其他go文件
	mkdir study
	cd study
	go mod init study
	edit main.go
	go get github.com/manifoldco/promptui
	mkdir third
	cd third
	edit select.go and prompt.go under third
		import study/third in main.go
	cd ..
	go run .

# resultChan := make(chan int)
在 Go 中，make 函数用于创建一个新的对象，例如切片、映射和通道等。在这里，make 函数被用于创建一个整数类型的通道（channel），并将其分配给变量 resultChan。
通道是一种用于在协程之间进行通信和同步的机制。在这里，我们创建了一个整数类型的通道，用于在不同的协程之间传递整数值。
具体来说，我们可以使用 <- 运算符向通道发送数据，例如：resultChan <- 42，这将把整数值 42 发送到 resultChan 通道中。我们也可以使用 <- 运算符从通道接收数据，例如：x := <- resultChan，这将等待通道中的下一个整数值，并将其赋值给变量 x。
在实际应用中，通道通常用于在不同的协程之间传递数据或信号，以实现协程之间的同步和通信。例如，在多个协程之间进行任务分配和结果汇总时，我们可以使用通道来传递任务和结果。

# cobra
```
var rootCmd = &cobra.Command{
	Use:   "go_study",
	Short: "go_study is a project for recording go learning process",
	Long: `go_study is a project for recording go learning process, 
				  Complete url is available at https://github.com/heshanu3d/go_study`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("run hugo...")
	},
}
```
cobra.Command 是一个结构体，代表一个命令，其各个属性含义如下：
Use 是命令的名称。
Short 代表当前命令的简短描述。
Long 表示当前命令的完整描述。
Run 属性是一个函数，当执行命令时会调用此函数。
rootCmd.Execute() 是命令的执行入口，其内部会解析 os.Args[1:] 参数列表（默认情况下是这样，也可以通过 Command.SetArgs 方法设置参数），然后遍历命令树，为命令找到合适的匹配项和对应的标志。