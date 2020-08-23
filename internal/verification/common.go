package verification

const (
	INT_MAX    = 1024
	INT_MIN    = -1024
	STRING_MAX = 20
	STRING_MIN = 0
)



type DataValidator interface {
	IsTypeValid(value interface{}) (bool, error)
	IsNullabilityValid(value interface{}) (bool, error)
	IsMaxValid(value interface{}) (bool, error)
	IsMinValid(value interface{}) (bool, error)
}

// Each schema IntegerSchema, StringSchema satisfies this
type SchemaRules interface {
	IsNullable() bool
	MaxValue() int
	MinValue() int
}
