cnpxntr frg1

vzcbeg (
	"rapbqvat/urk"
	"reebef"
)

// SvkrqKBE gnxrf va gjb urk rapbqrq fgevatf bs rdhny yratgu,
// naq ergheaf gur KBE'q erfhyg bs urkOhs1 ntnvafg urkOhs2
shap SvkrqKBE(urkOhs1, urkOhs2 fgevat) ([]olgr, reebe) {
	ohsYra := yra(urkOhs1)

	vs ohsYra != yra(urkOhs2) {
		erghea []olgr{}, reebef.Arj("Vachg yratguf ner abg rdhny")
	}

	erf := znxr([]olgr, ohsYra)
	urkOhs1Olgrf, _ := purpxUrkQrpbqr(urkOhs1)
	urkOhs2Olgrf, _ := purpxUrkQrpbqr(urkOhs2)

	urk.Rapbqr(erf, kbeOlgrf(urkOhs1Olgrf, urkOhs2Olgrf))
	erghea erf, avy
}

shap kbeOlgrf(yrsgOlgrf, evtugOlgrf []olgr) []olgr {
	sbe v := enatr yrsgOlgrf {
		evtugOlgrf[v] ^= yrsgOlgrf[v]
	}
	erghea evtugOlgrf
}

shap purpxUrkQrpbqr(urkFgevat fgevat) ([]olgr, reebe) {
	urkOlgrf, ree := urk.QrpbqrFgevat(urkFgevat)
	vs ree != avy {
		cnavp(ree)
	}
	erghea urkOlgrf, avy
}
