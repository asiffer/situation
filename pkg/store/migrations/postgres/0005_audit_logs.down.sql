DROP TRIGGER IF EXISTS "audit_endpoint_policies" ON "endpoint_policies";
DROP TRIGGER IF EXISTS "audit_flows" ON "flows";
DROP TRIGGER IF EXISTS "audit_user_applications" ON "user_applications";
DROP TRIGGER IF EXISTS "audit_users" ON "users";
DROP TRIGGER IF EXISTS "audit_application_endpoints" ON "application_endpoints";
DROP TRIGGER IF EXISTS "audit_applications" ON "applications";
DROP TRIGGER IF EXISTS "audit_packages" ON "packages";
DROP TRIGGER IF EXISTS "audit_network_interface_subnets" ON "network_interface_subnets";
DROP TRIGGER IF EXISTS "audit_network_interfaces" ON "network_interfaces";
DROP TRIGGER IF EXISTS "audit_disks" ON "disks";
DROP TRIGGER IF EXISTS "audit_gpus" ON "gpus";
DROP TRIGGER IF EXISTS "audit_cpus" ON "cpus";
DROP TRIGGER IF EXISTS "audit_machines" ON "machines";
DROP TRIGGER IF EXISTS "audit_subnetworks" ON "subnetworks";
DROP TRIGGER IF EXISTS "audit_agents" ON "agents";

DROP FUNCTION IF EXISTS public.audit_agent_write();
DROP TABLE IF EXISTS "audit_logs";