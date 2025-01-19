package handler

import (
	"fmt"
	"net/http"

	"github.com/antchfx/xmlquery"
)

// the youtube XML feed we're trying to filter
const reidFeed = "https://www.youtube.com/feeds/videos.xml?channel_id=UCbdbgaxLYyqYiSyGT3uDlFg"

// the prefix for all titles that are required
const requiredTitlePrefix = "BTD6 Advanced Challenge"

func Handler(w http.ResponseWriter, r *http.Request) {
	doc, err := xmlquery.LoadURL(reidFeed)
	if err != nil {
		fmt.Println(err)
		return
	}

	entries := xmlquery.Find(doc, "//entry/title[not(starts-with(text(), '"+requiredTitlePrefix+"'))]")

	for _, entry := range entries {
		xmlquery.RemoveFromTree(entry.Parent)
	}

	fmt.Fprint(w, doc.OutputXML(true))
}
