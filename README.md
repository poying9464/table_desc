# TABLE_DESC

#### 介绍
在工作中，我们经常需要将指定环境的数据库的表、字段以及备注信息导出为word文档，方便合作方查看。

table_desc是一个基于golang开发的数据库表信息导出工具。table_desc提供一个可视化操作界面，方便快捷的导出指定模式的表信息。

目前支持导出mysql, oracle, dms。

mysql仅支持8以后的版本。

数据库链接信息均未保存到服务器。

#### 软件架构

- 界面使用fyne框架编写

#### 使用说明

1.  `host` 数据库连接地址
2.  `port` 端口
3.  `username` 用户名
4.  `password` 密码
5.  `scheme` 数据库模式/库名

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

