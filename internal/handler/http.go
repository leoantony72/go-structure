package handler

import (
	"test/internal/middleware"
	"test/internal/ports"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testService ports.TestService
}

func NewTestHandler(r *gin.Engine, m *middleware.TestMiddleware, s ports.TestService) {
	handler := &TestHandler{
		testService: s,
	}
	r.GET("/", handler.getData)
}
