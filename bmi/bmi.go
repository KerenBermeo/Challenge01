package bmi

func Bmi(peso float64, altura float64) (float64, string) {
	var bmi float64
	var str string
	bmi = peso / (altura * altura)
	if bmi < 18.5 {
		str = "Tienes bajo peso, añade más patatas al caldo"
	} else if bmi >= 18.5 && bmi < 25 {
		str = "Tienes un peso normal, te tengo una sana envidia"
	} else if bmi >= 25 {
		str = "Tienes sobrepeso, lo sé, la pandemia nos ha afectado a todos"
	}
	return bmi, str
}
