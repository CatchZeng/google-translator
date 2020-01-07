package translator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/CatchZeng/gutils/file"
	"golang.org/x/text/language"
)

// Translate translate text
func Translate(text string, from language.Tag, to language.Tag) (string, error) {
	translated, err := translate(text, from.String(), to.String(), false, 2, 0)
	if err != nil {
		return "", err
	}

	return translated, nil
}

// TranslateFile translate file
func TranslateFile(src string, dstDir string, dstName string, override bool, from language.Tag, to language.Tag) error {
	if !file.Exists(src) {
		return fmt.Errorf("%s is not exists", src)
	}

	dst := path.Join(dstDir, dstName)
	if !override && file.Exists(dst) {
		return errors.New("The dst already exists")
	}

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	text := string(bytes)
	translated, err := Translate(text, from, to)
	if err != nil {
		return err
	}

	mode := file.Mode(src)
	if err := file.WriteStringToFile(translated, dst, mode); err != nil {
		return err
	}
	return nil
}
