package codicefiscale

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//ritorna le tre lettere del cognome
func GetCognome(cognome string) string {

	cognome = strings.ToLower(cognome)
	cons_num, cons := GetConsonanti(cognome)

	switch {
	case cons_num == 1:
		_, vows := GetVocali(cognome)
		return cons[0] + vows[0] + vows[1]
	case cons_num == 2:
		_, vows := GetVocali(cognome)
		return cons[0] + cons[1] + vows[0]
	case cons_num == 3, cons_num > 3:
		return cons[0] + cons[1] + cons[2]
	default:
		return "error"
	}

}

//ritorna le tre lettere del nome
func GetNome(nome string) string {

	nome = strings.ToLower(nome)
	cons_num, cons := GetConsonanti(nome)

	switch {
	case cons_num == 1:
		_, vows := GetVocali(nome)
		return cons[0] + vows[0] + vows[1]
	case cons_num == 2:
		_, vows := GetVocali(nome)
		return cons[0] + cons[1] + vows[0]
	case cons_num == 3:
		return cons[0] + cons[1] + cons[2]
	case cons_num == 4, cons_num > 4:
		return cons[0] + cons[2] + cons[3]
	default:
		return "error"
	}

}

//ritorna i numeri della data di nascita + genere
func GetDataNascita(data_di_nascita, genere string) string {

	//controlla se la data è valida
	data, err := time.Parse("2006-01-02", data_di_nascita)
	checkerr(err)

	map_file := make(map[string]string)

	gopath, err := exec.Command("go", "env", "GOPATH").Output()
	checkerr(err)

	path := string(gopath) + "/src/github.com/Ex0dIa-dev/codicefiscale"
	err = os.Chdir(path)
	checkerr(err)

	file, err := ioutil.ReadFile("data/mesi.json")
	checkerr(err)
	err = json.Unmarshal(file, &map_file)
	checkerr(err)

	lettera_mese := map_file[data.Month().String()]

	var giorno int

	if genere == "M" {
		giorno = data.Day()

	} else if genere == "F" {
		giorno = data.Day() + 40
	}

	anno := strconv.Itoa(data.Year())

	return anno[2:] + lettera_mese + strconv.Itoa(giorno)

}

//ritorna il codice della citta
func GetCodiceCitta(citta string) string {

	gopath, err := exec.Command("go", "env", "GOPATH").Output()
	checkerr(err)

	path := string(gopath) + "/src/github.com/Ex0dIa-dev/codicefiscale"
	err = os.Chdir(path)
	checkerr(err)

	file, err := os.Open("data/comuni.csv")
	checkerr(err)

	citta = strings.ToUpper(citta)
	reader := csv.NewReader(file)
	var codice_citta string

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if line[0] == citta {
			codice_citta = line[1]
		}
	}

	file.Close()
	return codice_citta
}

var tcf = map[string]int{
	"0": 1, "1": 0, "2": 5, "3": 7, "4": 9, "5": 13, "6": 15, "7": 17, "8": 19,
	"9": 21, "A": 1, "B": 0, "C": 5, "D": 7, "E": 9, "F": 13, "G": 15, "H": 17,
	"I": 19, "J": 21, "K": 2, "L": 4, "M": 18, "N": 20, "O": 11, "P": 3, "Q": 6, "R": 8,
	"S": 12, "T": 14, "U": 16, "V": 10, "W": 22, "X": 25, "Y": 24, "Z": 23,
}

var ordv = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9,
	"K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18,
	"T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25,
}

//ritorna il carattere di controllo
func GetCharControllo(cf string) string {

	if len(cf) != 15 {
		return "codice fiscale non valido"
	}

	cf = strings.ToUpper(cf)

	if regexp.MustCompile("[^a-zA-Z0-9]").MatchString(cf) {
		return ""
	}

	str := tcf[string(cf[14])]
	for i := 0; i <= 13; i += 2 {
		str += tcf[string(cf[i])] + ordv[string(cf[i+1])]
	}
	return string(rune(str%26) + rune('A'))

}
