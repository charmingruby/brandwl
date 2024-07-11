package search_adapter

import "github.com/charmingruby/brandwl/internal/domain/search/search_dto"

type GoogleAPIAdapter interface {
	SearchDomainConcurrents(dto search_dto.DomainConcurrentsResearchDTO)
}
