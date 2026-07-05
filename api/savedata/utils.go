package savedata

// Confirm that every migrator existing in the old save exists in the new save with identical timestamp
func ValidMigrators(new map[string]int, old map[string]int) bool {
	for migrator, timestamp := range old {
		newTimestamp := new[migrator]
		if newTimestamp != timestamp {
			return false
		}
	}
	return true
}
