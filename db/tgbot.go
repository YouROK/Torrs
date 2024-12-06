package db

import (
	bolt "go.etcd.io/bbolt"
)

func SaveTGFileID(fileID, tgID string) {
	db.Update(func(tx *bolt.Tx) error {
		ids, err := tx.CreateBucketIfNotExists([]byte("TGBotFileIDs"))
		if err != nil {
			return err
		}

		return ids.Put([]byte(fileID), []byte(tgID))
	})
}

func GetTGFileID(id string) string {
	ret := ""
	db.View(func(tx *bolt.Tx) error {
		ids := tx.Bucket([]byte("TGBotFileIDs"))
		if ids != nil {
			if b := ids.Get([]byte(id)); b != nil {
				ret = string(b)
			}
		}
		return nil
	})
	return ret
}
