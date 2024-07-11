package v1

import (
	docs "github.com/charmingruby/brandwl/docs"
	"github.com/charmingruby/brandwl/internal/domain/search/search_usecase"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewV1Handler(router *gin.Engine, searchUseCase search_usecase.SearchUseCase) *Handler {
	return &Handler{
		router:        router,
		searchUseCase: searchUseCase,
	}
}

type Handler struct {
	router        *gin.Engine
	searchUseCase search_usecase.SearchUseCase
}

func (h *Handler) Register() {
	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	{
		v1.GET("/welcome", h.welcomeEndpoint)
	}

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
