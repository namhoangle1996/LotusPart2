package utils

func IsImageType(contentType string) bool {
	var imageType = map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/gif":  true,
		"image/png":  true,
	}

	if _, ok := imageType[contentType]; ok {
		return true
	}

	return false
}
