# 余额功能清理变更记录

## 概述
从 Sing 项目中彻底删除所有与"用户余额"、"账户余额"相关的外部支付/计费功能残留代码。
核心内部计费（API Key 内部配额）保留，但用户可见的余额展示、充值、通知等功能全部移除。

## 变更批次

### Batch 1 - 前端余额UI组件删除
- 删除 UserBalanceModal.vue, UserBalanceHistoryModal.vue
- 删除 ProfileBalanceNotifyCard.vue
- 删除 Dashboard/Profile 中余额显示
- 删除 Admin 设置中余额通知配置
- 删除 common i18n 中 balance 相关词条
- 删除 ChannelMonitor 中配额/余额展示选项

### Batch 2 - 前端API/Types/Stores余额字段清理
- User 类型中 balance, frozen_balance, balance_notify_* 字段
- Settings 类型中 balance_low_notify_* 字段
- API endpoint 中 balance 相关接口定义
- Store 中 balance 相关状态

### Batch 3 - 后端余额路由/处理器删除
- Admin User balance 更新路由及 handler
- User balance 查询路由
- CN Provider balance 查询路由
- Balance notify handler

### Batch 4 - 后端余额服务删除
- BalanceNotifyService
- CNProviderBalanceService / CNProviderBalanceCheckService
- User UpdateBalance / currentBalance
- Admin User UpdateBalance
- 余额相关配置字段

### Batch 5 - 配置/迁移/Setup清理
- config.yaml / config.go 中 balance 相关配置
- setup.go 中 user_balance 默认值
- 已过期的 balance 相关迁移清理

### Batch 6 - 残余引用清理
- 测试文件更新
- Ent schema 中 balance 字段处理
- wire 依赖注入清理
