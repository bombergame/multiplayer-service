package args

import (
	"os"
	"regexp"
)

//GetString returns string command line argument
func GetString(name, defaultValue string) string {
	r, err := regexp.Compile(`^--(?P<name>\w+)=(?P<value>.+)$`)
	if err != nil {
		panic(err)
	}

	for _, arg := range os.Args {
		mp := parseGroups(r, arg)
		if mp == nil {
			continue
		}

		if mp["name"] == name {
			return mp["value"]
		}
	}
	return defaultValue
}

//GetFlag returns command line flag
func GetFlag(name string, defaultValue bool) bool {
	r, err := regexp.Compile(`^--(?P<name>\w+)$`)
	if err != nil {
		panic(err)
	}

	for _, arg := range os.Args {
		mp := parseGroups(r, arg)
		if mp == nil {
			continue
		}

		if mp["name"] == name {
			return true
		}
	}
	return defaultValue
}

func parseGroups(r *regexp.Regexp, str string) map[string]string {
	n := len(r.SubexpNames())

	m := r.FindStringSubmatch(str)
	if len(m) != n {
		return nil
	}

	mp := make(map[string]string)
	for i, k := range r.SubexpNames() {
		mp[k] = m[i]
	}

	return mp
}
