package match

import "strings"

const (
	// DeathMatch indicates the match is of
	// type  "Deathmatch"
	DeathMatch string = "dm"

	// CaptureTheFlag indicates the match is
	// of type "Capture the Flag"
	CaptureTheFlag string = "ctf"

	// LastManStanding indicates the match is
	// of type "Last Man Standing"
	LastManStanding string = "lms"

	// Duel indicates the match is of type "Duel"
	Duel string = "duel"
)

// IsValidGameType verifies that the value
// in the argument corresponds to a known
// game type
func IsValidGameType(gt string) (string, bool) {
	gt = strings.ToLower(gt)

	switch gt {
	case "deathmatch":
		return DeathMatch, true
	case "capture_the_flag":
		fallthrough
	case "capture the flag":
		return CaptureTheFlag, true
	case "last_man_standing":
		fallthrough
	case "last man standing":
		return LastManStanding, true
	case DeathMatch:
		fallthrough
	case CaptureTheFlag:
		fallthrough
	case LastManStanding:
		fallthrough
	case Duel:
		return gt, true
	default:
		return "", false
	}
}
