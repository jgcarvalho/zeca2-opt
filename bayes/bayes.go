package bayes

import (
	"math"

	"github.com/jgcarvalho/zeca2-opt/rule"
)

type PriorStates1 map[string]float64

type PriorStates2 map[string]map[string]float64

type PriorStates3 map[[3]string]map[string]float64

type PriorPatterns map[[3]string]map[[3]string]float64

type Likelihood map[string]map[[3]string]map[[3]string]float64

type Posterior rule.Rule

func CalcPriorStates1(correctStates []string) PriorStates1 {
	prior := make(map[string]float64)
	validstates := 0
	for _, v := range correctStates {
		if v != "#" {
			prior[v[1:3]] += 1.0
			validstates++
		}
	}
	for k, v := range prior {
		prior[k] = v / float64(validstates)
	}
	return prior
}

func CalcPriorStates2(correctStates []string) PriorStates2 {
	prior := make(map[string]map[string]float64)
	validstates := make(map[string]int)
	for _, v := range correctStates {
		if v != "#" {
			if _, ok := prior[v[0:1]]; !ok {
				prior[v[0:1]] = make(map[string]float64)
			}
			prior[v[0:1]][v[1:3]] += 1.0
			validstates[v[0:1]]++
		}
	}

	for k, v := range prior {
		for j, u := range v {
			prior[k][j] = u / float64(validstates[k])
		}
	}

	return prior
}

func CalcPriorStates3(correctStates []string) PriorStates3 {
	// priorAA := make(map[string]float64)
	prior := make(map[[3]string]map[string]float64)
	validstates := make(map[[3]string]int)
	var v [3]string
	for i := 1; i < len(correctStates)-1; i++ {
		if correctStates[i] != "#" {
			v = [3]string{correctStates[i-1][0:1], correctStates[i][0:1], correctStates[i+1][0:1]}
			if _, ok := prior[v]; !ok {
				prior[v] = make(map[string]float64)
			}
			prior[v][correctStates[i][1:3]] += 1.0
			validstates[v]++
		}
	}

	for k, v := range prior {
		for j, u := range v {
			prior[k][j] = u / float64(validstates[k])
		}
	}

	return prior
}

func CalcPriorPatterns(p map[[3]string]map[[3]string]float64) PriorPatterns {
	// normalize
	var total float64
	for _, v := range p {
		total = 0
		for _, u := range v {
			total += u
		}
		for j, u := range v {
			v[j] = u / total
		}
	}
	return p
}

func CalcLikelihood(l map[string]map[[3]string]map[[3]string]float64) Likelihood {
	var total float64
	for _, e := range l {
		for _, aa := range e {
			total = 0
			for _, freq := range aa {
				total += freq
			}
			for p, freq := range aa {
				aa[p] = freq / total
			}
		}
	}
	return l
}

func UpdateRule(r *rule.Rule, pS *PriorStates2, pP *PriorPatterns, like *Likelihood) {
	var posteriori float64
	var parcial float64
	for p := range *r {
		parcial = 0.0
		for s := range (*r)[p] {
			posteriori = (*pS)[s[0:1]][s[1:3]] // * like/pP
			posteriori *= (*like)[s][[3]string{p[0][0:1], p[1][0:1], p[2][0:1]}][[3]string{p[0][1:3], p[1][1:3], p[2][1:3]}]
			posteriori /= (*pP)[[3]string{p[0][0:1], p[1][0:1], p[2][0:1]}][[3]string{p[0][1:3], p[1][1:3], p[2][1:3]}]
			if math.IsNaN(posteriori) {
				posteriori = 0.001
			}
			if posteriori < 0.001 {
				posteriori = 0.001
			}
			parcial += posteriori
			// fmt.Println(s, posteriori)
			(*r)[p][s] = posteriori
		}
		for s := range (*r)[p] {
			(*r)[p][s] /= parcial
		}
	}
}
