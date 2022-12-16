package helper

import "regexp"

func ValidateImageUrl(imageUrl string) (bool, error) {
	isValid, err := regexp.MatchString(`^(?:https?://)?(?:[^/.\s]+\.)*com(?:/[^/\s]+)*/?$`, imageUrl)
	if err != nil {
		return false, err
	}
	return isValid, nil
}