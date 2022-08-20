{{ $embed := cembed 
    "title" "Administrator Help"
    "color" 16711680 
    "fields" (cslice 
	(sdict "name" "Moderation Commands" "value" "https://docs.yagpdb.xyz/tools-and-utilities/moderation\n_note: most moderation commands have a slash command equivalent_") 
	(sdict "name" "Custom Commands" "value" "• `-addReactionRoles <Channel> <MessageID> emoji1 role1 emoji2 role2 ... emojiN roleN`\n_example: -addReactionRoles #roles 1007928324271656933_ :slight_smile: <@&934742181808328776>\n• `-removeReactionRoles <Channel> <MessageID>`\n_example: -removeReactionRoles #roles 1007928324271656933_ \n• `-adminHelp`\n_displays this menu_") 
	
    ) 
}}
{{sendMessage nil $embed}}