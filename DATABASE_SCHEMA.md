# LifeMate AI Database Schema

## Database Choice

Recommended primary database:

- PostgreSQL for relational data
- pgvector for embeddings
- Redis for caching and queues
- S3-compatible object storage for documents and audio
- SQLite on device for offline mode

## Entity Overview

```text
users
profiles
family_groups
family_memberships
consents
reminders
medicine_logs
sos_contacts
sos_events
documents
document_summaries
translations
budgets
budget_items
daily_plans
plan_items
wellness_checkins
career_profiles
knowledge_sources
knowledge_chunks
ai_interactions
offline_sync_events
notifications
audit_logs
```

## PostgreSQL Schema

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  phone VARCHAR(32),
  email VARCHAR(255),
  password_hash TEXT,
  auth_provider VARCHAR(64),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE profiles (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  full_name VARCHAR(160),
  date_of_birth DATE,
  country VARCHAR(80),
  region VARCHAR(120),
  city VARCHAR(120),
  preferred_language VARCHAR(20) NOT NULL DEFAULT 'en',
  literacy_mode VARCHAR(32) DEFAULT 'standard',
  accessibility_settings JSONB DEFAULT '{}',
  medical_snapshot JSONB DEFAULT '{}',
  emergency_notes TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE family_groups (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(160) NOT NULL,
  owner_user_id UUID NOT NULL REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE family_memberships (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  family_group_id UUID NOT NULL REFERENCES family_groups(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role VARCHAR(40) NOT NULL,
  status VARCHAR(40) NOT NULL DEFAULT 'invited',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE(family_group_id, user_id)
);

CREATE TABLE consents (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  subject_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  grantee_user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  consent_type VARCHAR(80) NOT NULL,
  scope JSONB NOT NULL DEFAULT '{}',
  status VARCHAR(40) NOT NULL DEFAULT 'active',
  expires_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  revoked_at TIMESTAMPTZ
);

CREATE TABLE reminders (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  reminder_type VARCHAR(40) NOT NULL,
  title VARCHAR(200) NOT NULL,
  description TEXT,
  schedule_rule JSONB NOT NULL,
  timezone VARCHAR(80) NOT NULL,
  next_due_at TIMESTAMPTZ,
  status VARCHAR(40) NOT NULL DEFAULT 'active',
  caregiver_alert_enabled BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE medicine_logs (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  reminder_id UUID NOT NULL REFERENCES reminders(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  status VARCHAR(40) NOT NULL,
  scheduled_at TIMESTAMPTZ NOT NULL,
  confirmed_at TIMESTAMPTZ,
  notes TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sos_contacts (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(160) NOT NULL,
  phone VARCHAR(32) NOT NULL,
  relationship VARCHAR(80),
  priority INTEGER NOT NULL DEFAULT 1,
  can_receive_location BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sos_events (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  status VARCHAR(40) NOT NULL DEFAULT 'triggered',
  latitude DECIMAL(10, 7),
  longitude DECIMAL(10, 7),
  message TEXT,
  notified_contacts JSONB DEFAULT '[]',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  resolved_at TIMESTAMPTZ
);

CREATE TABLE documents (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  file_name VARCHAR(255) NOT NULL,
  file_type VARCHAR(80),
  storage_url TEXT NOT NULL,
  language VARCHAR(20),
  status VARCHAR(40) NOT NULL DEFAULT 'uploaded',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE document_summaries (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  document_id UUID NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
  summary TEXT NOT NULL,
  key_points JSONB DEFAULT '[]',
  action_items JSONB DEFAULT '[]',
  translated_language VARCHAR(20),
  model_provider VARCHAR(40),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE translations (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  source_language VARCHAR(20),
  target_language VARCHAR(20) NOT NULL,
  source_text TEXT NOT NULL,
  translated_text TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE budgets (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  month DATE NOT NULL,
  income_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
  savings_goal NUMERIC(12,2),
  notes TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE(user_id, month)
);

CREATE TABLE budget_items (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  budget_id UUID NOT NULL REFERENCES budgets(id) ON DELETE CASCADE,
  category VARCHAR(80) NOT NULL,
  planned_amount NUMERIC(12,2) NOT NULL,
  actual_amount NUMERIC(12,2) DEFAULT 0,
  item_type VARCHAR(40) NOT NULL DEFAULT 'expense'
);

CREATE TABLE daily_plans (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  plan_date DATE NOT NULL,
  summary TEXT,
  energy_level VARCHAR(40),
  created_by VARCHAR(40) DEFAULT 'ai',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE(user_id, plan_date)
);

CREATE TABLE plan_items (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  daily_plan_id UUID NOT NULL REFERENCES daily_plans(id) ON DELETE CASCADE,
  title VARCHAR(200) NOT NULL,
  item_type VARCHAR(60),
  starts_at TIMESTAMPTZ,
  ends_at TIMESTAMPTZ,
  status VARCHAR(40) DEFAULT 'pending'
);

CREATE TABLE wellness_checkins (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  mood_score INTEGER CHECK (mood_score BETWEEN 1 AND 10),
  stress_level VARCHAR(40),
  notes TEXT,
  risk_level VARCHAR(40) DEFAULT 'low',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE career_profiles (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  education_level VARCHAR(120),
  skills JSONB DEFAULT '[]',
  interests JSONB DEFAULT '[]',
  goals JSONB DEFAULT '[]',
  preferred_language VARCHAR(20),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE knowledge_sources (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title VARCHAR(255) NOT NULL,
  source_url TEXT,
  country VARCHAR(80),
  region VARCHAR(120),
  domain VARCHAR(80) NOT NULL,
  language VARCHAR(20) NOT NULL DEFAULT 'en',
  verified_by VARCHAR(160),
  last_verified_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE knowledge_chunks (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  source_id UUID NOT NULL REFERENCES knowledge_sources(id) ON DELETE CASCADE,
  title VARCHAR(255),
  content TEXT NOT NULL,
  embedding vector(1536),
  metadata JSONB DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE ai_interactions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id) ON DELETE SET NULL,
  session_id UUID,
  provider VARCHAR(40),
  model VARCHAR(120),
  intent VARCHAR(80),
  risk_level VARCHAR(40),
  input_language VARCHAR(20),
  output_language VARCHAR(20),
  prompt_tokens INTEGER,
  completion_tokens INTEGER,
  latency_ms INTEGER,
  user_feedback INTEGER,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE offline_sync_events (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  device_id VARCHAR(160) NOT NULL,
  event_type VARCHAR(80) NOT NULL,
  payload JSONB NOT NULL,
  status VARCHAR(40) NOT NULL DEFAULT 'pending',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  synced_at TIMESTAMPTZ
);

CREATE TABLE notifications (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  channel VARCHAR(40) NOT NULL,
  title VARCHAR(200),
  body TEXT,
  status VARCHAR(40) NOT NULL DEFAULT 'queued',
  sent_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE audit_logs (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  actor_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
  action VARCHAR(120) NOT NULL,
  entity_type VARCHAR(80),
  entity_id UUID,
  metadata JSONB DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

## Indexes

```sql
CREATE INDEX idx_reminders_user_due ON reminders(user_id, next_due_at);
CREATE INDEX idx_medicine_logs_user_time ON medicine_logs(user_id, scheduled_at);
CREATE INDEX idx_sos_events_user_time ON sos_events(user_id, created_at DESC);
CREATE INDEX idx_documents_user ON documents(user_id, created_at DESC);
CREATE INDEX idx_ai_interactions_user_time ON ai_interactions(user_id, created_at DESC);
CREATE INDEX idx_knowledge_sources_domain_region ON knowledge_sources(domain, country, region);
CREATE INDEX idx_knowledge_chunks_embedding ON knowledge_chunks USING ivfflat (embedding vector_cosine_ops);
```

## Offline SQLite Tables

```sql
CREATE TABLE local_reminders (
  id TEXT PRIMARY KEY,
  server_id TEXT,
  title TEXT NOT NULL,
  reminder_type TEXT NOT NULL,
  schedule_rule TEXT NOT NULL,
  next_due_at TEXT,
  status TEXT NOT NULL,
  updated_at TEXT NOT NULL,
  sync_status TEXT NOT NULL DEFAULT 'pending'
);

CREATE TABLE local_emergency_contacts (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  phone TEXT NOT NULL,
  priority INTEGER NOT NULL
);

CREATE TABLE local_saved_guidance (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  language TEXT,
  saved_at TEXT NOT NULL
);

CREATE TABLE local_offline_actions (
  id TEXT PRIMARY KEY,
  action_type TEXT NOT NULL,
  payload TEXT NOT NULL,
  created_at TEXT NOT NULL,
  sync_status TEXT NOT NULL DEFAULT 'pending'
);
```

## Privacy Notes

- Store medical data in encrypted columns where possible.
- Keep AI interaction logs minimal and avoid raw sensitive content by default.
- Do not expose private wellness notes to family without explicit consent.
- SOS events should be auditable but access-limited.
- Family dashboard should show statuses, not private conversations.

