package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Lavansai77/Linked-in-bot/internal/actions"
	"github.com/Lavansai77/Linked-in-bot/internal/auth"
	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/go-rod/rod"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load Environment
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Initialize Browser with Stealth Launcher
	// auth.Launcher is now defined in auth.go
	browser := rod.New().ControlURL(auth.Launcher()).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()

	// 3. Perform Authentication
	fmt.Println("🔑 Checking Authentication Status...")
	if err := auth.Login(page, os.Getenv("LINKEDIN_EMAIL"), os.Getenv("LINKEDIN_PASSWORD")); err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	// 4. Search and Collect
	keyword := "Software Recruiter"
	fmt.Printf("🔍 Searching for keyword: %s\n", keyword)
	profiles := actions.SearchAndCollect(page, keyword, 3) // Target 3 for a quick video demo

	if len(profiles) == 0 {
		fmt.Println("❌ No profiles found. LinkedIn might be blocking the search or selectors changed.")
		return
	}

	// 5. Connect Loop
	fmt.Printf("\n🤝 Found %d profiles. Starting connection sequence...\n", len(profiles))

	for i, url := range profiles {
		fmt.Printf("\n🕒 [%d/%d] Target: %s\n", i+1, len(profiles), url)

		// Navigate to the profile
		page.MustNavigate(url)
		stealth.RandomDelay(3, 6)

		// UPDATED: Changed from ConnectWithProfile to SendConnectionRequest
		err := actions.SendConnectionRequest(page, url, "Hi, I'd love to connect and learn more about your work!")

		if err != nil {
			fmt.Printf("⚠️ Skip: %v\n", err)
		} else {
			fmt.Println("✨ Success: Connection request sent.")
		}

		// Randomized pause to mimic human reading time
		stealth.RandomDelay(10, 15)
	}

	fmt.Println("\n🏁 Automation sequence finished successfully.")
}
