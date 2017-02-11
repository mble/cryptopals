FILES := $(wildcard set*/*.go)

decrypt:
	sed -i "" -f decrypt.sed $(FILES)

.PHONY: rot13
