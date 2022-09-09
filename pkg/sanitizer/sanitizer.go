package sanitizer

type Sanitizer interface {
	Sanitize(s string) string
}
