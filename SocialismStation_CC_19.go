{{$createdAt := div .User.ID 4194304 | add 1420070400000 | mult 1000000 | toDuration | (newDate 1970 1 1 0 0 0).Add }}
{{$embed := (cembed 
	"title" "Member joined"
     	"description" (printf "%s %dth to join\ncreated %s" .User.Mention .Guild.MemberCount (($createdAt.Format "Monday, January 2, 2006 at 3:04 AM")))
	"color" 16711680
	"author" (sdict
		"name" .User.String
		"icon_url" (.User.AvatarURL "256")
	)
	"footer" (sdict
		"text" (printf "ID: %s" (toString .User.ID))
	)
	"timestamp" currentTime
)}}

{{sendMessage 934711561719779398 $embed}}