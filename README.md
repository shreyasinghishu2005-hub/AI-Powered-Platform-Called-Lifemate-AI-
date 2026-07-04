# LifeMate AI

LifeMate AI is a multilingual personal AI companion for students, elderly people, working professionals, and rural users. It combines voice assistance, health and medicine reminders, emergency SOS, government scheme guidance, budgeting, translation, document summarization, career guidance, mental wellness support, offline-first access, daily planning, and a family dashboard.

This repository is a complete international AI hackathon package. It includes business strategy, product design, technical architecture, API plans, database schema, Figma-ready screen specs, pitch deck content, demo script, deployment architecture, roadmap, revenue model, and judge Q&A.

## Why It Matters

Many everyday digital services are still hard to use for people who face language barriers, low literacy, intermittent internet, disability, age-related challenges, or fragmented public information. LifeMate AI turns critical life tasks into simple conversations in the user's local language.

## Core Features

- Voice-based AI assistant
- Health and medicine reminders
- Government scheme guidance
- Emergency SOS and family alerts
- Financial budgeting
- AI document summarizer
- AI translator
- Career guidance
- Mental wellness companion
- Local language support
- Offline mode
- Personalized daily planning
- Family dashboard

## AI API Strategy

LifeMate AI can use OpenAI or Gemini behind a provider adapter:

- OpenAI Responses API for multimodal reasoning, tool calling, document understanding, planning, and summaries.
- OpenAI Realtime API for low-latency voice experiences.
- Gemini API for text, multimodal generation, translation, and regional provider redundancy.
- On-device speech and translation fallbacks for offline-first mode where available.

Recommended implementation keeps model IDs in environment variables:

```bash
AI_PROVIDER=openai
OPENAI_API_KEY=...
OPENAI_TEXT_MODEL=...
OPENAI_REALTIME_MODEL=...
GEMINI_API_KEY=...
GEMINI_TEXT_MODEL=...
```

## Project Files

- [MASTER_BLUEPRINT.md](MASTER_BLUEPRINT.md) - complete hackathon narrative and product blueprint
- [BUSINESS_PLAN.md](BUSINESS_PLAN.md) - market, competitors, go-to-market, revenue, impact
- [TECHNICAL_ARCHITECTURE.md](TECHNICAL_ARCHITECTURE.md) - system architecture and implementation plan
- [API_INTEGRATION.md](API_INTEGRATION.md) - OpenAI/Gemini integration, prompts, safety, tool calling
- [DATABASE_SCHEMA.md](DATABASE_SCHEMA.md) - PostgreSQL schema, offline sync, privacy notes
- [UI_UX_FIGMA.md](UI_UX_FIGMA.md) - Figma-ready screens, flows, components, accessibility
- [PITCH_DECK.md](PITCH_DECK.md) - 12-slide pitch deck content
- [PRESENTATION_SCRIPT.md](PRESENTATION_SCRIPT.md) - spoken pitch script
- [DEPLOYMENT_ARCHITECTURE.md](DEPLOYMENT_ARCHITECTURE.md) - cloud, edge, offline, observability
- [REVENUE_MODEL.md](REVENUE_MODEL.md) - pricing and unit economics
- [FUTURE_ROADMAP.md](FUTURE_ROADMAP.md) - 24-month roadmap
- [DEMO_WALKTHROUGH.md](DEMO_WALKTHROUGH.md) - 3-minute demo flow
- [JUDGE_QA.md](JUDGE_QA.md) - likely judge questions with strong answers
- [GITHUB_PROJECT_GUIDE.md](GITHUB_PROJECT_GUIDE.md) - repo structure and build guidance

## Suggested Prototype Stack

- Mobile app: React Native or Flutter
- Web dashboard: Next.js
- Backend: Node.js/NestJS or Python/FastAPI
- Database: PostgreSQL plus pgvector
- Cache and jobs: Redis, BullMQ/Celery
- Offline storage: SQLite on device
- Object storage: S3-compatible bucket
- Notifications: Firebase Cloud Messaging, SMS/WhatsApp provider
- AI orchestration: provider adapter for OpenAI and Gemini

## Render Deployment

This repository now includes a small dependency-free Go web server so Render's Go build command works:

```bash
go build -tags netgo -ldflags "-s -w" -o app
./app
```

The app reads the `PORT` environment variable required by Render, serves the LifeMate AI landing page at `/`, exposes `/health`, returns feature data at `/api/features`, and serves the hackathon documents from `/docs/{file}`.

## Winning Hackathon Angle

LifeMate AI is not another chatbot. It is a trusted life operations layer for people who are often excluded by language, literacy, connectivity, and complexity. The strongest demo shows one user speaking naturally in a local language and LifeMate handling a complete life task: understanding a government document, setting medicine reminders, creating a budget, alerting family, and working offline.
