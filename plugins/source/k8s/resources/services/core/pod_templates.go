package core

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PodTemplates() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_pod_templates",
		Resolver:  fetchPodTemplates,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.PodTemplate{},
			transformers.WithPrimaryKeys("UID")),
		Columns: schema.ColumnList{client.ContextColumn},
	}
}

func fetchPodTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().CoreV1().PodTemplates("")

	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
