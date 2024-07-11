package search_entity

import "github.com/oklog/ulid/v2"

func NewDomainResearchResult(domain, title, description string) *DomainResearchResult {
	res := DomainResearchResult{
		ID:          ulid.Make().String(),
		Domain:      domain,
		Title:       title,
		Description: description,
	}

	return &res
}

type DomainResearchResult struct {
	ID          string
	Domain      string
	Title       string
	Description string
}
