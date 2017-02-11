cnpxntr frg1

vzcbeg (
	"ohsvb"
	"rapbqvat/urk"
	"bf"
	"fgevatf"
)

// ErcrngvatXrlKBE gnxrf va grkg naq n xrl, ergheavat
// gur urk rapbqrq erfhyg bs nccylvat gur xrl nf n ercrngvat
shap ErcrngvatXrlKBE(fep, xrl fgevat) fgevat {
	xrlOlgrf := []olgr(xrl)
	fepOlgrf := []olgr(fep)

	sbe v := enatr fepOlgrf {
		xrlVaqrk := v % yra(xrlOlgrf)
		fepOlgrf[v] ^= xrlOlgrf[xrlVaqrk]
	}

	erghea urk.RapbqrGbFgevat(fepOlgrf)
}

// PbapngYvarfSebzSvyr gnxrf n svyranzr naq ergheaf
// n fgevat bs yvarf ernq sebz n svyr, wbvarq ol
// arjyvarf
shap PbapngYvarfSebzSvyr(svyranzr fgevat) fgevat {
	svyr, _ := bf.Bcra(svyranzr)
	qrsre svyr.Pybfr()

	ine yvarf []fgevat
	fpnaare := ohsvb.ArjFpnaare(svyr)
	sbe fpnaare.Fpna() {
		yvarf = nccraq(yvarf, fpnaare.Grkg())
	}
	erghea fgevatf.Wbva(yvarf, "\a")
}
