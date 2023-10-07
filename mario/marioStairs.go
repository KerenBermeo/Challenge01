package mario

func Stairs(num int) string {
	aux := ""
	for i := 0; i < num; i++ {
		for j := 0; j < i; j++ {
			aux = aux + "*"
		}
		aux = aux + "\n"
	}
	return aux
}
