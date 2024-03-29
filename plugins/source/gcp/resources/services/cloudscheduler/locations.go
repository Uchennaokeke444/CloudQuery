package cloudscheduler

import (
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	pb "google.golang.org/api/cloudscheduler/v1"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudscheduler_locations",
		Description: `https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudscheduler.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Location{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(true),
		},
		Relations: []*schema.Table{
			Jobs(),
		},
	}
}
