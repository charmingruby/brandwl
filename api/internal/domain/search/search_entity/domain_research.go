package search_entity

import "github.com/oklog/ulid/v2"

func NewDomainResearch(email string, terms []string) *DomainResearch {
	cs := DomainResearch{
		ID:     ulid.Make().String(),
		Email:  email,
		Terms:  terms,
		Result: nil,
	}

	return &cs
}

type DomainResearch struct {
	ID     string   `json:"id"`
	Email  string   `json:"email"`
	Terms  []string `json:"terms"`
	Result []DomainResearchResult
}
