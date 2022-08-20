{{$args := parseArgs 3 "Usage: <User> <Key> <List>" 
    (carg "userid" "userid")
    (carg "string" "key")
    (carg "string" "value")}}
{{$userKey := $args.Get 0}}
{{$key := $args.Get 1}}
{{$value := reSplit "," ($args.Get 2)}}
{{dbSet $userKey $key $value}}
 
