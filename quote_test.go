package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetQuote(t *testing.T) {

	quote, err := GetQuote()

	require.Nil(t, err)
	require.NotEmpty(t, quote)

	fmt.Print(*quote)
}
