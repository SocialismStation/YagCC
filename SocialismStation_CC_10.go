{{$args := parseArgs 3 "Usage: <Type> <User> <Key> <Value>" 	
	(carg "string" "type")
	(carg "userid" "userid")
	(carg "string" "key")
	(carg "string" "value")}}

{{$value := $args.Get 3}}
{{$type := $args.Get 0}}
{{if (eq "int" $type)}}
	{{$value = toInt $value}}
{{else if (eq "int64" $type)}}
	{{$value = toInt64 $value}}
{{else if (eq "float" $type)}}
	{{$value = toFloat $value}}
{{end}}

{{dbSet ($args.Get 1)  ($args.Get 2) $value}}
 
