package fake

import (
	"fmt"

	"github.com/charmingruby/brandwl/internal/domain/search/search_dto"
	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

func NewFakeGoogleAPI() *FakeGoogleAPI {
	return &FakeGoogleAPI{
		Items: []search_entity.DomainResearchResult{},
	}

}

type FakeGoogleAPI struct {
	Items []search_entity.DomainResearchResult
}

func (g *FakeGoogleAPI) GenerateFakeResults(amount int) {
	for i := 0; i < amount; i++ {
		res := search_entity.NewDomainResearchResult(
			fmt.Sprintf("domain%d.com", i),
			fmt.Sprintf("domain%d", i),
			fmt.Sprintf("domain%d description", i),
		)

		g.Items = append(g.Items, *res)
	}
}

func (g *FakeGoogleAPI) SearchDomainConcurrents(dto search_dto.DomainConcurrentsResearchDTO) ([]search_entity.DomainResearchResult, error) {
	return g.Items, nil
}
