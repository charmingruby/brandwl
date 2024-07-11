package search_usecase

import "github.com/charmingruby/brandwl/internal/domain/search/search_dto"

func (s *Suite) Test_SearchDomainConcurrentsUseCase() {
	s.Run("it should be able to search domain concurrents successfully", func() {
		concurrentResultsAmount := 4
		s.fakeGoogleAPI.GenerateFakeResults(concurrentResultsAmount)
		s.Equal(len(s.fakeGoogleAPI.Items), concurrentResultsAmount)

		dto := search_dto.DomainConcurrentsResearchDTO{
			Email: "domain.com",
			Terms: []string{"dummy"},
		}

		result, err := s.searchUseCase.SearchDomainConcurrentsUseCase(dto)

		s.NoError(err)
		s.Equal(result.ID, s.searchRepo.Items[0].ID)
		s.Equal(result.Email, s.searchRepo.Items[0].Email)
		s.Equal(result.Terms, s.searchRepo.Items[0].Terms)
		s.Equal(len(result.Result), concurrentResultsAmount)
	})
}
