# LifeMate AI Presentation Script

## 3-Minute Version

Hello judges. We are building LifeMate AI, a multilingual personal AI companion for everyday life.

The problem is simple but massive: daily life is becoming digital, but millions of people are still locked out by language, literacy, age, disability, confusing documents, and unreliable internet. An elderly person may miss medicine. A rural family may not understand which government scheme they qualify for. A student may need career guidance in their own language. A working professional may be overwhelmed by documents, budgeting, and planning.

LifeMate AI brings all of this into one voice-first companion.

Users can speak naturally in their local language. LifeMate can set health and medicine reminders, summarize documents, translate content, explain government schemes, create a daily plan, help with budgeting, guide careers, support mental wellness, and send SOS alerts to trusted family members.

What makes LifeMate different is that it is designed for real-world constraints. It works with voice, supports local languages, keeps essential reminders available offline, and gives families a consent-based dashboard. It also uses verified knowledge for sensitive areas like government schemes, and it has safety rules for health, finance, mental wellness, and emergencies.

The architecture is built around an AI orchestrator. The app connects to OpenAI or Gemini through a provider adapter. For factual guidance, we use retrieval-augmented generation from verified sources. For reminders, SOS, budgeting, and planning, the AI calls backend tools after user confirmation. Data is stored in PostgreSQL, embeddings in pgvector, and essential offline data in local SQLite.

Our demo shows a user speaking in a local language to set a diabetes medicine reminder for their mother. LifeMate confirms the schedule, saves it, and enables a caregiver alert. The user then asks about a government health scheme. LifeMate asks simple eligibility questions, explains required documents, and saves the checklist offline. Finally, we show document summarization, translation, SOS, and the family dashboard.

The business model is freemium for individuals, subscription-based for families, and licensed to clinics, schools, employers, NGOs, and governments.

Our north star metric is weekly completed life tasks per active user. We do not want users to just chat. We want them to take medicine on time, understand documents, access benefits, manage money, plan their day, and get help when it matters.

LifeMate AI is not just another chatbot. It is an everyday AI companion for people who need technology to become simpler, safer, and more human.

Thank you.

## 60-Second Version

LifeMate AI is a multilingual voice-first personal AI companion that helps people manage everyday life.

It is built for students, elderly people, working professionals, and rural users who need help with health reminders, medicine reminders, government schemes, emergency SOS, budgeting, document summaries, translation, career guidance, mental wellness, daily planning, and family support.

The key insight is that the next billion AI users may not use perfect English prompts. They need voice, local language, offline reliability, and trust.

LifeMate AI uses OpenAI or Gemini through a provider adapter, verified knowledge retrieval for factual guidance, and secure backend tools for reminders, SOS, budgeting, and planning. The family dashboard is consent-based, and offline mode keeps essential reminders and saved guidance available.

Our business model includes freemium users, family subscriptions, clinics, schools, employers, NGOs, and government deployments.

LifeMate AI turns generative AI into a practical life companion that helps people complete real tasks, not just ask questions.

## Closing Line

LifeMate AI brings AI from the browser into everyday life, in the user's language, with safety, trust, and human impact at the center.

