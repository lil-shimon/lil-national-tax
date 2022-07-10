package message

import (
	"github.com/lil-shimon/national-tax/question"
	"strings"
)

func FmtMsg(qs question.Question) string {
	split := strings.Split(qs.Question, " ")
	var msg string

	if len(split) > 1 {
		for _, v := range split {
			msg += v + "\n"
		}
	}
	return msg
}
