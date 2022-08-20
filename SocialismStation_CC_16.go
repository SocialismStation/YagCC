{{$embed := (cembed
	"title" "Help"
	"color" 16711680
	"fields" (cslice 
		(sdict "name" "Built-in Help" "value" "https://docs.yagpdb.xyz/commands/all-commands")
		(sdict "name" "Custom Commands" "value" "• `-sapplyResult <Right> <Auth> [Name]`\n_description: add your values to the server compass and returns the new compass_\n_note: when you see get your sapplyvalues results, in the url you will see right=x.xx and auth=x.xx; this corresponds to the right and auth arguments_\n• `-sapply`\n_description: returns the server compass_\n• `-bigemoji`\n_description: displays a big emoji or emojis of message reactions. see `-bigemoji help` for  more\n• `-?`\n_description: displays this menu_")
	)
)}}

{{sendMessage nil $embed}}