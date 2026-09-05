package handler

import (
	"strconv"

	"github.com/Wei-Shaw/sub2api/internal/handler/dto"
	"github.com/Wei-Shaw/sub2api/internal/pkg/response"
	middleware2 "github.com/Wei-Shaw/sub2api/internal/server/middleware"
	"github.com/Wei-Shaw/sub2api/internal/service"

	"github.com/gin-gonic/gin"
)

type MemoHandler struct {
	memoService *service.MemoService
}

func NewMemoHandler(memoService *service.MemoService) *MemoHandler {
	return &MemoHandler{memoService: memoService}
}

// List GET /api/v1/memos
func (h *MemoHandler) List(c *gin.Context) {
	subject, ok := middleware2.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "User not found in context")
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	items, err := h.memoService.List(c.Request.Context(), subject.UserID, limit, offset)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	out := make([]dto.Memo, 0, len(items))
	for i := range items {
		out = append(out, *dto.MemoFromService(&items[i]))
	}
	response.Success(c, out)
}

// Get GET /api/v1/memos/:id
func (h *MemoHandler) Get(c *gin.Context) {
	subject, ok := middleware2.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "User not found in context")
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	m, err := h.memoService.GetByID(c.Request.Context(), id, subject.UserID)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, dto.MemoFromService(m))
}

// Create POST /api/v1/memos
func (h *MemoHandler) Create(c *gin.Context) {
	subject, ok := middleware2.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "User not found in context")
		return
	}
	var req dto.CreateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
		return
	}
	input := &service.CreateMemoInput{
		Title:   req.Title,
		Content: req.Content,
		Pinned:  req.Pinned,
	}
	m, err := h.memoService.Create(c.Request.Context(), subject.UserID, input)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, dto.MemoFromService(m))
}

// Update PUT /api/v1/memos/:id
func (h *MemoHandler) Update(c *gin.Context) {
	subject, ok := middleware2.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "User not found in context")
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req dto.UpdateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
		return
	}
	input := &service.UpdateMemoInput{
		Title:   req.Title,
		Content: req.Content,
		Pinned:  req.Pinned,
	}
	m, err := h.memoService.Update(c.Request.Context(), id, subject.UserID, input)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, dto.MemoFromService(m))
}

// Delete DELETE /api/v1/memos/:id
func (h *MemoHandler) Delete(c *gin.Context) {
	subject, ok := middleware2.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "User not found in context")
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.memoService.Delete(c.Request.Context(), id, subject.UserID); err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, gin.H{"message": "deleted"})
}
