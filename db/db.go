package db

import (
	"encoding/binary"
	"encoding/hex"
	bolt "go.etcd.io/bbolt"
	"log"
	"path/filepath"
	"regexp"
	"time"
	"torrsru/models/fdb"
	"torrsru/web/global"
)

var (
	db *bolt.DB
	re = regexp.MustCompile(`.+((19|20)\d\d)`)
)

func Init() {
	d, err := bolt.Open(filepath.Join(global.PWD, "torrents.db"), 0o666, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		log.Fatalln("Error open db", err)
		return
	}
	db = d

	err = initIndex()
	if err != nil {
		log.Fatalln("Error open index", err)
		return
	}
}

func GetFileTime() int64 {
	var ft int64 = -1
	err := db.View(func(tx *bolt.Tx) error {
		sets := tx.Bucket([]byte("Settings"))
		if sets == nil {
			return nil
		}
		b := sets.Get([]byte("FileTime"))
		if b != nil {
			ft = int64(binary.LittleEndian.Uint64(b))
		}
		return nil
	})
	if err != nil {
		log.Println("Error get from db:", err)
	}
	return ft
}

func SetFileTime(fileTime int64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(fileTime))
	return db.Update(func(tx *bolt.Tx) error {
		sets, err := tx.CreateBucketIfNotExists([]byte("Settings"))
		if err != nil {
			return err
		}
		return sets.Put([]byte("FileTime"), b)
	})
}

func saveTorrents(cols []*fdb.Collection) error {
	return db.Update(func(tx *bolt.Tx) error {
		torrsb, err := tx.CreateBucketIfNotExists([]byte("Torrents"))
		if err != nil {
			return err
		}

		//save and index torrents
		batch := index.NewBatch()
		for _, col := range cols {
			for _, torr := range col.Value.Torrents {
				hash := torr.GetUnique()
				tb := torrsb.Bucket(hash)
				if tb != nil {
					ttmp, err := readTorrent(torrsb, hash)
					if err != nil {
						return err
					}
					torr = combineTorrents([]*fdb.Torrent{ttmp, torr})
				}
				err = saveTorrent(torrsb, torr)
				if err != nil {
					return err
				}

				err = batch.Index(hex.EncodeToString(hash), torr.Title)
				if err != nil {
					return err
				}
			}
		}
		return index.Batch(batch)
	})
}
