package codicefiscale

//ritorna numero consonanti e consonanti come array di stringhe
func GetConsonanti(str string) (int, []string) {

	counter := 0
	var Consonanti []string

	for _, char := range str {

		switch char {
		case 'b', 'c', 'd', 'f', 'g', 'h', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'z':
			counter++
			Consonanti = append(Consonanti, string(char))
		}

	}

	return counter, Consonanti
}

//ritorna numero vocali e vocali come array di stringhe
func GetVocali(str string) (int, []string) {

	counter := 0
	var Vocali []string

	for _, char := range str {

		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			counter++
			Vocali = append(Vocali, string(char))
		}

	}

	return counter, Vocali
}
