package auth

import (
	"encoding/json"
	"fmt"
	"os"

	// Updated Import Path
	"github.com/Lavansai77/Linked-in-bot/internal/stealth"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "cookies.json"

func Login(page *rod.Page, email, password string) error {
	if err := loadCookies(page); err == nil {
		fmt.Println("Reusing existing session...")
		page.MustNavigate("https://www.linkedin.com/feed/")
		if !isLoggedIn(page) {
			fmt.Println("Session expired. Logging in fresh...")
		} else {
			return nil
		}
	}

	fmt.Println("Navigating to login page...")
	page.MustNavigate("https://www.linkedin.com/login")

	// Input email using stealth typing
	page.MustElement("#username").MustInput("")
	stealth.TypeLikeHuman(page, email)

	stealth.RandomDelay(1, 2)

	// Input password
	page.MustElement("#password").MustInput(password)

	// FIX: Calculated Center coordinates manually
	loginBtn := page.MustElement("button[type=submit]")
	box := loginBtn.MustShape().Box()
	stealth.MoveMouseHumanLike(page, box.X+box.Width/2, box.Y+box.Height/2)
	loginBtn.MustClick()

	fmt.Println("Waiting for dashboard... Please solve any Captchas/2FA manually if they appear.")
	page.MustWaitIdle()

	return saveCookies(page)
}

func isLoggedIn(page *rod.Page) bool {
	has, _, _ := page.Has(".global-nav__me")
	return has
}

func saveCookies(page *rod.Page) error {
	cookies, err := page.Cookies(nil)
	if err != nil {
		return err
	}
	data, _ := json.Marshal(cookies)
	return os.WriteFile(cookieFile, data, 0644)
}

func loadCookies(page *rod.Page) error {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return err
	}
	// FIX: Use NetworkCookieParam for loading
	var cookies []*proto.NetworkCookieParam
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}
	return page.SetCookies(cookies)
}
