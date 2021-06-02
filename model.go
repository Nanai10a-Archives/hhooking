package hhooking

type SnowFlake uint64

// TODO: 汎用structの未定義によりOmitされます

type InteractionReponce struct {
    Type InteractionCallbackType `json:"type"`
    Data *InteractionApplicationCommandCallbackData `json:"data"`
}

type InteractionCallbackType int

const (
    IctPong InteractionCallbackType = 1
    IctChannelMessageWithSource InteractionCallbackType = 4
    IctDeferredChannelMessageWithSource InteractionCallbackType = 5
    IctDeferredUpdateMessage InteractionCallbackType = 6
    IctUpdateMessage InteractionCallbackType = 7
)

type InteractionApplicationCommandCallbackData struct {
    Tts bool `json:"tts"`
    Content *string `json:"content"`
    // Embeds *[]Embed `json:"embeds"`
    AllowedMentions *AllowedMention `json:"allowed_mentions"`
    Flags int `json:"flags"` // set 64
}

type AllowedMention interface {} // FIXME: これInterfaceなんだけどどうしよう

type MessageInteraction struct {
    Id SnowFlake `json:"id"`
    Type InteractionType `json:"type"`
    Name string `json:"name"`
    // User User `json:"user"`
}

type ApplicationCommand struct {
    Id SnowFlake `json:"id"`
    ApplicationId SnowFlake `json:"application_id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Options *[]ApplicationCommandOption `json:"options"`
    DefaultPermisson *bool `json:"default_permission"` // default: true
}

type ApplicationCommandOption struct {
    Type ApplicationCommandOptionType `json:"type"`
    Name string `json:"name"`
    Description string `json:"description"`
    Required *bool `json:"required"`
    Choices *[]ApplicationCommandOptionChoice `json:"choices"`
    Options *[]ApplicationCommandOption `json:"options"`
}

type ApplicationCommandOptionType int

// Acot: ApplicationCommandOptionType
const (
    AcotSubCommand ApplicationCommandOptionType = 1
    AcotSubCommandGroup ApplicationCommandOptionType= 2
    AcotString ApplicationCommandOptionType= 3
    AcotInteger ApplicationCommandOptionType= 4
    AcotBoolean ApplicationCommandOptionType= 5
    AcotUser ApplicationCommandOptionType= 6
    AcotChannel ApplicationCommandOptionType= 7
    AcotRole ApplicationCommandOptionType= 8
    AcotMentionable ApplicationCommandOptionType= 9
)

type ApplicationCommandOptionChoice struct {
    Name string `json:"name"`
    Value string `json:"value"` // string (others) or int (ACOT_INTEGER) // TODO: 合ってる？
}

type GuildApplicationCommandPermissions struct {
    Id SnowFlake `json:"id"`
    ApplicationId SnowFlake `json:"application_id"`
    GuiidId SnowFlake `json:"guild_id"`
    Permissons []ApplicationCommandPermissions `json:"permissons"`
}

type ApplicationCommandPermissions struct {
    Id SnowFlake `json:"id"`
    Type ApplicationCommandPermissonType `json:"type"`
    Permisson bool `json:"permisson"`
}

type ApplicationCommandPermissonType int

// Acpt: ApplicationCommandPermissonType
const (
    AcptRole = 1
    AcptUser = 2
)

type Interaction struct {
    Id SnowFlake `json:"id"`
    ApplicationId SnowFlake `json:"application_id"`
    Type InteractionType `json:"type"`
    Data *ApplicationCommandInteractionData `json:"data"`
    GuildId SnowFlake `json:"guild_id"`
    ChannelId SnowFlake `json:"channel_id"`
    // Member *Member `json:"member"`
    // User *User `json:"user"`
    Token string `json:"token"`
    Version int `json:"version"`
    // Message *Message `json:"message"`
}

type InteractionType int

// It: InteractionType
const (
    ItPing InteractionType = 1
    ItApplicationCommand InteractionType= 2
    ItMessageComponent InteractionType = 3
)

type ApplicationCommandInteractionData struct {
    Id SnowFlake `json:"id"`
    Name string `json:"name"`
    Resolved *ApplicationCommandInteractionDataResolved `json:"resolved"`
    Options *[]ApplicationCommandInteractionDataOption `json:"options"`
    CustomId string `json:"custom_id"`
    ComponentType int `json:"component_type"`
}

type ApplicationCommandInteractionDataResolved struct {
    // Users *map[SnowFlake]User `json:"users"`
    // Members *map[SnowFlake]GuildMember `json:"members"`
    // Roles *map[SnowFlake]Role `json:"roles"`
    // Channels *map[SnowFlake]GuildChannel `json:"channels"`
}

type ApplicationCommandInteractionDataOption struct {
    Name string `json:"name"`
    Type ApplicationCommandOptionType `json:"type"`
    Value *OptionType `json:"value"`
    Options *[]ApplicationCommandInteractionDataOption `json:"options"`
}

type OptionType string // FIXME: この型わかんない

