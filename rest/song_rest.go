package rest

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sath33sh/infra/db"
	"github.com/sath33sh/infra/log"
	"github.com/sath33sh/infra/util"
	"github.com/sath33sh/infra/wapi"
	"github.com/sath33sh/pattern/graph"
	"github.com/sath33sh/tunes/model/song"
	"net/http"
)

const MODULE = "rest"

func parseSongRequest(
	s *song.Song,
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
	decodeJson bool,
	validateId bool) (err error) {

	// Decode JSON input.
	if decodeJson {
		if err = wapi.DecodeJSON(r, s); err != nil {
			log.Errorf("Failed to decode song JSON: %v", err)
			return util.ErrJsonDecode
		}
	}

	// Get ID from params, and validate if required.
	s.Id = graph.NodeId(params.ByName("id"))
	if validateId && len(s.Id) == 0 {
		log.Errorf("Invalid song ID")
		return util.ErrInvalidInput
	}

	return nil
}

func CreateSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var s song.Song
	var err error

	if err = parseSongRequest(&s, w, r, params, true, false); err != nil {
		wapi.ReturnError(w, r, err)
		return
	}

	// Create.
	if err = s.Create(); err != nil {
		log.Errorf("Failed to create song %s: %v", s.Name, err)
		wapi.ReturnError(w, r, err)
	} else {
		log.Debugf(MODULE, "Created song %s", s.Id)
		wapi.ReturnOk(w, r, &s)
	}
}

func ShowSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var s song.Song
	var err error

	if err = parseSongRequest(&s, w, r, params, false, true); err != nil {
		wapi.ReturnError(w, r, err)
		return
	}

	// Get.
	if err = s.Show(); err != nil {
		log.Errorf("Failed to get song %s: %v", s.Id, err)
		wapi.ReturnError(w, r, err)
	} else {
		log.Debugf(MODULE, "Show song %s", s.Id)
		wapi.ReturnOk(w, r, &s)
	}
}

func UpdateSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var s song.Song
	var err error

	if err = parseSongRequest(&s, w, r, params, true, true); err != nil {
		wapi.ReturnError(w, r, err)
		return
	}

	// Get.
	if err = s.Update(); err != nil {
		log.Errorf("Failed to update song %s: %v", s.Id, err)
		wapi.ReturnError(w, r, err)
	} else {
		log.Debugf(MODULE, "Updated song %s", s.Id)
		wapi.ReturnOk(w, r, &s)
	}
}

func DeleteSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var s song.Song
	var err error

	if err = parseSongRequest(&s, w, r, params, false, true); err != nil {
		wapi.ReturnError(w, r, err)
		return
	}

	// Get.
	if err = s.Delete(); err != nil {
		log.Errorf("Failed to delete song %s: %v", s.Id, err)
		wapi.ReturnError(w, r, err)
	} else {
		log.Debugf(MODULE, "Deleted song %s", s.Id)
		wapi.ReturnOk(w, r, &s)
	}
}

func ListSongs(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var qr song.SongQueryResult
	var err error

	// Parse limit and offset.
	limit, offset, err := db.ParsePageArgs(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	if err != nil {
		wapi.ReturnError(w, r, err)
		return
	}

	// Query songs.
	var size int
	size, err = qr.Query("", limit, offset)
	if err != nil {
		log.Errorf("Query songs failed: %v", err)
		wapi.ReturnError(w, r, err)
	} else {
		log.Debugf(MODULE, "List %d songs", size)
		wapi.ReturnOk(w, r, &qr)
	}
}

func InitSong() {
	// Debug enable.
	log.EnableDebug(MODULE)

	// Register REST callbacks.
	wapi.POST("/v1.0/song/create", CreateSong)
	wapi.GET("/v1.0/song/show/:id", ShowSong)
	wapi.POST("/v1.0/song/update/:id", UpdateSong)
	wapi.POST("/v1.0/song/delete/:id", DeleteSong)
	wapi.GET("/v1.0/song/list", ListSongs)
}
