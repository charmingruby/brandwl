package search_entity

import "github.com/oklog/ulid/v2"

func NewDomainResearchResult(domain string) *DomainResearchResult {
	res := DomainResearchResult{
		ID:     ulid.Make().String(),
		Domain: domain,
	}

	return &res
}

type DomainResearchResult struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
}
