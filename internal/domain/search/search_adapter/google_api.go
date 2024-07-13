package search_adapter

import (
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

type GoogleAPIAdapter interface {
	SearchDomain(terms []string) ([]search_entity.DomainResearchResult, error)
}
