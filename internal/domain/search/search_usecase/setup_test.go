package search_usecase

import (
	"testing"

	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
	"github.com/charmingruby/brandwl/tests/fake"
	"github.com/charmingruby/brandwl/tests/inmemory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	searchRepo    *inmemory.SearchInMemoryRepository
	fakeGoogleAPI *fake.FakeGoogleAPI
	searchUseCase *SearchUseCaseRegistry
}

// initial setup
func (s *Suite) SetupSuite() {
	s.searchRepo = inmemory.NewSearchInMemoryRepository()
	s.fakeGoogleAPI = fake.NewFakeGoogleAPI()
	s.searchUseCase = NewSearchUseCase(s.searchRepo, s.fakeGoogleAPI)
}

func (s *Suite) SetupTest() {
	s.fakeGoogleAPI.Items = []search_entity.DomainResearchResult{}
	s.searchRepo.Items = []search_entity.DomainResearch{}
}

func (s *Suite) TearDownTest() {
	s.fakeGoogleAPI.Items = []search_entity.DomainResearchResult{}
	s.searchRepo.Items = []search_entity.DomainResearch{}
}

func (s *Suite) SetupSubTest() {
	s.fakeGoogleAPI.Items = []search_entity.DomainResearchResult{}
	s.searchRepo.Items = []search_entity.DomainResearch{}
}

func (s *Suite) TearDownSubTest() {
	s.fakeGoogleAPI.Items = []search_entity.DomainResearchResult{}
	s.searchRepo.Items = []search_entity.DomainResearch{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
