package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input" // Added this to solve the "undefined: input" error
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
)

// InitializePage creates a browser instance with stealth properties and randomized viewport.
func InitializePage(browser *rod.Browser) *rod.Page {
	// Mandatory Requirement: Browser Fingerprint Masking
	page := stealth.MustPage(browser)

	// Technique: Viewport Randomization
	viewports := [][]int{{1920, 1080}, {1440, 900}, {1366, 768}}
	chosen := viewports[rand.Intn(len(viewports))]
	page.MustSetViewport(chosen[0], chosen[1], 1, false)

	return page
}

// MoveMouseHumanLike implements Quadratic Bézier curves with natural jitter.
func MoveMouseHumanLike(page *rod.Page, targetX, targetY float64) {
	start := page.Mouse.Position()

	// Technique: Natural Overshoot & Micro-corrections
	controlX := (start.X+targetX)/2 + float64(rand.Intn(120)-60)
	controlY := (start.Y+targetY)/2 + float64(rand.Intn(120)-60)

	steps := 25 + rand.Intn(15)
	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		// Bézier Curve Formula: B(t) = (1-t)²P₀ + 2(1-t)tP₁ + t²P₂
		x := math.Pow(1-t, 2)*start.X + 2*(1-t)*t*controlX + math.Pow(t, 2)*targetX
		y := math.Pow(1-t, 2)*start.Y + 2*(1-t)*t*controlY + math.Pow(t, 2)*targetY

		// Add tiny jitter (micro-corrections)
		x += rand.Float64()*2 - 1
		y += rand.Float64()*2 - 1

		page.Mouse.MoveTo(proto.Point{X: x, Y: y})

		// Technique: Variable speed (mimics human acceleration/deceleration)
		time.Sleep(time.Duration(rand.Intn(8)+3) * time.Millisecond)
	}
}



// TypeLikeHuman simulates uneven typing rhythm on the active page.
func TypeLikeHuman(page *rod.Page, text string) {
	for _, char := range text {
		// Cast the rune to the input.Key type specifically
		page.Keyboard.MustType(input.Key(char))
		
		// Randomized delay between 70ms and 250ms
		ms := rand.Intn(180) + 70
		if char == ' ' {
			ms += 100 
		}
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

// TypeLikeHumanInElement targets a specific element directly.
func TypeLikeHumanInElement(el *rod.Element, text string) {
	for _, char := range text {
		el.MustInput(string(char))
		time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)
	}
}

// RandomDelay mimics human cognitive "think time".
func RandomDelay(min, max int) {
	if min >= max {
		time.Sleep(time.Duration(min) * time.Second)
		return
	}
	seconds := rand.Intn(max-min) + min
	time.Sleep(time.Duration(seconds) * time.Second)
}

// RandomScroll simulates natural reading behavior.
func RandomScroll(page *rod.Page) {
	distance := float64(rand.Intn(500) + 300)
	page.Mouse.Scroll(0, distance, 12)
	RandomDelay(1, 2)

	// Occasional tiny scroll back up (as if re-reading)
	if rand.Float32() > 0.7 {
		page.Mouse.Scroll(0, -70, 5)
	}
}