package grunt

import "github.com/Gophercraft/core/format/tag"

type OS tag.Tag

func (os OS) String() string {
	return tag.Tag(os).String()
}

var (
	Windows = OS(tag.Make("\x00Win"))
	MacOS   = OS(tag.Make("\x00OSX"))
	Linux   = OS(tag.Make("Linx"))
)

type Architecture tag.Tag

func (arch Architecture) String() string {
	return tag.Tag(arch).String()
}

var (
	X86     = Architecture(tag.Make("\x00x86"))
	X64     = Architecture(tag.Make("\x00x64"))
	PowerPC = Architecture(tag.Make("\x00PPC"))
	ARMv6   = Architecture(tag.Make("ARM6"))
	ARMv7   = Architecture(tag.Make("ARM7"))
)

type Locale tag.Tag

func (locale Locale) String() string {
	return tag.Tag(locale).String()
}

// ...

var (
	Locale_enUS = Locale(tag.Make("enUS"))
)

type Program tag.Tag

var (
	WoW = Program(tag.Make("\x00WoW"))
)

func (program Program) String() string {
	return tag.Tag(program).String()
}
