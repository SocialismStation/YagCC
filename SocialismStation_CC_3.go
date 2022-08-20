{{$args := parseArgs 3 "Usage: <User> <Key> <Value>" 
    (carg "userid" "userid")
    (carg "string" "key")
    (carg "string" "value")}}
{{dbSet ($args.Get 0)  ($args.Get 1) ($args.Get 2)}}
 
