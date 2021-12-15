// Licensed under the Open Software License version 3.0

package nekos

import (
	"fmt"
	"net/url"
)

// RandomHentaiGif returns the url of a hentai NSFW gif
func RandomHentaiGif() string {
	res := req("img/Random_hentai_gif")
	return res.Url
}

// Pussy returns the url of a pussy NSFW image/gif
func Pussy() string {
	res := req("img/pussy")
	return res.Url
}

// LewdNekoGif returns the url of a neko NSFW gif
func LewdNekoGif() string {
	res := req("img/nsfw_neko_gif")
	return res.Url
}

// LewdNeko returns the url of a neko NSFW image
func LewdNeko() string {
	res := req("img/lewd")
	return res.Url
}

// Lesbian returns the url of a lesbian NSFW image/gif
func Lesbian() string {
	res := req("img/les")
	return res.Url
}

// Kuni returns the url of a kuni NSFW image/gif
func Kuni() string {
	res := req("img/kuni")
	return res.Url
}

// Cumsluts returns the url of a cumslut NSFW image/gif
func Cumsluts() string {
	res := req("img/cum")
	return res.Url
}

// Classic returns the url of the classic endpoint NSFW image/gif
func Classic() string {
	res := req("img/classic")
	return res.Url
}

// Boobs returns the url of a boobs NSFW image/gif
func Boobs() string {
	res := req("img/boobs")
	return res.Url
}

// BJ returns the url of a bj NSFW image/gif
func BJ() string {
	res := req("img/bj")
	return res.Url
}

// Anal returns the url of an anal NSFW image/gif
func Anal() string {
	res := req("img/anal")
	return res.Url
}

// LewdAvatar returns the url of an avatar NSFW image/gif
func LewdAvatar() string {
	res := req("img/nsfw_avatar")
	return res.Url
}

// Yuri returns the url of a yuri NSFW image/gif
func Yuri() string {
	res := req("img/yuri")
	return res.Url
}

// Trap returns the url of a trap NSFW image/gif
func Trap() string {
	res := req("img/trap")
	return res.Url
}

// Tits returns the url of a tits NSFW image/gif
func Tits() string {
	res := req("img/tits")
	return res.Url
}

// GirlSoloGif returns the url of a solo girl NSFW gif
func GirlSoloGif() string {
	res := req("img/solog")
	return res.Url
}

// GirlSolo returns the url of a solo girl NSFW image
func GirlSolo() string {
	res := req("img/solos")
	return res.Url
}

// PussyWankGif returns the url of a pussy masturbation NSFW gif
func PussyWankGif() string {
	res := req("img/pwankg")
	return res.Url
}

// PussyArt returns the url of a pussy art NSFW image/gif
func PussyArt() string {
	res := req("img/pussy_jpg")
	return res.Url
}

// LewdKemonomimi returns the url of a kemonomimi NSFW image/gif
func LewdKemonomimi() string {
	res := req("img/lewdkemo")
	return res.Url
}

// Kitsune returns the url of a kitsune NSFW image/gif
func Kitsune() string {
	res := req("img/lewdk")
	return res.Url
}

// Keta returns the url of a keta NSFW image/gif
func Keta() string {
	res := req("img/keta")
	return res.Url
}

// LewdHolo returns the url of a holo NSFW image/gif
func LewdHolo() string {
	res := req("img/hololewd")
	return res.Url
}

// HoloEro returns the url of a holo ero NSFW image/gif
func HoloEro() string {
	res := req("img/holoero")
	return res.Url
}

// Hentai returns the url of a hentai NSFW image/gif
func Hentai() string {
	res := req("img/hentai")
	return res.Url
}

// Futanari returns the url of a futa NSFW image/gif
func Futanari() string {
	res := req("img/futanari")
	return res.Url
}

// Femdom returns the url of a femdom NSFW image/gif
func Femdom() string {
	res := req("img/femdom")
	return res.Url
}

// FeetGif returns the url of a feet NSFW gif
func FeetGif() string {
	res := req("img/feetg")
	return res.Url
}

// EroFeet returns the url of an ero feet NSFW image/gif
func EroFeet() string {
	res := req("img/erofeet")
	return res.Url
}

// Feet returns the url of a feet NSFW image
func Feet() string {
	res := req("img/feet")
	return res.Url
}

// Ero returns the url of an ero NSFW image/gif
func Ero() string {
	res := req("img/ero")
	return res.Url
}

// EroKitsune returns the url of an ero kitsune NSFW image/gif
func EroKitsune() string {
	res := req("img/erok")
	return res.Url
}

// EroKemonomimi returns the url of an ero kemonomimi NSFW image/gif
func EroKemonomimi() string {
	res := req("img/erokemo")
	return res.Url
}

// EroNeko returns the url of an ero neko NSFW image/gif
func EroNeko() string {
	res := req("img/eron")
	return res.Url
}

// EroYuri returns the url of an ero yrui NSFW image/gif
func EroYuri() string {
	res := req("img/eroyuri")
	return res.Url
}

// CumArts returns the url of a cum arts NSFW image/gif
func CumArts() string {
	res := req("img/cum_jpg")
	return res.Url
}

// BlowJob returns the url of a blowjob NSFW image/gif
func BlowJob() string {
	res := req("img/blowjob")
	return res.Url
}

// Spank returns the url of a spank NSFW image/gif
func Spank() string {
	res := req("img/spank")
	return res.Url
}

// Gasm returns the url of a gasm NSFW image
func Gasm() string {
	res := req("img/gasm")
	return res.Url
}

// Tickle returns the url of a tickle image/gif
func Tickle() string {
	res := req("img/tickle")
	return res.Url
}

// Slap returns the url of a slap image/gif
func Slap() string {
	res := req("img/slap")
	return res.Url
}

// Poke returns the url of a poke image/gif
func Poke() string {
	res := req("img/poke")
	return res.Url
}

// Pat returns the url of a pat image/gif
func Pat() string {
	res := req("img/pat")
	return res.Url
}

// Neko returns the url of a neko image
func Neko() string {
	res := req("img/neko")
	return res.Url
}

// Meow returns the url of a cat image/gif
func Meow() string {
	res := req("img/meow")
	return res.Url
}

// Lizard returns the url of a lizard image
func Lizard() string {
	res := req("img/lizard")
	return res.Url
}

// Kiss returns the url of a kiss image/gif
func Kiss() string {
	res := req("img/kiss")
	return res.Url
}

// Hug returns the url of a hug image/gif
func Hug() string {
	res := req("img/hug")
	return res.Url
}

// FoxGirl returns the url of a fox girl image/gif
func FoxGirl() string {
	res := req("img/fox_girl")
	return res.Url
}

// Feed returns the url of a feeding image/gif
func Feed() string {
	res := req("img/feed")
	return res.Url
}

// Cuddle returns the url of a cuddle image/gif
func Cuddle() string {
	res := req("img/cuddle")
	return res.Url
}

// Why returns a question
func Why() string {
	res := req("why")
	return res.Why
}

// CatText returns a cat kaomoji
func CatText() string {
	res := req("cat")
	return res.Cat
}

// Owoify returns the owoified version of the provided string
func Owoify(text string) string {
	target := fmt.Sprintf("owoify?text=%s", url.QueryEscape(text))
	res := req(target)
	return res.Owo
}

// EightBall returns an Ball struct that contains the response of a magic 8ball and a corresponding image
func EightBall() Ball {
	res := req("8ball")
	ball := Ball{res.Response, res.Url}
	return ball
}

// Fact returns a fact
func Fact() string {
	res := req("fact")
	return res.Fact
}

// NekoGif returns the url of a neko gif
func NekoGif() string {
	res := req("img/ngif")
	return res.Url
}

// Kemonomimi returns the url of a kemonomimi image/gif
func Kemonomimi() string {
	res := req("img/kemonomimi")
	return res.Url
}

// Holo returns the url of a holo image/gif
func Holo() string {
	res := req("img/holo")
	return res.Url
}

// Smug returns the url of a smug image/gif
func Smug() string {
	res := req("img/smug")
	return res.Url
}

// Baka returns the url of a baka image/gif
func Baka() string {
	res := req("img/baka")
	return res.Url
}

// Woof returns the url of a dog image/gif
func Woof() string {
	res := req("img/woof")
	return res.Url
}

// Spoiler returns a spoiler per letter for discord markdown version of the provided string
func Spoiler(text string) string {
	target := fmt.Sprintf("spoiler?text=%s", url.QueryEscape(text))
	res := req(target)
	return res.Owo
}

// Wallpaper returns the url of a wallpaper
func Wallpaper() string {
	res := req("img/wallpaper")
	return res.Url
}

// Goose returns the url of a goose image
func Goose() string {
	res := req("img/goose")
	return res.Url
}

// Gecg returns the url of a genetically engineered catgirl image
func Gecg() string {
	res := req("img/gecg")
	return res.Url
}

// Avatar returns the url of an avatar image
func Avatar() string {
	res := req("img/avatar")
	return res.Url
}

// Waifu returns the url of a waifu image
func Waifu() string {
	res := req("img/waifu")
	return res.Url
}
