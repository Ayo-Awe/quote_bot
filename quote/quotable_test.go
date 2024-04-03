package quote

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetQuote(t *testing.T) {
	provider := NewQuotableProvider()

	t.Run("should successfully return a random quote", func(t *testing.T) {
		quote, err := provider.GetQuote("")
		require.NoError(t, err)

		require.NotEmpty(t, quote)
	})

	t.Run("should successfully return a quote in the given category", func(t *testing.T) {
		quote, err := provider.GetQuote("technology")
		require.NoError(t, err)

		require.NotEmpty(t, quote)
		require.Contains(t, quote.Categories, "Technology")
	})
}

func TestSearchQuote(t *testing.T) {
	provider := NewQuotableProvider()

	t.Run("should successfully find matching quote", func(t *testing.T) {
		// Visit https://api.quotable.io/search/quotes?query=every+good+technology+is+basically+magic&fields=content for details
		expectedQuoteID := "T6AMWsNRE5"

		quote, err := provider.Search("every good technology is basically magic")
		require.NoError(t, err)

		require.NotEmpty(t, quote)
		require.Equal(t, quote.ID, expectedQuoteID)
	})

}
