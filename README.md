# nodebus

在一处管理多台服务器

1. 安装方式

```shell
git clone https://github.com/runoneall/nodebus.git
cd nodebus
go mod tidy
go build
```

编译完成后，即可使用 `./nodebus` 运行程序。

2. 添加节点

```shell
./nodebus add
```

按照提示输入内容即可

3. 列出节点

```shell
./nodebus list
```

4. 删除节点

```shell
./nodebus --node <node1>,<node2>,... del
```

其中 `<node1>,<node2>,...` 为要删除的节点名称，多个节点名称用 `,` 分隔。

5. 批量执行命令

```shell
./nodebus --node <node1>,<node2>,... run <command>
```

其中 `<node1>,<node2>,...` 为要执行命令的节点名称，多个节点名称用 `,` 分隔。`<command>` 为要执行的命令。
