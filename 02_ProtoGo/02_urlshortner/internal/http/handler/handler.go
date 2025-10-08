package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/02_protogo/02_urlshortner/internal/config"
	inmemdb "github.com/02_protogo/02_urlshortner/internal/db/Inmemdb"
	"github.com/02_protogo/02_urlshortner/internal/shortner"
	"github.com/02_protogo/02_urlshortner/internal/types"
)

func CreateShortUrl(cfg *config.Config) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var req types.CreateShortUrlRequest;

		err := json.NewDecoder(r.Body).Decode(&req);
		if err != nil {
			w.WriteHeader(http.StatusBadRequest);
			w.Write([]byte("Invalid JSON BODy"));
			return;
		}

		if req.URL==""{
			w.WriteHeader(http.StatusBadRequest);
			w.Write([]byte("EmptyJSON BODy"));
			return;
		}

		sortUrlID, err := shortner.GenerateString();
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError);
			w.Write([]byte("Server unable to generate Short url"));
			return;
		}
		inmemdb.DbMapMutex.Lock();
		defer inmemdb.DbMapMutex.Unlock();

		inmemdb.Dbmap[sortUrlID] = types.UrlNode{
			ID:          sortUrlID,
			OriginalUrl: req.URL,
			ShortUrl:    sortUrlID,
		}

		res:=struct{
			ShortUrl string `json:"shorturl"`
		}{
			ShortUrl: fmt.Sprintf("http://%s:%d/%s",cfg.ServerIp,cfg.Port,sortUrlID),
		}
		
		w.Header().Set("Content-Type","application/json");
		w.WriteHeader(http.StatusOK);
		json.NewEncoder(w).Encode(&res);

	})
}


func RedirectHits(w http.ResponseWriter,r *http.Request){
	id:=r.PathValue("id");

	inmemdb.DbMapMutex.Lock();
	defer inmemdb.DbMapMutex.Unlock();

	node,ok:=inmemdb.Dbmap[id];

	if !ok{
		http.NotFound(w,r);
	}
	
	http.Redirect(w,r,node.OriginalUrl,http.StatusFound);

}
