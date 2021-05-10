# What is goraml?
It's a raml parser using go base on the go-yaml.Because there are few repository for go to parse the raml, so I want to create one myself.

# Status
This repo is unstable for now(without testing), everything is in progress.
Now, it only supports for a basic usecase:
Get the basic resource information such as the method, query, headers, request body, response also support the included example file.

# 什么是goraml
goraml是基于go-yaml开发的用于解析raml的包。因目前github上没有一个适合于go用于解析raml的包，所以决定自己写一个。

# 包的状态
该包仍处于开发阶段，十分不稳定。目前支持基本的功能，如：获取api resource的基础信息，如请求方法，query参数，header参数，请求体和返回体等，同时还支持包含include的示例。
