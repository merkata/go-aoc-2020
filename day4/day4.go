package day4

import (
	"regexp"
	"strconv"
	"strings"
)

// PassportFields are all fields that need to be set in a passport
// for it to be considered valid.
type PassportFields struct {
	BYR bool // (Birth Year)
	IYR bool // (Issue Year)
	EYR bool // (Expiration Year)
	HGT bool // (Height)
	HCL bool // (Hair Color)
	ECL bool // (Eye Color)
	PID bool // (Passport ID)
}

// SetField marks a field in a passport as present.
func (pf *PassportFields) SetField(a []string, validation bool) {
	switch a[0] {
	case "byr":
		if validation {
			birth, _ := strconv.Atoi(a[1])
			pf.BYR = birth >= 1920 && birth <= 2002
		} else {
			pf.BYR = true
		}
	case "iyr":
		if validation {
			issue, _ := strconv.Atoi(a[1])
			pf.IYR = issue >= 2010 && issue <= 2020
		} else {
			pf.IYR = true
		}
	case "eyr":
		if validation {
			expire, _ := strconv.Atoi(a[1])
			pf.EYR = expire >= 2020 && expire <= 2030
		} else {
			pf.EYR = true
		}
	case "hgt":
		if validation {
			measure := a[1][len(a[1])-2:]
			height := a[1][:len(a[1])-2]
			hv, _ := strconv.Atoi(height)
			if measure == "cm" {
				pf.HGT = hv >= 150 && hv <= 193
			} else if measure == "in" {
				pf.HGT = hv >= 59 && hv <= 76
			} else {
				pf.HGT = false
			}
		} else {
			pf.HGT = true
		}
	case "hcl":
		if validation {
			r := regexp.MustCompile("#[0-9a-f]{6}")
			pf.HCL = r.MatchString(a[1])
		} else {
			pf.HCL = true
		}
	case "ecl":
		if validation {
			pf.ECL = a[1] == "amb" || a[1] == "blu" || a[1] == "brn" || a[1] == "gry" || a[1] == "grn" || a[1] == "hzl" || a[1] == "oth"
		} else {
			pf.ECL = true
		}
	case "pid":
		if validation {
			if len(a[1]) != 9 {
				pf.PID = false
			} else {
				_, err := strconv.Atoi(a[1])
				if err != nil {
					pf.PID = false
				} else {
					pf.PID = true
				}
			}
		} else {
			pf.PID = true
		}
	}
}

// Valid returns whether a passport is valid or not.
func (pf *PassportFields) Valid() bool {
	return pf.BYR && pf.IYR && pf.EYR && pf.HGT && pf.HCL && pf.ECL && pf.PID
}

// ValidPassports renders a list of passports and returns how many are valid.
func ValidPassports(input []string, validation bool) int {
	var valid int
	passport := PassportFields{}
	for _, line := range input {
		if line == "" {
			if passport.Valid() {
				valid++
			}
			passport = PassportFields{}
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				key := strings.Split(field, ":")
				passport.SetField(key, validation)
			}
		}
	}
	if passport.Valid() {
		valid++
	}
	return valid
}
