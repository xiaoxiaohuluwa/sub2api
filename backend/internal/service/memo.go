package service

import (
	"context"

	infraerrors "github.com/Wei-Shaw/sub2api/internal/pkg/errors"
)

var (
	ErrMemoNotFound    = infraerrors.NotFound("MEMO_NOT_FOUND", "memo not found")
	ErrMemoInputNeeded = infraerrors.BadRequest("MEMO_INPUT_REQUIRED", "memo input is required")
	ErrMemoTitle       = infraerrors.BadRequest("MEMO_TITLE_INVALID", "memo title is required and must be 200 characters or less")
	ErrMemoContent     = infraerrors.BadRequest("MEMO_CONTENT_REQUIRED", "memo content is required")
)

type Memo struct {
	ID        int64
	UserID    int64
	Title     string
	Content   string
	Pinned    bool
	CreatedAt string
	UpdatedAt string
}

type MemoRepository interface {
	Create(ctx context.Context, m *Memo) error
	GetByID(ctx context.Context, id, userID int64) (*Memo, error)
	Update(ctx context.Context, id, userID int64, m *Memo) (*Memo, error)
	Delete(ctx context.Context, id, userID int64) error
	List(ctx context.Context, userID int64, limit, offset int) ([]Memo, error)
}
