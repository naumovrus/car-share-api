package validator

import "time"

func ValidateAge(dateOfBirth time.Time) bool {
	year := time.Hour * 8760
	return dateOfBirth.Add(time.Duration(year)).Compare(time.Now()) >= 0
}
