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
