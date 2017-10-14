package core

// GameType is a type designating the type of a game/match
type GameType string

// DeathMatch indicates the match is of type "deathmatch"
const DeathMatch GameType = "dm"

// CaptureTheFlag indicates the match is of type "capture the flag"
const CaptureTheFlag GameType = "ctf"

// LastManStanding indicates the match is of type "last man standing"
const LastManStanding GameType = "lms"

// Duel indicates the match is of type "duel" (1 on 1)
const Duel GameType = "duel"

// InvalidGameType defines an invalid value for the GameType type
const InvalidGameType = ""

// IsValidGameType checks that the value of 'gt'
// is indeed a valid game type designator
func IsValidGameType(gt string) (GameType, bool) {
	switch GameType(gt) {
	case DeathMatch:
		return DeathMatch, true
	case CaptureTheFlag:
		return CaptureTheFlag, true
	case LastManStanding:
		return LastManStanding, true
	case Duel:
		return Duel, true
	default:
		return InvalidGameType, false
	}
}

// GameTypeToString returns a string representation
// of the game type value
func GameTypeToString(gt GameType) string {
	return string(gt)
}
