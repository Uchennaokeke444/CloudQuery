package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func blob_services() *schema.Table {
	return &schema.Table{
		Name:                 "azure_storage_blob_services",
		Resolver:             fetchBlobServices,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/storagerp/blob-services/list?tabs=HTTP#blobserviceproperties",
		Transform:            transformers.TransformWithStruct(&armstorage.BlobServiceProperties{}, transformers.WithPrimaryKeys("ID"), transformers.WithSkipFields("BlobServiceProperties")),
		Columns: schema.ColumnList{
			client.SubscriptionID,
			{
				Name:     "properties",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("BlobServiceProperties"),
			},
		},
	}
}
