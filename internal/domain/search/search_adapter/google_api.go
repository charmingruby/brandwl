package search_adapter

import (
	"github.com/charmingruby/brandwl/internal/domain/search/search_dto"
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

type GoogleAPIAdapter interface {
	SearchDomainConcurrents(dto search_dto.DomainConcurrentsResearchDTO) ([]search_entity.DomainResearchResult, error)
}
