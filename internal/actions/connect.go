package actions

import (
	"fmt"
	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/go-rod/rod"
)

func SendConnectionRequest(page *rod.Page, profileURL string, note string) error {
	fmt.Printf("Navigating to profile: %s\n", profileURL)
	page.MustNavigate(profileURL)
	stealth.RandomDelay(4, 7)

	// Look for the Connect button
	// LinkedIn often hides Connect behind 'More' for certain profiles
	btn, err := page.Element(`button[aria-label^="Invite"], button[aria-label^="Connect"]`)
	
	if err != nil {
		fmt.Println("Connect button hidden. Opening 'More' menu...")
		moreBtn := page.MustElement(`button[aria-label="More actions"]`)
		box := moreBtn.MustShape().Box()
		stealth.MoveMouseHumanLike(page, box.X+box.Width/2, box.Y+box.Height/2)
		moreBtn.MustClick()
		
		stealth.RandomDelay(1, 2)
		btn = page.MustElement(`div[aria-label^="Invite"], div[aria-label^="Connect"]`)
	}

	// Move and Click
	box := btn.MustShape().Box()
	stealth.MoveMouseHumanLike(page, box.X+box.Width/2, box.Y+box.Height/2)
	btn.MustClick()
	stealth.RandomDelay(2, 3)

	// Handling the Note Modal
	if note != "" {
		addNote, err := page.Element(`button[aria-label="Add a note"]`)
		if err == nil {
			addNote.MustClick()
			stealth.RandomDelay(1, 2)
			
			// Use our stealth typing in the text area
			textArea := page.MustElement("#custom-message")
			textArea.MustClick()
			stealth.TypeLikeHuman(page, note)
			
			stealth.RandomDelay(1, 2)
			page.MustElement(`button[aria-label="Send now"]`).MustClick()
		}
	} else {
		// Just send without a note
		page.MustElement(`button[aria-label="Send without a note"]`).MustClick()
	}

	fmt.Println("Request successfully sent.")
	return nil
}