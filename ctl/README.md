# LeetCode-Go ctl

## 配置方法

1. 在`.gitignore`中，添加一行`*.toml`
2. 在`LeetCode-Go`目录下，添加文本文件`config.toml`。
3. 把以下内容复制到`config.toml`中。
4. 把`config.toml`中的`test`分别修改为你的 leetcode `用户名`和`密码`。
5. 去 leetcode 登录后，把网站 cookie 复制 (需要复制 csrftoken 和 LEETCODE_SESSION) 并替换 `config.toml`中的`Cookie`。

```toml
Username="test"
Password="test"
Cookie="csrftoken=XXXXXXXXX; LEETCODE_SESSION=YYYYYYYY;"
CSRFtoken="ZZZZZZZZ"
```

## PDF 生成

用 `leetcode-go pdf` 命令先生成书籍内容的合并版本 pdf.md，再用 vscode 或者其他支持 toc 目录生成的工具，生成 toc。最后用 Typora 把 md 文件转换成 pdf。就可以发布 release 新版本了。