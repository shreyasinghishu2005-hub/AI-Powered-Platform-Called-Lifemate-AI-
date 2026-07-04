# LifeMate AI Technical Architecture

## Architecture Goals

- Voice-first and multilingual
- Offline-first for essential tasks
- Safe AI orchestration
- Provider-agnostic AI layer for OpenAI and Gemini
- Privacy-preserving user and family data
- Scalable for B2C and institutional deployments
- Auditable for high-risk flows

## High-Level System

```text
Mobile App / Web Dashboard
        |
API Gateway
        |
Auth + Consent Service
        |
LifeMate Orchestrator
        |
-------------------------------------------------
| AI Provider Adapter | Tool Router | Safety Layer |
-------------------------------------------------
        |
Domain Services
-------------------------------------------------
| Reminders | SOS | Schemes | Budget | Documents |
| Translate | Career | Wellness | Planning | Family |
-------------------------------------------------
        |
Data Layer
-------------------------------------------------
| PostgreSQL | pgvector | Redis | Object Storage |
| Local SQLite on device | Analytics Warehouse |
-------------------------------------------------
```

## Client Applications

### Mobile App

Recommended:

- React Native or Flutter
- Offline SQLite
- Local notifications
- Push notifications
- Audio recording and playback
- Camera/document upload
- Device location for SOS

Key modules:

- Voice assistant
- Daily plan
- Reminders
- Documents
- Budget
- Family sharing
- Offline queue
- Emergency SOS

### Web Dashboard

Recommended:

- Next.js
- Role-based access
- Family and caregiver dashboard
- Institutional admin dashboard
- Analytics and user support tools

## Backend Services

### API Gateway

Responsibilities:

- Authentication
- Rate limiting
- Request validation
- Routing to services
- Audit logging

### LifeMate Orchestrator

The central brain of the platform.

Responsibilities:

- Detect language
- Classify intent
- Retrieve user context
- Select AI provider/model
- Call tools
- Apply safety policies
- Return multimodal response
- Store relevant memory with consent

### Domain Services

#### Reminder Service

- Medicine schedules
- Health appointments
- Daily tasks
- Recurring reminders
- Missed reminder escalation
- Offline sync

#### SOS Service

- Emergency contact management
- Location sharing
- Medical snapshot
- SOS event audit log
- SMS/push/WhatsApp escalation

#### Government Scheme Service

- Verified scheme database
- Eligibility questionnaire
- Required documents checklist
- Application guidance
- Citation-backed answers

#### Budget Service

- Income and expense categories
- Monthly budget plans
- Savings goals
- Low-risk financial education
- Spending insights

#### Document Service

- Upload PDFs/images
- OCR
- Summarization
- Translation
- Action extraction
- Deadline extraction

#### Translation Service

- Text translation
- Voice translation
- Saved phrases
- Low-bandwidth mode

#### Career Guidance Service

- Skill profile
- Career goals
- Learning roadmap
- Resume guidance
- Interview practice

#### Mental Wellness Service

- Mood check-ins
- Reflective conversation
- Breathing and grounding exercises
- Crisis signal detection
- Escalation resources

#### Family Service

- Family groups
- Consent settings
- Shared reminders
- Alert acknowledgements
- Privacy-preserving summary view

## AI Provider Adapter

Design a provider-neutral interface:

```ts
interface AIProvider {
  generateText(input: GenerateTextInput): Promise<GenerateTextResult>;
  summarizeDocument(input: DocumentInput): Promise<DocumentSummary>;
  translate(input: TranslateInput): Promise<TranslateResult>;
  classifyIntent(input: IntentInput): Promise<IntentResult>;
  streamVoiceSession(input: VoiceSessionInput): AsyncIterable<VoiceEvent>;
}
```

Provider implementations:

- `OpenAIProvider`
- `GeminiProvider`
- `OfflineProvider`

## Model Routing

| Task | Model Strategy |
| --- | --- |
| Simple intent classification | Smaller fast model |
| Government scheme guidance | Retrieval + stronger reasoning model |
| Medical reminder creation | Deterministic parser + AI confirmation |
| Document summarization | Multimodal/document model |
| Voice conversation | Realtime speech model |
| Translation | Translation model or multilingual LLM |
| Crisis detection | Classifier + conservative escalation |
| Offline reminder | Local deterministic logic |

## Retrieval-Augmented Generation

Use RAG for any factual domain that changes or affects user decisions:

- Government schemes
- Institution policies
- Health education content
- Career resources
- Financial literacy content

Pipeline:

1. Ingest verified documents.
2. Chunk and embed content.
3. Store embeddings in pgvector.
4. Retrieve top passages by semantic search.
5. Ask LLM to answer only from retrieved context.
6. Return citations and confidence.
7. Route low-confidence answers to "I need more information" or human support.

## Offline Mode

Offline mode should support:

- Medicine reminders
- Daily plan
- Saved emergency contacts
- SOS SMS fallback where network allows
- Saved government scheme checklists
- Saved translations
- Voice notes
- Local data capture for later sync

Offline architecture:

```text
Device SQLite
  - profiles_cache
  - reminders_cache
  - offline_actions
  - saved_guidance
  - emergency_contacts
  - sync_metadata
```

When connectivity returns:

1. Sync offline actions.
2. Resolve conflicts by timestamp and action type.
3. Upload logs securely.
4. Refresh verified content packs.

## Safety Layer

Safety checks run before and after AI generation.

Pre-generation:

- Identify high-risk domain
- Load user consent
- Load verified context
- Restrict tools by domain
- Require confirmation for external actions

Post-generation:

- Check for medical diagnosis claims
- Check for financial guarantees
- Check for legal certainty
- Check for crisis language
- Add appropriate disclaimers
- Escalate when needed

## Data Privacy

Principles:

- Collect minimum data required
- Encrypt sensitive data at rest and in transit
- Store family data only with explicit consent
- Use role-based access
- Provide delete/export options
- Separate raw AI logs from personal identifiers
- Mask sensitive data in analytics

## Suggested Tech Stack

| Layer | Recommended Stack |
| --- | --- |
| Mobile | React Native or Flutter |
| Web | Next.js, TypeScript |
| Backend | FastAPI or NestJS |
| Database | PostgreSQL, pgvector |
| Cache/Queue | Redis, BullMQ or Celery |
| Object Storage | S3-compatible |
| Search | pgvector first, OpenSearch later |
| Auth | OAuth, phone OTP, passkeys |
| Notifications | Firebase Cloud Messaging, SMS provider |
| AI | OpenAI + Gemini provider adapter |
| Observability | OpenTelemetry, Prometheus/Grafana, Sentry |

## MVP Build Plan

### Week 1

- Auth and profiles
- Chat interface
- OpenAI/Gemini adapter
- Basic text assistant
- Reminder CRUD

### Week 2

- Voice input/output
- Document upload and summary
- Translation
- Government scheme RAG prototype

### Week 3

- SOS mock flow
- Budget planner
- Family dashboard
- Offline reminder simulation
- Safety policies

### Week 4

- Demo polish
- Analytics
- Error handling
- Pitch deck
- Security review

## Scalability

To scale:

- Split domain services by traffic.
- Use queues for long-running document jobs.
- Cache RAG retrieval results.
- Precompute embeddings.
- Add read replicas for analytics.
- Use regional object storage.
- Add per-tenant isolation for institutions.

## API References Used

- OpenAI Responses API: https://platform.openai.com/docs/api-reference/responses
- OpenAI Realtime API: https://platform.openai.com/docs/guides/realtime
- Gemini API text generation: https://ai.google.dev/gemini-api/docs/text-generation

