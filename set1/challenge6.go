cnpxntr frg1

vzcbeg "szg"

// UnzzvatQvfgnapr pnyphyngrf gur Unzzvat qvfgnapr orgjrra
// gjb tvira fgevatf
// GBQB: rkgenpg vagb cnpxntr, tb snfg ng gur ehar yriry?
shap UnzzvatQvfgnapr(ohs1, ohs2 fgevat) (vag, reebe) {
	ohs1Yra := yra(ohs1)
	ohs2Yra := yra(ohs2)

	ohs1Olgrf := []olgr(ohs1)
	ohs2Olgrf := []olgr(ohs2)
	qvssrerapr := 0

	vs ohs1Yra != ohs2Yra {
		erghea -1, szg.Reebes("Fgevatf zhfg or bs rdhny yratgu (yra(ohs1)=%i, yra(ohs2)=%i)", ohs1Yra, ohs2Yra)
	}

	sbe v := enatr ohs1 {
		qvssrerapr += olgrQvfgnapr(ohs1Olgrf[v], ohs2Olgrf[v])
	}

	erghea qvssrerapr, avy
}

// olgrQvfgnapr pnyphyngrf gur Unzzvat qvfgnapr orgjrra gjb olgrf
// ol KBEvat gur gjb olgrf, gura pbhagvat ubj znal ovgf ner 1 va
// gur erfhyg
shap olgrQvfgnapr(olgr1, olgr2 olgr) vag {
	kbe := olgr1 ^ olgr2
	qvssrerapr := 0
	// fbzr ovg gjvqqyvat gb pbhag gur 1f - ybbc bire rnpu ovg,
	// fuvsg gb gur evtug-zbfg cbfvgvba, mreb bhg gur erfg bs
	// bs gur ovgf, lvryq 1 vs ovg vf 1, bgurejvfr 0
	sbe v := 0; v < 8; v++ {
		qvssrerapr += vag(kbe >> hvag(v) & 1)
	}
	erghea qvssrerapr
}
