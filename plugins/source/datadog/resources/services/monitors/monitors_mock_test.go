package monitors

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildMonitorsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockMonitorsAPIClient(ctrl)
	d := mocks.NewMockDowntimesAPIClient(ctrl)
	services := client.DatadogServices{
		MonitorsAPI:  m,
		DowntimesAPI: d,
	}

	var monitor datadogV1.Monitor
	err := faker.FakeObject(&monitor)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	monitor.Deleted.Set(&now)
	priority := int64(123)
	monitor.Priority.Set(&priority)
	monitor.AdditionalProperties = map[string]any{"key": "value"}
	arr := []string{"test"}
	monitor.RestrictedRoles.Set(&arr)

	m.EXPECT().ListMonitorsWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(monitor))

	var dt []datadogV1.Downtime
	err = faker.FakeObject(&dt)
	if err != nil {
		t.Fatal(err)
	}
	i64val := int64(123)
	i32val := int32(123)
	textVal := "test string"
	c := datadogV1.NewDowntimeChild()
	c.AdditionalProperties = map[string]any{"key": "value"}
	dt[0].ActiveChild.Set(c)
	dt[0].Canceled.Set(&i64val)
	dt[0].End.Set(&i64val)
	dt[0].MonitorId.Set(&i64val)
	dt[0].ParentId.Set(&i64val)
	r := datadogV1.NewDowntimeRecurrence()
	r.AdditionalProperties = map[string]any{"key": "value"}
	dt[0].Recurrence.Set(r)
	dt[0].UpdaterId.Set(&i32val)
	dt[0].Message.Set(&textVal)
	dt[0].AdditionalProperties = map[string]any{"key": "value"}

	d.EXPECT().ListMonitorDowntimes(gomock.Any(), gomock.Any()).Return(dt, nil, nil)

	return services
}

func TestMonitors(t *testing.T) {
	client.DatadogMockTestHelper(t, Monitors(), buildMonitorsMock, client.TestOptions{})
}
