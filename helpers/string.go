package helpers

func StringIfEmpty(v, vIfEmpty string) string {
	if v == "" {
		return vIfEmpty
	}
	return v
}
