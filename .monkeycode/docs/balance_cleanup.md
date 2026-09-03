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

### Batch 3 - 后端余额路由/处理器/服务删除（已推送 05e76f5e0）
- `user_handler.go` 移除 `UpdateBalanceRequest` struct、`UpdateBalance`、`GetBalanceHistory` 方法
- `admin.go` 路由移除 `POST /:id/balance` 和 `GET /:id/balance-history`
- `admin_service.go` 接口移除 `UpdateUserBalance`、`GetUserBalanceHistory`
- `admin_user.go` 移除上述方法及其全部内部辅助函数（`tryAccrueAffiliateRebateForAdminRecharge`、`getAllUserBalanceHistory`、`listRedeemBalanceHistoryForMerge`、`listAffiliateBalanceHistoryForMerge`、`listAffiliateBalanceHistory`、`countAffiliateBalanceHistory`、`mergeBalanceHistoryCodes`、`redeemCodeHistoryTime`）
- `dto/types.go` User DTO 移除 `Balance`、`FrozenBalance`、`BalanceNotifyEnabled/Threshold/ThresholdType/ExtraEmails`、`TotalRecharged` 字段
- `dto/mappers.go` 移除对应映射行
- `gateway_handler.go` 用户信息响应移除 `remaining`、`balance` 字段
- `admin_service_stub_test.go` 移除 `GetUserBalanceHistory` 桩方法
- `admin_basic_handlers_test.go` 移除 balance 路由测试

### Batch 4 - 后端用户余额通知配置与API清理（已推送 319a5a3a9->当前批次）
- `user_handler.go` 移除 `SendNotifyEmailCode`/`VerifyNotifyEmail`/`RemoveNotifyEmail`/`ToggleNotifyEmail` handler 及相关 request struct
- `user_handler.go` `UpdateProfileRequest` 移除 `BalanceNotifyEnabled`/`BalanceNotifyThreshold`
- `user_service.go` `UpdateProfileRequest` 移除 `BalanceNotifyEnabled`/`BalanceNotifyThreshold`
- `user_service.go` `UserUpdateFields` 移除 `BalanceNotifySettings`/`BalanceNotifyExtraEmails`
- `user_service.go` 移除 `SendNotifyEmailCode`/`VerifyAndAddNotifyEmail`/`addOrVerifyNotifyEmail`/`RemoveNotifyEmail`/`ToggleNotifyEmail` 及辅助函数
- `user_service.go` 移除 `maxNotifyEmails`/`notifyCodeUserRateLimit`/`notifyCodeUserRateWindow` 常量
- `user_service.go` 移除 `notifyVerifyEmailTemplate` HTML 模板
- `routes/user.go` 移除 `/notify-email` 路由组
- `gateway_usage_billing.go` 移除日志中的 `notify_enabled`/`threshold` 字段

### Batch 5 - 配置/Setup清理（待提交）
- setup 生成的配置文件移除 `user_balance` 输出项
- `backend/config.yaml` 和 `deploy/config.example.yaml` 移除 `user_balance`
- 保留 `DefaultConfig.UserBalance` 及历史迁移，供现有内部用户账本和已部署数据库兼容

### Batch 6 - 残余引用清理（进行中）
- 删除已移除管理员余额服务对应的测试和测试桩
- 删除基础 handler 测试中已失效的 `/balance` 请求
- 保留内部账本、余额缓存、账户配额通知及其 Ent 字段，维持 API Key 计费链路
