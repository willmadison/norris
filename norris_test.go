package norris_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/norris"
)

func TestFetchChuckNorrisFactCategories(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `["explicit","dev","movie","food","celebrity","science"]`)
	}))
	defer ts.Close()

	c := norris.New(ts.URL)

	categories, err := c.Categories()

	assert.Nil(t, err)
	assert.NotEmpty(t, categories)
	assert.Equal(t, 6, len(categories))
}

func TestFetchChuckNorrisFact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"category":null,
			"icon_url": "https:\/\/assets.chucknorris.host\/img\/avatar\/chuck-norris.png",
			"id": "yu8u7ggkqwwe9dj4rqj3qw",
			"url": "https:\/\/api.chucknorris.io\/jokes\/yu8u7ggkqwwe9dj4rqj3qw",
			"value":"Chuck Norris can taste lies."
		}`)
	}))
	defer ts.Close()

	c := norris.New(ts.URL)

	fact, err := c.Fact()

	assert.Nil(t, err)
	assert.Equal(t, "Chuck Norris can taste lies.", fact.Value)
	assert.Empty(t, fact.Category)
}

func TestFetchChuckNorrisFactByCategory(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"category": [
			  "food"
			],
			"icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
			"id": "e690xhz_te2hnf7nk7ppfq",
			"url": "https://api.chucknorris.io/jokes/e690xhz_te2hnf7nk7ppfq",
			"value": "Chuck Norris eats steak for every single meal. Most times he forgets to kill the cow."
		  }`)
	}))
	defer ts.Close()

	c := norris.New(ts.URL)

	fact, err := c.Categorized(norris.Category("food"))

	assert.Nil(t, err)
	assert.Equal(t, "Chuck Norris eats steak for every single meal. Most times he forgets to kill the cow.", fact.Value)
	assert.NotEmpty(t, fact.Category)
	assert.Equal(t, "food", fact.Category[0])
}

func TestNorrisIntegration(t *testing.T) {
	t.Skip()
	c := norris.New("https://api.chucknorris.io")

	fact, err := c.Fact()

	assert.Nil(t, err)
	assert.NotEmpty(t, fact.Value)

	categories, err := c.Categories()
	assert.Nil(t, err)

	for _, category := range categories {
		fact, err := c.Categorized(category)

		assert.Nil(t, err)
		assert.NotEmpty(t, fact.Value)

		if len(fact.Category) > 0 {
			assert.Equal(t, category, norris.Category(fact.Category[0]))
		}
	}
}
