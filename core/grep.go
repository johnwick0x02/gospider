package core

import (
	"regexp"
	"strings"
)

const SUBRE = `(?i)(([a-zA-Z0-9]{1}|[_a-zA-Z0-9]{1}[_a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1})[.]{1})+`

var AWSS3 = regexp.MustCompile(``)

// SubdomainRegex returns a Regexp object initialized to match
// subdomain names that end with the domain provided by the parameter.
func subdomainRegex(domain string) *regexp.Regexp {
	// Change all the periods into literal periods for the regex
	d := strings.Replace(domain, ".", "[.]", -1)
	return regexp.MustCompile(SUBRE + d)
}

func GetSubdomains(source, domain string) []string {
	var subs []string
	re := subdomainRegex(domain)
	for _, match := range re.FindAllStringSubmatch(source, -1) {
		subs = append(subs, CleanSubdomain(match[0]))
	}
	return subs
}

func GetAWSS3(source string) []string {
	var aws []string
	for _, match := range AWSS3.FindAllStringSubmatch(source, -1) {
		aws = append(aws, DecodeChars(match[0]))
	}
	return aws
}
