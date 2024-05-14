package db

import (
	"database/sql"
	"testing"

	"context"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccommodation(t *testing.T) Accommodation {
	area := createRandomLocation(t)
	Accommodationtype := createRandomAccommodationType(t)
	phoneNumber, err := util.GenerateRandomPhoneNumber()
	require.NoError(t, err)

	arg := CreateAccommodationParams{
		AccommodationName:        util.RandomString(6),
		Pernight:                 util.RandomInt(1, 1000),
		AccommodationType:        Accommodationtype.ID,
		AccommodationDescription: util.RandomString(20),
		Webaddress:               util.RandomWebsite(),
		Emailaddress:             util.GenerateRandomEmail(),
		Phonenumber:              phoneNumber,
		Area:                     area.ID,
	}

	Accommodation, err := testQueries.CreateAccommodation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Accommodation)

	require.Equal(t, arg.AccommodationName, Accommodation.AccommodationName)
	require.Equal(t, arg.Pernight, Accommodation.Pernight)
	require.Equal(t, Accommodationtype.ID, Accommodation.AccommodationType)
	require.Equal(t, arg.AccommodationDescription, Accommodation.AccommodationDescription)
	require.Equal(t, arg.Webaddress, Accommodation.Webaddress)
	require.Equal(t, arg.Emailaddress, Accommodation.Emailaddress)
	require.Equal(t, arg.Phonenumber, Accommodation.Phonenumber)
	require.Equal(t, area.ID, Accommodation.Area)

	return Accommodation
}

func TestCreateAccommodation(t *testing.T) {
	createRandomAccommodation(t)
}

func TestGetAccommodation(t *testing.T) {
	Accommodation1 := createRandomAccommodation(t)
	Accommodation2, err := testQueries.GetAccommodation(context.Background(), Accommodation1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Accommodation2)

	require.Equal(t, Accommodation1.ID, Accommodation2.ID)
	require.Equal(t, Accommodation1.AccommodationName, Accommodation2.AccommodationName)
	require.Equal(t, Accommodation1.Pernight, Accommodation2.Pernight)
	require.Equal(t, Accommodation1.AccommodationType, Accommodation2.AccommodationType)
	require.Equal(t, Accommodation1.AccommodationDescription, Accommodation2.AccommodationDescription)
	require.Equal(t, Accommodation1.Webaddress, Accommodation2.Webaddress)
	require.Equal(t, Accommodation1.Emailaddress, Accommodation2.Emailaddress)
	require.Equal(t, Accommodation1.Phonenumber, Accommodation2.Phonenumber)
	require.Equal(t, Accommodation1.Area, Accommodation2.Area)
}

func TestUpdateAccommodation(t *testing.T) {

	Accommodationtype := createRandomAccommodationType(t)

	Accommodation1 := createRandomAccommodation(t)

	arg := UpdateAccommodationParams{
		ID:                Accommodation1.ID,
		AccommodationName: Accommodation1.AccommodationName,
		Pernight:          Accommodation1.Pernight,
		AccommodationType: Accommodationtype.ID,
		Webaddress:        Accommodation1.Webaddress,
		Emailaddress:      Accommodation1.Emailaddress,
		Phonenumber:       Accommodation1.Phonenumber,
		Area:              Accommodation1.Area,
	}
	Accommodation2, err := testQueries.UpdateAccommodation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Accommodation2)

	require.Equal(t, Accommodation1.ID, Accommodation2.ID)
	require.Equal(t, arg.AccommodationName, Accommodation2.AccommodationName)
	require.Equal(t, Accommodation1.Area, Accommodation2.Area)
	require.Equal(t, arg.AccommodationType, Accommodation2.AccommodationType)
	require.Equal(t, arg.Webaddress, Accommodation2.Webaddress)
	require.Equal(t, Accommodation1.Emailaddress, Accommodation2.Emailaddress)
	require.Equal(t, Accommodation1.Phonenumber, Accommodation2.Phonenumber)
	require.Equal(t, Accommodation1.Area, Accommodation2.Area)

}
func TestDeleteAccommodation(t *testing.T) {
	Accommodation1 := createRandomAccommodation(t)
	err := testQueries.DeleteAccommodation(context.Background(), Accommodation1.ID)
	require.NoError(t, err)

	Accommodation2, err := testQueries.GetAccommodation(context.Background(), Accommodation1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, Accommodation2)
}

func TestListAccommodations(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccommodation(t)
	}
	arg := ListAccommodationParams{
		Limit:  5,
		Offset: 0,
	}
	Accommodations, err := testQueries.ListAccommodation(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Accommodations, 5)

	for _, Accommodation := range Accommodations {
		require.NotEmpty(t, Accommodation)
	}
}

func TestGetAccommodationByArea(t *testing.T) {
	Accommodation1 := createRandomAccommodation(t)
	area := Accommodation1.Area

	locations, err := testQueries.GetAccommodationByLocation(context.Background(), area)

	require.NoError(t, err)
	require.NotEmpty(t, locations)

	for _, Accommodation := range locations {
		require.Equal(t, Accommodation.Area, area)
	}
}
