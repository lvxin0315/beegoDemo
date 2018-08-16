## 跨平台编译

go build main.go 主入口文件编译

go run main.go 编译并启动

64位linux    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

64位windows  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

64位mac  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build

dep 包管理工具

bee beego框架出的创建、热编译、开发等功能工具

bee run -gendoc=true -downdoc=true

-gendoc=true 生成api文档

-downdoc=true 加载swagger静态资源内容