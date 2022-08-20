{{$userid := (parseArgs 1 "Usage: <User>" (carg "userid" "user")).Get 0}}
 
{{/* iterate over entries */}}
{{range dbGetPattern $userid "%" 100 0}}
	{{- printf "%s: %s" .Key .Value}}
{{else}}
	Empty
{{end}}