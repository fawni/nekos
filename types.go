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

// Ball struct that contains 8ball response and a corresponding image
type Ball struct {
	Text  string
	Image string
}
