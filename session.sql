-- Drop all tables
DO $$ DECLARE
  r RECORD;
BEGIN
  FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
    EXECUTE 'DROP TABLE ' || quote_ident(r.tablename) || ' CASCADE';
  END LOOP;
END $$;

CREATE TABLE sessions (
    hash TEXT NOT NULL UNIQUE,
    id INT NOT NULL,
    data JSONB,
    expiry TIMESTAMPTZ NOT NULL DEFAULT (CURRENT_DATE + INTERVAL '1 month')
);

-- run a cron script daily to clear all cookies older than a month

-- be very sure to 
