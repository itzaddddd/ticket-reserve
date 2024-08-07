CREATE TABLE IF NOT EXISTS events (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  name TEXT NOT NULL,
  quota BIGINT NOT NULL,
  remain_quota BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_events_name ON events (name);
CREATE INDEX IF NOT EXISTS idx_events_deleted_at ON events (deleted_at);

CREATE TABLE ticket_reserve_logs (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    user_id INT NOT NULL,
    event_id INT NOT NULL,
    number_reserved BIGINT NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE INDEX idx_ticket_reserve_logs_event_id ON ticket_reserve_logs (event_id);
CREATE INDEX idx_ticket_reserve_logs_user_id ON ticket_reserve_logs (user_id);
CREATE INDEX idx_ticket_reserve_logs_deleted_at ON ticket_reserve_logs (deleted_at);
