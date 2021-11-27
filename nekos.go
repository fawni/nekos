// Licensed under the Open Software License version 3.0

package nekos

import (
	"fmt"
	"net/url"
)

// found in https://github.com/Nekos-life/nekos-dot-life/blob/master/endpoints.json
// nsfw
func RandomHentaiGif() string {
	res := DoRequest("img/Random_hentai_gif")
	return res.Url
}

func Pussy() string {
	res := DoRequest("img/pussy")
	return res.Url
}

func LewdNekoGif() string {
	res := DoRequest("img/nsfw_neko_gif")
	return res.Url
}

func LewdNeko() string {
	res := DoRequest("img/lewd")
	return res.Url
}

func Lesbian() string {
	res := DoRequest("img/les")
	return res.Url
}

func Kuni() string {
	res := DoRequest("img/kuni")
	return res.Url
}

func Cumsluts() string {
	res := DoRequest("img/cum")
	return res.Url
}

func Classic() string {
	res := DoRequest("img/classic")
	return res.Url
}

func Boobs() string {
	res := DoRequest("img/boobs")
	return res.Url
}

func BlowJobGif() string {
	res := DoRequest("img/bj")
	return res.Url
}

func Anal() string {
	res := DoRequest("img/anal")
	return res.Url
}

func LewdAvatar() string {
	res := DoRequest("img/nsfw_avatar")
	return res.Url
}

func Yuri() string {
	res := DoRequest("img/yuri")
	return res.Url
}

func Trap() string {
	res := DoRequest("img/trap")
	return res.Url
}

func Tits() string {
	res := DoRequest("img/tits")
	return res.Url
}

func GirlSoloGif() string {
	res := DoRequest("img/solog")
	return res.Url
}

func GirlSolo() string {
	res := DoRequest("img/solos")
	return res.Url
}

func PussyWankGif() string {
	res := DoRequest("img/pwankg")
	return res.Url
}

func PussyArt() string {
	res := DoRequest("img/pussy_jpg")
	return res.Url
}

func LewdKemonomimi() string {
	res := DoRequest("img/lewdkemo")
	return res.Url
}

func Kitsune() string {
	res := DoRequest("img/lewdk")
	return res.Url
}

func Keta() string {
	res := DoRequest("img/keta")
	return res.Url
}

func LewdHolo() string {
	res := DoRequest("img/hololewd")
	return res.Url
}

func HoloEro() string {
	res := DoRequest("img/holoero")
	return res.Url
}

func Hentai() string {
	res := DoRequest("img/hentai")
	return res.Url
}

func Futanari() string {
	res := DoRequest("img/futanari")
	return res.Url
}

func Femdom() string {
	res := DoRequest("img/femdom")
	return res.Url
}

func FeetGif() string {
	res := DoRequest("img/feetg")
	return res.Url
}

func EroFeet() string {
	res := DoRequest("img/erofeet")
	return res.Url
}

func Feet() string {
	res := DoRequest("img/feet")
	return res.Url
}

func Ero() string {
	res := DoRequest("img/ero")
	return res.Url
}

func EroKitsune() string {
	res := DoRequest("img/erok")
	return res.Url
}

func EroKemonomimi() string {
	res := DoRequest("img/erokemo")
	return res.Url
}

func EroNeko() string {
	res := DoRequest("img/eron")
	return res.Url
}

func EroYuri() string {
	res := DoRequest("img/eroyuri")
	return res.Url
}

func CumArts() string {
	res := DoRequest("img/cum_jpg")
	return res.Url
}

func BlowJob() string {
	res := DoRequest("img/blowjob")
	return res.Url
}
func Spank() string {
	res := DoRequest("img/spank")
	return res.Url
}
func Gasm() string {
	res := DoRequest("img/gasm")
	return res.Url
}

// sfw
func Tickle() string {
	res := DoRequest("img/tickle")
	return res.Url
}

func Slap() string {
	res := DoRequest("img/slap")
	return res.Url
}

func Poke() string {
	res := DoRequest("img/poke")
	return res.Url
}

func Pat() string {
	res := DoRequest("img/pat")
	return res.Url
}

func Neko() string {
	res := DoRequest("img/neko")
	return res.Url
}

func Meow() string {
	res := DoRequest("img/meow")
	return res.Url
}

func Lizard() string {
	res := DoRequest("img/lizard")
	return res.Url
}

func Kiss() string {
	res := DoRequest("img/kiss")
	return res.Url
}

func Hug() string {
	res := DoRequest("img/hug")
	return res.Url
}

func FoxGirl() string {
	res := DoRequest("img/fox_girl")
	return res.Url
}

func Feed() string {
	res := DoRequest("img/feed")
	return res.Url
}

func Cuddle() string {
	res := DoRequest("img/cuddle")
	return res.Url
}

func Why() string {
	res := DoRequest("why")
	return res.Why
}

func CatText() string {
	res := DoRequest("cat")
	return res.Cat
}

func Owoify(text string) string {
	target := fmt.Sprintf("owoify?text=%s", url.QueryEscape(text))
	res := DoRequest(target)
	return res.Owo
}

func EightBall() (string, string) {
	res := DoRequest("8ball")
	return res.Response, res.Url
}

func Fact() string {
	res := DoRequest("fact")
	return res.Fact
}

func NekoGif() string {
	res := DoRequest("img/ngif")
	return res.Url
}

func Kemonomimi() string {
	res := DoRequest("img/kemonomimi")
	return res.Url
}

func Holo() string {
	res := DoRequest("img/holo")
	return res.Url
}

func Smug() string {
	res := DoRequest("img/smug")
	return res.Url
}

func Baka() string {
	res := DoRequest("img/baka")
	return res.Url
}

func Woof() string {
	res := DoRequest("img/woof")
	return res.Url
}

func Spoiler(text string) string {
	target := fmt.Sprintf("spoiler?text=%s", url.QueryEscape(text))
	res := DoRequest(target)
	return res.Owo
}

func Wallpaper() string {
	res := DoRequest("img/wallpaper")
	return res.Url
}

func Goose() string {
	res := DoRequest("img/goose")
	return res.Url
}

func Gecg() string {
	res := DoRequest("img/gecg")
	return res.Url
}

func Avatar() string {
	res := DoRequest("img/avatar")
	return res.Url
}

func Waifu() string {
	res := DoRequest("img/waifu")
	return res.Url
}
