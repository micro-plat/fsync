### fsync 文件服务
#### client端
1. 提供http,rpc文件上传服务
2. 定时将文件同步到`server`端
3. 支持同步到多个`server`端


#### server端
1. 提供http文件下载服务
2. 提供http,rpc文件上传服务
3. 接收来自`client`端发送的文件，并保存到指定目录
4. 对频繁下载的文件进行缓存，降低文件下载时间


![同步服务](https://github.com/micro-plat/fsync/blob/master/fsync.png)