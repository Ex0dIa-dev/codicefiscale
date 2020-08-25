package codicefiscale

import (
	"log"
	"strings"
)

func GeneraCodiceFiscale(cognome, nome, data_di_nascita, genere, citta string) string {

	lettere_cognome := strings.ToUpper(GetCognome(cognome))
	lettere_nome := strings.ToUpper(GetNome(nome))
	lettere_data := GetDataNascita(data_di_nascita, genere)
	codice_citta := GetCodiceCitta(citta)

	codicefiscale := lettere_cognome + lettere_nome + lettere_data + codice_citta
	char_controllo := GetCharControllo(codicefiscale)
	codicefiscale = codicefiscale + char_controllo

	return codicefiscale
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
