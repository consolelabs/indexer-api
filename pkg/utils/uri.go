package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// StandardizeUri return standardized uri, remove trailing slash, replace ipfs link and so on
func StandardizeUri(s, ipfsServer string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "ipfs://") {
		s = strings.Replace(s, "ipfs://", fmt.Sprintf("https://%s/ipfs/", ipfsServer), -1)
	}
	return s
}

func StandardizeSortQuery(sortQ string) string {
	if sortQ == "" {
		return sortQ
	}

	f := func(c rune) bool {
		return c == ','
	}
	sorts := strings.FieldsFunc(sortQ, f)

	re, err := regexp.Compile(`[^\w|-]`)
	if err != nil {
		return ""
	}

	for i := range sorts {
		sort := re.ReplaceAllString(sorts[i], "")
		operator := "ASC"
		if sort[0] == '-' {
			operator = "DESC"
			sort = strings.Replace(sort, "-", "", 1)
		}
		sorts[i] = fmt.Sprintf("%s %s", sort, operator)
	}

	return strings.Join(sorts, ",")
}
