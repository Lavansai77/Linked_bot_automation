package main

import (
	"log"
	"os"

	"github.com/Lavansai77/Linked-in-bot/internal/actions"
	"github.com/Lavansai77/Linked-in-bot/internal/auth"
	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/Lavansai77/Linked-in-bot/internal/storage"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher" // Added this import
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

	// 2. Initialize Persistent Storage (SQLite)
	storage.InitDB()

	// 3. Initialize Browser with Launcher to avoid Windows Defender "leakless" error
	// .Leakless(false) prevents the creation of the .exe that Defender flags
	l := launcher.New().
		Leakless(false).
		Headless(false). // Set to true if you want it to run in the background
		MustLaunch()

	browser := rod.New().ControlURL(l).MustConnect()
	defer browser.MustClose()

	// 4. Initialize Stealth Page
	page := stealth.InitializePage(browser)

	// 5. Authentication
	email := os.Getenv("LINKEDIN_EMAIL")
	pass := os.Getenv("LINKEDIN_PASSWORD")

	log.Println("Step 1: Starting Authentication...")
	if err := auth.Login(page, email, pass); err != nil {
		log.Fatalf("Critical Error during login: %v", err)
	}

	// 6. Search & Scrape
	searchKeyword := "Software Engineer"
	targetLimit := 3
	log.Printf("Step 2: Searching for '%s' (Limit: %d)...", searchKeyword, targetLimit)

	profiles := actions.SearchAndCollect(page, searchKeyword, targetLimit)

	// 7. Process Connections with Duplicate Prevention
	log.Println("Step 3: Processing profiles with personalized notes...")
	for _, url := range profiles {
		// Check SQLite if we've handled this person before
		if storage.WasContacted(url) {
			log.Printf("Skipping [Already Contacted]: %s", url)
			continue
		}

		log.Printf("Interacting with: %s", url)

		// Personalize your message
		note := "Hi! I noticed your profile in the Software Engineering space and would love to connect."

		err := actions.SendConnectionRequest(page, url, note)
		if err != nil {
			log.Printf("Could not send request to %s: %v", url, err)
			continue
		}

		// Save to database only after successful action
		storage.MarkAsContacted(url)
		log.Printf("Success! Record saved for: %s", url)

		// Technique: Human-like 'Think Time' between different profiles
		stealth.RandomDelay(10, 20)
	}

	log.Println("Automation Task Complete. Closing browser in 5 seconds...")
	stealth.RandomDelay(5, 6)
}
