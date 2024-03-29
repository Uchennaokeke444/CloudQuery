-- uses view which uses aws_security_group_ingress_rules.sql query
insert into aws_policy_results
WITH IndividualRuleStatus AS (
  SELECT
      account_id,
    id as resource_id,
     case when
        (ip = '0.0.0.0/0' or ip = '::/0')
        and ((from_port is null and to_port is null) -- all ports
        or 20 between from_port and to_port
        or 21 between from_port and to_port
        or 22 between from_port and to_port
        or 23 between from_port and to_port
        or 25 between from_port and to_port
        or 110 between from_port and to_port
        or 135 between from_port and to_port
        or 143 between from_port and to_port
        or 445 between from_port and to_port
        or 1433 between from_port and to_port
        or 1434 between from_port and to_port
        or 3000 between from_port and to_port
        or 3306 between from_port and to_port
        or 3389 between from_port and to_port
        or 4333 between from_port and to_port
        or 5000 between from_port and to_port
        or 5432 between from_port and to_port
        or 5500 between from_port and to_port
        or 5601 between from_port and to_port
        or 8080 between from_port and to_port
        or 8088 between from_port and to_port
        or 8888 between from_port and to_port
        or 9200 between from_port and to_port
        or 9300 between from_port and to_port)
        then 'fail'
        else 'pass'
    end as status
  FROM view_aws_security_group_ingress_rules
)

SELECT
	:'execution_time' as execution_time,
	:'framework' as framework,
	:'check_id' as check_id,
	'Aggregates rules of security groups with ports and IPs including ipv6' as title,	
    account_id,
    resource_id,
    CASE
      WHEN SUM(CASE WHEN status = 'fail' THEN 1 ELSE 0 END) > 0 THEN 'fail'
      ELSE 'pass'
    END as status
  FROM IndividualRuleStatus
  GROUP BY account_id, resource_id