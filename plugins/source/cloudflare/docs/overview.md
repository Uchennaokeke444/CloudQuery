---
name: Cloudflare
stage: GA
title: Cloudflare Source Plugin
description: CloudQuery Cloudflare Plugin documentation
---
# Cloudflare Source Plugin

:badge

The CloudQuery Cloudflare plugin pulls configuration out of Cloudflare resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Query Examples

### Find all zones with `dev_mode` enabled

```sql copy
SELECT id, account_id, host_name, name, original_ns FROM cloudflare_zones WHERE dev_mode = true;
```

### Find all DNS records

```sql copy
SELECT id, account_id, zone_id, name, type FROM cloudflare_dns_records;
```
