## 说明

英华学堂网课助手客户端版本  
作者: Pluto  
禁止二次售卖, 仅供学习与交流  
代刷慕课的请不要使用, 我讨厌代刷慕课的    

## 演示

![](docs/preview_1.gif)

### web端

![](docs/preview_2.png)

## 使用

* 在`releases`页面下载最新版, 到电脑  
  + `windows`系统请下载文件名带`windows`的, 别傻乎乎的下载`linux`的
  + 直达: <https://github.com/aoaostar/mooc/releases/latest>   

* 配置好`config.json`, 直接双击运行`运行.bat`即可，窗口不能关闭, 只能最小化
  + 关闭了还问我怎么没有了, 程序都没有运行了, 怎么可能还有  
* 也可以双击运行`后台运行.bat`让程序在后台运行 ( 没有窗口 )  
* 打开 <http://127.0.0.1:10086> 可以在浏览器查看服务状态  
* 如果想要结束后台程序请运行`结束.bat`结束后台程序  

### linux系统
自己琢磨, 不教

### 配置

> 使用`mooc.yinghuaonline.com`时`school_id`为必填项  
> 使用自己学校的平台可以不填，默认为`0` (推荐)   
> `server`网页端地址, `:10086`=> `127.0.0.1:10086` ( 不懂就不要改 )  
> `limit`协程数, 支持多门课程一起刷, 拉满 ( 填数字就行了, 99就行了 ) 可以以最快速度刷完 (推荐拉满)  
> JSON编辑工具: <https://tool.aoaostar.com/json>

```json
{
  "global": {
    "server": ":10086",
    "limit": 3
  },
  "users": [
    {
      "base_url": "https://mooc.yinghuaonline.com/",
      "school_id": 0,
      "username": "username",
      "password": "password"
    }
  ]
}
```

#### 支持多账号

```json
{
  "global": {
    "server": ":10086",
    "limit": 3
  },
  "users": [
    {
      "base_url": "https://mooc.yinghuaonline.com/",
      "school_id": 0,
      "username": "username1",
      "password": "password1"
    },
    {
      "base_url": "https://mooc.yinghuaonline.com/",
      "school_id": 0,
      "username": "username2",
      "password": "password2"
    }
  ]
}
```