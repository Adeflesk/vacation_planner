package db

import (
	"database/sql"
	"testing"

	"context"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomFood(t *testing.T) Food {
	area := createRandomLocation(t)
	foodtype := createRandomFoodType(t)

	arg := CreateFoodParams{
		Name:       util.RandomString(6),
		Area:       area.ID,
		FoodType:   foodtype.ID,
		Webaddress: util.RandomWebsite(),
	}

	Food, err := testQueries.CreateFood(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Food)

	require.Equal(t, arg.Name, Food.Name)
	require.Equal(t, area.ID, Food.Area)
	require.Equal(t, foodtype.ID, Food.FoodType)
	require.Equal(t, arg.Webaddress, Food.Webaddress)

	return Food
}

func TestCreateFood(t *testing.T) {
	createRandomFood(t)
}

func TestGetFood(t *testing.T) {
	Food1 := createRandomFood(t)
	Food2, err := testQueries.GetFood(context.Background(), Food1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Food2)

	require.Equal(t, Food1.ID, Food2.ID)
	require.Equal(t, Food1.Name, Food2.Name)
	require.Equal(t, Food1.Area, Food2.Area)
	require.Equal(t, Food1.FoodType, Food2.FoodType)
	require.Equal(t, Food1.Webaddress, Food2.Webaddress)
}

func TestUpdateFood(t *testing.T) {

	foodtype := createRandomFoodType(t)

	Food1 := createRandomFood(t)

	arg := UpdateFoodParams{
		ID:         Food1.ID,
		Name:       util.RandomString(6),
		Area:       Food1.Area,
		FoodType:   foodtype.ID,
		Webaddress: Food1.Webaddress,
	}
	Food2, err := testQueries.UpdateFood(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Food2)

	require.Equal(t, Food1.ID, Food2.ID)
	require.Equal(t, arg.Name, Food2.Name)
	require.Equal(t, Food1.Area, Food2.Area)
	require.Equal(t, arg.FoodType, Food2.FoodType)
	require.Equal(t, arg.Webaddress, Food2.Webaddress)

}
func TestDeleteFood(t *testing.T) {
	Food1 := createRandomFood(t)
	err := testQueries.DeleteFood(context.Background(), Food1.ID)
	require.NoError(t, err)

	Food2, err := testQueries.GetFood(context.Background(), Food1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, Food2)
}

func TestListFods(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomFood(t)
	}
	arg := ListfoodParams{
		Limit:  5,
		Offset: 0,
	}
	foods, err := testQueries.Listfood(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, foods, 5)

	for _, food := range foods {
		require.NotEmpty(t, food)
	}
}

func TestGetFoodByArea(t *testing.T) {
	Food1 := createRandomFood(t)
	area := Food1.Area

	locations, err := testQueries.GetfoodByLocation(context.Background(), area)

	require.NoError(t, err)
	require.NotEmpty(t, locations)

	for _, Food := range locations {
		require.Equal(t, Food.Area, area)
	}
}
