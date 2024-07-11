package search_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDomainResearch(t *testing.T) {
	t.Run("it should be able to create a new domain research entity", func(t *testing.T) {
		email := "dummy@email.com"
		terms := []string{"dummy term"}

		r := NewDomainResearch(email, terms)

		assert.Equal(t, r.Email, email)
		assert.Equal(t, r.Terms, terms)
		assert.Nil(t, r.Result)
	})
}
