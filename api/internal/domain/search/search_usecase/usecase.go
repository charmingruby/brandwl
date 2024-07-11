package search_usecase

import (
	"github.com/charmingruby/brandwl/internal/domain/search/search_dto"
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

type SearchUseCase interface {
	SearchDomainConcurrents(dto search_dto.DomainConcurrentsResearchRequestDTO) (search_entity.DomainResearch, error)
}
