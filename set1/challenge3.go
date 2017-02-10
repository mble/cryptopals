cnpxntr frg1

vzcbeg (
	"olgrf"
	"fgevatf"
)

// SerdFpber fpberf gur pbagragf bs n fgevat onfrq ba gur
// serdhrapl bs pbzzba Ratyvfu punenpgref va neenl serdPunef
// Uvture serdhrapl zrnaf n orggre fpber, 0 vs abar ner cerfrag
shap SerdFpber(f fgevat) vag {
	fpber := 0
	serdPunef := [13]fgevat{"r", "g", "n", "b", "v", "a", "f", "u", "e", "q", "y", "h", " "}

	sbe v := enatr serdPunef {
		fpber += fgevatf.Pbhag(fgevatf.GbYbjre(f), serdPunef[v])
	}

	erghea fpber
}

// Qrpelcg1OlgrKBE qrpelcgf n fyvpr bs olgrf rapelcgrq
// ol KBEvat ntnvafg n fvatyr olgr, ol vgrengvat bire nyy
// 256 olgrf, KBEvat, naq purpxvat vs gur erfhyg ybbxf yvxr
// na Ratyvfu jbeq
shap Qrpelcg1OlgrKBE(ohs []olgr) []olgr {
	ine qrpelcgrqOhs []olgr
	gbcFpber := 0

	sbe v := 0; v <= 256; v++ {
		pnaqvqngr := olgrf.Znc(kbeShap(ehar(v)), ohs)
		fpber := SerdFpber(fgevat(pnaqvqngr))
		vs fpber > gbcFpber {
			qrpelcgrqOhs = pnaqvqngr
			gbcFpber = fpber
		}
	}
	erghea qrpelcgrqOhs
}

// kbeShap ergheaf n shapgvba gung KBEf n fvatyr olgr
// ntnvafg nabgure
shap kbeShap(v ehar) shap(ehar) ehar {
	erghea shap(pune ehar) ehar {
		erghea pune ^ v
	}
}
