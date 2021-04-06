package utils

func PointerInterface(intrface interface{}) *interface{} {
	i := intrface
	return &i
}

func PointerSlice(intrfaces []interface{}) []*interface{} {
	var newIntrfaces []*interface{}
	for _, intrface := range intrfaces {
		newIntrfaces = append(newIntrfaces, PointerInterface(intrface))
	}
	return newIntrfaces
}
