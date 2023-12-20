package fsx

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func StorageVirtualMachines() *schema.Table {
	tableName := "aws_fsx_storage_virtual_machines"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachine.html`,
		Resolver:    fetchFsxStorageVirtualMachines,
		Transform:   transformers.TransformWithStruct(&types.StorageVirtualMachine{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "fsx"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ResourceARN"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchFsxStorageVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceFsx).Fsx
	input := fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeStorageVirtualMachinesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx, func(options *fsx.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- result.StorageVirtualMachines
	}
	return nil
}
