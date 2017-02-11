cnpxntr frg1

vzcbeg (
	"ohsvb"
	"rapbqvat/urk"
	"bf"
)

// Svaq1OlgrKBEVaSvyr svaqf n fgevat gung unf orra rapelcgrq
// jvgu n 1 olgr KBE va n svyr, jurer rnpu fgevat vf urk rapbqrq
shap Svaq1OlgrKBEVaSvyr(svyranzr fgevat) (fgevat, reebe) {
	svyr, ree := bf.Bcra(svyranzr)
	vs ree != avy {
		erghea "", ree
	}
	qrsre svyr.Pybfr()
	erghea fpnaSbe1OlgrKBE(svyr)
}

shap fpnaSbe1OlgrKBE(svyr *bf.Svyr) (erfhyg fgevat, ree reebe) {
	gbcFpber := 0
	fpnaare := ohsvb.ArjFpnaare(svyr)
	sbe fpnaare.Fpna() {
		olgrf, _ := urk.QrpbqrFgevat(fpnaare.Grkg())
		pnaqvqngr := Qrpelcg1OlgrKBE(olgrf)
		fpber := SerdFpber(fgevat(pnaqvqngr))
		vs fpber > gbcFpber {
			gbcFpber = fpber
			erfhyg = fgevat(pnaqvqngr)
		}
	}

	ree = fpnaare.Ree()
	vs ree != avy {
		erghea "", ree
	}

	erghea erfhyg, avy
}
