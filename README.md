# 泰拉瑞亚自动端口代理 terraria-auto-port

## 功能

当玩家通过`terraria-auto-port`加入时，代理端会自动区分玩家的客户端类型(Vanilla/tModLoader)，然后将玩家重定向到对应的服务端，从而实现原版服务器和tModLoader服务器共用一个端口。

## 那有啥用啊？

如果你有一个原版服务器和一个tModLoader服务器，你就可以用它来合并端口，让两个服务器都能使用7777。

> [!NOTE]
> 其实事实上并没有啥用其实，只是Cai用来练习Go的很随便的一个项目。

## 配置

```jsonc
{
  "listen_port": 7777, // 代理端侦听端口
  "vanilla_address": "127.0.0.1:8888", // 原版服务器地址
  "tModLoader_address": "127.0.0.1:9999" // tModLoader服务器地址
}
```