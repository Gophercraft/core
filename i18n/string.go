// Internationalization
package i18n

import (
	"fmt"
)

type Locale uint8

const (
	English Locale = iota
	Korean
	French
	German
	SimplifiedChinese
	TraditionalChinese
	PeninsularSpanish
	LatinAmericanSpanish
	Russian
	Japanese
	Portuguese
	Italian
	Unk12
	Unk13
	Unk14
	Unk15
	MaxLocale
)

func (l Locale) String() string {
	lc, err := l.EncodeWord()
	if err != nil {
		return err.Error()
	}

	return lc
}

func (l *Locale) EncodeWord() (string, error) {
	switch *l {
	case English:
		return "enUS", nil
	case Korean:
		return "koKR", nil
	case French:
		return "frFR", nil
	case German:
		return "deDE", nil
	case SimplifiedChinese:
		return "zhCN", nil
	case TraditionalChinese:
		return "zhTW", nil
	case PeninsularSpanish:
		return "esES", nil
	case LatinAmericanSpanish:
		return "esMX", nil
	case Russian:
		return "ruRU", nil
	case Portuguese:
		return "ptPT", nil
	case Italian:
		return "itIT", nil
	default:
		return "", fmt.Errorf("unknown locale %d", l)
	}
}

func LocaleFromString(text string) (l Locale, err error) {
	switch text {
	case "enUS", "enGB":
		l = English
	case "koKR":
		l = Korean
	case "frFR":
		l = French
	case "deDE":
		l = German
	case "zhCN":
		l = SimplifiedChinese
	case "zhTW":
		l = TraditionalChinese
	case "esES":
		l = PeninsularSpanish
	case "esMX":
		l = LatinAmericanSpanish
	case "ruRU":
		l = Russian
	case "jaJP":
		l = Japanese
	case "ptPT", "ptBR":
		l = Portuguese
	case "itIT":
		l = Italian
	default:
		err = fmt.Errorf("unknown locale id %s", text)
	}

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

func (str Text) GetLocalized(locale Locale) string {
	if str == nil {
		return ""
	}

	lString, ok := str[locale]
	if !ok {
		for x := English; x < MaxLocale; x++ {
			if str, ok := str[x]; ok {
				return str
			}
		}

		return "<no localized strings in i18n.String>"
	}

	return lString
}

func GetEnglish(str string) Text {
	if str == "" {
		return nil
	}

	return Text{English: str}
}
