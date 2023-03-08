package tile

// Attributes
// These are used with tiles to help determine interaction logic.
// We should ALWAYS include a ; at the end of an attribute.
type Attributes string

const (
	ATTR_EXAMPLE  Attributes = "SOME_VALUE;"
)
func GenerateAttributes(attrs ...Attributes) string {
	attributes := ""
	for _,attr := range attrs {
		attributes += string(Attributes(attr))
	}
	return attributes
}