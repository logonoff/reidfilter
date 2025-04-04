package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/antchfx/xmlquery"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	doc, err := xmlquery.LoadURL(os.Getenv("REID_FEED"))
	if err != nil {
		fmt.Println(err)
		return
	}

	link := xmlquery.Find(doc, "//entry/title[starts-with(text(), '"+os.Getenv("VIDEO_TITLE_PREFIX")+"')]")[0].Parent.SelectElements("link")[0].SelectAttr("href")

	// return http 307 with link
	http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}
