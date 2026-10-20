package gauge

import (
	"fmt"
	"math"

	"github.com/egoist/mygo/ui"
)

// Zone is one band of the GaugeZones: everything below To in the color,
// the last zone always running to Max.
type Zone struct {
	// To is the upper cutoff of the zone in domain units. It is ignored
	// for the last zone.
	To float64
	// Color fills the zone.
	Color ui.Color
}

// ZoneColor is the color of the zone a value falls in, or fallback when
// there are no zones. It walks the same cutoffs Zones draws with: each
// zone runs up from the one before, so out-of-order cutoffs are clamped
// into the band they share rather than leaving a gap, and a reading
// outside the domain takes the first or the last zone. Use it to tint a
// needle or a readout by the zone the reading is in.
func ZoneColor(zones []Zone, min, max, value float64, fallback ui.Color) ui.Color {
	if len(zones) == 0 {
		return fallback
	}
	from := min
	for i := range zones[:len(zones)-1] {
		to := clamp(zones[i].To, from, max)
		if value < to {
			return zones[i].Color
		}
		from = to
	}
	return zones[len(zones)-1].Color
}

// compassPoints are the sixteen points of a compass rose, in order.
var compassPoints = [...]string{
	"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE",
	"S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
}

// compassLabel formats a heading in degrees as a compass point when it
// lands on one of the sixteen, and as degrees otherwise, so a 0–360
// domain reads like a bezel: N, 30°, 60°, E, and so on.
func compassLabel(degrees float64) string {
	// Fold the heading into [0, 360) first: a bearing of -10 is 350, and
	// math.Mod would keep the sign and read the far side of the rose.
	turn := math.Mod(math.Mod(degrees, 360)+360, 360)
	index := int(math.Round(turn/22.5)) % 16
	if math.Abs(turn-float64(index)*22.5) < 0.5 {
		return compassPoints[index]
	}
	return fmt.Sprintf("%d°", int(math.Round(turn)))
}

// clockLabel formats a value in hours as a clock numeral, so 0 and 12
// both read "12".
func clockLabel(hours float64) string {
	step := (int(math.Round(hours))%12 + 12) % 12
	if step == 0 {
		return "12"
	}
	return fmt.Sprintf("%d", step)
}

// padTwo writes n with at least two digits.
func padTwo(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprintf("%d", n)
}

// clockTime formats hours past midnight as a 24-hour clock time: 6.5
// reads "06:30". Negative hours reach back into the day before, and past
// 24 into the day after.
func clockTime(hours float64) string {
	minutes := int(math.Round(math.Mod(hours, 24) * 60))
	if minutes < 0 {
		minutes += 24 * 60
	}
	return padTwo(minutes/60%24) + ":" + padTwo(minutes%60)
}

// durationLabel formats a duration in hours as "7h 32m". Hours going
// back past midnight stay negative on the hours alone, so the minutes
// never carry a sign of their own: -0.5 is "-0h 30m", not "0h -30m".
func durationLabel(hours float64) string {
	minutes := int(math.Round(hours * 60))
	sign := ""
	if minutes < 0 {
		sign = "-"
		minutes = -minutes
	}
	return fmt.Sprintf("%s%dh %02dm", sign, minutes/60, minutes%60)
}
