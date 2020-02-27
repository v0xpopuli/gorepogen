package entity

type (
	Field struct {
		VarName string
		VarType string
	}

	Definition map[string][]Field
)
