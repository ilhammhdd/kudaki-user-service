package redisearch

import "strings"

type Sanitizer interface {
	Set(string)
	Sanitize() string
	UnSanitize() string
}

type RedisearchText string

func (rt *RedisearchText) Set(s string) {
	*rt = RedisearchText(s)
}

func (rt *RedisearchText) Sanitize() string {
	replacer := strings.NewReplacer(`,`, `\,`, `.`, `\.`, `<`, `\<`, `>`, `\>`, `{`, `\{`, `}`, `\}`, `[`, `\[`, `]`, `\]`, `"`, `\"`, `'`, `\'`, `:`, `\:`, `;`, `\;`, `!`, `\!`, `@`, `\@`, `#`, `\#`, `$`, `\$`, `%`, `\%`, `^`, `\^`, `&`, `\&`, `*`, `\*`, `(`, `\(`, `)`, `\)`, `-`, `\-`, `+`, `\+`, `=`, `\=`, `~`, `\~`)
	return replacer.Replace(string(*rt))
}

func (rt *RedisearchText) UnSanitize() string {
	replacer := strings.NewReplacer(`\,`, `,`, `\.`, `.`, `\<`, `<`, `\>`, `>`, `\{`, `{`, `\}`, `}`, `\[`, `[`, `\]`, `]`, `\"`, `"`, `\'`, `'`, `\:`, `:`, `\;`, `;`, `\!`, `!`, `\@`, `@`, `\#`, `#`, `\$`, `$`, `\%`, `%`, `\^`, `^`, `\&`, `&`, `\*`, `*`, `\(`, `(`, `\)`, `)`, `\-`, `-`, `\+`, `+`, `\=`, `=`, `\~`, `~`)
	return replacer.Replace(string(*rt))
}
