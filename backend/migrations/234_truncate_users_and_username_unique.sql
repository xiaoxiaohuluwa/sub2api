-- 234: 移除邮箱登录，切换为 username 登录。
--
-- 背景：
--   产品决定彻底移除邮箱登录与注册，只保留 username + password 登录；
--   用户账户数据整体清空重建（仅由管理员 bootstrap 创建唯一账号）。
--
-- 1) 清空 users 表及其外键依赖表，配合 RESTART IDENTITY 重置自增序列。
--    该迁移在事务内执行，TRUNCATE CASCADE 会级联清空引用 users 的关联数据。
-- 2) 为 username 创建部分唯一索引：
--    - 仅约束 "未软删除且 username 非空" 的行（软删除后可复用用户名）；
--    - email 的部分唯一索引（016 迁移）保留，依然作为内部字段使用。
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username_unique_key
    ON users (username)
    WHERE deleted_at IS NULL AND username <> '';