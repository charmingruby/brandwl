package template

import (
	"fmt"

	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
)

func NewConcurrentsResearchResultTemplate(research search_entity.DomainResearch) string {
	spacing := "<br>"
	header := "<h1>Concurrents Research Result</h1>"

	var terms string
	for idx, t := range research.Terms {
		if idx+1 != len(research.Terms) {
			terms += fmt.Sprintf("%s, ", t)
		} else {
			terms += t
		}
	}

	termsHeader := "<h2>Terms</h2>"

	resultsHeader := "<h2>Domain results</h2>"

	var results string
	if len(research.Result) != 0 {
		for idx, r := range research.Result {
			var result string
			result += "[ " + r.Domain + " ]"
			if idx != len(research.Result)-1 {
				result += spacing
			}

			results += result
		}
	} else {
		results += "No concurrent domain found"
	}

	content := header + termsHeader + terms + spacing + resultsHeader + results

	return content
}
