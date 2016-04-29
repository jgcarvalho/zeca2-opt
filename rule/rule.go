package rule

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pattern [3]string

type Probabilities map[string]float64

type Rule map[Pattern]Probabilities

func Read(fn string) Rule {
	r := make(Rule)
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("ERROR: reading rule", err)
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rep := strings.NewReplacer("[", " ", "]", " ", "->", " ", "{", " ", "}", " ", ":", " ", ",", " ")
	var ln, c, rn string
	var s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11 string
	var p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11 float64

	for scanner.Scan() {
		fmt.Sscanf(rep.Replace(scanner.Text()), "%s %s %s %s %f %s %f %s %f %s %f %s %f %s %f %s %f %s %f %s %f %s %f %s %f",
			&ln, &c, &rn, &s1, &p1, &s2, &p2, &s3, &p3, &s4, &p4, &s5, &p5, &s6, &p6, &s7, &p7, &s8, &p8, &s9, &p9, &s10, &p10, &s11, &p11)
		r[Pattern{ln, c, rn}] = Probabilities{s1: p1, s2: p2, s3: p3, s4: p4, s5: p5, s6: p6, s7: p7, s8: p8, s9: p9, s10: p10, s11: p11}
	}
	return r
}
