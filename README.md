# LinkedIn Automation Proof-of-Concept

A sophisticated Go-based automation tool built with the Rod library, focusing on stealth and anti-detection engineering.

## 🛠 Tech Stack
- **Language**: Go 1.25+
- **Browser Automation**: [Rod](https://github.com/go-rod/rod)
- **Stealth**: [Rod-Stealth](https://github.com/go-rod/stealth)
- **Database**: SQLite3 (for state persistence)

## 🕵️ Stealth & Anti-Detection Implementation
This project implements **8 unique stealth techniques** to mimic human behavior:

1. **Quadratic Bézier Curves**: Mouse movements follow natural arcs rather than robotic straight lines.
2. **Micro-Corrections**: Random "jitter" added to cursor paths to simulate human hand instability.
3. **Fingerprint Masking**: Spoofing `navigator.webdriver` and browser hardware signatures via `rod-stealth`.
4. **Natural Typing Rhythm**: Variable delays (70ms-250ms) between keystrokes with extended pauses for spaces.
5. **Human-like Scrolling**: Non-linear scrolling speeds with occasional "read-back" upward movements.
6. **Viewport Randomization**: Randomizes screen dimensions on launch to avoid consistent bot footprints.
7. **Session Persistence**: Utilizes cookie injection to maintain logged-in states and avoid repetitive login flags.
8. **Activity Scheduling**: Randomized "think-time" delays between high-level navigation actions.

## 🚀 Getting Started
1. Rename `.env.example` to `.env` and add your LinkedIn credentials.
2. Run `go mod tidy` to install dependencies.
3. Execute `go run cmd/main.go`.