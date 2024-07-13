package scrap

type SearchResult struct {
	Meta               Meta              `json:"meta"`
	UserCreditsInfo    UserCreditsInfo   `json:"user_credits_info"`
	MenuItems          []MenuItem        `json:"menu_items"`
	PeopleAlsoAskedFor []PeopleAlsoAsked `json:"peopleAlsoAskedFor"`
	OrganicResults     []OrganicResult   `json:"organic_results"`
	RelatedSearches    []RelatedSearch   `json:"relatedSearches"`
}

type Meta struct {
	Q      string `json:"q"`
	Gl     string `json:"gl"`
	Hl     string `json:"hl"`
	Domain string `json:"domain"`
	ApiKey string `json:"api_key"`
}

type UserCreditsInfo struct {
	Quota        int `json:"quota"`
	Requests     int `json:"requests"`
	RequestsLeft int `json:"requests_left"`
}

type MenuItem struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Position int    `json:"position"`
}

type PeopleAlsoAsked struct {
	Question      string `json:"question"`
	ID            string `json:"id"`
	Rank          int    `json:"rank"`
	DisplayedLink string `json:"displayed_link"`
	Link          string `json:"link"`
	Title         string `json:"title"`
	Answer        string `json:"answer"`
}

type OrganicResult struct {
	Title             string             `json:"title"`
	DisplayedLink     string             `json:"displayed_link"`
	Snippet           string             `json:"snippet"`
	Link              string             `json:"link"`
	ExtendedSitelinks []ExtendedSitelink `json:"extended_sitelinks"`
	Rank              int                `json:"rank"`
}

type ExtendedSitelink struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type RelatedSearch struct {
	Query string `json:"query"`
	Link  string `json:"link"`
}
