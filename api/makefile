# 命令
help:
	@echo 'Usage:'
	@echo '     db 生成sql执行代码'
	@echo '     api 根据api文件生成go-zero api代码'
	@echo '     dev 运行'
db:
	gentool -dsn 'root:123456@tcp(192.168.1.13:3306)/bk?charset=utf8mb4&parseTime=True&loc=Local' -outPath './bkmodel/dao/query'
.PHONY:api
api:
	goctl api go -api project.api -dir ./ -style gozero
dev:
	go run backend.go -f etc/backend.yaml