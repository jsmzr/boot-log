# botostrap-log

[![Build Status](https://github.com/jsmzr/boot-log/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-log/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-log/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-log)

提供统一的 log 门面，旨在简化日志的接入、迁移工作，通过引入不同的适配器（实现）来实现日志功能。

## 使用说明

基于门面完成适配，在项目中显示将适配器初始化后，通过门面方法即可调用即可。

