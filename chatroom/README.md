# ChartRoom

- 主要功能
    聊天室在线人数、在线列表
    昵称、头像设置
    发送/接收消息
    进入、退出聊天室
    不同消息类型的支持：文字、表情、图片、语音、视频等
    敏感词过滤
    黑名单/禁言
    聊天室信息
    创建、修改聊天室
    聊天室信息存储和历史消息查询
    ...

## 基于tcp的聊天室

- 如果大量用户广播延迟问题
- 消息的堵塞

## 目录

- tcp 实现了tcp版本的wesocket
- webscocket 使用不同的库，性能上对比，似乎gobwas 是比较优于其他
- chatroom 是完整版本的main
- logic 存放核心业务逻辑和service 目录类似
- server 存放server相关代码，类似于存放controller 的代码
- template 存放静态代码文件

```md
.
├── README.md
├── cmd
│   ├── chatroom
│   ├── tcp
│   │   ├── client
│   │   │   ├── client.go
│   │   │   └── client_test.go
│   │   └── server
│   │       ├── server.go
│   │       └── server_test.go
│   └── websocket
│       ├── gobwas
│       │   ├── client
│       │   │   ├── clinet.go
│       │   │   └── clinet_test.go
│       │   └── server
│       │       └── server.go
│       └── nhooyr
│           ├── client
│           │   └── client.go
│           └── server
│               └── server.go
├── logic
├── server
└── template
```