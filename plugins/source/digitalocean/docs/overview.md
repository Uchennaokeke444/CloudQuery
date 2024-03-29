---
name: DigitalOcean
stage: GA
title: DigitalOcean Source Plugin
description: CloudQuery DigitalOcean Plugin documentation
---

# DigitalOcean Source Plugin

:badge

The CloudQuery DigitalOcean plugin pulls configuration from DigitalOcean and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Configuration

The following example config sets up the DigitalOcean plugin, and connects it to a destination:

:configuration

## DigitalOcean Spec

- `token` (`string`, optional, default: `DIGITALOCEAN_TOKEN` or `DIGITALOCEAN_ACCESS_TOKEN` env variable value)

  DigitalOcean API token.

- `spaces_regions` (`[]string`, optional, default: `["nyc3", "sfo3", "ams3", "sgp1", "fra1", "syd1"]`)

  Regions for `digitalocean_spaces_*` tables.

- `spaces_access_key` (`string`, optional, default: `SPACES_SECRET_ACCESS_KEY` env variable value)

  Secret access key for `digitalocean_spaces_*` tables.

- `spaces_access_key_id` (`string`, optional, default: `SPACES_ACCESS_KEY_ID` env variable value)

  Access key ID for `digitalocean_spaces_*` tables.

- `spaces_debug_logging` (`bool`, optional, default: `false`)

  Enables debug logging for `digitalocean_spaces_*` tables.

- `concurrency` (`int`, optional, default: `10000`)

  A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

## Query Examples

### Find public facing spaces

```sql copy
--  public facing spaces are accessible by anyone, easily query which space is public facing in your account
SELECT bucket->>'Name',location,public FROM digitalocean_spaces WHERE public = true;
```

### List Droplets with public facing ipv4 or ipv6

```sql copy
-- Find any droplets that have a public ipv6 or ipv4 IP
SELECT id, name, v4->>'ip_address' AS address_v4, v4->>'netmask' AS netmask_v4, v4->>'gateway' AS gateway_v4,
       v6->>'ip_address' AS address_v6, v6->>'netmask' AS netmask_v6, v6->>'gateway' AS gateway_v6
FROM
  (SELECT id,name,v4,NULL as v6 FROM digitalocean_droplets CROSS JOIN JSONB_ARRAY_ELEMENTS(digitalocean_droplets.networks->'v4') AS v4
  UNION
  SELECT id,name,NULL as v4,v6 FROM digitalocean_droplets CROSS JOIN JSONB_ARRAY_ELEMENTS(digitalocean_droplets.networks->'v6') AS v6) AS union_v46
WHERE v4->>'type' = 'public' OR v6->>'type' = 'public';
```

### List Apps and their alert rules

```sql copy
-- List Apps and their alert rules
SELECT apps.live_url, alerts.spec->>'rule' AS rule
FROM digitalocean_apps apps
LEFT JOIN digitalocean_apps_alerts alerts ON apps._cq_id = alerts._cq_parent_id;
```
