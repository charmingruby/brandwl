package inmemory

import "github.com/charmingruby/brandwl/internal/domain/search/search_entity"

func NewSearchInMemoryRepository() *SearchInMemoryRepository {
	return &SearchInMemoryRepository{
		Items: []search_entity.DomainResearch{},
	}
}

type SearchInMemoryRepository struct {
	Items []search_entity.DomainResearch
}

func (r *SearchInMemoryRepository) Store(dr *search_entity.DomainResearch) error {
	r.Items = append(r.Items, *dr)
	return nil
}
