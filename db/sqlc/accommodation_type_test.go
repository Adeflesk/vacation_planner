package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccommodationType(t *testing.T) AccommodationType {
	atype := util.RandomString(9)
	accommodation_type, err := testQueries.CreateAccommodation_type(context.Background(), atype)

	require.NoError(t, err)
	require.NotEmpty(t, accommodation_type)

	require.Equal(t, accommodation_type.Type, atype)
	return accommodation_type

}

func TestAccommodationType(t *testing.T) {
	createRandomAccommodationType(t)
}

func TestListAllAccommodationTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccommodationType(t)
	}
	arg := ListAccommodation_typesParams{
		Limit:  5,
		Offset: 0,
	}
	accommodationtypes, err := testQueries.ListAccommodation_types(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accommodationtypes, 5)

	for _, accommodation_type := range accommodationtypes {
		require.NotEmpty(t, accommodation_type)
	}

}

func TestDeleteAccommodationType(t *testing.T) {
	accommodation_type1 := createRandomAccommodationType(t)
	err := testQueries.DeleteAccommodation_type(context.Background(), accommodation_type1.ID)
	require.NoError(t, err)

	accommodation_type2, err := testQueries.GetAccommodation_type(context.Background(), accommodation_type1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accommodation_type2)
}

func TestUpdateAccommodationType(t *testing.T) {
	accommodation_type1 := createRandomAccommodationType(t)

	arg := UpdateAccommodation_typeParams{
		ID:   accommodation_type1.ID,
		Type: util.RandomString(9),
	}
	accommodation_type2, err := testQueries.UpdateAccommodation_type(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accommodation_type2)

	require.Equal(t, accommodation_type1.ID, accommodation_type2.ID)
	require.Equal(t, arg.Type, accommodation_type2.Type)

}
