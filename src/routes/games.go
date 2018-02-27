package routes

import (
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"context"
	"ppio-web/models"
	"reflect"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
)

func getGamesHandler(client *elastic.Client) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var games []byte
		ctx := context.Background()
		docMatchQuery := elastic.NewTermQuery("_type", "doc")
		scroller := client.
			Scroll().
			Index("games").
			Type("doc").
			Query(docMatchQuery).
			Sort("time", true).
			Size(1)

		docs := 0
		for {
			res, err := scroller.Do(ctx)
			if err == io.EOF {
				// No more documents
				break
			}
			for _, hit := range res.Hits.Hits {
				item := make(map[string]interface{})
				_ := json.Unmarshal(*hit.Source, &item)
				docs++
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(games)
	}

	return http.HandlerFunc(fn)
}
