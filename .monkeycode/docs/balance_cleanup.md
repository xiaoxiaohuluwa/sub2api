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

### Batch 9 - 管理端仪表盘清理（已推送 3feaca3b0）
- 删除快捷操作按钮和 Token 计费展示

### Batch 10 - 运维监控清理（已推送 a5ed92609）
- 删除系统日志板块

### Batch 11 - 管理端渠道管理板块删除（已推送 ae15ef6bc）
- 删除管理端渠道管理入口和页面
- 保留 api/admin/channels 和 components/admin/channel/（GroupsView 复用）

### Batch 12 - 用户侧渠道状态板块删除（已推送 e8d37742d）
- 删除用户侧渠道状态板块
- 保留 ProviderIcon 和 features/channel-monitor-v2/ 目录（当时仍被引用）

### Batch 13 - 前端渠道监控系统删除（已推送 892fc9f98）
- 删除孤儿组件：components/admin/monitor/(11文件)、features/channel-monitor-v2/(12文件)、MonitorQuotaView、ProviderIcon
- 删除渠道监控API：api/admin/channelMonitor.ts、channelMonitorTemplate.ts、api/channelMonitorV2.ts
- 清理 api/admin/channels.ts：删除CRUD函数，保留getModelDefaultPricing/ChannelModelPricing
- 清理 utils/featureFlags.ts：移除channelMonitor FeatureFlag及6个函数
- 清理 SettingsView.vue：删除channel_monitor_*设置卡片、类型、默认值、加载/保存逻辑
- 清理 stores/app.ts、types/index.ts、api/admin/settings.ts：移除channel_monitor_*
- 删除渠道监控i18n和孤儿测试
- 52 files changed, 3 insertions(+), 8189 deletions(-)

### Batch 14 - 后端渠道监控系统删除（已推送 0fc49c52a）
- 删除15个service文件（channel_monitor_*.go）
- 删除6个handler文件（channel_handler.go、channel_monitor_handler.go 等）
- 删除9个repository/domain/schema文件
- 删除21个测试文件
- 删除20个migration文件
- 删除32个ent生成代码文件+4个目录
- 清理 wire.go、handler.go、admin.go、user.go 路由
- 清理 ops_cleanup_service.go、server_timing.go、migrations_runner.go
- 清理 setting_handler*.go、settings.go、setting_parse.go 等11个settings文件
- 清理 domain_constants.go 常量
- 重新生成 ent 和 wire 代码
- 136 files changed, 207 insertions(+), 41705 deletions(-)

### Batch 15 - 前端用户余额功能删除（已推送 76c8add84 + f2e3bf49e）
- KeyUsageView 删除 wallet balance；UsersView 删除 balance_platform_quota 列
- AnnouncementTargetingEditor 删除 balance 条件；EmailTemplateEditor 删除 balance 邮件模板
- UserCreateModal 删除初始余额输入；AnnouncementReadStatusDialog 删除 balance 列
- OpsSettingsDialog 删除 ignore_insufficient_balance_errors
- api/admin/ops.ts、api/admin/settings.ts 删除对应字段
- 7 files changed, 8 insertions(+), 41 deletions(-)

### Batch 16 - 后端 BalanceNotifyService + 余额设置删除（已推送 6902d98b1 + 36c024ba3）
- 删除 balance_notify_service.go 和3个测试文件（1386行）
- 清理 user.go、api_key_auth_cache.go、gateway_usage_billing.go、gateway_service.go 等
- 删除 SettingKeyDefaultBalance 和 SettingKeyBalanceLowNotify* 常量
- 删除 GetDefaultBalance 函数及 auth_service/admin_user 调用者
- 18 files changed + 15 files changed，共 -1728 lines

### Batch 17 - 后端内部计费核心清理（已推送 2f972e776）
- 删除 UsageBillingCommand 的 BalanceCost/SubscriptionCost/APIKeyQuotaCost/APIKeyRateLimitCost 4个字段
- 删除 UsageBillingApplyResult 的 NewBalance/BalanceOverdrafted 字段
- 删除 applyUsageBillingEffects 中的4个扣费分支（subscription/balance/APIKeyQuota/APIKeyRateLimit）
- 删除扣费函数：incrementUsageBillingSubscription、deductUsageBillingBalance、incrementUsageBillingAPIKeyQuota、incrementUsageBillingAPIKeyRateLimit
- 清理 gateway_usage_billing.go：删除4个cost赋值、syncBalanceCacheAfterDeduction、APIKeyQuotaExhausted处理、shouldDeductAPIKeyQuota
- 简化 finalizePostUsageBilling 仅保留平台配额累加
- 简化 legacy postUsageBilling 路径
- 删除4个余额扣费测试
- 保留：AccountQuotaCost（供应商账号配额）、平台配额、用量日志
- 4 files changed, 6 insertions(+), 369 deletions(-)

### Batch 18a - CheckBillingEligibility 简化（已推送 e4bb52f55）
- 删除 CheckBillingEligibility 的 subscription/balance 资格检查
- 函数签名去掉 subscription 参数，更新21个调用点（14个handler文件）
- 删除未使用函数：checkBalanceEligibility、checkSubscriptionEligibility、minimumBalanceReserve、balanceBelowEligibilityThreshold
- 保留：平台配额检查、API Key 限流、RPM 限流
- 15 files changed, 24 insertions(+), 122 deletions(-)

### Batch 18b-e - 计费套餐集群删除（进行中）
- subscription/payment/redeem/affiliate/promo 5个子系统构成强耦合"计费套餐"集群
- 通过 AuthService(注册邀请码)、PaymentService、wire.go、BillingCacheService、APIKeyService 深度耦合
- 需整体删除约90+文件并重构上述枢纽
- 依赖 UserSubscription 实体（Batch 21 处理 ent schema）
- 保留：antigravity（外部供应商网关）、供应商账号配额、平台配额、用量统计
