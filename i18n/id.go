package i18n

type ID uint8

const (
	ID_English ID = iota
	ID_Korean
	ID_French
	ID_German
	ID_SimplifiedChinese
	ID_TraditionalChinese
	ID_PeninsularSpanish
	ID_LatinAmericanSpanish
	ID_Russian
	ID_Japanese
	ID_Portuguese
	ID_Italian
	ID_Unk12
	ID_Unk13
	ID_Unk14
	ID_Unk15
	ID_Max
)

var (
	id_locale_translation = []Locale{
		English,
		Korean,
		French,
		German,
		SimplifiedChinese,
		TraditionalChinese,
		PeninsularSpanish,
		LatinAmericanSpanish,
		Russian,
		Japanese,
		BrazilianPortuguese,
		Italian,
		Unknown,
		Unknown,
		Unknown,
		Unknown,
	}
)

func (id ID) Locale() Locale {
	index := int(id)
	found_exact := index < len(id_locale_translation)
	if !found_exact {
		return English
	}
	exact_locale := id_locale_translation[index]

	if exact_locale == Unknown {
		return English
	}

	return exact_locale
}
