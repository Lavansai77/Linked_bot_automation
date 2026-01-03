package actions

import (
	"fmt"
	"strings"

	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/go-rod/rod"
)

func SearchAndCollect(page *rod.Page, keyword string, targetCount int) []string {
	var profileURLs []string
	searchURL := fmt.Sprintf("https://www.linkedin.com/search/results/people/?keywords=%s", keyword)

	fmt.Printf("🔍 Starting search for: %s\n", keyword)
	page.MustNavigate(searchURL)
	page.MustElement(".reusable-search__entity-result-list").MustWaitVisible()

	for len(profileURLs) < targetCount {
		fmt.Printf("📜 Scrolling to load more results (Current count: %d)...\n", len(profileURLs))
		stealth.RandomScroll(page)
		stealth.RandomDelay(2, 4)

		elements, err := page.Elements("span.entity-result__title-text a.app-aware-link")
		if err != nil || len(elements) == 0 {
			fmt.Println("⚠️ No more profile elements found on this page.")
			break
		}

		fmt.Println("📦 Extracting profile URLs from DOM...")
		for _, el := range elements {
			url, _ := el.Attribute("href")
			if url != nil && strings.Contains(*url, "/in/") {
				cleanURL := strings.Split(*url, "?")[0]
				if !contains(profileURLs, cleanURL) {
					profileURLs = append(profileURLs, cleanURL)
					fmt.Printf("✅ Found: %s\n", cleanURL)
				}
			}
			if len(profileURLs) >= targetCount {
				fmt.Printf("🏁 Target of %d reached!\n", targetCount)
				return profileURLs
			}
		}

		// Pagination
		nextBtn, err := page.Element("button[aria-label='Next']")
		if err == nil {
			fmt.Println("➡️ Navigating to the next page of results...")
			nextBtn.MustClick()
			page.MustWaitIdle()
			stealth.RandomDelay(2, 3)
		} else {
			fmt.Println("🔚 Reached the end of the search results.")
			break
		}
	}
	return profileURLs
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
