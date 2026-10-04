CREATE TABLE public.audit_logs (
    id               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    occurred_at      TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    txid             BIGINT      NOT NULL DEFAULT txid_current(),
    db_user          TEXT        NOT NULL DEFAULT session_user,
    application_name TEXT,
    schema_name      TEXT NOT NULL,
    table_name       TEXT NOT NULL,
    operation        TEXT NOT NULL CHECK (operation IN ('INSERT', 'UPDATE', 'DELETE')),
    old_data         JSONB,
    new_data         JSONB
);

CREATE OR REPLACE FUNCTION public.audit_agent_write()
RETURNS trigger
LANGUAGE plpgsql
SECURITY DEFINER
SET search_path = pg_catalog, pg_temp
AS $$
DECLARE
    v_old jsonb;
    v_new jsonb;
BEGIN
    -- Gate on role here if you only want to audit a specific service:
    -- IF session_user <> 'agent_service' THEN RETURN NULL; END IF;

    IF TG_OP <> 'INSERT' THEN v_old := to_jsonb(OLD); END IF;
    IF TG_OP <> 'DELETE' THEN v_new := to_jsonb(NEW); END IF;

    IF TG_OP = 'UPDATE' AND v_old = v_new THEN
        RETURN NULL;
    END IF;

    INSERT INTO public.audit_logs
        (application_name, schema_name, table_name, operation, old_data, new_data)
    VALUES
        (NULLIF(current_setting('application_name', true), ''),
         TG_TABLE_SCHEMA, TG_TABLE_NAME, TG_OP, v_old, v_new);

    RETURN NULL;
END;
$$;


CREATE OR REPLACE TRIGGER "audit_agents"
    AFTER INSERT OR UPDATE OR DELETE ON "agents"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_subnetworks"
    AFTER INSERT OR UPDATE OR DELETE ON "subnetworks"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_machines"
    AFTER INSERT OR UPDATE OR DELETE ON "machines"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_cpus"
    AFTER INSERT OR UPDATE OR DELETE ON "cpus"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_gpus"
    AFTER INSERT OR UPDATE OR DELETE ON "gpus"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_disks"
    AFTER INSERT OR UPDATE OR DELETE ON "disks"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_network_interfaces"
    AFTER INSERT OR UPDATE OR DELETE ON "network_interfaces"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_network_interface_subnets"
    AFTER INSERT OR UPDATE OR DELETE ON "network_interface_subnets"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_packages"
    AFTER INSERT OR UPDATE OR DELETE ON "packages"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_applications"
    AFTER INSERT OR UPDATE OR DELETE ON "applications"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_application_endpoints"
    AFTER INSERT OR UPDATE OR DELETE ON "application_endpoints"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_users"
    AFTER INSERT OR UPDATE OR DELETE ON "users"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_user_applications"
    AFTER INSERT OR UPDATE OR DELETE ON "user_applications"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_flows"
    AFTER INSERT OR UPDATE OR DELETE ON "flows"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();
CREATE OR REPLACE TRIGGER "audit_endpoint_policies"
    AFTER INSERT OR UPDATE OR DELETE ON "endpoint_policies"
    FOR EACH ROW EXECUTE FUNCTION public.audit_agent_write();