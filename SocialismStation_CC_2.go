{{$args := parseArgs 3 "Usage: <User> <Key>" 
    (carg "userid" "userid")
    (carg "string" "key")
}}
{{(dbGet ($args.Get 0)  ($args.Get 1)).Value}}
 
