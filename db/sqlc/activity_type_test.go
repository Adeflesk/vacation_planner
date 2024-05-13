package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomActivityType(t *testing.T) ActivityType {
	atype := util.RandomString(9)
	activity_type, err := testQueries.Createactivity_type(context.Background(), atype)

	require.NoError(t, err)
	require.NotEmpty(t, activity_type)

	require.Equal(t, activity_type.Name, atype)
	return activity_type

}

func TestActivityType(t *testing.T) {
	createRandomActivityType(t)
}

func TestListAllActivityTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomActivityType(t)
	}
	arg := Listactivity_typesParams{
		Limit:  5,
		Offset: 0,
	}
	ActivityTypes, err := testQueries.Listactivity_types(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ActivityTypes, 5)

	for _, activity_type := range ActivityTypes {
		require.NotEmpty(t, activity_type)
	}

}

func TestDeleteActivityType(t *testing.T) {
	activity_type1 := createRandomActivityType(t)
	err := testQueries.Deleteactivity_type(context.Background(), activity_type1.ID)
	require.NoError(t, err)

	activity_type2, err := testQueries.Getactivity_type(context.Background(), activity_type1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, activity_type2)
}

func TestUpdateActivityType(t *testing.T) {
	activity_type1 := createRandomActivityType(t)

	arg := Updateactivity_typeParams{
		ID:   activity_type1.ID,
		Name: util.RandomString(9),
	}
	activity_type2, err := testQueries.Updateactivity_type(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, activity_type2)

	require.Equal(t, activity_type1.ID, activity_type2.ID)
	require.Equal(t, arg.Name, activity_type2.Name)

}
