package stealth

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// MoveMouseHumanLike simulates human motor skills using a Bézier curve.
func MoveMouseHumanLike(page *rod.Page, targetX, targetY float64) {
	mouse := page.Mouse
	currX, currY := 0.0, 0.0

	controlX := (currX+targetX)/2 + float64(rand.Intn(100)-50)
	controlY := (currY+targetY)/2 + float64(rand.Intn(100)-50)

	steps := 15
	for i := 1; i <= steps; i++ {
		t := float64(i) / float64(steps)
		x := (1-t)*(1-t)*currX + 2*(1-t)*t*controlX + t*t*targetX
		y := (1-t)*(1-t)*currY + 2*(1-t)*t*controlY + t*t*targetY

		mouse.MustMoveTo(x, y)
		time.Sleep(time.Duration(rand.Intn(15)+10) * time.Millisecond)
	}
}

// TypeLikeHuman mimics natural typing cadence with variable delays.
func TypeLikeHuman(page *rod.Page, text string) {
	for _, char := range text {
		page.Keyboard.MustType(input.Key(char))
		time.Sleep(time.Duration(rand.Intn(200)+50) * time.Millisecond)
	}
}

// RandomScroll simulates a human mouse wheel scroll to load dynamic content.
// Essential for triggering LinkedIn's infinite scroll/lazy loading.
func RandomScroll(page *rod.Page) {
	// Scroll distance between 300 and 700 pixels
	scrollAmount := rand.Intn(400) + 300

	// Scroll in 5 small "ticks" to look like a human using a mouse wheel
	steps := 5
	for i := 0; i < steps; i++ {
		page.Mouse.MustScroll(0, float64(scrollAmount/steps))
		time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
	}
	fmt.Println("📜 Dynamic content triggered via stealth scroll.")
}

func RandomDelay(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Second)
}
