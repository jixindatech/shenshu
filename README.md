# 神荼
中文 | [English](./README_EN.md)
web防火墙管理端

##安装
1. 前端用vue，后端用golang，正常安装即可

## 配置文件说明
1. auth.json 是基于rbac的权限，写在了配置文件中，默认只有admin角色，可以自定义配置
其它角色和对应的权限
2. basic_model.conf 和basic_policy.csv 是基于casbin的后端接口权限校验，
其它角色可以自己独立配置。
3. 关于config.yaml，database是mysql存放所有的配置数据， redis是缓存，神荼网关会从redis里面定时拉取
配置数据，elasticsearch是存放神荼网关的事件，跟神荼网关相关的redis和elasticsearch配置要跟网关一致。
4. adminpassword 是管理员密码，如果存在这个字段，启动时就会修改管理员密码为这个配置，
启动之后，可以删除这个字段即可。

##Contributing
欢迎issue和star

## Discussion Group
QQ群: 254210748

## License
Unlicense



