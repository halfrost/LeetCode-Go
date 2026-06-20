#!/usr/bin/env bash
#
# 生成单一、合法的 Go 覆盖率文件 coverage.txt。
#
# 旧写法是对每个包分别 -coverprofile 再 cat 追加，产出的文件顶部有一个空行、
# 中间还夹着多行重复的 "mode: atomic" 头。旧的 Travis bash 上传器能容忍，
# 但新版 Codecov 上传器解析更严格，会把这种文件算成 0% 覆盖率。
#
# Go 1.10+ 支持对多个包一次性 -coverprofile，直接产出单个合法 profile。
set -e

go test -covermode=atomic -coverprofile=coverage.txt ./leetcode/...
