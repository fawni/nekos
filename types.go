// Licensed under the Open Software License version 3.0

package nekos

// Response struct that conatinas all possible returned keys of different endpoints
type Response struct {
	Cat      string
	Fact     string
	Owo      string
	Response string
	Url      string
	Why      string
}

// 8ball struct that contains its text and a corresponding text
type Ball struct {
	Text  string
	Image string
}
