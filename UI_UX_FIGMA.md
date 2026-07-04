# LifeMate AI UI/UX and Figma Screen Specification

## Design Principle

LifeMate AI should feel calm, trustworthy, accessible, and human. It is not a futuristic gadget app. It is a life companion for real users who may be stressed, elderly, low-literacy, busy, or using a local language.

## Visual Direction

- Tone: warm, clear, respectful
- Layout: task-first, voice-first, low clutter
- Accessibility: large touch targets, high contrast, readable text
- Colors: balanced palette, not dominated by one hue
- Typography: simple sans-serif, strong hierarchy
- Motion: gentle confirmations, no distracting animation

## Suggested Palette

- Trust Blue: `#2563EB`
- Care Green: `#16A34A`
- Alert Red: `#DC2626`
- Warm Yellow: `#F59E0B`
- Ink: `#111827`
- Soft Background: `#F8FAFC`
- Surface: `#FFFFFF`
- Border: `#E5E7EB`

## Components

- Voice orb button with microphone icon
- Bottom navigation
- Language switcher
- Reminder card
- SOS button
- Family alert status
- Daily plan timeline
- Budget bar
- Document summary panel
- Translation split panel
- Consent toggle
- Offline mode banner
- Accessibility mode switch

## Main Navigation

Tabs:

1. Home
2. Reminders
3. Assistant
4. Documents
5. Family

Secondary:

- Budget
- Career
- Wellness
- Government schemes
- Settings

## Screen 1: Onboarding - Welcome

Purpose:

Introduce LifeMate AI and let users choose language and access mode.

Elements:

- LifeMate AI logo
- Language picker
- Large "Speak to start" button
- Text option: "Type instead"
- Accessibility options: larger text, voice-only mode

Primary action:

Continue

Figma notes:

- Mobile frame 390 x 844
- Large centered microphone action
- Keep copy minimal

## Screen 2: Profile Setup

Purpose:

Collect only essential personalization details.

Fields:

- Name
- Preferred language
- User type: Student, Elderly, Professional, Rural user, Caregiver
- Country/region
- Emergency contact prompt

AI behavior:

Ask one question at a time in voice mode.

## Screen 3: Home Dashboard

Purpose:

Show the user's day at a glance.

Sections:

- Greeting in preferred language
- Voice assistant button
- Today's reminders
- Daily plan
- SOS shortcut
- Offline status
- Suggested task: "Check eligible schemes", "Summarize document", or "Plan budget"

Key UI:

- Persistent microphone button
- Small offline badge
- Clear next action

## Screen 4: Voice Assistant

Purpose:

Main conversational interface.

Elements:

- Live transcript
- AI response card
- Voice waveform
- Language indicator
- Suggested quick actions
- Tool confirmation sheet

Example quick actions:

- Set medicine reminder
- Translate this
- Summarize document
- Find government schemes
- Plan my day

## Screen 5: Medicine Reminder

Purpose:

Create and manage medicine reminders.

Elements:

- Medicine name
- Time and schedule
- Instructions
- Caregiver alert toggle
- Dose taken / skipped buttons
- Missed reminder history

Safety:

- The app must not recommend dosage.
- If dosage is missing, ask user to follow prescription.

## Screen 6: Emergency SOS

Purpose:

Help user contact trusted people quickly.

Elements:

- Large SOS button
- Countdown before sending
- Cancel button
- Share location toggle
- Emergency contacts list
- Medical snapshot preview

Interaction:

Hold for 3 seconds or tap plus confirmation.

## Screen 7: Government Scheme Guidance

Purpose:

Guide user through eligibility and documents.

Elements:

- Search or voice query
- Region selector
- Eligibility questions
- Scheme cards
- Required documents checklist
- Official link or verification note
- Save offline button

## Screen 8: Financial Budgeting

Purpose:

Create a simple budget without intimidating finance language.

Elements:

- Income input
- Expense categories
- Savings goal
- Budget health indicator
- AI suggestions
- Monthly summary

Safety:

- No investment guarantees.
- Focus on budgeting education.

## Screen 9: AI Document Summarizer

Purpose:

Upload or scan documents and get simple explanations.

Elements:

- Upload/scan button
- Document preview
- Summary
- Key points
- Action items
- Deadlines
- Translate button
- Save/share options

## Screen 10: Translator

Purpose:

Translate text or speech.

Elements:

- Source language
- Target language
- Input text/voice
- Output translation
- Speaker playback
- Save phrase

## Screen 11: Career Guidance

Purpose:

Help students and workers plan career steps.

Elements:

- Skills profile
- Interest selector
- Career suggestions
- Learning roadmap
- Resume tips
- Interview practice

## Screen 12: Mental Wellness Companion

Purpose:

Offer supportive check-ins and safe escalation.

Elements:

- Mood scale
- "Talk to LifeMate" button
- Breathing exercise
- Journal prompt
- Trusted contact shortcut

Safety:

- Crisis language triggers emergency support flow.
- Family sharing requires explicit consent.

## Screen 13: Family Dashboard

Purpose:

Support loved ones without violating privacy.

Elements:

- Family member cards
- Reminder status
- Missed medicine alerts
- SOS alerts
- Consent settings
- Check-in message button

Privacy:

- Show "medicine reminder missed" instead of private conversation content.
- User controls what family sees.

## Screen 14: Offline Mode

Purpose:

Make the app useful without stable internet.

Elements:

- Offline banner
- Available offline features
- Saved guidance
- Local reminders
- Pending sync queue
- Last updated timestamp

## Figma Frames to Create

Recommended frames:

1. Mobile onboarding
2. Mobile home
3. Mobile voice assistant
4. Mobile reminder creation
5. Mobile SOS
6. Mobile government scheme
7. Mobile document summary
8. Mobile budget
9. Mobile mental wellness
10. Mobile offline mode
11. Web family dashboard
12. Web institution dashboard

## Prototype Flow for Hackathon

1. Start on Home.
2. Tap microphone.
3. Ask for medicine reminder in a local language.
4. Confirm reminder.
5. Ask for government scheme guidance.
6. Upload document.
7. Show summary and translation.
8. Trigger SOS demo.
9. Open Family Dashboard.
10. Toggle offline mode.

## Accessibility Checklist

- Minimum 44 x 44 touch targets
- High contrast text
- Voice-first navigation
- Text alternative for voice
- Simple words
- One task per screen
- Confirmation for risky actions
- Local language support
- Large text mode
- Low bandwidth mode

