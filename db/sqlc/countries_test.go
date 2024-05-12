package db

import (
	"testing"

	"context"

	"github.com/stretchr/testify/require"
)

func createRandomCountry(t *testing.T) Country {
	arg := CreateCountryParams{
		Name:          util.createRandomCountry(),
		ContinentName: util.creatRandomContinent(),
	}

	country, err := testQueries.CreateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, country)

	require.Equal(t, arg.Name, country.Name)
	require.Equal(t, arg.ContinentName, country.ContinentName)

	return country
}
