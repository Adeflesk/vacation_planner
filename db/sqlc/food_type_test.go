package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomFoodType(t *testing.T) FoodType {
	ftype := util.RandomString(9)
	foodType, err := testQueries.Createfood_type(context.Background(), ftype)

	require.NoError(t, err)
	require.NotEmpty(t, foodType)

	require.Equal(t, foodType.Type, ftype)
	return foodType

}

func TestCreateFoodType(t *testing.T) {
	createRandomFoodType(t)
}

func TestListAllFoodTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomFoodType(t)
	}
	arg := Listfood_typesParams{
		Limit:  5,
		Offset: 0,
	}
	foodtypes, err := testQueries.Listfood_types(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, foodtypes, 5)

	for _, foofoodtype := range foodtypes {
		require.NotEmpty(t, foofoodtype)
	}

}

func TestDeleteFoodType(t *testing.T) {
	foodtype1 := createRandomFoodType(t)
	err := testQueries.Deletefood_type(context.Background(), foodtype1.ID)
	require.NoError(t, err)

	foodType2, err := testQueries.Getfood_type(context.Background(), foodtype1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, foodType2)
}

func TestUpdateFoodType(t *testing.T) {
	foodtype1 := createRandomFoodType(t)

	arg := Updatefood_typeParams{
		ID:   foodtype1.ID,
		Type: util.RandomString(9),
	}
	foodtype2, err := testQueries.Updatefood_type(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, foodtype2)

	require.Equal(t, foodtype1.ID, foodtype2.ID)
	require.Equal(t, arg.Type, foodtype2.Type)

}
