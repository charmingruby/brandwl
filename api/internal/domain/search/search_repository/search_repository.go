package search_repository

import "github.com/charmingruby/brandwl/internal/domain/search/search_entity"

type DomainResearchRepository interface {
	Store(r *search_entity.DomainResearch) error
}
