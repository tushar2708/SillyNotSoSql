package database

// Each value Integer, String satisfies this interface
type DataType interface {
	IsValidValue(value interface{}) (bool, error)
}
