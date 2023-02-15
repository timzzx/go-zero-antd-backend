# GoDockerDev
golang docker-compose dockerfile

项目目录
```
.
├── LICENSE
├── README.md
├── docker-compose.yml
├── mysql
│   └── Dockerfile
└── redis
    └── Dockerfile
```

启动
``` 
docker-compose up -d
```

测试mysql redis 连接成功

停止
```
docker-compose stop
```

删除
```
docker-compose rm
```
只要data目录不删除数据就会保留