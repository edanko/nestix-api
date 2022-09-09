package sanitizer

type NopSanitizer struct{}

func (NopSanitizer) Sanitize(s string) string {
	return s
}
