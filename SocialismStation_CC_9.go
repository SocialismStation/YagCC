{{/* GLOBAL STORE 1 (REACTION ROLES MESSAGE STORE) SCHEMA
Reaction Roles Message Store is Map<ChannelID, Number>
0 holds the current count
*/}}

{{/* GLOBAL STORE N > 1 (REACTION ROLES REACTION LIST STORE) SCHEMA
Reaction Roles Reaction List Store is Map<MessageId, [Emoji, RoleID][]>
*/}}

{{/*parse messageId as userid for int64 typing*/}}
{{$args := parseArgs 3 "Usage: <ChannelID> <MessageId> <...:emoji: roleName>"
	(carg "channel" "channelID")
	(carg "userid" "messageID")
	(carg "string" "roleList")
}}

{{$channel := ($args.Get 0).ID}}
{{$message := ($args.Get 1)}}
{{$roleList := reSplit " " ($args.Get 2)}}

{{/*split list into list of 2x lists*/}}
{{$pairedRoleList := cslice}}
{{$start := false}}
{{$endIndex := div (len $roleList) 2}}
{{$roleOnlyList := cslice}}
{{range (seq 0 $endIndex)}}
	{{$cursor := mult . 2}}
	{{$current := index $roleList $cursor}}
	{{$next := getRole (index $roleList (add $cursor 1))}}
	{{$next = structToSdict $next}}
	{{$next = $next.Get "ID"}}
	{{$roleOnlyList = $roleOnlyList.Append $current}}
	{{$pairedRoleList = $pairedRoleList.Append (cslice $current $next)}}
{{end}}

{{/*add reactions to the message and the database. added under userid 1, channelid, messageid*/}}
{{/*handle most common cases efficiently [1..9]*/}}
{{$roleCount := len $roleOnlyList}}
{{if (eq 1 $roleCount)}}
    {{addMessageReactions $channel $message (index $roleOnlyList 0)}}
{{else if (eq 2 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1)}}
{{else if (eq 3 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2)}}
{{else if (eq 4 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2) (index $roleOnlyList 3)}}
{{else if (eq 5 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2) (index $roleOnlyList 3) (index $roleOnlyList 4)}}
{{else if (eq 6 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2) (index $roleOnlyList 3) (index $roleOnlyList 4) (index $roleOnlyList 5)}}
{{else if (eq 7 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2) (index $roleOnlyList 3) (index $roleOnlyList 4) (index $roleOnlyList 5) (index $roleOnlyList 6)}}
{{else if (eq 8 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2) (index $roleOnlyList 3) (index $roleOnlyList 4) (index $roleOnlyList 5) (index $roleOnlyList 6) (index $roleOnlyList 7)}}
{{else if (eq 9 $roleCount)}}
	{{addMessageReactions $channel $message (index $roleOnlyList 0) (index $roleOnlyList 1) (index $roleOnlyList 2) (index $roleOnlyList 3) (index $roleOnlyList 4) (index $roleOnlyList 5) (index $roleOnlyList 6) (index $roleOnlyList 7) (index $roleOnlyList 8)}}
{{else}}
	{{range $roleOnlyList}}
        {{addMessageReactions $channel $message .}}
	{{end}}
{{end}}

{{/* see if our channel is mapped to a store, if not add it */}}
{{$channelKey := toInt64 ((dbGet 1 $channel).Value)}}
{{if (not $channelKey)}} 
	{{$channelKey = toInt64 (dbIncr 1 0 1)}}
	{{dbSet 1 $channel $channelKey}}
{{end}}

{{/*set the reaction roles for this channel message*/}}
{{dbSet $channelKey $message $pairedRoleList}}