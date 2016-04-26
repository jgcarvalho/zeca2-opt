package bayes

import (
	"testing"

	"github.com/jgcarvalho/zeca2-opt/bayes"
)

func Test(t *testing.T) {
	end := []string{"#", "M??", "T??", "E??", "A??", "A??", "A??", "Q??", "P??", "H??", "A??", "L??", "P??", "A??", "D??", "A??", "P??", "D??", "I??", "A??", "P??", "E??", "R??", "Db0", "La0", "Ld0", "Sa0", "Kd0", "Fd0", "Da0", "Ga0", "La0", "Ia0", "Aa0", "Ea0", "Ra0", "Qa0", "Ka0", "La0", "La0", "Da0", "Sd0", "Gd1", "Vp0", "Tp0", "D??", "Pd0", "Fd0", "Ad0", "Ip0", "Vb0", "M??", "Ea0", "Qb0", "Vb0", "Ka0", "Sb0", "Pa0", "Td0", "Eb0", "Ab0", "Vb0", "Ib0", "Ra1", "Gd1", "Kb0", "Dp0", "Tb0", "Ip0", "Lb0", "L??", "G??", "Tb0", "Yd0", "Ng0", "Ya0", "Md0", "Ga1", "Md0", "Ta0", "Fd0", "Dp0", "Pa0", "Da0", "Va0", "Ia0", "Aa0", "Aa0", "Ga0", "Ka0", "Ea0", "Aa0", "La0", "Ea0", "Kd0", "Fd0", "Gp1", "Sd0", "Gb0", "Td0", "Na1", "Gb0", "Sp0", "Ra0", "Md0", "L??", "Nd0", "Gd0", "Tb0", "Fp0", "Ha0", "Da0", "Ha0", "Ma0", "Ea0", "Va0", "Ea0", "Qa0", "Aa0", "La0", "Ra0", "Da0", "Fa0", "Yd0", "Ga1", "Tb0", "Td0", "Gb0", "A??", "Ib0", "Vb0", "Fb0", "Sd0", "Tb0", "Ga0", "Ya0", "Ma0", "Aa0", "Na0", "La0", "Ga0", "Ia0", "Ia0", "Sd0", "Td0", "La0", "Ab0", "G??", "Kp0", "Gd1", "E??", "Yb0", "Vb0", "Ib0", "Lb0", "Dp0", "Aa0", "Dd0", "Sp0", "Hp0", "Aa0", "Sa0", "Ia0", "Ya0", "Da0", "Ga0", "Ca0", "Qa0", "Qa0", "Gp0", "Nd0", "Ap0", "Eb0", "Ib0", "Vb0", "Rp0", "Fb0", "Rp0", "Hp0", "Na1", "Sb0", "Va0", "Ea0", "Da0", "La0", "Da0", "Ka0", "Ra0", "La0", "Ga0", "Rd0", "Lp0", "Pp0", "Ka0", "Ea0", "Pp0", "Ap0", "Kb0", "Lb0", "Vb0", "Vb0", "Lb0", "Eb0", "Ga0", "Vb0", "Yb0", "Sa0", "Ma0", "Ld0", "Gd1", "Dp0", "Ib0", "Ap0", "Pg0", "La0", "Ka0", "Ea0", "Ma0", "Va0", "Aa0", "Va0", "Aa0", "Ka0", "Ka0", "Hd0", "Ga1", "Ap0", "Mp0", "Vb0", "Lb0", "Vb0", "Db0", "E??", "Aa0", "Hd0", "S??", "Mp0", "Gd1", "F??", "Fb0", "G??", "Pa0", "Nd0", "Ga1", "Ra0", "G??", "Va0", "Ya0", "Ea0", "Aa0", "Qd0", "Gd1", "Ld0", "Ep0", "Gd1", "Qd0", "Ib0", "Da0", "Fb0", "Vb0", "Vb0", "Gb0", "Tb0", "Fd0", "S??", "S??", "Va0", "Gd1", "T??", "Vb0", "G??", "Gb0", "Fb0", "Vb0", "Vb0", "Sb0", "Nd0", "Hp0", "Pa0", "Kd0", "Fa0", "Ed0", "Aa0", "Va0", "Ra0", "Ld0", "A??", "Cb0", "Ra0", "Pd0", "Ya0", "Ia0", "F??", "Tb0", "Ap0", "Sp0", "Lp0", "Pp0", "Pa0", "Sa0", "Va0", "Va0", "Aa0", "Ta0", "Aa0", "Ta0", "Ta0", "Sa0", "Ia0", "Ra0", "Ka0", "La0", "Ma0", "T??", "Ag0", "Ha0", "Ea0", "Ka0", "Ra0", "Ea0", "Ra0", "La0", "Wa0", "Sa0", "Na0", "Aa0", "Ra0", "Aa0", "La0", "Ha0", "Ga0", "Ga0", "La0", "Ka0", "Aa0", "Ma0", "Gd1", "Fp0", "Rb0", "Lb0", "Gd0", "Tb0", "Ea0", "Tb0", "C??", "Dd0", "Sb0", "Ad0", "Ib0", "Vb0", "Ab0", "Vb0", "Mb0", "Lb0", "Ea0", "Db0", "Qa0", "Ea0", "Qa0", "Aa0", "Aa0", "Ma0", "Ma0", "Wa0", "Qa0", "Aa0", "La0", "La0", "Da0", "Gd0", "Gd1", "Lb0", "Yb0", "Vb0", "Nb0", "Mb0", "Ab0", "Rb0", "Pp0", "Pd0", "Aa0", "Tb0", "Pp0", "Ap0", "Gd1", "Tb0", "Fb0", "Lb0", "Lb0", "Rb0", "Cb0", "Sb0", "Ip0", "Cb0", "Aa0", "Ed0", "Hp0", "Tp0", "Pa0", "Aa0", "Qa0", "Ia0", "Qa0", "Ta0", "Va0", "La0", "Ga0", "Ma0", "Fa0", "Qa0", "Aa0", "Aa0", "Ga0", "Ra0", "Aa0", "Vd0", "Gd1", "Va0", "I??", "G??", "L??", "E??", "#"}
	prior1 := bayes.CalcPriorStates1(end)
	t.Log(prior1)
	prior2 := bayes.CalcPriorStates2(end)
	t.Log(prior2)
	prior3 := bayes.CalcPriorStates3(end)
	t.Log(prior3)
}
