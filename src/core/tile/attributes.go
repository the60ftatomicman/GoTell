package tile

// Attributes
// These are used with tiles to help determine interaction logic.
// The core file will provide an example of how to define the attributes later on
type Attributes string

const ATTRIBUTE_DELIMITER string = ";"

// Try to name your attributes with ATTR_
// DO NOT ADD MORE HERE! make a new const with attributes elsewhere
const (
	ATTR_EXAMPLE1  Attributes = "EXAMPLE1;"  // You can define it with a semi-colon
	ATTR_EXAMPLE2  Attributes = "EXAMPLE2"   // ... or not. We'll add it anywho!
)


//Generates an attribute string for a tile.
func GenerateAttributes(attrs ...Attributes) string {
	attributes := ""
	for _,attr := range attrs {
		//TOD -- add panic here to test if we have a delimiter in the middle of a string
		attributes += string(Attributes(attr))
		if (attributes[len(attributes)-1:] != ATTRIBUTE_DELIMITER) {
			attributes += ATTRIBUTE_DELIMITER
		}
	}
	return attributes
}