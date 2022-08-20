{{/* political compass data is stored per used in 123456 */}}
{{$entries := dbGetPattern 123456 "%" 100 0}}

{{$apiUrl := "https://www.politicalcompass.org/crowdchart2?spots="}}
{{range $entries}}
	{{$right := .Value.Get "right"}}
	{{$auth := .Value.Get "auth"}}
	{{$name := .Value.Get "name"}}
	{{$apiUrl = joinStr "" $apiUrl (printf "%s|%s|%s," $right $auth $name)}}
{{end}}

{{/* remove trailing comma in url and escape url */}}
{{if (gt (len $entries) 0)}} 
	{{$apiUrlLength := len $apiUrl}}
	{{$apiUrl = slice $apiUrl 0 (sub $apiUrlLength 1)}}
{{end}}

{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png"}}
{{$guildIconUrl := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=256"}}

{{$embed := (cembed 
	"title" "Political Compass"
	"url" $apiUrl
	"timestamp" currentTime
	"footer" (sdict
		"text" "Generated"
		"icon_url" $guildIconUrl)
)}}

{{sendMessage nil $embed}}