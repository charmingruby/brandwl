package inmemory

import "github.com/charmingruby/brandwl/internal/domain/search/search_entity"

func NewDomainResearchInMemoryRepository() *DomainResearchInMemoryRepository {
	return &DomainResearchInMemoryRepository{
		Items: []search_entity.DomainResearch{},
	}
}

type DomainResearchInMemoryRepository struct {
	Items []search_entity.DomainResearch
}

func (r *DomainResearchInMemoryRepository) Store(dr *search_entity.DomainResearch) error {
	r.Items = append(r.Items, *dr)
	return nil
}
