package yevpckg

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
)

type errCatcher struct {
	argS string
	msg  string
}

func (e *errCatcher) Error() string {
	return fmt.Sprintf("%s -> %s", e.argS, e.msg)
}

// MyCustomIsDbURLRegex - just try regex in go, not used
func MyCustomIsDbURLRegex(u string) (string, error) {
	reg, err := regexp.Compile(`(\w+):/{2}([\w-]+:[\w-]+)@{1}([\w-]+):(\d+)/{1}.*`)
	if err != nil {
		log.Fatal("Cannot compile regex")
	}
	if reg.MatchString(u) {
		return "OK", nil
	}
	return "BAD", &errCatcher{u, "catch db url parse error"}
}

// IsDbURLValid descr
func IsDbURLValid(u string) (string, error) {
	ul, err := url.Parse(u)
	if err != nil {
		log.Fatal("Cannot parse url link")
	}
	okSl := []string{}
	for k, v := range urlParser(ul) {
		if !v {
			return "BAD", &errCatcher{k, "catch db url parse error"}
		}
		okSl = append(okSl, k)
	}
	return fmt.Sprintf("%v", okSl), nil
}

func urlParser(u *url.URL) map[string]bool {
	fMap := make(map[string]bool)
	if u.Scheme == "" {
		fMap["Scheme"] = false
	} else {
		fMap["Scheme"] = true
	}

	if u.Host == "" {
		fMap["Host"] = false
	} else {
		fMap["Host"] = true
	}
	//if u.User == nil {
	//	fMap["User"] = false
	//} else {
	//	fMap["User"] = true
	//}
	return fMap
}
