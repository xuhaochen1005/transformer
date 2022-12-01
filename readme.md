# 明珠电气变压器设计系统

## 项目内容说明

### 后端
本项目的后端采用Golang语言开发，涉及到的工具包有
* [jwt-go](https://github.com/dgrijalva/jwt-go)
* [gorm ](https://github.com/go-gorm/gorm)
* [casbin](https://github.com/casbin/casbin)
* [gin](https://github.com/gin-gonic/gin)

### 前端
本项目前端采用vue.js,具体框架采用 [vue-element-admin]


#### 前端待完成
 * [ ] 用户登录页
 * [ ] 用户密码重置、找回  
 * [ ] 系统首页
 * [ ] 用户页面
   * [ ] 用户详情页【个人信息修改】
   * [ ] 用户列表页【增、删、改、查】
 * [ ] 部门管理页
   * [ ] 部门详情页面【部门信息修改】
   * [ ] 部门列表页面【增、删、改、查】
 
 * [ ] 权限管理页
   * [ ] 权限列表页面【增、删、改、查】
   * [ ] 用户权限调整
 
 * [ ] 标准库维护页面
   * [ ] 标准库列表【增、删、改、查】
   
#### 后端待完成
待完成  
 * [X] jwt 验证集成
 * [X] logurs 日志库集成
 * [X] trace 请求追踪功能
 * [X] 后台跨域问题解决  
 * [X] gorm sql语句的记录与数据库配置优化
 * [X] Error错误处理
 * [X] 配置文件解析，go-yaml进行配置文件的解析
 * [ ] pdf、execl、word等文件格式的导出
 * [X] 基于rbac的权限控制 https://github.com/casbin/casbin
    * [X] RBAC模型 【针对用户；支持多角色和角色继承】
    * [X] 整合casbin和restful权限控制 【参考https://github.com/casbin/casbin/issues/177】  
    * [X] 资源模型 【针对资源】
    * [X] root用户【超级管理员】
 * [X] 业务审批流
 * [X] 通知服务；站内通知或者短信、邮箱等
 * [ ] 计算算法库，数学公式维护
 * [ ] 标准值表维护
 * [ ] 前端开发 vue+element-ui
 * [ ] 程序部署方案确定【保证双机高可用】
 

参考
1. https://github.com/casbin/gorm-adapter
2. https://github.com/casbin/casbin
3. https://github.com/sirupsen/logrus
4. https://github.com/appleboy/gin-jwt
 
