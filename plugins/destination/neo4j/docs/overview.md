---
name: Neo4j
title: Neo4j Destination Plugin
description: CloudQuery Neo4j destination plugin documentation
---
# Neo4j Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to a Neo4j database.

Supported database (tested) versions (We use the [official Neo4j Go driver](https://github.com/neo4j/neo4j-go-driver#neo4j-and-bolt-protocol-versions)):

- Neo4j >= 4.4

As a side note graph databases can be quite useful for various networking use-cases, visualization, for read-teams, blue-teams and more.

## Configuration

### Example

:configuration

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

The Neo4j destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes). 

### Plugin Spec

This is the (nested) spec used by the Neo4j destination Plugin.

- `connection_string` (`string`) (required)

  Connection string to connect to the database. This can be a URL or a DSN, as per official [neo4j docs](https://neo4j.com/docs/browser-manual/current/operations/dbms-connection/#uri-scheme).

  - `"bolt://localhost:7687"`
  - `"neo4j://localhost:7687"`

- `username` (`string`) (required)

  Username to connect to the database.

- `password` (`string`) (required)

  Password to connect to the database.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Number of records to batch together before sending to the database.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (= 4 MiB))

  Number of bytes (as Arrow buffer size) to batch together before sending to the database.
