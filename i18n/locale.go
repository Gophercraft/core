// package i18n implements Internationalization and localized text types
package i18n

import (
	"fmt"
	"sort"

	"github.com/Gophercraft/core/format/tag"
)

type Locale tag.Tag

var (
	English              = MakeLocale("enUS")
	Korean               = MakeLocale("koKR")
	French               = MakeLocale("frFR")
	German               = MakeLocale("deDE")
	SimplifiedChinese    = MakeLocale("zhCN")
	TraditionalChinese   = MakeLocale("zhTW")
	PeninsularSpanish    = MakeLocale("esES")
	LatinAmericanSpanish = MakeLocale("esMX")
	Russian              = MakeLocale("ruRU")
	Japanese             = MakeLocale("jaJP")
	BrazilianPortuguese  = MakeLocale("ptBR")
	EuropeanPortuguese   = MakeLocale("ptPT")
	Italian              = MakeLocale("itIT")
	Unknown              = MakeLocale("xxXX")
)

func MakeLocale(text string) Locale {
	return Locale(tag.Make(text))
}

func (l Locale) String() string {
	return tag.Tag(l).String()
}

func (l Locale) EncodeWord() (string, error) {
	return l.String(), nil
}

func LocaleFromString(text string) (l Locale, err error) {
	if len(text) != 4 {
		err = fmt.Errorf("i18n: locale tag must be 4 bytes long")
		return
	}

	l = MakeLocale(text)
	return
}

// Note: refers to the Locale identifier, not the encoding of text in the language.
func (l *Locale) DecodeWord(word string) error {
	lc, err := LocaleFromString(word)
	if err != nil {
		return err
	}
	*l = lc
	return nil
}

type Text map[Locale]string

func (str Text) String() string {
	if len(str) == 0 {
		return "<empty>"
	}

	return str.GetLocalized(English)
}

func (text Text) GetLocalized(locale Locale) string {
	if text == nil {
		return ""
	}

	// Look for an exact match
	exact_match, found_exact := text[locale]
	if found_exact {
		return exact_match
	}

	// Look for other localizations in the same language group
	language_group := locale.String()[0:2]

	var available_locales []Locale
	for k := range text {
		available_locales = append(available_locales, k)
	}

	// Make deterministic
	sort.Slice(available_locales, func(i, j int) bool {
		return available_locales[i] < available_locales[j]
	})

	for i := range available_locales {
		available_locale := available_locales[i]
		available_language_group := locale.String()[0:2]
		if available_language_group == language_group {
			// found a bad but acceptable localization
			return text[available_locale]
		}
	}

	// failed to find acceptable localization
	return "<no acceptable localized strings in i18n.Text>"
}

func GetEnglish(str string) Text {
	if str == "" {
		return nil
	}

	return Text{English: str}
}
