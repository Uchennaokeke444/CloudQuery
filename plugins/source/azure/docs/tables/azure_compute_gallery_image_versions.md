# Table: azure_compute_gallery_image_versions

This table shows data for Azure Compute Gallery Image Versions.

https://learn.microsoft.com/en-us/rest/api/compute/gallery-image-versions/list-by-gallery-image?tabs=HTTP#galleryimageversion

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_gallery_images](azure_compute_gallery_images.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|