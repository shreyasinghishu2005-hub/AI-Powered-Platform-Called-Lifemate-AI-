# LifeMate AI API Integration

## Provider Strategy

LifeMate AI should support both OpenAI and Gemini through a provider adapter. This avoids vendor lock-in, improves regional availability, and lets the platform route each task to the best model for cost, latency, and quality.

Environment variables:

```bash
AI_PROVIDER=openai
OPENAI_API_KEY=...
OPENAI_TEXT_MODEL=...
OPENAI_REALTIME_MODEL=...
GEMINI_API_KEY=...
GEMINI_TEXT_MODEL=...
EMBEDDING_MODEL=...
```

## OpenAI Integration

Use cases:

- General reasoning
- Tool calling
- Document summarization
- Planning
- Multilingual conversation
- Voice assistant through Realtime API

### Text and Tool Calling Flow

```ts
import OpenAI from "openai";

const client = new OpenAI({ apiKey: process.env.OPENAI_API_KEY });

export async function runLifeMateAgent({
  userMessage,
  profile,
  language,
  retrievedContext,
  tools
}) {
  const response = await client.responses.create({
    model: process.env.OPENAI_TEXT_MODEL,
    input: [
      {
        role: "system",
        content: buildSystemPrompt({ language, profile })
      },
      {
        role: "user",
        content: userMessage
      },
      {
        role: "developer",
        content: `Use this verified context when relevant:\n${retrievedContext}`
      }
    ],
    tools
  });

  return response;
}
```

### Realtime Voice Flow

Use the Realtime API for low-latency speech-to-speech sessions.

Client flow:

1. Mobile app asks backend for an ephemeral session token.
2. Backend creates a Realtime session with the OpenAI API.
3. Mobile app streams microphone audio to the Realtime connection.
4. Assistant streams voice audio back.
5. Tool calls are routed through backend services.

Backend session endpoint:

```ts
app.post("/api/voice/session", async (req, res) => {
  const session = await fetch("https://api.openai.com/v1/realtime/sessions", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${process.env.OPENAI_API_KEY}`,
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      model: process.env.OPENAI_REALTIME_MODEL,
      voice: "alloy",
      instructions: "You are LifeMate AI, a safe multilingual personal companion."
    })
  });

  res.json(await session.json());
});
```

## Gemini Integration

Use cases:

- Multilingual text generation
- Translation
- Document/image understanding where suitable
- Provider redundancy
- Region-specific deployments

```ts
import { GoogleGenAI } from "@google/genai";

const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

export async function geminiGenerate(prompt: string) {
  const response = await ai.models.generateContent({
    model: process.env.GEMINI_TEXT_MODEL,
    contents: prompt
  });

  return response.text;
}
```

## Tool Calling Design

LifeMate AI should call backend tools for actions that change state.

Available tools:

```json
[
  {
    "name": "create_medicine_reminder",
    "description": "Create a medicine reminder after user confirmation.",
    "parameters": {
      "type": "object",
      "properties": {
        "medicine_name": { "type": "string" },
        "dosage": { "type": "string" },
        "schedule": { "type": "string" },
        "start_date": { "type": "string" },
        "caregiver_alert": { "type": "boolean" }
      },
      "required": ["medicine_name", "schedule"]
    }
  },
  {
    "name": "send_sos_alert",
    "description": "Send SOS alert with user location after confirmation.",
    "parameters": {
      "type": "object",
      "properties": {
        "location": { "type": "string" },
        "message": { "type": "string" },
        "contacts": {
          "type": "array",
          "items": { "type": "string" }
        }
      },
      "required": ["message", "contacts"]
    }
  },
  {
    "name": "lookup_government_scheme",
    "description": "Search verified scheme knowledge base.",
    "parameters": {
      "type": "object",
      "properties": {
        "country": { "type": "string" },
        "region": { "type": "string" },
        "need": { "type": "string" },
        "user_profile": { "type": "object" }
      },
      "required": ["country", "need"]
    }
  },
  {
    "name": "create_budget_plan",
    "description": "Create a personal monthly budget.",
    "parameters": {
      "type": "object",
      "properties": {
        "monthly_income": { "type": "number" },
        "expenses": { "type": "array" },
        "goal": { "type": "string" }
      },
      "required": ["monthly_income"]
    }
  },
  {
    "name": "create_daily_plan",
    "description": "Create a personalized daily plan.",
    "parameters": {
      "type": "object",
      "properties": {
        "date": { "type": "string" },
        "tasks": { "type": "array" },
        "energy_level": { "type": "string" },
        "constraints": { "type": "array" }
      },
      "required": ["date"]
    }
  }
]
```

## Prompt Architecture

### System Prompt

```text
You are LifeMate AI, a multilingual personal AI companion.
Your job is to help users with daily life tasks in a simple, respectful, and safe way.

Principles:
- Reply in the user's preferred language unless they ask otherwise.
- Use plain language suitable for low-literacy users.
- Ask one question at a time when collecting information.
- Do not claim to be a doctor, lawyer, financial advisor, or emergency service.
- For health symptoms, encourage professional medical help.
- For emergencies, help contact trusted people or local emergency services.
- For government schemes, rely on verified context and mention uncertainty.
- For mental wellness crisis signals, encourage immediate human support and crisis resources.
- Confirm before sending messages, creating reminders, or sharing location.
- Respect privacy and family consent.
```

### Government Scheme Prompt

```text
Answer only from the verified scheme context below.
If the context does not contain the answer, say you need updated local information.
Explain:
1. What the scheme is
2. Who may qualify
3. Documents needed
4. Next steps
5. Where the user should verify or apply

Use short sentences and the user's language.
```

### Medicine Reminder Prompt

```text
Extract reminder details from the user's request.
Do not provide dosage advice.
If dosage or schedule is unclear, ask the user to confirm.
Before creating a reminder, repeat the schedule and ask for confirmation.
```

### Mental Wellness Prompt

```text
Be warm, calm, and supportive.
Do not diagnose.
If the user mentions self-harm, suicide, abuse, immediate danger, or inability to stay safe:
- Ask if they are in immediate danger.
- Encourage contacting local emergency services or a trusted person now.
- Offer to alert an emergency contact if consent is available.
- Keep the response short and direct.
```

## RAG Pipeline

### Ingestion

```ts
async function ingestKnowledgeDocument(document) {
  const text = await extractText(document);
  const chunks = chunkText(text, { maxTokens: 700, overlap: 120 });

  for (const chunk of chunks) {
    const embedding = await createEmbedding(chunk.text);
    await db.knowledge_chunks.insert({
      source_id: document.sourceId,
      title: document.title,
      language: document.language,
      country: document.country,
      region: document.region,
      content: chunk.text,
      embedding
    });
  }
}
```

### Retrieval

```ts
async function retrieveContext(query, filters) {
  const embedding = await createEmbedding(query);
  return db.query(`
    SELECT title, content, source_url
    FROM knowledge_chunks
    WHERE country = $1
    ORDER BY embedding <=> $2
    LIMIT 6
  `, [filters.country, embedding]);
}
```

## Safety Classifier

Every user input should be classified:

```json
{
  "domain": "health | finance | legal | emergency | wellness | general",
  "risk_level": "low | medium | high | critical",
  "requires_human_escalation": false,
  "requires_confirmation": true,
  "allowed_tools": ["create_medicine_reminder"]
}
```

High-risk handling:

- Health diagnosis request: provide general education and suggest clinician.
- Financial investment request: provide budgeting education, no guaranteed returns.
- Legal/government certainty: provide verified guidance and official verification step.
- Emergency: prioritize human contact and local emergency services.
- Mental health crisis: crisis escalation and trusted contact.

## Example End-to-End Request

User:

```text
Mujhe maa ke liye diabetes medicine ka reminder lagana hai, subah nashta ke baad.
```

Pipeline:

1. Detect language: Hindi.
2. Intent: medicine reminder.
3. Extract entities: user wants reminder for mother, diabetes medicine, after breakfast.
4. Missing medicine name: ask follow-up.
5. Confirm schedule.
6. Create reminder.
7. Ask whether family alert should be enabled.

Response:

```text
Theek hai. Kripya medicine ka naam bata dijiye. Main nashta ke baad roz reminder laga dunga.
```

## API References

- OpenAI Responses API: https://platform.openai.com/docs/api-reference/responses
- OpenAI Realtime API: https://platform.openai.com/docs/guides/realtime
- Gemini API text generation: https://ai.google.dev/gemini-api/docs/text-generation

