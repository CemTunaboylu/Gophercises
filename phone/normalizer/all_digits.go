package normalizer 


import "strings"

// func Normalize(string) takes an arbitrary format phone number and gets rid of all the formatting leaving only digits. -> transforming in ########## format and returns it as a string 
func Normalize(phone_num string ) string {

	var sb strings.Builder

	for _, r := range phone_num {
		if r >= '0' && r <= '9' {
			sb.WriteRune(r)
		}	
	}

	return sb.String()
}
