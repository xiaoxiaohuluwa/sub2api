package dto

import "github.com/Wei-Shaw/sub2api/internal/service"

type Memo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Pinned    bool   `json:"pinned"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateMemoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Pinned  bool   `json:"pinned"`
}

type UpdateMemoRequest struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
	Pinned  *bool   `json:"pinned,omitempty"`
}

func MemoFromService(m *service.Memo) *Memo {
	if m == nil {
		return nil
	}
	return &Memo{
		ID:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		Pinned:    m.Pinned,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
