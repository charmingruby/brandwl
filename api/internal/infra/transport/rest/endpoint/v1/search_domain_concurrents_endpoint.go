package v1

import (
	"fmt"
	"log/slog"

	"github.com/charmingruby/brandwl/internal/domain/search/search_dto"
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
	"github.com/charmingruby/brandwl/pkg/mailer/template"
	"github.com/charmingruby/brandwl/pkg/response"
	"github.com/gin-gonic/gin"
)

type SearchDomainConcurrentsRequest struct {
	Email string   `json:"email" binding:"required"`
	Terms []string `json:"terms" binding:"required"`
}

type SearchDomainConcurrentsResponse struct {
	Message string                         `json:"message"`
	Data    []search_entity.DomainResearch `json:"data"`
}

// Search Domain Concurrents godoc
//
//	@Summary		Search Domain Concurrents
//	@Description	Search Domain Concurrents
//	@Tags			Search
//	@Accept			json
//	@Produce		json
//	@Param			request	body		 SearchDomainConcurrentsRequest  	true	"Search Domain Concurrents Payload"
//	@Success		200	{object}	SearchDomainConcurrentsResponse
//	@Failure		500	{object}	response.Response
//	@Router			/search [post]
func (h *Handler) searchDomainConcurrentsEndpoint(c *gin.Context) {
	var req SearchDomainConcurrentsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewPayloadError(c, err)
		return
	}

	dto := search_dto.DomainConcurrentsResearchDTO{
		Email: req.Email,
		Terms: req.Terms,
	}

	result, err := h.searchUseCase.SearchDomainConcurrentsUseCase(dto)
	if err != nil {
		response.NewInternalServerError(c, err)
		return
	}

	if err := h.mailer.SendEmail(
		req.Email, "Brand Monitor Concurrents Search", template.NewConcurrentsResearchResultTemplate(*result),
	); err != nil {
		slog.Error(fmt.Sprintf("[MAILER] Unable to mail to `%s` with concurrents search result", req.Email))
	}

	response.NewOkResponse(c, "concurrents result", result)
}
