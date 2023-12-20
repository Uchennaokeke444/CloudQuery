package vmmigration

import (
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func TargetProjects() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vmmigration_target_projects",
		Description: `https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.targetProjects`,
		Resolver:    fetchTargetProjects,
		Multiplex:   client.ProjectMultiplexEnabledServices("vmmigration.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TargetProject{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
