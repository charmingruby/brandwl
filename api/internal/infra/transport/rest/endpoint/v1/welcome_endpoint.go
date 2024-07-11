package v1

import (
	"github.com/charmingruby/brandwl/pkg/response"
	"github.com/gin-gonic/gin"
)

// Welcome godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Router			/welcome [get]
func (h *Handler) welcomeEndpoint(c *gin.Context) {
	response.NewOkResponse(c, "OK!", nil)
}
