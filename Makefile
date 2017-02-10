FILES := $(wildcard set*/*)

decrypt:
	sed -i "" -f decrypt.sed $(FILES)

.PHONY: rot13
