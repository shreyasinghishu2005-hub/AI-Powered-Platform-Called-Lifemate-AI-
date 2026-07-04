package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type feature struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type pageData struct {
	Features []feature
	Docs     []string
	Year     int
}

var features = []feature{
	{"Voice AI Assistant", "Talk naturally in local languages and get guided help for everyday life tasks."},
	{"Health and Medicine Reminders", "Create reminders, track missed doses, and alert caregivers with consent."},
	{"Government Scheme Guidance", "Explain eligibility, documents, and next steps using verified public-service content."},
	{"Emergency SOS", "Send location-aware SOS alerts to trusted family contacts during urgent moments."},
	{"Financial Budgeting", "Build simple monthly budgets and savings plans in plain language."},
	{"AI Document Summarizer", "Turn complex PDFs, notices, and forms into short action checklists."},
	{"AI Translator", "Translate text and voice across local languages for daily communication."},
	{"Career Guidance", "Help students and workers discover paths, skills, resumes, and interview practice."},
	{"Mental Wellness Companion", "Offer calm check-ins, grounding exercises, and safe escalation when needed."},
	{"Offline Mode", "Keep reminders, emergency contacts, saved guidance, and daily plans available offline."},
	{"Personalized Daily Planning", "Convert goals, appointments, and reminders into a clear daily plan."},
	{"Family Dashboard", "Let caregivers support loved ones through consent-based alerts and status updates."},
}

var docs = []string{
	"README.md",
	"MASTER_BLUEPRINT.md",
	"BUSINESS_PLAN.md",
	"TECHNICAL_ARCHITECTURE.md",
	"API_INTEGRATION.md",
	"DATABASE_SCHEMA.md",
	"UI_UX_FIGMA.md",
	"PITCH_DECK.md",
	"PRESENTATION_SCRIPT.md",
	"DEPLOYMENT_ARCHITECTURE.md",
	"REVENUE_MODEL.md",
	"FUTURE_ROADMAP.md",
	"DEMO_WALKTHROUGH.md",
	"JUDGE_QA.md",
	"GITHUB_PROJECT_GUIDE.md",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/features", featuresHandler)
	mux.HandleFunc("/docs/", docsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           securityHeaders(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("LifeMate AI server listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := pageData{
		Features: features,
		Docs:     docs,
		Year:     time.Now().Year(),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := homeTemplate.Execute(w, data); err != nil {
		http.Error(w, "Could not render LifeMate AI page", http.StatusInternalServerError)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "LifeMate AI",
	})
}

func featuresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(features)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/docs/")
	if name == "" || strings.Contains(name, "/") || strings.Contains(name, "\\") {
		http.NotFound(w, r)
		return
	}

	allowed := false
	for _, doc := range docs {
		if name == doc {
			allowed = true
			break
		}
	}
	if !allowed {
		http.NotFound(w, r)
		return
	}

	content, err := os.ReadFile(filepath.Clean(name))
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not read %s", name), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(content)
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("X-Frame-Options", "DENY")
		next.ServeHTTP(w, r)
	})
}

var homeTemplate = template.Must(template.New("home").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>LifeMate AI | Multilingual Personal AI Companion</title>
  <meta name="description" content="LifeMate AI is a multilingual personal AI companion for health reminders, government schemes, SOS, budgeting, documents, translation, career guidance, wellness, offline mode, and family support.">
  <style>
    :root {
      color-scheme: light;
      --ink: #111827;
      --muted: #4b5563;
      --line: #d9e2ec;
      --surface: #ffffff;
      --soft: #f7fafc;
      --blue: #2563eb;
      --green: #16a34a;
      --amber: #d97706;
      --red: #dc2626;
      --teal: #0f766e;
    }
    * { box-sizing: border-box; }
    body {
      margin: 0;
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      color: var(--ink);
      background: var(--soft);
      line-height: 1.55;
    }
    a { color: inherit; }
    .hero {
      min-height: 92vh;
      display: grid;
      align-items: center;
      padding: 28px clamp(18px, 5vw, 72px);
      background:
        linear-gradient(120deg, rgba(17,24,39,.83), rgba(15,118,110,.76)),
        url("https://images.unsplash.com/photo-1576091160399-112ba8d25d1d?auto=format&fit=crop&w=1800&q=80") center/cover;
      color: white;
    }
    nav {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      display: flex;
      justify-content: space-between;
      gap: 16px;
      padding: 20px clamp(18px, 5vw, 72px);
      font-weight: 700;
    }
    nav span:last-child {
      font-weight: 500;
      opacity: .86;
    }
    .hero-content {
      width: min(860px, 100%);
      padding-top: 54px;
    }
    .eyebrow {
      display: inline-flex;
      align-items: center;
      gap: 8px;
      margin: 0 0 18px;
      font-size: .9rem;
      font-weight: 700;
      text-transform: uppercase;
      letter-spacing: 0;
    }
    h1 {
      margin: 0;
      font-size: clamp(3rem, 9vw, 7.5rem);
      line-height: .95;
      letter-spacing: 0;
    }
    .hero p {
      width: min(680px, 100%);
      margin: 24px 0 0;
      font-size: clamp(1.05rem, 2vw, 1.35rem);
      color: rgba(255,255,255,.92);
    }
    .actions {
      display: flex;
      flex-wrap: wrap;
      gap: 12px;
      margin-top: 30px;
    }
    .button {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      min-height: 46px;
      padding: 12px 18px;
      border-radius: 8px;
      background: white;
      color: var(--ink);
      text-decoration: none;
      font-weight: 800;
    }
    .button.secondary {
      background: rgba(255,255,255,.14);
      color: white;
      border: 1px solid rgba(255,255,255,.42);
    }
    main section {
      padding: 68px clamp(18px, 5vw, 72px);
    }
    .section-inner {
      max-width: 1160px;
      margin: 0 auto;
    }
    .section-title {
      max-width: 780px;
      margin-bottom: 30px;
    }
    h2 {
      margin: 0 0 10px;
      font-size: clamp(2rem, 5vw, 3.4rem);
      line-height: 1.05;
      letter-spacing: 0;
    }
    .section-title p {
      margin: 0;
      color: var(--muted);
      font-size: 1.08rem;
    }
    .grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
      gap: 14px;
    }
    .card {
      min-height: 164px;
      padding: 20px;
      border: 1px solid var(--line);
      border-radius: 8px;
      background: var(--surface);
    }
    .card strong {
      display: block;
      margin-bottom: 8px;
      font-size: 1.06rem;
    }
    .card p {
      margin: 0;
      color: var(--muted);
    }
    .band {
      background: white;
      border-block: 1px solid var(--line);
    }
    .workflow {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(190px, 1fr));
      gap: 12px;
      counter-reset: step;
    }
    .step {
      position: relative;
      padding: 18px;
      border-radius: 8px;
      background: #f8fafc;
      border: 1px solid var(--line);
    }
    .step:before {
      counter-increment: step;
      content: counter(step);
      display: inline-grid;
      place-items: center;
      width: 30px;
      height: 30px;
      margin-bottom: 14px;
      border-radius: 999px;
      background: var(--blue);
      color: white;
      font-weight: 800;
    }
    .docs {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
      gap: 10px;
      padding: 0;
      list-style: none;
    }
    .docs a {
      display: block;
      padding: 14px 16px;
      border: 1px solid var(--line);
      border-radius: 8px;
      background: white;
      text-decoration: none;
      font-weight: 700;
    }
    .metrics {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
      gap: 14px;
    }
    .metric {
      padding: 22px;
      background: var(--ink);
      color: white;
      border-radius: 8px;
    }
    .metric strong {
      display: block;
      font-size: 2rem;
      line-height: 1;
    }
    footer {
      padding: 28px clamp(18px, 5vw, 72px);
      color: var(--muted);
      background: white;
      border-top: 1px solid var(--line);
    }
    @media (max-width: 680px) {
      nav {
        position: static;
        padding: 18px;
        background: var(--ink);
      }
      nav span:last-child { display: none; }
      .hero {
        min-height: 88vh;
        padding: 24px 18px;
      }
      .hero-content { padding-top: 0; }
      .button { width: 100%; }
      main section { padding: 48px 18px; }
    }
  </style>
</head>
<body>
  <header class="hero">
    <nav>
      <span>LifeMate AI</span>
      <span>Voice-first multilingual life companion</span>
    </nav>
    <div class="hero-content">
      <p class="eyebrow">Health | Safety | Documents | Money | Family</p>
      <h1>LifeMate AI</h1>
      <p>A multilingual personal AI companion that helps students, elderly people, working professionals, and rural users manage everyday life through voice, local language support, offline mode, and trusted family assistance.</p>
      <div class="actions">
        <a class="button" href="#features">Explore Features</a>
        <a class="button secondary" href="/docs/PITCH_DECK.md">Open Pitch Deck</a>
      </div>
    </div>
  </header>

  <main>
    <section id="features">
      <div class="section-inner">
        <div class="section-title">
          <h2>Built For Real Life</h2>
          <p>LifeMate AI turns generative AI into completed daily tasks, with safety, consent, and accessibility designed from the start.</p>
        </div>
        <div class="grid">
          {{range .Features}}
          <article class="card">
            <strong>{{.Title}}</strong>
            <p>{{.Description}}</p>
          </article>
          {{end}}
        </div>
      </div>
    </section>

    <section class="band">
      <div class="section-inner">
        <div class="section-title">
          <h2>AI Architecture</h2>
          <p>OpenAI or Gemini can power the assistant through a provider adapter, while verified knowledge retrieval handles sensitive factual guidance.</p>
        </div>
        <div class="workflow">
          <div class="step">User speaks or uploads a document in their preferred language.</div>
          <div class="step">LifeMate detects intent, language, risk, and needed context.</div>
          <div class="step">The AI orchestrator retrieves verified knowledge and calls safe tools.</div>
          <div class="step">Reminders, SOS, summaries, translations, and plans sync across app and family dashboard.</div>
        </div>
      </div>
    </section>

    <section>
      <div class="section-inner">
        <div class="section-title">
          <h2>Hackathon Package</h2>
          <p>Open the complete business, technical, UI, pitch, roadmap, and judge-prep documents directly from this deployed app.</p>
        </div>
        <ul class="docs">
          {{range .Docs}}
          <li><a href="/docs/{{.}}">{{.}}</a></li>
          {{end}}
        </ul>
      </div>
    </section>

    <section class="band">
      <div class="section-inner">
        <div class="section-title">
          <h2>Impact Metrics</h2>
          <p>The product is measured by completed life tasks, not just chats.</p>
        </div>
        <div class="metrics">
          <div class="metric"><strong>20s</strong>Medicine reminder demo target</div>
          <div class="metric"><strong>12+</strong>Core life assistance features</div>
          <div class="metric"><strong>24mo</strong>Future roadmap included</div>
          <div class="metric"><strong>100%</strong>Docs included for judges</div>
        </div>
      </div>
    </section>
  </main>

  <footer>
    LifeMate AI (c) {{.Year}}. Designed for international AI hackathon deployment on Render.
  </footer>
</body>
</html>`))
