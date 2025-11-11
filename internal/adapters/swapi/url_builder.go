package swapi

import "fmt"

// BuildURL constructs a SWAPI URL with endpoint, page, and optional search
func BuildURL(baseURL, endpoint string, page int, search string) string {
	url := fmt.Sprintf("%s/%s/?page=%d", baseURL, endpoint, page)
	if search != "" {
		url += fmt.Sprintf("&search=%s", search)
	}
	return url
}
