package models

type Auth interface {
	NewAuth(...interface{}) interface{}
}
