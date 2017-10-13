package core

// GameType is a type designating the type of a game/match
type GameType string

// DeathMatch indicates the match is of type "deathmatch"
const DeathMatch GameType = "dm"

// CaptureTheFlag indicates the match is of type "capture the flag"
const CaptureTheFlag GameType = "ctf"

// LastManStanding indicates the match is of type "last man standing"
const LastManStanding GameType = "lms"
