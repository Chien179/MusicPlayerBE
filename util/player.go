package util

const (
	prev = "prev"
	next = "next"
)

// Direction checks if the player's direction corrects
func IsValidDirection(currency string) bool {
	switch currency {
	case prev, next:
		return true
	}

	return false
}
