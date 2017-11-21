package player

import "strconv"

// StringToID converts a string into
// an equivalent value of the type ID
func StringToID(strID string) (ID, error) {
	id, err := strconv.Atoi(strID)
	return ID(id), err
}

// IDToString converts a player's id
// into its string equivalent
func IDToString(id ID) string {
	return strconv.Itoa(int(id))
}
