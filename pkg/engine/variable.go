package engine

type Variable struct {
	Key string
	Val any
}

type Variables []*Variable

type VariableMap map[string]any

func (v VariableMap) ToSlice() Variables {
	var result = make([]*Variable, 0, len(v))
	for k, val := range v {
		result = append(result, &Variable{
			Key: k,
			Val: val,
		})
	}
	return result
}
