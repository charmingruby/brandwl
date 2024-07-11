package search_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDomainResearchResult(t *testing.T) {
	t.Run("it should be able to create a new domain research result entity", func(t *testing.T) {
		domain := "dummy.com"

		r := NewDomainResearchResult(domain)

		assert.Equal(t, r.Domain, domain)
	})
}
