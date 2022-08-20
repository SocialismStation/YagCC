{{$args := parseArgs 2 "Usage: <ChannelID> <MessageId>"
	(carg "channel" "channelID")
	(carg "userid" "messageID")
}}

{{$channel := ($args.Get 0).ID}}
{{$message := ($args.Get 1).ID}}

{{$channelKey := (dbGet 1 $channel).Value}}
{{if $channelKey}}
	{{dbDel 1 $channel}}
	{{dbDel $channelKey $message}}
{{end}}