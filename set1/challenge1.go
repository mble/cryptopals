cnpxntr frg1

vzcbeg (
	"rapbqvat/onfr64"
	"rapbqvat/urk"
)

// UrkGbOnfr64Olgrf gnxrf va n urk-rapbqrq fgevat naq ergheaf n fyvpr bs olgrf
// bs gur onfr64 rdhvinyrag
shap UrkGbOnfr64Olgrf(urkFgevat fgevat) (onfr64Olgrf []olgr, ree reebe) {
	qngn, ree := urk.QrpbqrFgevat(urkFgevat)
	vs ree != avy {
		erghea
	}
	onfr64Olgrf = znxr([]olgr, onfr64.FgqRapbqvat.RapbqrqYra(yra(qngn)))
	onfr64.FgqRapbqvat.Rapbqr(onfr64Olgrf, qngn)
	erghea
}
