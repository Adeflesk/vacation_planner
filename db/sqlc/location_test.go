package db

import (
	"database/sql"
	"testing"

	"context"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomLocation(t *testing.T) Location {
	arg1 := CreateCountryParams{
		Name:          util.RandomString(6),
		ContinentName: util.RandomContinents(),
	}
	country, err := testQueries.CreateCountry(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, country)

	arg2 := CreateLocationParams{
		LocationName:        util.RandomString(6),
		LocationDescription: util.RandomString(45),
		CountryID:           country.ID,
	}

	location, err := testQueries.CreateLocation(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, location)

	require.Equal(t, arg2.LocationName, location.LocationName)
	require.Equal(t, arg2.LocationDescription, location.LocationDescription)
	require.Equal(t, arg2.CountryID, country.ID)

	return location
}

func TestCreateLocation(t *testing.T) {
	createRandomLocation(t)
}

func TestGetLocation(t *testing.T) {
	location1 := createRandomLocation(t)
	location2, err := testQueries.GetLocation(context.Background(), location1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, location2)

	require.Equal(t, location1.ID, location2.ID)
	require.Equal(t, location1.LocationName, location2.LocationName)
	require.Equal(t, location1.LocationDescription, location2.LocationDescription)
	require.Equal(t, location1.CountryID, location2.CountryID)
}

func TestUpdateLocation(t *testing.T) {
	location1 := createRandomLocation(t)

	arg := UpdateLocationParams{
		ID:                  location1.ID,
		LocationName:        util.RandomString(6),
		LocationDescription: util.RandomString(24),
		CountryID:           location1.CountryID,
	}

	location2, err := testQueries.UpdateLocation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, location2)

	require.Equal(t, location1.ID, location2.ID)
	require.Equal(t, arg.LocationName, location2.LocationName)

}
func TestDeleteLocation(t *testing.T) {
	location1 := createRandomLocation(t)
	err := testQueries.DeleteLocation(context.Background(), location1.ID)
	require.NoError(t, err)

	location2, err := testQueries.GetLocation(context.Background(), location1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, location2)
}

func TestListlocations(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomLocation(t)
	}
	arg := ListLocationsParams{
		Limit:  5,
		Offset: 0,
	}
	locations, err := testQueries.ListLocations(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, locations, 5)

	for _, location := range locations {
		require.NotEmpty(t, location)
	}
}

func TestGetlocationsByCountry(t *testing.T) {
	location1 := createRandomLocation(t)
	countryID := location1.CountryID

	locations, err := testQueries.GetLocationsByCountry(context.Background(), countryID)

	require.NoError(t, err)
	require.NotEmpty(t, locations)

	for _, location := range locations {
		require.Equal(t, location.CountryID, countryID)
	}
}
