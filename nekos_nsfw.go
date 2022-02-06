// Licensed under the Open Software License version 3.0

package nekos

// RandomHentaiGif returns the url of a hentai NSFW gif
func RandomHentaiGif() (string, error) {
	res, err := req("img/Random_hentai_gif")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Pussy returns the url of a pussy NSFW image/gif
func Pussy() (string, error) {
	res, err := req("img/pussy")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// LewdNekoGif returns the url of a neko NSFW gif
func LewdNekoGif() (string, error) {
	res, err := req("img/nsfw_neko_gif")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// LewdNeko returns the url of a neko NSFW image
func LewdNeko() (string, error) {
	res, err := req("img/lewd")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Lesbian returns the url of a lesbian NSFW image/gif
func Lesbian() (string, error) {
	res, err := req("img/les")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Kuni returns the url of a kuni NSFW image/gif
func Kuni() (string, error) {
	res, err := req("img/kuni")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Cumsluts returns the url of a cumslut NSFW image/gif
func Cumsluts() (string, error) {
	res, err := req("img/cum")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Classic returns the url of the classic endpoint NSFW image/gif
func Classic() (string, error) {
	res, err := req("img/classic")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Boobs returns the url of a boobs NSFW image/gif
func Boobs() (string, error) {
	res, err := req("img/boobs")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// BlowJobGif returns the url of a bj NSFW gif
func BlowJobGif() (string, error) {
	res, err := req("img/bj")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Anal returns the url of an anal NSFW image/gif
func Anal() (string, error) {
	res, err := req("img/anal")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// LewdAvatar returns the url of an avatar NSFW image/gif
func LewdAvatar() (string, error) {
	res, err := req("img/nsfw_avatar")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Yuri returns the url of a yuri NSFW image/gif
func Yuri() (string, error) {
	res, err := req("img/yuri")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Trap returns the url of a trap NSFW image/gif
func Trap() (string, error) {
	res, err := req("img/trap")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Tits returns the url of a tits NSFW image/gif
func Tits() (string, error) {
	res, err := req("img/tits")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// GirlSoloGif returns the url of a solo girl NSFW gif
func GirlSoloGif() (string, error) {
	res, err := req("img/solog")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// GirlSolo returns the url of a solo girl NSFW image
func GirlSolo() (string, error) {
	res, err := req("img/solos")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// PussyWankGif returns the url of a pussy masturbation NSFW gif
func PussyWankGif() (string, error) {
	res, err := req("img/pwankg")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// PussyArt returns the url of a pussy art NSFW image/gif
func PussyArt() (string, error) {
	res, err := req("img/pussy_jpg")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// LewdKemonomimi returns the url of a kemonomimi NSFW image/gif
func LewdKemonomimi() (string, error) {
	res, err := req("img/lewdkemo")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Kitsune returns the url of a kitsune NSFW image/gif
func Kitsune() (string, error) {
	res, err := req("img/lewdk")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Keta returns the url of a keta NSFW image/gif
func Keta() (string, error) {
	res, err := req("img/keta")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// LewdHolo returns the url of a holo NSFW image/gif
func LewdHolo() (string, error) {
	res, err := req("img/hololewd")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// HoloEro returns the url of a holo ero NSFW image/gif
func HoloEro() (string, error) {
	res, err := req("img/holoero")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Hentai returns the url of a hentai NSFW image/gif
func Hentai() (string, error) {
	res, err := req("img/hentai")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Futanari returns the url of a futa NSFW image/gif
func Futanari() (string, error) {
	res, err := req("img/futanari")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Femdom returns the url of a femdom NSFW image/gif
func Femdom() (string, error) {
	res, err := req("img/femdom")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// FeetGif returns the url of a feet NSFW gif
func FeetGif() (string, error) {
	res, err := req("img/feetg")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// EroFeet returns the url of an ero feet NSFW image/gif
func EroFeet() (string, error) {
	res, err := req("img/erofeet")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Feet returns the url of a feet NSFW image
func Feet() (string, error) {
	res, err := req("img/feet")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Ero returns the url of an ero NSFW image/gif
func Ero() (string, error) {
	res, err := req("img/ero")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// EroKitsune returns the url of an ero kitsune NSFW image/gif
func EroKitsune() (string, error) {
	res, err := req("img/erok")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// EroKemonomimi returns the url of an ero kemonomimi NSFW image/gif
func EroKemonomimi() (string, error) {
	res, err := req("img/erokemo")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// EroNeko returns the url of an ero neko NSFW image/gif
func EroNeko() (string, error) {
	res, err := req("img/eron")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// EroYuri returns the url of an ero yrui NSFW image/gif
func EroYuri() (string, error) {
	res, err := req("img/eroyuri")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// CumArts returns the url of a cum arts NSFW image/gif
func CumArts() (string, error) {
	res, err := req("img/cum_jpg")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// BlowJob returns the url of a blowjob NSFW image
func BlowJob() (string, error) {
	res, err := req("img/blowjob")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Spank returns the url of a spank NSFW image/gif
func Spank() (string, error) {
	res, err := req("img/spank")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}

// Gasm returns the url of a gasm NSFW image
func Gasm() (string, error) {
	res, err := req("img/gasm")
	if err != nil {
		return "", err
	}
	return res.Url, nil
}
