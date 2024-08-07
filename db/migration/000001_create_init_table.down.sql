DROP INDEX IF EXISTS idx_ticket_reserve_logs_event_id;
DROP INDEX IF EXISTS idx_ticket_reserve_logs_user_id;
DROP INDEX IF EXISTS idx_ticket_reserve_logs_deleted_at;

DROP TABLE IF EXISTS ticket_reserve_logs;

DROP INDEX IF EXISTS idx_events_name;
DROP INDEX IF EXISTS idx_events_deleted_at;

DROP TABLE IF EXISTS events;