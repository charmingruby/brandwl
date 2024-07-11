package search_usecase

import (
	"github.com/charmingruby/brandwl/internal/domain/search/search_adapter"
	"github.com/charmingruby/brandwl/internal/domain/search/search_dto"
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
	"github.com/charmingruby/brandwl/internal/domain/search/search_repository"
)

type SearchUseCase interface {
	SearchDomainConcurrents(dto search_dto.DomainConcurrentsResearchDTO) (search_entity.DomainResearch, error)
}

func NewSearchUseCase(searchRepo search_repository.DomainResearchRepository, googleAPI *search_adapter.GoogleAPIAdapter) *SearchUseCaseRegistry {
	return &SearchUseCaseRegistry{
		SearchRepository: searchRepo,
		GoogleAPI:        googleAPI,
	}
}

type SearchUseCaseRegistry struct {
	SearchRepository search_repository.DomainResearchRepository
	GoogleAPI        *search_adapter.GoogleAPIAdapter
}
