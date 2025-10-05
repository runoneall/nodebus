# nodebus

在一处管理多台服务器

## 安装方式

```shell
git clone https://github.com/runoneall/nodebus.git
cd nodebus
go mod tidy
go build
```

编译完成后，即可使用 `./nodebus` 运行程序。

## 基础用法

### 添加节点

```shell
./nodebus add
```

按照提示输入内容即可

### 列出节点

```shell
./nodebus list
```

### 删除节点

```shell
./nodebus --node <node1>,<node2>,... del
```

其中 `<node1>,<node2>,...` 为要删除的节点名称，多个节点名称用 `,` 分隔。

### 批量执行命令

```shell
./nodebus --node <node1>,<node2>,... run <command>
```

其中 `<node1>,<node2>,...` 为要执行命令的节点名称，多个节点名称用 `,` 分隔。 `<command>` 为要执行的命令。

### docker 管理

```shell
./nodebus --node <node1>,<node2>,... docker <command>
```

其中 `<node1>,<node2>,...` 为要执行命令的节点名称，多个节点名称用 `,` 分隔。 `<command>` 为要执行的命令，支持全部 docker 命令

## 高级用法

[nodebus v3版本发布！ | Runoneall](https://oneall.eu.org/2025/09/20/nodebus-v3%E7%89%88%E6%9C%AC%E5%8F%91%E5%B8%83%EF%BC%81/)

[nodebus新功能: cfgcenter | Runoneall](https://oneall.eu.org/2025/10/01/nodebus%E6%96%B0%E5%8A%9F%E8%83%BD-cfgcenter/)

[nodebus cfgcenter从http迁移到ipc通信 | Runoneall](https://oneall.eu.org/2025/10/01/nodebus-cfgcenter%E4%BB%8Ehttp%E8%BF%81%E7%A7%BB%E5%88%B0ipc%E9%80%9A%E4%BF%A1/)

[nodebus新功能：x11转发 | Runoneall](https://oneall.eu.org/2025/10/05/nodebus%E6%96%B0%E5%8A%9F%E8%83%BD%EF%BC%9Ax11%E8%BD%AC%E5%8F%91/)
