package db

import (
	"encoding/hex"
	"fmt"
	"github.com/blevesearch/bleve"
	bolt "go.etcd.io/bbolt"
	"log"
	"path/filepath"
	"torrsru/models/fdb"
	"torrsru/web/global"
)

var (
	index bleve.Index
)

func initIndex() error {
	mappings := bleve.NewIndexMapping()
	var err error
	index, err = bleve.Open(filepath.Join(global.PWD, "index.db"))
	if err != nil {
		index, err = bleve.NewUsing(filepath.Join(global.PWD, "index.db"), mappings, "scorch", "scorch", nil)
	}
	return err
}

func RebuildIndex() error {
	var err error
	mappings := bleve.NewIndexMapping()
	index, err = bleve.NewUsing(filepath.Join(global.PWD, "index.db"), mappings, "scorch", "scorch", nil)
	if err != nil {
		return err
	}
	err = db.View(func(tx *bolt.Tx) error {
		torrsB := tx.Bucket([]byte("Torrents"))
		if torrsB == nil {
			return nil
		}

		batch := index.NewBatch()
		err := torrsB.ForEach(func(uniTorr, _ []byte) error {
			torrB := torrsB.Bucket(uniTorr)
			if torrB == nil {
				return fmt.Errorf("Error in db struct")
			}

			title := string(torrB.Get([]byte("title")))

			if title != "" {
				if batch.Size() >= 100000 {
					err := index.Batch(batch)
					if err != nil {
						return err
					}
					log.Println("Indexed torrents:", batch.Size())
					batch = index.NewBatch()
				} else {
					err := batch.Index(hex.EncodeToString(uniTorr), title)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if batch.Size() > 0 {
			err = index.Batch(batch)
			if err != nil {
				return err
			}
			log.Println("Indexed torrents:", batch.Size())
		}
		return err
	})
	return err
}

func Search(query string) ([]*fdb.Torrent, error) {
	q := bleve.NewMatchQuery(query)
	searchRequest := bleve.NewSearchRequest(q)
	searchRequest.Size = 1000
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	var list []*fdb.Torrent
	err = db.View(func(tx *bolt.Tx) error {
		torrsB := tx.Bucket([]byte("Torrents"))
		if torrsB == nil {
			return nil
		}
		for _, hit := range searchResult.Hits {
			if hit.Score > 2.0 {
				buf, err := hex.DecodeString(hit.ID)
				if err != nil {
					return err
				}
				torr, err := readTorrent(torrsB, buf)
				if err != nil {
					return err
				}
				list = append(list, torr)
			}
		}
		return nil
	})
	return list, err
}
