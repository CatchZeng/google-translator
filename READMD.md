# google-translator

> Google Translate inspired by [google-translate-api](https://github.com/matheuss/google-translate-api) and [gtranslate](https://github.com/bregydoc/gtranslate).

## Install

```shell
go get github.com/CatchZeng/google-translator
```

## Usage

```go
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
	//translated: 你好，世界
}
```
