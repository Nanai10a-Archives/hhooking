package hhooking

type SnowFlake string

// TODO: 汎用structの未定義によりOmitされます

type InteractionReponse struct {
	Type InteractionCallbackType                    `json:"type"`
	Data *InteractionApplicationCommandCallbackData `json:"data,omitempty"`
}

type InteractionCallbackType int

const (
	IctPong                             InteractionCallbackType = 1
	IctChannelMessageWithSource         InteractionCallbackType = 4
	IctDeferredChannelMessageWithSource InteractionCallbackType = 5
	IctDeferredUpdateMessage            InteractionCallbackType = 6
	IctUpdateMessage                    InteractionCallbackType = 7
)

type InteractionApplicationCommandCallbackData struct {
	Tts     bool    `json:"tts"`
	Content *string `json:"content,omitempty"`
	// Embeds *[]Embed `json:"embeds,omitempty"`
	AllowedMentions *AllowedMention `json:"allowed_mentions,omitempty"`
	Flags           int             `json:"flags"` // set 64
}

type AllowedMention interface{} // FIXME: これInterfaceなんだけどどうしよう

type MessageInteraction struct {
	Id   SnowFlake       `json:"id"`
	Type InteractionType `json:"type"`
	Name string          `json:"name"`
	// User User `json:"user"`
}

type ApplicationCommand struct {
	Id               SnowFlake                   `json:"id"`
	ApplicationId    SnowFlake                   `json:"application_id"`
	Name             string                      `json:"name"`
	Description      string                      `json:"description"`
	Options          *[]ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermisson *bool                       `json:"default_permission,omitempty"` // default: true
}

type ApplicationCommandOption struct {
	Type        ApplicationCommandOptionType      `json:"type"`
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Required    *bool                             `json:"required,omitempty"`
	Choices     *[]ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options     *[]ApplicationCommandOption       `json:"options,omitempty"`
}

type ApplicationCommandOptionType int

// Acot: ApplicationCommandOptionType
const (
	AcotSubCommand      ApplicationCommandOptionType = 1
	AcotSubCommandGroup ApplicationCommandOptionType = 2
	AcotString          ApplicationCommandOptionType = 3
	AcotInteger         ApplicationCommandOptionType = 4
	AcotBoolean         ApplicationCommandOptionType = 5
	AcotUser            ApplicationCommandOptionType = 6
	AcotChannel         ApplicationCommandOptionType = 7
	AcotRole            ApplicationCommandOptionType = 8
	AcotMentionable     ApplicationCommandOptionType = 9
)

type ApplicationCommandOptionChoice struct {
	Name  string `json:"name"`
	Value string `json:"value"` // string (others) or int (ACOT_INTEGER) // TODO: 合ってる？
}

type GuildApplicationCommandPermissions struct {
	Id            SnowFlake                       `json:"id"`
	ApplicationId SnowFlake                       `json:"application_id"`
	GuiidId       SnowFlake                       `json:"guild_id"`
	Permissons    []ApplicationCommandPermissions `json:"permissons"`
}

type ApplicationCommandPermissions struct {
	Id        SnowFlake                       `json:"id"`
	Type      ApplicationCommandPermissonType `json:"type"`
	Permisson bool                            `json:"permisson"`
}

type ApplicationCommandPermissonType int

// Acpt: ApplicationCommandPermissonType
const (
	AcptRole = 1
	AcptUser = 2
)

type Interaction struct {
	Id            SnowFlake                          `json:"id"`
	ApplicationId SnowFlake                          `json:"application_id"`
	Type          InteractionType                    `json:"type"`
	Data          *ApplicationCommandInteractionData `json:"data,omitempty"`
	GuildId       *SnowFlake                         `json:"guild_id"`
	ChannelId     *SnowFlake                         `json:"channel_id"`
	// Member *Member `json:"member,omitempty"`
	// User *User `json:"user,omitempty"`
	Token   string `json:"token"`
	Version int    `json:"version"`
	// Message *Message `json:"message,omitempty"`
}

type InteractionType int

// It: InteractionType
const (
	ItPing               InteractionType = 1
	ItApplicationCommand InteractionType = 2
	ItMessageComponent   InteractionType = 3
)

type ApplicationCommandInteractionData struct {
	Id            SnowFlake                                  `json:"id"`
	Name          string                                     `json:"name"`
	Resolved      *ApplicationCommandInteractionDataResolved `json:"resolved,omitempty"`
	Options       *[]ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	CustomId      string                                     `json:"custom_id"`
	ComponentType int                                        `json:"component_type"`
}

type ApplicationCommandInteractionDataResolved struct {
	// Users *map[SnowFlake]User `json:"users,omitempty"`
	// Members *map[SnowFlake]GuildMember `json:"members,omitempty"`
	// Roles *map[SnowFlake]Role `json:"roles,omitempty"`
	// Channels *map[SnowFlake]GuildChannel `json:"channels,omitempty"`
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    ApplicationCommandOptionType               `json:"type"`
	Value   *OptionType                                `json:"value,omitempty"`
	Options *[]ApplicationCommandInteractionDataOption `json:"options,omitempty"`
}

type OptionType string // FIXME: この型わかんない
