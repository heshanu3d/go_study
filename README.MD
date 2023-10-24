
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
	