{{/* Uses the information stored in the database via the addReaction commands. See schema defined in top of file for more info */}}
{{/* This command will trigger up to two database lookups for every emoji added or removed! It is important to only run this on whitelisted 
channels. Preferably, channels with new emojis disabled for maximum kindness to yag */}}
{{$emoji := .Reaction.Emoji.APIName}}
{{$channel := .ReactionMessage.ChannelID}}
{{$message := .ReactionMessage.ID}}


{{$channelKey := toInt64 ((dbGet 1 $channel).Value)}}

{{$reactionKeys := false}}
{{if $channelKey}}
	{{$reactionKeys = (dbGet $channelKey $message).Value}}
{{end}}

{{$role := false}}
{{if $reactionKeys}}
	{{$i := 0}}
	{{while (and (not $role) (lt $i (len $reactionKeys)))}}
		{{$rolePair := index $reactionKeys $i}}
		{{$roleEmoji := index $rolePair 0}}
		{{if (eq $emoji $roleEmoji)}}
			{{$role = index $rolePair 1}}
		{{end}}
		{{$i = add 1 $i}}
	{{end}}
{{end}}

{{/* if a role actually exists for this message there is work to be done */}}
{{if $role}}
	{{if .ReactionAdded}}
		{{giveRoleID .User.ID $role}}
	{{else}}
		{{takeRoleID .User.ID $role}}
	{{end}}
{{end}}
