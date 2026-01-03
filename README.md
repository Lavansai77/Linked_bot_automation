https://www.loom.com/share/5caf7a950de744b69d4079c1a772ea37
 
 
 
 
 # 🚀 LinkedIn Automation Proof-of-Concept (Go-Based)

A sophisticated automation engine built with **Golang** and the **Rod** library. This project focuses on high-level stealth engineering, anti-detection techniques, and persistent state management to simulate human-like networking activity.

---

## 🛠 Tech Stack
* **Language:** Go 1.25+
* **Browser Automation:** [Rod](https://github.com/go-rod/rod)
* **Stealth Logic:** Custom implementation (Bézier curves & Fingerprint masking)
* **Database:** SQLite (Pure-Go driver for cross-platform persistence)
* **Configuration:** Dotenv (.env) for secure credential management

---

## 🕵️‍♂️ Stealth & Anti-Detection Implementation
This project implements **8 unique stealth techniques** to bypass heuristic detection patterns:

1. **Quadratic Bézier Curves:** Mouse movements follow natural arcs rather than robotic straight-line paths.
2. **Micro-Corrections:** Randomized "jitter" added to cursor paths to simulate human hand instability.
3. **Fingerprint Masking:** Spoofing `navigator.webdriver` and browser hardware signatures.
4. **Natural Typing Rhythm:** Variable delays (70ms–250ms) between keystrokes with extended pauses for spaces.
5. **Human-like Scrolling:** Non-linear scrolling speeds with occasional "read-back" upward movements.
6. **Viewport Randomization:** Randomizes screen dimensions on launch to avoid consistent bot footprints.
7. **Session Persistence:** Utilizes cookie injection to maintain logged-in states and avoid repetitive login flags.
8. **Activity Scheduling:** Randomized “think-time” delays between high-level navigation actions.

---

## 📦 Architecture Overview
The project follows a modular Go structure for scalability and maintainability:

* `cmd/main.go`: The application entry point and execution orchestrator.
* `internal/stealth/`: Core engine for human-mimicry and browser fingerprint masking.
* `internal/storage/`: SQLite implementation for tracking interaction history (prevents duplicate messaging).
* `internal/actions/`: High-level automation logic (Searching, Scraping, Connecting).
* `internal/auth/`: Secure login handling and session management.

---

## 🚀 Getting Started

### 1. Prerequisites
* [Go](https://go.dev/dl/) installed on your machine.
* A valid LinkedIn account.

### 2. Setup
1. Clone the repository:
   ```bash
   git clone [https://github.com/YOUR_USERNAME/Linked-in-bot.git](https://github.com/YOUR_USERNAME/Linked-in-bot.git)
   cd Linked-in-bot
