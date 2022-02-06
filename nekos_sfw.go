// Licensed under the Open Software License version 3.0

package nekos

import (
	"fmt"
	"net/url"
)

// Tickle returns the url of a tickle image/gif
func Tickle() (string, error) {
	res, err := req("img/tickle")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Slap returns the url of a slap image/gif
func Slap() (string, error) {
	res, err := req("img/slap")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Poke returns the url of a poke image/gif
func Poke() (string, error) {
	res, err := req("img/poke")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Pat returns the url of a pat image/gif
func Pat() (string, error) {
	res, err := req("img/pat")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Neko returns the url of a neko image
func Neko() (string, error) {
	res, err := req("img/neko")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Meow returns the url of a cat image/gif
func Meow() (string, error) {
	res, err := req("img/meow")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Lizard returns the url of a lizard image
func Lizard() (string, error) {
	res, err := req("img/lizard")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Kiss returns the url of a kiss image/gif
func Kiss() (string, error) {
	res, err := req("img/kiss")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Hug returns the url of a hug image/gif
func Hug() (string, error) {
	res, err := req("img/hug")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// FoxGirl returns the url of a fox girl image/gif
func FoxGirl() (string, error) {
	res, err := req("img/fox_girl")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Feed returns the url of a feeding image/gif
func Feed() (string, error) {
	res, err := req("img/feed")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Cuddle returns the url of a cuddle image/gif
func Cuddle() (string, error) {
	res, err := req("img/cuddle")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Why returns a question
func Why() (string, error) {
	res, err := req("why")
	if err != nil {
		return "", err
	}
	return res.Why, nil
}

// CatText returns a cat kaomoji
func CatText() (string, error) {
	res, err := req("cat")
	if err != nil {
		return "", err
	}
	return res.Cat, nil
}

// Owoify returns the owoified version of the provided string
func Owoify(text string) (string, error) {
	target := fmt.Sprintf("owoify?text=%s", url.QueryEscape(text))
	res, err := req(target)
	if err != nil {
		return "", err
	}
	return res.Owo, nil
}

// EightBall returns an Ball struct that contains the response of a magic 8ball and a corresponding image
func EightBall() (Ball, error) {
	res, err := req("8ball")
	if err != nil {
		return Ball{}, err
	}
	return Ball{res.Response, res.Url}, nil
}

// Fact returns a fact
func Fact() (string, error) {
	res, err := req("fact")
	if err != nil {
		return "", err
	}
	return res.Fact, nil
}

// NekoGif returns the url of a neko gif
func NekoGif() (string, error) {
	res, err := req("img/ngif")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Kemonomimi returns the url of a kemonomimi image/gif
func Kemonomimi() (string, error) {
	res, err := req("img/kemonomimi")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Holo returns the url of a holo image/gif
func Holo() (string, error) {
	res, err := req("img/holo")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Smug returns the url of a smug image/gif
func Smug() (string, error) {
	res, err := req("img/smug")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Baka returns the url of a baka image/gif
func Baka() (string, error) {
	res, err := req("img/baka")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Woof returns the url of a dog image/gif
func Woof() (string, error) {
	res, err := req("img/woof")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Spoiler returns a spoiler per letter for discord markdown version of the provided string
func Spoiler(text string) (string, error) {
	target := fmt.Sprintf("spoiler?text=%s", url.QueryEscape(text))
	res, err := req(target)
	if err != nil {
		return "", err
	}
	return res.Owo, nil
}

// Wallpaper returns the url of a wallpaper
func Wallpaper() (string, error) {
	res, err := req("img/wallpaper")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Goose returns the url of a goose image
func Goose() (string, error) {
	res, err := req("img/goose")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Gecg returns the url of a genetically engineered catgirl image
func Gecg() (string, error) {
	res, err := req("img/gecg")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Avatar returns the url of an avatar image
func Avatar() (string, error) {
	res, err := req("img/avatar")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Waifu returns the url of a waifu image
func Waifu() (string, error) {
	res, err := req("img/waifu")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}
