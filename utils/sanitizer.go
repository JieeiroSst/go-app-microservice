package utils

import "github.com/microcosm-cc/bluemonday"

func SanitizeString(str string) string {
	ugcPolicy := bluemonday.UGCPolicy()
	return ugcPolicy.Sanitize(str)
}
