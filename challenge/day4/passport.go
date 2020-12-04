package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	tokenBirthYear      = "byr"
	tokenIssueYear      = "iyr"
	tokenExpirationYear = "eyr"
	tokenHeight         = "hgt"
	tokenHairColor      = "hcl"
	tokenEyeColor       = "ecl"
	tokenPassportId     = "pid"
	tokenCountryId      = "cid"

	keySeparator      = " "
	keyValueSeparator = ":"

	heightUnitCentimeters = "cm"
	heightUnitInches      = "in"
)

var (
	hairColorRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eyeColorRegex  = regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
)

type passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	ID             string
	CountryID      string
}

func parsePassport(from <-chan string) passport {
	result := passport{}

	for line := range from {
		if line == "" {
			return result
		}

		for _, token := range strings.Split(line, keySeparator) {
			kv := strings.Split(token, keyValueSeparator)
			if len(kv) != 2 {
				panic(fmt.Errorf("invalid token: %s", token))
			}

			switch kv[0] {
			case tokenBirthYear:
				result.BirthYear = kv[1]
			case tokenIssueYear:
				result.IssueYear = kv[1]
			case tokenExpirationYear:
				result.ExpirationYear = kv[1]
			case tokenHeight:
				result.Height = kv[1]
			case tokenHairColor:
				result.HairColor = kv[1]
			case tokenEyeColor:
				result.EyeColor = kv[1]
			case tokenPassportId:
				result.ID = kv[1]
			case tokenCountryId:
				result.CountryID = kv[1]
			}
		}
	}

	return result
}

func (p passport) valid() bool {
	return p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.ID != ""
}

func (p passport) strictlyValid() bool {
	if byr, err := strconv.Atoi(p.BirthYear); p.BirthYear == "" || err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	if iyr, err := strconv.Atoi(p.IssueYear); p.IssueYear == "" || err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr, err := strconv.Atoi(p.ExpirationYear); p.ExpirationYear == "" || err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	switch {
	case strings.HasSuffix(p.Height, heightUnitCentimeters):
		if height, err := strconv.Atoi(strings.TrimSuffix(p.Height, heightUnitCentimeters)); err != nil || height < 150 || height > 193 {
			return false
		}
	case strings.HasSuffix(p.Height, heightUnitInches):
		if height, err := strconv.Atoi(strings.TrimSuffix(p.Height, heightUnitInches)); err != nil || height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}

	if !hairColorRegex.MatchString(p.HairColor) {
		return false
	}

	if !eyeColorRegex.MatchString(p.EyeColor) {
		return false
	}

	if _, err := strconv.Atoi(p.ID); err != nil || len(p.ID) != 9 {
		return false
	}

	return true
}
