package reader

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"
)

// Position 0 in vars is always the constant 1
func ReadVariables(file_name string) []string {

	var (
		err error
		r rune
		v []rune
		vars []string
	)

	file, err := os.Open(file_name)
	if err != nil {
		log.Println(">>>>>>>>>>> error opening file: " + err.Error())
		return vars
	}
	defer file.Close()

	reader := bufio.NewReader(file)


	vars = append(vars, "1")


	// Throw away everything before '{'
	for {
		r, _, err = reader.ReadRune()
		if err != nil && err != io.EOF {
			log.Println(">>>>>>>>>>> error: " + err.Error())
			break
		}
		if err == io.EOF {
			break
		}

		if r == 123 /* { */ {
			break
		}
	}

	if r == 123 {
		for {
			r, _, err = reader.ReadRune()
			if err != nil && err != io.EOF {
				log.Println(">>>>>>>>>>> error: " + err.Error())
				break
			}
			if err == io.EOF {
				break
			}

			if r == 125 /* } */ {
				break
			}

			if unicode.IsSpace(r) || r == 92 /* backslash */ {
				continue
			} else if r == 44 /* comma */ {
				vars = append(vars, string(v))
				v = []rune{}
			} else {
				v = append(v, r)
			}
		}
		vars = append(vars, string(v))
	}

	return vars
}
