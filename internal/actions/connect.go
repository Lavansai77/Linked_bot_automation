package actions

import (
	"fmt"

	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/go-rod/rod"
)

func SendConnectionRequest(page *rod.Page, profileURL string, note string) error {
	fmt.Printf("👤 Processing Profile: %s\n", profileURL)
	page.MustNavigate(profileURL)

	// Ensure profile action bar is loaded
	page.MustElement(".pvs-profile-actions").MustWaitVisible()
	stealth.RandomDelay(2, 4)

	// Strategy: Attempt to find 'Connect' button. If hidden, look in 'More'
	var btn *rod.Element
	var err error

	// Try Regex search for Connect/Invite text
	btn, err = page.ElementR("button", "/(Connect|Invite)/")

	if err != nil {
		fmt.Println("ℹ️ Connect button hidden. Invoking 'More' menu fallback...")
		// Target the 'More' button within the profile header
		moreBtn, moreErr := page.Element(".pvs-profile-actions__action button[aria-label^='More']")
		if moreErr != nil {
			return fmt.Errorf("failed to locate action buttons")
		}

		moreBtn.MustClick()
		stealth.RandomDelay(1, 2)

		// Look for Connect inside the dropdown
		btn, err = page.Element(`div[role="button"][aria-label^="Connect"]`)
		if err != nil {
			return fmt.Errorf("connection restricted by user privacy settings")
		}
	}

	// Move and Click
	box := btn.MustShape().Box()
	stealth.MoveMouseHumanLike(page, box.X+box.Width/2, box.Y+box.Height/2)
	btn.MustClick()

	// Modal Handling
	stealth.RandomDelay(1, 2)
	if note != "" {
		// Use Regex to find the "Add a note" button in the modal
		addNote, err := page.ElementR("button", "Add a note")
		if err == nil {
			addNote.MustClick()
			stealth.TypeLikeHuman(page, note)
			page.MustElementR("button", "Send").MustClick()
		}
	} else {
		// Fallback to "Send without a note" or "Send now"
		sendBtn, _ := page.ElementR("button", "/(Send now|Send without a note)/")
		if sendBtn != nil {
			sendBtn.MustClick()
		}
	}

	return nil
}
