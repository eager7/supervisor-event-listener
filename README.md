# supervisor-event-listener
Supervisor事件通知

## 简介
Supervisor是*nix环境下的进程管理工具, 可以把前台进程转换为守护进程, 当进程异常退出时自动重启.  
supervisor-event-listener监听进程异常退出事件, 并发送通知.

### 源码安装
* `go get -u github.com/eager7/supervisor-event-listener`

## Supervisor配置
```ini
[eventlistener:supervisor-event-listener]
command=/path/to/listener
events=PROCESS_STATE
......
```
