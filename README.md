# ua-http

> 通 User-Agent 判断设备型号、浏览器信息等内容

```shell
$ curl --compressed "https://api.gaozihang.com/ua-http/?ua=Mozilla%2F5.0%20(iPhone%3B%20CPU%20iPhone%20OS%2014_0_1%20like%20Mac%20OS%20X)%20AppleWebKit%2F605.1.15%20(KHTML,%20like%20Gecko)%20Mobile%2F15E148%20MicroMessenger%2F7.0.17(0x1700112a)%20NetType%2FWIFI%20Language%2Fzh_CN%22"

// return
{
  user-agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.1 (0x1700112a) NetType/WIFI Language/zh_CN"",
  platform: "Mobile",
  os: "IOS",
  device: "iPhone",
  browser: "Wechart"
}
```

配置信息写在 [config](config.toml) 文件中, 目前主要区分三种类型: 系统类型/设备类型/浏览器类型。

```toml
# OS type
os = [
  { name = "IOS", regexp = "\\(iPhone|iPad;( U;)? CPU" },
  { name = "Android", regexp = "Android|Adr" },
  ···
]

# device type
device = [
  { name = "OPPO", regexp = "; OPPO" },
  { name = "RedMi", regexp = "; (redmi|HM NOTE)" },
  ···
]

# Browser type
browser = [
  { name = "QQ", regexp = "QQBrowser" },
  { name = "Wechart", regexp = "MicroMessenger" },
  ···
]
```

每一种类型都会逐行进行匹配, 如果命中就直接返回当前命中信息, 不会再匹配之后的配置文件。

**因为个人能力有限, 有些 `User-Agent` 的匹配信息可能有问题, 欢迎大家提交 PR 修改错误的匹配信息。**
