---
title: Audit logs
summary: Track PostgreSQL changes made by Situation agents
---

Situation can keep a row-level audit log for writes made to its PostgreSQL database. The log records the database session's `application_name` together with the operation and the row values before and after the change.

!!! info "Important"
    The audit feature is **PostgreSQL-only**. It is installed by the `0005_audit_logs` migration when database migrations are run from the agent; SQLite deployments are not affected.

## Coverage

An agent migration creates the `audit_logs` table and attaches triggers to the Situation tables: `agents`, `subnetworks`, `machines`, `cpus`, `gpus`, `disks`, `network_interfaces`, `network_interface_subnets`, `packages`, `applications`, `application_endpoints`, `users`, `user_applications`, `flows`, and `endpoint_policies`.

Each changed row creates one audit entry:

- `INSERT`: `new_data` contains the inserted row; `old_data` is `NULL`.
- `UPDATE`: both `old_data` and `new_data` contain the row state.
- `DELETE`: `old_data` contains the deleted row; `new_data` is `NULL`.

An entry also contains `occurred_at`, `application_name`, `schema_name`, `table_name`, and `operation`. The audit insert runs in the same transaction as the original write: if the transaction rolls back, its audit entries roll back too. Changes to `audit_logs` itself are not audited.

## Agent identity

The PostgreSQL connection sets `application_name` to the agent identifier. The trigger stores that value in each audit entry and skips writes made by sessions whose `application_name` is empty.

!!! warning "Attribution only"
	Any client can set or spoof `application_name`. It is useful for attribution, but it is not an authentication or authorization mechanism. Use PostgreSQL roles and connection controls to enforce trusted access boundaries.

## Querying entries

For example, inspect the most recent changes attributed to one agent:

```sql
SELECT occurred_at, table_name, operation, old_data, new_data
FROM audit_logs
WHERE application_name = '<agent-id>'
ORDER BY occurred_at DESC
LIMIT 100;
```

To inspect changes to one table, filter by `table_name`:

```sql
SELECT occurred_at, application_name, operation, old_data, new_data
FROM audit_logs
WHERE table_name = 'machines'
ORDER BY occurred_at DESC
LIMIT 100;
```


## Retention

Audit entries are kept until they are explicitly deleted; the migration included in the agent does not schedule automatic cleanup. A simple setup is to choose a retention period and run a cleanup query once a day from the database host or an existing scheduler. For example, this deletes entries older than 90 days:

```sql
DELETE FROM public.audit_logs
WHERE occurred_at < now() - INTERVAL '90 days';
```

### Using `pg_cron`

[`pg_cron`](https://github.com/citusdata/pg_cron) runs scheduled SQL inside PostgreSQL. It must be installed on the database server for the matching PostgreSQL major version.

!!! warning "Server configuration required"
	Add `pg_cron` to `shared_preload_libraries` in `postgresql.conf`, preserving any libraries already listed, set the database that will hold the extension metadata, and restart PostgreSQL. Managed database services may require enabling the extension through their control panel or a parameter group.

For example:

```conf
shared_preload_libraries = 'pg_cron'
cron.database_name = 'situation'
cron.timezone = 'UTC'
```

Then, as a PostgreSQL administrator, connect to the `situation` database and install the extension:

```sql
CREATE EXTENSION IF NOT EXISTS pg_cron;
```

The following example gives a dedicated role only the permissions needed to delete expired audit entries and schedule its own job. Set up authentication for this role according to the server's `pg_hba.conf` and pg_cron connection mode.

```sql
CREATE ROLE audit_cleanup LOGIN;
GRANT USAGE ON SCHEMA public, cron TO audit_cleanup;
GRANT SELECT, DELETE ON TABLE public.audit_logs TO audit_cleanup;

SELECT cron.schedule(
		'purge-situation-audit-logs',
		'15 3 * * *',
		$$DELETE FROM public.audit_logs
			WHERE occurred_at < now() - INTERVAL '90 days'$$
);
```

The job runs daily at 03:15 in the configured `cron.timezone`. Change the interval and schedule to fit the retention requirement.

