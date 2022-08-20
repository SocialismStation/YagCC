{{/* SCHEMA GLOBAL STORE 123456 (SAPPLY RESULTS) 
the sapply results store is a map of userid to {right: int, auth: int, name: string}
*/}}
{{$STORE_KEY := 123456}}
{{$MIN_INPUT := -10.00}}
{{$MAX_INPUT := 10.00}}

{{$args := parseArgs 2 "<Right> <Auth> [Username]" 
	(carg "float" "right")
	(carg "float" "auth")
	(carg "string" "name")

}}

{{$right := $args.Get 0}}
{{$auth := $args.Get 1}}

{{$error := false}}
{{/* validate right input */}}
{{if (or (lt $right $MIN_INPUT) (gt $right $MAX_INPUT))}}
	{{$error = true}}
	{{printf "right value %d is out of allowed range [-10, 10]" $right}}
{{else if (or (lt $auth $MIN_INPUT) (gt $auth $MAX_INPUT))}}
	{{$error = true}}
	{{printf "auth value %d is out of allowed range [-10, 10]" $auth}}
{{end}}

{{$right := printf "%.2f" $right}}
{{$auth := printf "%.2f" $auth}}

{{/* Determine the name to use */}}
{{$argLen := len .CmdArgs}}
{{$name := false}}
{{if (gt $argLen 2)}}
	{{$name = ""}}
	{{range $index, $value := .CmdArgs}}
		{{if (lt $index 2)}}
			{{continue}}
		{{end}}
		{{$name = joinStr " " $name $value}}
	{{end}}
{{end}}

{{if (not $name)}}
	{{$member := getMember .User.ID}}
	{{if $member}}
		{{$name = $member.Nick}}
	{{end}}
	{{if (not $name)}}
		{{$name = .User.Username}}
	{{end}}
{{end}}

{{/* validate the username */}}
{{$error := false}}
{{if (not (reFind "^\\w+$" $name))}}
	{{$error = true}}
	{{if (eq 2 $argLen)}}
		{{- printf "Name %q contains special characters. Provide an alphanmumeric name as the third argument" $name}}
	{{else}}
		{{- printf "Name %q is not alphanumeric" $name}}
	{{end}}
{{end}}

{{/* Set the users results in the database */}}
{{if (not $error)}}
	{{$result := sdict "right" $right "auth" $auth "name" $name}}
	{{dbSet 123456 .User.ID $result}}

	{{/* Display the new compass link */}}
	{{execCC 15 nil 3 ""}}
{{end}}