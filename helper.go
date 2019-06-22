package main

type Data struct {
	Message string `json:"message"`
}

// Reverse a string
func (data *Data) reverse() {
	chars := []rune(data.Message)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	data.Message = string(chars)
}
