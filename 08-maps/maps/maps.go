package maps

import "fmt"

func Map() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["AWS"])

	websites["Linkedin"] = "https://linkedin.com"
	delete(websites, "Google")
	fmt.Println(websites)
}
