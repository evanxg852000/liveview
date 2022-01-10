package core

type Transformer interface {
	Transform(input interface{}, context map[string]interface{}) interface{}
}
