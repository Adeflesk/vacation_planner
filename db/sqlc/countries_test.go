package db

import (
	"database/sql"
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

func TestUpdateCountry(t *testing.T) {
	country1 := createRandomCountry(t)

	arg := UpdateCountryParams{
		ID:            country1.ID,
		Name:          util.RandomString(6),
		ContinentName: country1.ContinentName,
	}
	country2, err := testQueries.UpdateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, country2)

	require.Equal(t, country1.ID, country2.ID)
	require.Equal(t, arg.Name, country2.Name)

}
func TestDeleteCountry(t *testing.T) {
	country1 := createRandomCountry(t)
	err := testQueries.DeleteCountry(context.Background(), country1.ID)
	require.NoError(t, err)

	country2, err := testQueries.GetCountry(context.Background(), country1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, country2)
}

func TestListCountries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCountry(t)
	}
	arg := ListCountriesParams{
		Limit:  5,
		Offset: 0,
	}
	countries, err := testQueries.ListCountries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, countries, 5)

	for _, country := range countries {
		require.NotEmpty(t, country)
	}
}

func TestGetCountriesByContinent(t *testing.T) {
	country1 := createRandomCountry(t)
	continentName := country1.ContinentName

	countries, err := testQueries.GetCountriesByContinent(context.Background(), continentName)

	require.NoError(t, err)
	require.NotEmpty(t, countries)

	for _, country := range countries {
		require.Equal(t, country.ContinentName, continentName)
	}
}
