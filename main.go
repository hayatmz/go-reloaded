package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Convertit une chapine hexadécimale en décimal
func Hex(hexadecimal string) string {
   
   decimal, err := strconv.ParseInt(hexadecimal, 16, 64)
   if err != nil {
      panic(err)
   }

   return strconv.Itoa(int(decimal))
}

// Convertit une chaîne binaire en décimal
func Bin(decimal string) string {

	binaire, err := strconv.ParseInt(decimal, 2, 64)
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(int(binaire))
}

// Met un mot en majuscule
func Up(s string) string {
	return strings.ToUpper(s)
}

// Met un mot en minuscule
func Low(s string) string {
	return strings.ToLower(s)
}

// cMet un mot en capitales
func Cap(s string) string {
	return strings.Title(s)
}

// Vérifie si une lettre est une voyelle ou un h
func EstVoyelleOuH(lettre rune) bool {
	if lettre == 'a' || lettre == 'e' || lettre == 'i' || lettre == 'o' || lettre == 'u' || lettre == 'h'||
	lettre == 'A' || lettre == 'E' || lettre == 'I' || lettre == 'O' || lettre == 'U' || lettre == 'H' {
		return true
	} else {
		return false
	}
}
/* Version retravaillée de la fontion :
func EstVoyelleOuH(lettre rune) bool {
	return strings.ContainsRune("aAeEiIoOuUhH", lettre)
	}
*/

func main() {
// Lecture du fichier passé en argument 
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v", err)
		return
	}
	
	str := string(input)

// Conversion du contenu en tableau de mots
	output := strings.Split(str," ")

	// Boucle sur chaque mot pour appliquer les transformations
	for index, element := range output {
		switch element{
		case "(hex)" :
			output[index-1] = Hex(output[index-1])
		case "(bin)" :
			output[index-1] = Bin(output[index-1])
		case "(up)" :
			output[index-1] = Up(output[index-1])
		case "(low)" :
			output[index-1] = Low(output[index-1])
		case "(cap)" :
			output[index-1] = Cap(output[index-1])
		
		// Majuscules sur x mots avant le tag
		case "(up,":
			next := output[index+1]
			idx := next[:len(next)-1]
			loop, err := strconv.Atoi(idx)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du tag.", element + next)
				panic(err)
			}
			for i :=1; i<= loop; i++ {
				output[index-i] = Up(output[index-i])
			}

		// Minuscules sur x mots avant le tag
		case "(low,":
			next := output[index+1] // je récupère le string après (low,
			idx := next[:len(next)-1] // je récupère tous les numéro avant la fin du tag
			loop, err := strconv.Atoi(idx) // je converti mon num string en int
			if err != nil {
				fmt.Println("Erreur lors de la lecture du tag.", element + next)
				panic(err)
			}
			for i := 1; i <= loop; i++ { // je me sers du int pour itérer x fois sur les mots précèdents
				output[index-i] = Low(output[index-i])
			}

		// Capitalisation sur x mots avant le tag
		case "(cap,":
			next := output[index+1]
			idx := next[:len(next)-1]
			loop, err := strconv.Atoi(idx)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du tag.", element + next)
				panic(err)
			}
			for i := 1; i <= loop; i++ {
				output[index-i] = Cap(output[index-i])
			}

		// Remplace "a" par "an" si le mot suivant commence par voyelle ou h
		case "a":
			next := output[index+1]
			if EstVoyelleOuH(rune(next[0])) {
				output[index] = "an"
				}
		
		// Remplace "an" par "a" si le mot suivant ne commence pas par une voyelle ou h
		case "an":
			next := output[index+1]
			if !EstVoyelleOuH(rune(next[0])) {
				output[index] = "a"
				}
		}
	}
	
		// Reconvertit le tableau en string
		Noutput := strings.Join(output, " ")

		// Nettoie les tag de transformation
		for i := 1; i < 10; i++ {
			Noutput = strings.ReplaceAll(Noutput, "(cap, " + strconv.Itoa(int(i)) + ")", "")
			Noutput = strings.ReplaceAll(Noutput, "(low, " + strconv.Itoa(int(i)) + ")", "")
			Noutput = strings.ReplaceAll(Noutput, "(up, " + strconv.Itoa(int(i)) + ")", "")
		}
		Noutput = strings.ReplaceAll(Noutput, "(cap)", "")
		Noutput = strings.ReplaceAll(Noutput, "(low)", "")
		Noutput = strings.ReplaceAll(Noutput, "(up)", "")
		Noutput = strings.ReplaceAll(Noutput, "(bin)", "")
		Noutput = strings.ReplaceAll(Noutput, "(hex)", "")
		
		// Ajuste la ponctuation
		Noutput = strings.ReplaceAll(Noutput, " .", ".")
		Noutput = strings.ReplaceAll(Noutput, " ,", ", ")
		Noutput = strings.ReplaceAll(Noutput, " ;", "; ")
		Noutput = strings.ReplaceAll(Noutput, " !", "!")
		Noutput = strings.ReplaceAll(Noutput, " ...", "...")
		Noutput = strings.ReplaceAll(Noutput, "' ", "'")
		Noutput = strings.ReplaceAll(Noutput, " '", "'")
		Noutput = strings.ReplaceAll(Noutput, "  ", " ")
		Noutput = strings.ReplaceAll(Noutput, " :", ": ")
		Noutput = strings.ReplaceAll(Noutput, " ?", "?")

		// Ecrit le résultat dans un fichier de sortie (sample.txt)
		errr := os.WriteFile(os.Args[2], []byte(Noutput), 0644)
		if errr != nil {
			fmt.Println("Erreur lors de la création de sample.txt", errr)
			return
		}
}