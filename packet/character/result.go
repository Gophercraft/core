package character

import "github.com/Gophercraft/core/vsn"

//go:generate gcraft_stringer -type=Result
type Result uint8

const (
	CharCreateInProgress Result = iota
	CharCreateSuccess
	CharCreateError
	CharCreateFailed
	CharCreateNameInUse
	CharCreateDisabled
	CharCreatePvpTeamsViolation
	CharCreateServerLimit
	CharCreateAccountLimit
	CharCreateServerQueue
	CharCreateOnlyExisting
	CharCreateExpansion
	CharCreateExpansionClass
	CharCreateLevelRequirement
	CharCreateUniqueClassLimit
	CharCreateCharacterInGuild
	CharCreateRestrictedRaceclass
	CharCreateCharacterChooseRace
	CharCreateCharacterArenaLeader
	CharCreateCharacterDeleteMail
	CharCreateCharacterSwapFaction
	CharCreateCharacterRaceOnly
	CharCreateCharacterGoldLimit
	CharCreateForceLogin
	CharNameSuccess
	CharNameFailure
	CharNameNoName
	CharNameTooShort
	CharNameTooLong
	CharNameInvalidCharacter
	CharNameMixedLanguages
	CharNameProfane
	CharNameReserved
	CharNameInvalidApostrophe
	CharNameMultipleApostrophes
	CharNameThreeConsecutive
	CharNameInvalidSpace
	CharNameConsecutiveSpaces
	CharNameRussianConsecutiveSilentCharacters
	CharNameRussianSilentCharacterAtBeginningOrEnd
	CharNameDeclensionDoesntMatchBaseName
	CharDeleteInProgress
	CharDeleteSuccess
	CharDeleteFailed
	CharDeleteFailedLockedForTransfer
	CharDeleteFailedGuildLeader
	CharDeleteFailedArenaCaptain
	CharCreateTimeout
	CharCreateThrottle
	CharCreateAlliedRaceAchievement
	CharCreateLevelRequirementDemonHunter
	CharCreateCharacterInCommunity
)

type ResultDescriptor map[Result]uint8

var ResultDescriptors = map[vsn.BuildRange]ResultDescriptor{
	{0, 3368}: {
		CharCreateInProgress: 0x27,
		CharCreateSuccess:    0x28,
		CharCreateError:      0x29,
		CharCreateFailed:     0x2A,
		CharCreateNameInUse:  0x2B,
		CharCreateDisabled:   0x2C,
		CharDeleteInProgress: 0x2D,
		CharDeleteSuccess:    0x2E,
		CharDeleteFailed:     0x2F,
	},

	{5875, 6005}: {
		CharCreateInProgress:        0x2D,
		CharCreateSuccess:           0x2E,
		CharCreateError:             0x2F,
		CharCreateFailed:            0x30,
		CharCreateNameInUse:         0x31,
		CharCreateDisabled:          0x3A,
		CharCreatePvpTeamsViolation: 0x33,
		CharCreateServerLimit:       0x34,
		CharCreateAccountLimit:      0x35,
		CharCreateServerQueue:       0x30, /// UNSURE
		CharCreateOnlyExisting:      0x30, /// UNSURE

		CharDeleteInProgress:              0x38,
		CharDeleteSuccess:                 0x39,
		CharDeleteFailed:                  0x3A,
		CharDeleteFailedLockedForTransfer: 0x3A, /// UNSURE
		CharDeleteFailedGuildLeader:       0x3A, /// UNSURE

		/*CHAR_LOGIN_IN_PROGRESS                                 : 0x3B,
		  CHAR_LOGIN_SUCCESS                                     : 0x3C,
		  CHAR_LOGIN_NO_WORLD                                    : 0x3D,
		  CHAR_LOGIN_DUPLICATE_CHARACTER                         : 0x3E,
		  CHAR_LOGIN_NO_INSTANCES                                : 0x3F,
		  CHAR_LOGIN_FAILED                                      : 0x40,
		  CHAR_LOGIN_DISABLED                                    : 0x41,
		  CHAR_LOGIN_NO_CHARACTER                                : 0x42,
		  CHAR_LOGIN_LOCKED_FOR_TRANSFER                         : 0x40, /// UNSURE
		  CHAR_LOGIN_LOCKED_BY_BILLING                           : 0x40, /// UNSURE*/

		CharNameSuccess:                                0x50,
		CharNameFailure:                                0x4F,
		CharNameNoName:                                 0x43,
		CharNameTooShort:                               0x44,
		CharNameTooLong:                                0x45,
		CharNameInvalidCharacter:                       0x46,
		CharNameMixedLanguages:                         0x47,
		CharNameProfane:                                0x48,
		CharNameReserved:                               0x49,
		CharNameInvalidApostrophe:                      0x4A,
		CharNameMultipleApostrophes:                    0x4B,
		CharNameThreeConsecutive:                       0x4C,
		CharNameInvalidSpace:                           0x4D,
		CharNameConsecutiveSpaces:                      0x4E,
		CharNameRussianConsecutiveSilentCharacters:     0x4E, /// UNSURE
		CharNameRussianSilentCharacterAtBeginningOrEnd: 0x4E, /// UNSURE
		CharNameDeclensionDoesntMatchBaseName:          0x4E, /// UNSURE
	},

	{6180, 8606}: {
		CharCreateInProgress:        0x2E,
		CharCreateSuccess:           0x2F,
		CharCreateError:             0x30,
		CharCreateFailed:            0x31,
		CharCreateNameInUse:         0x32,
		CharCreateDisabled:          0x33,
		CharCreatePvpTeamsViolation: 0x34,
		CharCreateServerLimit:       0x35,
		CharCreateAccountLimit:      0x36,
		CharCreateServerQueue:       0x37,
		CharCreateOnlyExisting:      0x38,
		CharCreateExpansion:         0x39,

		CharDeleteInProgress:              0x3A,
		CharDeleteSuccess:                 0x3B,
		CharDeleteFailed:                  0x3C,
		CharDeleteFailedLockedForTransfer: 0x3D,
		CharDeleteFailedGuildLeader:       0x3E,
		CharDeleteFailedArenaCaptain:      0x3F,

		CharNameSuccess:                                0x4A,
		CharNameFailure:                                0x4B,
		CharNameNoName:                                 0x4C,
		CharNameTooShort:                               0x4D,
		CharNameTooLong:                                0x4E,
		CharNameInvalidCharacter:                       0x4F,
		CharNameMixedLanguages:                         0x50,
		CharNameProfane:                                0x51,
		CharNameReserved:                               0x52,
		CharNameInvalidApostrophe:                      0x53,
		CharNameMultipleApostrophes:                    0x54,
		CharNameThreeConsecutive:                       0x55,
		CharNameInvalidSpace:                           0x56,
		CharNameConsecutiveSpaces:                      0x57,
		CharNameRussianConsecutiveSilentCharacters:     0x58,
		CharNameRussianSilentCharacterAtBeginningOrEnd: 0x59,
		CharNameDeclensionDoesntMatchBaseName:          0x5A,
	},

	{vsn.V3_0_2, vsn.V3_3_5a}: {
		CharCreateInProgress:           46,
		CharCreateSuccess:              47,
		CharCreateError:                48,
		CharCreateFailed:               49,
		CharCreateNameInUse:            50,
		CharCreateDisabled:             51,
		CharCreatePvpTeamsViolation:    52,
		CharCreateServerLimit:          53,
		CharCreateAccountLimit:         54,
		CharCreateServerQueue:          55,
		CharCreateOnlyExisting:         56,
		CharCreateExpansion:            57,
		CharCreateExpansionClass:       58,
		CharCreateLevelRequirement:     59,
		CharCreateUniqueClassLimit:     60,
		CharCreateCharacterInGuild:     61,
		CharCreateRestrictedRaceclass:  62,
		CharCreateCharacterChooseRace:  63,
		CharCreateCharacterArenaLeader: 64,
		CharCreateCharacterDeleteMail:  65,
		CharCreateCharacterSwapFaction: 66,
		CharCreateCharacterRaceOnly:    67,
		CharCreateCharacterGoldLimit:   68,
		CharCreateForceLogin:           69,

		CharDeleteInProgress:              70,
		CharDeleteSuccess:                 71,
		CharDeleteFailed:                  72,
		CharDeleteFailedLockedForTransfer: 73,
		CharDeleteFailedGuildLeader:       74,
		CharDeleteFailedArenaCaptain:      75,

		CharNameSuccess:                                87,
		CharNameFailure:                                88,
		CharNameNoName:                                 89,
		CharNameTooShort:                               90,
		CharNameTooLong:                                91,
		CharNameInvalidCharacter:                       92,
		CharNameMixedLanguages:                         93,
		CharNameProfane:                                94,
		CharNameReserved:                               95,
		CharNameInvalidApostrophe:                      96,
		CharNameMultipleApostrophes:                    97,
		CharNameThreeConsecutive:                       98,
		CharNameInvalidSpace:                           99,
		CharNameConsecutiveSpaces:                      100,
		CharNameRussianConsecutiveSilentCharacters:     101,
		CharNameRussianSilentCharacterAtBeginningOrEnd: 102,
		CharNameDeclensionDoesntMatchBaseName:          103,
	},
}
