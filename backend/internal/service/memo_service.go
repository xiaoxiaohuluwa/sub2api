package service

import (
	"context"
	"strings"
)

type MemoService struct {
	memoRepo MemoRepository
}

func NewMemoService(memoRepo MemoRepository) *MemoService {
	return &MemoService{memoRepo: memoRepo}
}

type CreateMemoInput struct {
	Title   string
	Content string
	Pinned  bool
}

type UpdateMemoInput struct {
	Title   *string
	Content *string
	Pinned  *bool
}

func (s *MemoService) Create(ctx context.Context, userID int64, input *CreateMemoInput) (*Memo, error) {
	if input == nil {
		return nil, ErrMemoInputNeeded
	}
	title := strings.TrimSpace(input.Title)
	content := strings.TrimSpace(input.Content)
	if title == "" || len(title) > 200 {
		return nil, ErrMemoTitle
	}
	if content == "" {
		return nil, ErrMemoContent
	}
	m := &Memo{
		UserID:  userID,
		Title:   title,
		Content: content,
		Pinned:  input.Pinned,
	}
	if err := s.memoRepo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *MemoService) GetByID(ctx context.Context, id, userID int64) (*Memo, error) {
	return s.memoRepo.GetByID(ctx, id, userID)
}

func (s *MemoService) Update(ctx context.Context, id, userID int64, input *UpdateMemoInput) (*Memo, error) {
	if input == nil {
		return nil, ErrMemoInputNeeded
	}
	m := &Memo{}
	if input.Title != nil {
		title := strings.TrimSpace(*input.Title)
		if title == "" || len(title) > 200 {
			return nil, ErrMemoTitle
		}
		m.Title = title
	}
	if input.Content != nil {
		content := strings.TrimSpace(*input.Content)
		if content == "" {
			return nil, ErrMemoContent
		}
		m.Content = content
	}
	if input.Pinned != nil {
		m.Pinned = *input.Pinned
	}
	return s.memoRepo.Update(ctx, id, userID, m)
}

func (s *MemoService) Delete(ctx context.Context, id, userID int64) error {
	return s.memoRepo.Delete(ctx, id, userID)
}

func (s *MemoService) List(ctx context.Context, userID int64, limit, offset int) ([]Memo, error) {
	return s.memoRepo.List(ctx, userID, limit, offset)
}
