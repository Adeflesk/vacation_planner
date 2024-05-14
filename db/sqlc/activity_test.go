package db

import (
	"database/sql"
	"fmt"
	"testing"

	"context"

	"github.com/Adeflesk/vacation_planner/util"
	"github.com/stretchr/testify/require"
)

func createRandomActivity(t *testing.T) Activity {
	area := createRandomLocation(t)
	Activitytype := createRandomActivityType(t)

	arg := CreateActivityParams{
		ActivityName:  util.RandomString(6),
		ActivityType:  Activitytype.ID,
		Description:   util.RandomString(20),
		Webaddress:    util.RandomWebsite(),
		TimeAllocated: util.RandomInt(1, 10),
		Area:          area.ID,
	}

	Activity, err := testQueries.CreateActivity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Activity)

	require.Equal(t, arg.ActivityName, Activity.ActivityName)
	require.Equal(t, Activitytype.ID, Activity.ActivityType)
	require.Equal(t, arg.Description, Activity.Description)
	require.Equal(t, arg.Webaddress, Activity.Webaddress)
	require.Equal(t, arg.TimeAllocated, Activity.TimeAllocated)
	require.Equal(t, area.ID, Activity.Area)

	return Activity
}

func TestCreateActivity(t *testing.T) {
	createRandomActivity(t)
}

func TestGetActivity(t *testing.T) {
	Activity1 := createRandomActivity(t)
	Activity2, err := testQueries.GetActivity(context.Background(), Activity1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Activity2)

	require.Equal(t, Activity1.ID, Activity2.ID)
	require.Equal(t, Activity1.ActivityName, Activity2.ActivityName)
	require.Equal(t, Activity1.ActivityType, Activity2.ActivityType)
	require.Equal(t, Activity1.Description, Activity2.Description)
	require.Equal(t, Activity1.Webaddress, Activity2.Webaddress)
	require.Equal(t, Activity1.TimeAllocated, Activity2.TimeAllocated)
	require.Equal(t, Activity1.Area, Activity2.Area)
}

func TestUpdateActivity(t *testing.T) {

	Activitytype := createRandomActivityType(t)

	Activity1 := createRandomActivity(t)

	arg := UpdateActivityParams{
		ID:            Activity1.ID,
		ActivityName:  Activity1.ActivityName,
		ActivityType:  Activitytype.ID,
		Webaddress:    Activity1.Webaddress,
		TimeAllocated: Activity1.TimeAllocated,
		Area:          Activity1.Area,
	}
	Activity2, err := testQueries.UpdateActivity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Activity2)

	require.Equal(t, Activity1.ID, Activity2.ID)
	require.Equal(t, arg.ActivityName, Activity2.ActivityName)
	require.Equal(t, arg.ActivityType, Activity2.ActivityType)
	require.Equal(t, arg.Webaddress, Activity2.Webaddress)
	require.Equal(t, arg.TimeAllocated, Activity2.TimeAllocated)
	require.Equal(t, Activity1.Area, Activity2.Area)

}
func TestDeleteActivity(t *testing.T) {
	Activity1 := createRandomActivity(t)
	err := testQueries.DeleteActivity(context.Background(), Activity1.ID)
	require.NoError(t, err)

	Activity2, err := testQueries.GetActivity(context.Background(), Activity1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, Activity2)
}

func TestListActivitys(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomActivity(t)
	}
	arg := ListActivityParams{
		Limit:  5,
		Offset: 0,
	}
	Activitys, err := testQueries.ListActivity(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Activitys, 5)

	for _, Activity := range Activitys {
		require.NotEmpty(t, Activity)
	}
}

func TestGetActivityByArea(t *testing.T) {
	Activity1 := createRandomActivity(t)
	fmt.Println(Activity1.Area)
	area := Activity1.Area

	locations, err := testQueries.GetActivityByLocation(context.Background(), area)

	require.NoError(t, err)
	require.NotEmpty(t, locations)

	for _, Activity := range locations {
		require.Equal(t, Activity.Area, area)
	}
}
