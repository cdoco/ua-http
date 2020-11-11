# ua-http

> 通 UA 判断设备型号、浏览器信息等内容

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
