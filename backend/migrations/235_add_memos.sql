-- 235: Add memos table for user personal notes
CREATE TABLE IF NOT EXISTS memos (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    pinned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS memos_user_id_idx ON memos (user_id);
CREATE INDEX IF NOT EXISTS memos_pinned_updated_at_idx ON memos (pinned DESC, updated_at DESC);
