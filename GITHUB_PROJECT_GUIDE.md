# LifeMate AI GitHub Project Guide

## Suggested Repository Structure

```text
lifemate-ai/
  apps/
    mobile/
    web-dashboard/
  services/
    api/
    orchestrator/
    notification-worker/
    document-worker/
  packages/
    ai-provider/
    shared-types/
    ui/
    safety/
  infrastructure/
    docker-compose.yml
    k8s/
    terraform/
  docs/
    business-plan.md
    technical-architecture.md
    api-integration.md
    database-schema.md
    pitch-deck.md
  scripts/
    ingest-knowledge.ts
    seed-demo-data.ts
  README.md
```

## Recommended README Sections

- Problem
- Solution
- Features
- Demo video
- Architecture diagram
- Tech stack
- AI APIs
- Safety and privacy
- Setup
- Environment variables
- Roadmap
- Team

## Example Setup Commands

```bash
git clone https://github.com/your-team/lifemate-ai.git
cd lifemate-ai
cp .env.example .env
npm install
npm run dev
```

## Example Environment File

```bash
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/lifemate
REDIS_URL=redis://localhost:6379
AI_PROVIDER=openai
OPENAI_API_KEY=
OPENAI_TEXT_MODEL=
OPENAI_REALTIME_MODEL=
GEMINI_API_KEY=
GEMINI_TEXT_MODEL=
S3_BUCKET=
FCM_SERVER_KEY=
SMS_PROVIDER_API_KEY=
```

## Demo Data

Seed:

- User Asha
- Family member Priya
- Medicine reminder
- Sample government scheme
- Sample document
- Budget categories
- Emergency contacts

## GitHub Issues for Hackathon Team

### Frontend

- Build onboarding
- Build home dashboard
- Build voice assistant screen
- Build reminders screen
- Build SOS screen
- Build document summary screen
- Build family dashboard
- Build offline mode banner

### Backend

- Create auth endpoints
- Create profile API
- Create reminder API
- Create SOS API
- Create document upload API
- Create AI orchestrator endpoint
- Create scheme retrieval endpoint
- Create family dashboard endpoint

### AI

- Build provider adapter
- Add OpenAI implementation
- Add Gemini implementation
- Add intent classifier
- Add safety classifier
- Add prompt templates
- Add RAG retrieval

### Data

- Create PostgreSQL migrations
- Add pgvector
- Seed scheme data
- Add AI interaction logs
- Add offline sync events

### Pitch

- Record demo
- Finalize pitch deck
- Prepare judge Q&A
- Prepare impact metrics

## Pull Request Checklist

- Feature works on mobile and desktop where relevant
- Sensitive actions require confirmation
- No private data exposed in logs
- Errors are user-friendly
- Offline behavior considered
- Accessibility checked
- Tests or demo script updated

## License Suggestion

Use MIT for hackathon code unless partner or institution rules require a different license.

