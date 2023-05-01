package libs

type GameEventChannel string

const (
	GameManagementEvent GameEventChannel = "game-management"
	NodeManagementEvent GameEventChannel = "node-management"
	NetworkEvent        GameEventChannel = "network"
	FileEvent           GameEventChannel = "file"
	CommandEvent        GameEventChannel = "command"
	MusicEvent          GameEventChannel = "music"
	SettingsEvent       GameEventChannel = "settings"
	ChatEvent           GameEventChannel = "chat"
	PlayerEvent         GameEventChannel = "player"
	CoreSystemEvent     GameEventChannel = "core-system"
)
