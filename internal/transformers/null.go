package transformers

// A `Null` tranformer.
type Null struct{}

// NewNull constructs a null tranformer that does nothing.
func NewNull() Null {
	return Null{}
}

// Transform just returns the original input
func (t *Null) Transform(input interface{}, context map[string]interface{}) interface{} {
	return input
}
