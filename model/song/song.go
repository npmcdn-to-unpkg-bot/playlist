package song

import (
	//"encoding/json"
	"fmt"
	"github.com/sath33sh/infra/db"
	"github.com/sath33sh/infra/log"
	"github.com/sath33sh/infra/util"
	"github.com/sath33sh/pattern/graph"
)

// Module name.
const MODULE = "song"

// Document type for song.
const OBJ_SONG db.ObjType = "song"

// Song.
type Song struct {
	graph.Node
	Artist  graph.Node `json:"artist,omitempty"`  // Artist.
	Rank    int        `json:"rank,omitempty"`    // Rank.
	Price   string     `json:"price,omitempty"`   // Price.
	Deleted bool       `json:"deleted,omitempty"` // Delete marker.
}

// Database interface methods.
func (s *Song) GetMeta() db.ObjMeta {
	return db.ObjMeta{
		Bucket: db.DEFAULT_BUCKET,
		Type:   OBJ_SONG,
		Id:     string(s.Id),
	}
}

func (s *Song) SetType() {
	s.Type = OBJ_SONG
}

func (s *Song) Validate() (err error) {
	if s.Name == "" {
		log.Errorf("Invalid name")
		return util.ErrInvalidInput
	}

	return nil
}

func (s *Song) Create() (err error) {
	if err = s.Validate(); err != nil {
		return err
	}

	if s.Id, err = graph.NewNodeId(); err != nil {
		log.Errorf("Failed to allocate ID: %v", err)
		return err
	}

	if err = db.Upsert(s, 0); err != nil {
		return err
	}

	return nil
}

func (s *Song) Update() (err error) {
	if err = s.Validate(); err != nil {
		return err
	}

	var ls Song
	ls.Id = s.Id
	lock, err := db.GetLock(&ls)
	if err != nil {
		log.Errorf("GetLock() failed: song %s: %v", s.Id, err)
		return err
	}

	ls.Name = s.Name

	db.WriteUnlock(&ls, lock, 0)

	return nil
}

func (s *Song) Delete() (err error) {
	var ls Song
	ls.Id = s.Id
	lock, err := db.GetLock(&ls)
	if err != nil {
		return err
	}

	ls.Deleted = true

	db.WriteUnlock(&ls, lock, 0)

	return nil
}

func (s *Song) Show() (err error) {
	err = db.Get(s)
	if err != nil {
		return err
	}

	if s.Deleted {
		return util.ErrNotFound
	}

	return nil
}

// Song query result.
type SongQueryResult struct {
	Results    []Song `json:"results"`    // Results is a list of songs.
	NextOffset string `json:"nextOffset"` // Next offset.
	PrevOffset string `json:"prevOffset"` // Previous offset.
}

func (qr *SongQueryResult) GetRowPtr(index int) interface{} {
	if index < len(qr.Results) {
		return &qr.Results[index]
	} else if index == len(qr.Results) {
		qr.Results = append(qr.Results, Song{})
		return &qr.Results[index]
	} else {
		return nil
	}
}

// Query songs.
func (qr *SongQueryResult) Query(filterStmt string, limit, offset int) (size int, err error) {
	// N1QL query statement.
	queryStmt := fmt.Sprintf("SELECT `%s`.* FROM `%s` WHERE `type`=\"%s\" AND `deleted` IS MISSING %s",
		db.BucketName(db.DEFAULT_BUCKET), db.BucketName(db.DEFAULT_BUCKET), OBJ_SONG, filterStmt)

	size, err = db.ExecPagedQuery(db.DEFAULT_BUCKET, qr, queryStmt, limit, offset)
	if err != nil {
		return size, err
	}

	qr.Results = qr.Results[:size]
	qr.PrevOffset = fmt.Sprintf("%d", offset)
	qr.NextOffset = fmt.Sprintf("%d", offset+size)

	return size, err
}

// Count.
func Count(filterStmt string) (count int, err error) {
	// N1QL query statement.
	queryStmt := fmt.Sprintf("SELECT COUNT(*) FROM `%s` WHERE `type`=\"%s\" AND `deleted` IS MISSING %s",
		db.BucketName(db.DEFAULT_BUCKET), OBJ_SONG, filterStmt)

	count, err = db.ExecCount(db.DEFAULT_BUCKET, queryStmt)
	if err != nil {
		log.Errorf("Failed to get count: %v", err)
		return 0, err
	}

	return count, nil
}

// Song iterator.
func ForEach(filterStmt string, cb func(*Song)) {
	var err error

	// N1QL query statement.
	queryStmt := fmt.Sprintf("SELECT `%s`.* FROM `%s` WHERE `type`=\"%s\" AND `deleted` IS MISSING %s",
		db.BucketName(db.DEFAULT_BUCKET), db.BucketName(db.DEFAULT_BUCKET), OBJ_SONG, filterStmt)

	size := db.QUERY_LIMIT_MAX
	offset := 0
	for size == db.QUERY_LIMIT_MAX {
		var qr SongQueryResult
		size, err = db.ExecPagedQuery(db.DEFAULT_BUCKET, &qr, queryStmt, size, offset)
		if err != nil {
			return
		}

		for index := 0; index < size; index++ {
			cb(&qr.Results[index])
		}
		offset += size
	}
}
