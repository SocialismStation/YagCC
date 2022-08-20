{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png"}}
{{$guildIconUrl := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=256"}}
{{$embed := (cembed
	"author" (sdict
		"name" .User.Username
		"icon_url" (.User.AvatarURL "256")
	)
	"title" "Member Joined"
	"description" (printf "Welcome to SocialismStation %s!\nDon't be shy, say hello!" .User.Mention)
	"color" 16711680
	"fields" (cslice
		(sdict
			"name" "New User Guid"
			"value" "• **Roles** - <#934726001848614932>\n• **Command Guide** - `-?`"
		)
	)
	"footer" (sdict
		"text" "Joined on"
		"icon_url" $guildIconUrl
	)
	"timestamp" currentTime
)}}

{{sendMessage nil $embed}}