package repository

import (
	"context"
	"time"

	dbent "github.com/Wei-Shaw/sub2api/ent"
	"github.com/Wei-Shaw/sub2api/ent/memo"
	"github.com/Wei-Shaw/sub2api/internal/service"
)

type memoRepository struct {
	client *dbent.Client
}

func NewMemoRepository(client *dbent.Client) service.MemoRepository {
	return &memoRepository{client: client}
}

func (r *memoRepository) Create(ctx context.Context, m *service.Memo) error {
	client := clientFromContext(ctx, r.client)
	created, err := client.Memo.Create().
		SetUserID(m.UserID).
		SetTitle(m.Title).
		SetContent(m.Content).
		SetPinned(m.Pinned).
		Save(ctx)
	if err != nil {
		return err
	}
	m.ID = created.ID
	m.CreatedAt = created.CreatedAt.Format(time.RFC3339)
	m.UpdatedAt = created.UpdatedAt.Format(time.RFC3339)
	return nil
}

func (r *memoRepository) GetByID(ctx context.Context, id, userID int64) (*service.Memo, error) {
	created, err := r.client.Memo.Query().
		Where(
			memo.IDEQ(id),
			memo.UserIDEQ(userID),
		).
		Only(ctx)
	if err != nil {
		return nil, service.ErrMemoNotFound
	}
	return toServiceMemo(created), nil
}

func (r *memoRepository) Update(ctx context.Context, id, userID int64, m *service.Memo) (*service.Memo, error) {
	builder := r.client.Memo.UpdateOneID(id).
		Where(memo.UserIDEQ(userID))
	if m.Title != "" {
		builder = builder.SetTitle(m.Title)
	}
	if m.Content != "" {
		builder = builder.SetContent(m.Content)
	}
	builder = builder.SetPinned(m.Pinned)

	updated, err := builder.Save(ctx)
	if err != nil {
		return nil, service.ErrMemoNotFound
	}
	return toServiceMemo(updated), nil
}

func (r *memoRepository) Delete(ctx context.Context, id, userID int64) error {
	err := r.client.Memo.DeleteOneID(id).
		Where(memo.UserIDEQ(userID)).
		Exec(ctx)
	if err != nil {
		return service.ErrMemoNotFound
	}
	return nil
}

func (r *memoRepository) List(ctx context.Context, userID int64, limit, offset int) ([]service.Memo, error) {
	if limit <= 0 {
		limit = 100
	}
	if limit > 500 {
		limit = 500
	}
	items, err := r.client.Memo.Query().
		Where(memo.UserIDEQ(userID)).
		Order(
			dbent.Desc(memo.FieldPinned),
			dbent.Desc(memo.FieldUpdatedAt),
		).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]service.Memo, 0, len(items))
	for _, item := range items {
		result = append(result, *toServiceMemo(item))
	}
	return result, nil
}

func toServiceMemo(m *dbent.Memo) *service.Memo {
	return &service.Memo{
		ID:        m.ID,
		UserID:    m.UserID,
		Title:     m.Title,
		Content:   m.Content,
		Pinned:    m.Pinned,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
	}
}
