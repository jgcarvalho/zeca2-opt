package ca

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jgcarvalho/zeca2-opt/bayes"
	"github.com/jgcarvalho/zeca2-opt/rule"
)

type Config struct {
	InitState []string
	EndState  []string
	// 	TransStates    []string `toml:"transition-states"`
	// 	Hydrophobicity string   `toml:"hydrophobicity"`
	// 	R              int      `toml:"r"`
	Steps int `toml:"steps"`
	// Consensus int `toml:"consensus"`
	IgnoreSteps int `toml:"ignore-steps"`
}

func (conf Config) Run(rule rule.Rule) (bayes.PriorPatterns, bayes.Likelihood) {
	rand.Seed(time.Now().UTC().UnixNano())
	priorPatterns := make(map[[3]string]map[[3]string]float64)
	likelihood := make(map[string]map[[3]string]map[[3]string]float64)
	var init, end, previous, current []string
	init = make([]string, len(conf.InitState))
	end = make([]string, len(conf.EndState))
	copy(init, conf.InitState)
	copy(end, conf.EndState)
	if len(init) != len(end) {
		panic("Init and End States have diffent lenghts")
	}
	previous = make([]string, len(init))
	copy(previous, init)
	current = make([]string, len(init))

	// set begin and end equals to # (static states)
	current[0], current[len(init)-1] = "###", "###"

	fmt.Println("INIT")
	fmt.Println(init)
	fmt.Println("END")
	fmt.Println(end)

	for i := 0; i < conf.Steps; i++ {

		if i%2 == 0 {
			step(&previous, &current, &init, &end, &rule, &priorPatterns, &likelihood)
			fmt.Println(current)
		} else {
			step(&current, &previous, &init, &end, &rule, &priorPatterns, &likelihood)
			fmt.Println(previous)
		}
	}
	p := bayes.CalcPriorPatterns(priorPatterns)
	l := bayes.CalcLikelihood(likelihood)
	return p, l
}

func step(previous, current, init, end *[]string, ru *rule.Rule, priorP *(map[[3]string]map[[3]string]float64), likelihood *(map[string]map[[3]string]map[[3]string]float64)) {
	var rnd float64
	var state string
	var patAA [3]string
	var patSS [3]string
	for c := 1; c < len(*init)-1; c++ {
		if (*previous)[c] == "###" {
			(*current)[c] = "###"
		} else {
			rnd = rand.Float64()
			for k, v := range (*ru)[rule.Pattern{(*previous)[c-1], (*previous)[c], (*previous)[c+1]}] {
				if v > rnd {
					state = k
					break
				} else {
					rnd -= v
				}
				// parece que a somatoria de v não dá 1.0 sempre (preciso corrigir)
				// isso resulta em estado "", para resolver TEMPORARIAMENTE damos um state qualquer (nesse caso o primeiro)
				if state == "" {
					state = k
				}
			}
			(*current)[c] = state

			if state == "" || state == " " {
				fmt.Println("ERRO", (*previous)[c-1], (*previous)[c], (*previous)[c+1])
			}

			// prior patterns probabilities
			patAA = [3]string{(*previous)[c-1][0:1], (*previous)[c][0:1], (*previous)[c+1][0:1]}
			patSS = [3]string{(*previous)[c-1][1:3], (*previous)[c][1:3], (*previous)[c+1][1:3]}

			if _, ok := (*priorP)[patAA]; !ok {
				(*priorP)[patAA] = make(map[[3]string]float64)
			}
			(*priorP)[patAA][patSS] += 1.0

			// likelihood
			if _, ok := (*likelihood)[(*end)[c]]; !ok {
				(*likelihood)[(*end)[c]] = make(map[[3]string]map[[3]string]float64)
			}
			if _, ok := (*likelihood)[(*end)[c]][patAA]; !ok {
				(*likelihood)[(*end)[c]][patAA] = make(map[[3]string]float64)
			}
			(*likelihood)[(*end)[c]][patAA][patSS] += 1.0
		}
	}
}
