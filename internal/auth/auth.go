package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Lavansai77/Linked-in-bot/internal/stealth"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "cookies.json"

func Launcher() string {
	l := launcher.New().
		Headless(false).
		Devtools(false).
		Leakless(false)

	l.Set("start-maximized")
	l.Set("no-sandbox", "true")
	l.Set("disable-infobars", "true")

	fmt.Println("🚀 Launching browser with native window settings...")
	return l.MustLaunch()
}

func Login(page *rod.Page, email, password string) error {
	_ = page.SetViewport(nil)

	if err := loadCookies(page); err == nil {
		fmt.Println("🔄 Session found. Checking authentication status...")
		page.MustNavigate("https://www.linkedin.com/feed/")

		if waitLoggedIn(page, 5) {
			fmt.Println("✅ Success: Session restored.")
			return nil
		}
		fmt.Println("❌ Session expired. Proceeding to manual login.")
	}

	fmt.Println("🌐 Navigating to LinkedIn Login page...")
	page.MustNavigate("https://www.linkedin.com/login")

	fmt.Println("⌨️ Typing credentials...")
	page.MustElement("#username").MustWaitVisible().MustInput("")
	stealth.TypeLikeHuman(page, email)

	page.MustElement("#password").MustWaitVisible().MustInput(password)

	fmt.Println("🖱️ Moving mouse to 'Sign In' button...")
	loginBtn := page.MustElement("button[type=submit]")
	box := loginBtn.MustWaitVisible().MustShape().Box()
	stealth.MoveMouseHumanLike(page, box.X+box.Width/2, box.Y+box.Height/2)

	fmt.Println("✨ Clicked Sign In! Waiting for dashboard to load...")
	loginBtn.MustClick()

	// Wait for the URL to change away from /login
	page.MustWaitIdle()

	// Use the new waitLoggedIn function with a 10-second timeout
	if waitLoggedIn(page, 10) {
		fmt.Println("🔓 Login Successful! Dashboard reached.")
		return saveCookies(page)
	}

	return fmt.Errorf("login failed: layout check timed out (dashboard not found)")
}

// waitLoggedIn attempts to find the login indicator multiple times before giving up
func waitLoggedIn(page *rod.Page, timeoutSeconds int) bool {
	for i := 0; i < timeoutSeconds; i++ {
		// Look for the navigation bar or the "Me" menu
		has, _, _ := page.Has(".global-nav__me")
		if has {
			return true
		}
		// If not found, wait 1 second and try again
		time.Sleep(1 * time.Second)
	}
	return false
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
	var cookies []*proto.NetworkCookieParam
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}
	return page.SetCookies(cookies)
}
