package scrap

import (
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
	"github.com/oklog/ulid/v2"
)

func SearchResultParser(
	result *SearchResult,
) []search_entity.DomainResearchResult {
	var parsedResults []search_entity.DomainResearchResult
	for _, r := range result.OrganicResults {
		pr := search_entity.DomainResearchResult{
			ID:     ulid.Make().String(),
			Domain: r.Link,
		}
		parsedResults = append(parsedResults, pr)
	}
	return parsedResults
}
