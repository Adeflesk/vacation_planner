package db

import (
	"testing"

	"context"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomCountry(t *testing.T) Country {
	arg := CreateCountryParams{
		Name:          util.RandomCountry(),
		ContinentName: util.RandomContinents(),
	}

	country, err := testQueries.CreateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, country)

	require.Equal(t, arg.Name, country.Name)
	require.Equal(t, arg.ContinentName, country.ContinentName)

	return country
}

func TestCreateCountry(t *testing.T) {
	createRandomCountry(t)
}

func TestGetCountry(t *testing.T) {
	country1 := createRandomCountry(t)
	country2, err := testQueries.GetCountry(context.Background(), country1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, country2)

	require.Equal(t, country1.ID, country2.ID)
	require.Equal(t, country1.Name, country2.Name)
	require.Equal(t, country1.ContinentName, country2.ContinentName)
}
