package config

type WorldVarType int

const (
	Bool WorldVarType = iota
	Uint
	Int
	Float
	String
)

var WorldVarList = map[string]WorldVarType{
	// the XP multiplier
	"XP.Rate": Uint,

	// The number of active connections that can be opened per IP address.
	"Network.ConnectionsPerAddress": Int,

	// The number of active sessions allowed per game account.
	// if <= 0 : theoretically infinite, new connection throttling/rate limiting notwithstanding
	"Network.SessionsPerAccount": Int,

	// The number of players will normally be logged in without entering a waiting queue.
	// if <= 0, an infinite amount of players can join this server (until memory is exhausted).
	"Network.MaxPlayers": Int,

	// If true, when players try to log in after the number of in-world players > MaxPlayers, they will be placed in a waiting queue.
	// If false, players will simply be booted from the server if the limit is reached.
	"Network.WaitQueue": Bool,

	// The typical distance at which objects become no longer visible.
	// Some objects can be flagged to disregard this parameter.
	// Some maps will disregard this parameter
	"Sync.VisibilityRange": Float,

	// If true, players will be flagged for PvP upon entering contested/enemy territory.
	"PvP.CombatantUponEnteringContestedTerritory": Bool,

	// If true, players are always flagged for PvP and can kill other players,
	// even if they are on the same team!
	// Players in guilds form a faction that cannot fight eachother.
	"PvP.Deathmatch": Bool,
	// If true, players are divided into a red and blue team that can attack eachother
	// If false, world peace has been achieved. Red and blue cannot normally fight.
	"PvP.AtWar": Bool,
	// If true, players can temporarily set their differences aside and join together in raid and party groups.
	"PvP.CrossFactionGroups": Bool,
	// If true, players can forsake their default racial affiliation and fight for the enemy.
	"PvP.AllowTreason": Bool,

	// If true, players will be restricted to the languages they understand.
	"Chat.LanguageBarrier": Bool,
	// If true, and LanguageBarrier is true, players who speak similar/related languages
	// will be able to understand eachother when those languages are selected.
	"Chat.MutualIntelligibility": Bool,
	// Host an channel where all players can chat. Can be integrated with Discord.
	"Chat.HostAnnounceChannel": Bool,

	// If true, players will be allowed to register normally blocked race-class combinations.
	"Char.AllowSpecialCombos": Bool,

	// The level characters will be after first login
	"Char.StartLevel": Int,

	// Show cinematic on first login
	"Char.Cinematic":     Bool,
	"Char.StartPosition": String,

	// If true, binding items become bound to a player character for its lifetime, and cannot be traded
	"Item.Binding": Bool,

	// Manipulates the chance gradient of loot quality you will receive.
	"Item.LootQuality": Float,

	// If true, item restrictions according to race, level, skill, required spells, locations etc
	// Are removed
	"Item.IgnoreRestrictions": Bool,

	"Weather.On":    Bool,
	"Weather.Crazy": Bool,
}
