{{ $ROLE_TRUNCATION_THRESHOLD := 30 }}

{{$roleMentions := cslice}}
{{range .Member.Roles }} 
	{{- $roleMentions = $roleMentions.Append (print "<@&" . ">") -}} 
{{end}}
{{$roleDisplay := joinStr ", " $roleMentions.StringSlice}}

{{$joinedAt := (.Member.JoinedAt)}}
{{if $joinedAt}}
	{{$joinedAt = $joinedAt.Parse.Format "Jan 02, 2006 3:04 AM"}}
{{end}}

{{$embed := (cembed
	"author" (sdict 
		"name" .User.String
		"icon_url" (.User.AvatarURL "256")
	)
	"title" "Member left"
      	"description" (printf "%s joined %s\n**Roles:** %s" .User.Mention $joinedAt $roleDisplay)
)}}

{{sendMessage 934711561719779398 $embed}}