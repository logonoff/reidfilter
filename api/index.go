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

	entries := xmlquery.Find(doc, "//entry/title[not(starts-with(text(), '"+os.Getenv("VIDEO_TITLE_PREFIX")+"'))]")

	for _, entry := range entries {
		xmlquery.RemoveFromTree(entry.Parent)
	}

	fmt.Fprint(w, doc.OutputXML(true))
}
