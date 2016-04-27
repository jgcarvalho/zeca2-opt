package ca

import (
	"fmt"
	"testing"

	"github.com/jgcarvalho/zeca2-opt/bayes"
	"github.com/jgcarvalho/zeca2-opt/ca"
	"github.com/jgcarvalho/zeca2-opt/rule"
)

func TestRun(t *testing.T) {
	r := rule.Read("/home/jgcarvalho/gocode/src/github.com/jgcarvalho/zeca2-opt/AAA.probs2")
	// fmt.Println(r)
	conf := ca.Config{
		InitState:   []string{"###", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "###"},
		EndState:    []string{"###", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Ad0", "Ad0", "Ab0", "Ab0", "Ab0", "Ab0", "Ab0", "Ab0", "Ab0", "Ab0", "Ab0", "Ab0", "###"},
		Steps:       5000,
		IgnoreSteps: 1,
	}
	p1 := bayes.CalcPriorStates2(conf.EndState)
	p2, l := conf.Run(r)
	// fmt.Println(p2)
	fmt.Println("Likelihood")
	fmt.Println(l)
	// fmt.Println(p1)
	// fmt.Println(r)
	// bayes.UpdateRule(&r, &p1, &p2, &l)
	// p2, l = conf.Run(r)

	for i := 0; i < 1000; i++ {
		bayes.UpdateRule(&r, &p1, &p2, &l)
		p2, l = conf.Run(r)
	}

	// fmt.Println(r)

}
