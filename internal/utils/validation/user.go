package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRole(fl validator.FieldLevel) bool {
	var validRoles = []int{1, 2, 3} //1: admin, 2: mentor, 3: user
	role := int(fl.Field().Int())

	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}

func ValidateSpecialFeatures(fl validator.FieldLevel) bool {
	validFeatures := map[string]struct{}{
		"Trailers":          {},
		"Commentaries":      {},
		"Deleted Scenes":    {},
		"Behind the Scenes": {},
	}

	features := fl.Field().String()
	featureList := strings.Split(features, ",")

	for _, feature := range featureList {
		if _, exists := validFeatures[feature]; !exists {
			return false
		}
	}

	return true
}
