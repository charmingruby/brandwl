package search_usecase

import (
	"fmt"
	"log/slog"

	"github.com/charmingruby/brandwl/internal/core"
	"github.com/charmingruby/brandwl/internal/domain/search/search_dto"
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

func (uc *SearchUseCaseRegistry) SearchDomainConcurrentsUseCase(dto search_dto.DomainConcurrentsResearchDTO) (*search_entity.DomainResearch, error) {
	research := search_entity.NewDomainResearch(dto.Email, dto.Terms)

	result, err := uc.GoogleAPI.SearchDomainConcurrents(dto)
	if err != nil {
		slog.Error(fmt.Sprintf("[Search Domain Concurrents UseCase] Google API: %s", err.Error()))
		return nil, core.NewInternalErr("search domain concurrents") // no explicitation of a infra error
	}

	research.Result = result

	if err := uc.SearchRepository.Store(research); err != nil {
		slog.Error(fmt.Sprintf("[Search Domain Concurrents UseCase] Search repository save: %s", err.Error()))
		return nil, core.NewInternalErr("search domain concurrents") // no explicitation of a infra error
	}

	return research, nil
}
