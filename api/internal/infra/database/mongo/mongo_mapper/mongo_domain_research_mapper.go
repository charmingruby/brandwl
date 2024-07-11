package mongo_mapper

import "github.com/charmingruby/brandwl/internal/domain/search/search_entity"

func DomainResearchToMongo(dr *search_entity.DomainResearch) *MongoDomainResearch {
	convResults := []MongoDomainResearchResult{}

	for _, r := range dr.Result {
		convResult := MongoDomainResearchResult{
			ID:          r.ID,
			Domain:      r.Domain,
			Title:       r.Title,
			Description: r.Description,
		}
		convResults = append(convResults, convResult)
	}

	mdr := MongoDomainResearch{
		ID:     dr.ID,
		Email:  dr.Email,
		Terms:  dr.Terms,
		Result: convResults,
	}

	return &mdr
}

type MongoDomainResearch struct {
	ID     string                      `json:"id" bson:"id"`
	Email  string                      `json:"email" bson:"email"`
	Terms  []string                    `json:"terms" bson:"terms"`
	Result []MongoDomainResearchResult `json:"result" bson:"result"`
}

type MongoDomainResearchResult struct {
	ID          string `json:"id" bson:"_id"`
	Domain      string `json:"domain" bson:"domain"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"id"`
}
