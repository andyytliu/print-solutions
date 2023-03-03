package reader

import (
	"bufio"
	"io"
	"log"
	"strings"
	"strconv"
)

func ReadSolutions(reader *bufio.Reader,
	writer *bufio.Writer, vars []string) {

	var (
		err error
	)

	_, err = writer.WriteString("{ ")
	if err != nil {
		log.Println(">>>>>>>>>>> error: " + err.Error())
	}

	for {
		var (
			isPrefix = true
			line, ln []byte
		)

		for isPrefix && err == nil {
			ln, isPrefix, err = reader.ReadLine()
			line = append(line, ln...)
		}
		
		if err != nil && err != io.EOF {
			log.Println(">>>>>>>>>>> error: " + err.Error())
			break
		}
		if isPrefix {
			log.Println(">>>>>>>>>>> error: line only partially read!")
			break
		}

		fields := strings.Fields(string(line))

		if len(fields) == 0 {
			if err == io.EOF {
				break
			} else {
				continue
			}
		}
		if len(fields) % 2 != 1 {
			log.Println(">>>>>>>>>>> error: non-odd entries in solution; correct format: 'idx idx_1 coef_1 idx_2 coef_2 ...'")
			break
		}

		idx1, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			log.Println(">>>>>>>>>>> error when parsing index 1 in solution: " +
				fields[0] + ". " + err.Error())
		}

		writer.WriteString(vars[idx1] + " -> ")

		// Empty solution is just 0
		if len(fields) == 1 {
			writer.WriteString("0")
		}

		for i := 0; i < len(fields) / 2; i++ {
			idx2, err := strconv.ParseInt(fields[2*i + 1], 10, 64)
			if err != nil {
				log.Println(">>>>>>>>>>> error when parsing index 2 in solution: " +
					fields[2*i + 1] + ". " + err.Error())
			}

			coef, err := strconv.ParseInt(fields[2*i + 2], 10, 64)
			if err != nil {
				log.Println(">>>>>>>>>>> error when parsing coef in solution: " +
					fields[2*i + 2] + ". " + err.Error())
			}

			out := "+(" + strconv.FormatInt(int64(coef), 10) + ")"
			if idx2 != 0 {
				out += ("*" + vars[idx2])
			}
			writer.WriteString(out)
		}
		writer.WriteString(",\n")


		if err == io.EOF {
			break
		}
	}

	_, err = writer.WriteString(" }\n")
	if err != nil {
		log.Println(">>>>>>>>>>> error: " + err.Error())
	}
	writer.Flush()
}
