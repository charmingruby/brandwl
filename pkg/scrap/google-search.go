package scrap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

const (
	DOMAIN_API_URL = "https://api.serpdog.io"
)

func NewGoogleSearch(key string) *GoogleSearch {
	url := fmt.Sprintf("%s/search?api_key=%s", DOMAIN_API_URL, key)

	return &GoogleSearch{
		BaseURL: url,
	}
}

type GoogleSearch struct {
	BaseURL string
}

func (g *GoogleSearch) SearchDomain(terms []string) ([]search_entity.DomainResearchResult, error) {
	var query string
	for idx, term := range terms {
		query += term

		if idx != len(terms)-1 {
			query += "+"
		}
	}

	url := fmt.Sprintf("%s&q=%s",
		g.BaseURL,
		query,
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response SearchResult
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	parsedResults := SearchResultParser(&response)

	return parsedResults, nil
}
