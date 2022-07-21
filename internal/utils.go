package internal

func IsvalidMembership(membershipType string) bool {
	for _, value := range validMemberships {
		if value == membershipType {
			return false
		}
	}
	return true
}
