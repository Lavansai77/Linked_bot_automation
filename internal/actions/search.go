package actions

import (
	"fmt"
	"strings"

	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/go-rod/rod"
)

// SearchAndCollect finds profile URLs based on a keyword
func SearchAndCollect(page *rod.Page, keyword string, targetCount int) []string {
	var profileURLs []string
	searchURL := fmt.Sprintf("https://www.linkedin.com/search/results/people/?keywords=%s", keyword)

	page.MustNavigate(searchURL)
	stealth.RandomDelay(3, 5)

	for len(profileURLs) < targetCount {
		stealth.RandomScroll(page)

		// Target links specifically in the search results
		elements, err := page.Elements(".entity-result__title-text a.app-aware-link")
		if err != nil {
			break
		}

		for _, el := range elements {
			url := el.MustAttribute("href")
			if url != nil {
				// Clean the URL (remove tracking ?miniProfileUrn...)
				cleanURL := strings.Split(*url, "?")[0]
				profileURLs = append(profileURLs, cleanURL)
			}
			if len(profileURLs) >= targetCount {
				break
			}
		}

		// Pagination logic
		if len(profileURLs) < targetCount {
			nextBtn, err := page.Element("button[aria-label='Next']")
			if err != nil {
				fmt.Println("No more pages available.")
				break
			}

			box := nextBtn.MustShape().Box()
			stealth.MoveMouseHumanLike(page, box.X+box.Width/2, box.Y+box.Height/2)
			nextBtn.MustClick()

			stealth.RandomDelay(3, 6)
		}
	}
	return profileURLs
}
