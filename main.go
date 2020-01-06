package main

import (
	"log"

	"github.com/CatchZeng/google-translator/translator"
	"golang.org/x/text/language"
)

func main() {
	translatedText, err := translator.Translate("Hello World", language.English, language.SimplifiedChinese)
	if err != nil {
		panic(err)
	}
	log.Printf("translated: %s", translatedText)
}
