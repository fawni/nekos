// Licensed under the Open Software License version 3.0

package nekos

import (
	"fmt"
	"net/url"
)

// found in https://github.com/Nekos-life/nekos-dot-life/blob/master/endpoints.json
// nsfw
func RandomHentaiGif() string {
	res := req("img/Random_hentai_gif")
	return res.Url
}

func Pussy() string {
	res := req("img/pussy")
	return res.Url
}

func LewdNekoGif() string {
	res := req("img/nsfw_neko_gif")
	return res.Url
}

func LewdNeko() string {
	res := req("img/lewd")
	return res.Url
}

func Lesbian() string {
	res := req("img/les")
	return res.Url
}

func Kuni() string {
	res := req("img/kuni")
	return res.Url
}

func Cumsluts() string {
	res := req("img/cum")
	return res.Url
}

func Classic() string {
	res := req("img/classic")
	return res.Url
}

func Boobs() string {
	res := req("img/boobs")
	return res.Url
}

func BlowJobGif() string {
	res := req("img/bj")
	return res.Url
}

func Anal() string {
	res := req("img/anal")
	return res.Url
}

func LewdAvatar() string {
	res := req("img/nsfw_avatar")
	return res.Url
}

func Yuri() string {
	res := req("img/yuri")
	return res.Url
}

func Trap() string {
	res := req("img/trap")
	return res.Url
}

func Tits() string {
	res := req("img/tits")
	return res.Url
}

func GirlSoloGif() string {
	res := req("img/solog")
	return res.Url
}

func GirlSolo() string {
	res := req("img/solos")
	return res.Url
}

func PussyWankGif() string {
	res := req("img/pwankg")
	return res.Url
}

func PussyArt() string {
	res := req("img/pussy_jpg")
	return res.Url
}

func LewdKemonomimi() string {
	res := req("img/lewdkemo")
	return res.Url
}

func Kitsune() string {
	res := req("img/lewdk")
	return res.Url
}

func Keta() string {
	res := req("img/keta")
	return res.Url
}

func LewdHolo() string {
	res := req("img/hololewd")
	return res.Url
}

func HoloEro() string {
	res := req("img/holoero")
	return res.Url
}

func Hentai() string {
	res := req("img/hentai")
	return res.Url
}

func Futanari() string {
	res := req("img/futanari")
	return res.Url
}

func Femdom() string {
	res := req("img/femdom")
	return res.Url
}

func FeetGif() string {
	res := req("img/feetg")
	return res.Url
}

func EroFeet() string {
	res := req("img/erofeet")
	return res.Url
}

func Feet() string {
	res := req("img/feet")
	return res.Url
}

func Ero() string {
	res := req("img/ero")
	return res.Url
}

func EroKitsune() string {
	res := req("img/erok")
	return res.Url
}

func EroKemonomimi() string {
	res := req("img/erokemo")
	return res.Url
}

func EroNeko() string {
	res := req("img/eron")
	return res.Url
}

func EroYuri() string {
	res := req("img/eroyuri")
	return res.Url
}

func CumArts() string {
	res := req("img/cum_jpg")
	return res.Url
}

func BlowJob() string {
	res := req("img/blowjob")
	return res.Url
}
func Spank() string {
	res := req("img/spank")
	return res.Url
}
func Gasm() string {
	res := req("img/gasm")
	return res.Url
}

// sfw
func Tickle() string {
	res := req("img/tickle")
	return res.Url
}

func Slap() string {
	res := req("img/slap")
	return res.Url
}

func Poke() string {
	res := req("img/poke")
	return res.Url
}

func Pat() string {
	res := req("img/pat")
	return res.Url
}

func Neko() string {
	res := req("img/neko")
	return res.Url
}

func Meow() string {
	res := req("img/meow")
	return res.Url
}

func Lizard() string {
	res := req("img/lizard")
	return res.Url
}

func Kiss() string {
	res := req("img/kiss")
	return res.Url
}

func Hug() string {
	res := req("img/hug")
	return res.Url
}

func FoxGirl() string {
	res := req("img/fox_girl")
	return res.Url
}

func Feed() string {
	res := req("img/feed")
	return res.Url
}

func Cuddle() string {
	res := req("img/cuddle")
	return res.Url
}

func Why() string {
	res := req("why")
	return res.Why
}

func CatText() string {
	res := req("cat")
	return res.Cat
}

func Owoify(text string) string {
	target := fmt.Sprintf("owoify?text=%s", url.QueryEscape(text))
	res := req(target)
	return res.Owo
}

func EightBall() Ball {
	res := req("8ball")
	ball := Ball{res.Response, res.Url}
	return ball
}

func Fact() string {
	res := req("fact")
	return res.Fact
}

func NekoGif() string {
	res := req("img/ngif")
	return res.Url
}

func Kemonomimi() string {
	res := req("img/kemonomimi")
	return res.Url
}

func Holo() string {
	res := req("img/holo")
	return res.Url
}

func Smug() string {
	res := req("img/smug")
	return res.Url
}

func Baka() string {
	res := req("img/baka")
	return res.Url
}

func Woof() string {
	res := req("img/woof")
	return res.Url
}

func Spoiler(text string) string {
	target := fmt.Sprintf("spoiler?text=%s", url.QueryEscape(text))
	res := req(target)
	return res.Owo
}

func Wallpaper() string {
	res := req("img/wallpaper")
	return res.Url
}

func Goose() string {
	res := req("img/goose")
	return res.Url
}

func Gecg() string {
	res := req("img/gecg")
	return res.Url
}

func Avatar() string {
	res := req("img/avatar")
	return res.Url
}

func Waifu() string {
	res := req("img/waifu")
	return res.Url
}
