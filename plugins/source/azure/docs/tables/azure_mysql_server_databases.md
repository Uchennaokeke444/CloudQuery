# Table: azure_mysql_server_databases

This table shows data for Azure MySQL Server Databases.

https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/databases/list-by-server?tabs=HTTP#database

The primary key for this table is **id**.

## Relations

This table depends on [azure_mysql_servers](azure_mysql_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|