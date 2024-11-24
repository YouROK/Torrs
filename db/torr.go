package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"torrsru/db/utils"
	"torrsru/models/fdb"
)

func saveTorrent(parent *bolt.Bucket, torr *fdb.Torrent) error {
	torrB, err := parent.CreateBucketIfNotExists(torr.GetUnique())
	if err != nil {
		return err
	}

	// Записываем основные поля
	err = torrB.Put([]byte("size"), utils.I2B(torr.Size))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("quality"), utils.I2B(torr.Quality))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("videotype"), []byte(torr.Videotype))
	if err != nil {
		return err
	}

	// Записываем Voices в отдельный бакет
	if len(torr.Voices) > 0 {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err = enc.Encode(torr.Voices)
		if err != nil {
			return err
		}
		err = torrB.Put([]byte("voices"), buf.Bytes())
		if err != nil {
			return err
		}
	}

	// Аналогично для сезонов
	if len(torr.Seasons) > 0 {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err = enc.Encode(torr.Seasons)
		if err != nil {
			return err
		}
		err = torrB.Put([]byte("seasons"), buf.Bytes())
		if err != nil {
			return err
		}
	}

	err = torrB.Put([]byte("trackerName"), []byte(torr.TrackerName))
	if err != nil {
		return err
	}

	// Сохраняем Types
	if len(torr.Types) > 0 {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err = enc.Encode(torr.Types)
		if err != nil {
			return err
		}
		err = torrB.Put([]byte("types"), buf.Bytes())
		if err != nil {
			return err
		}
	}

	err = torrB.Put([]byte("url"), []byte(torr.URL))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("title"), []byte(torr.Title))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("sid"), utils.I2B(torr.Sid))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("pir"), utils.I2B(torr.Pir))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("sizeName"), []byte(torr.SizeName))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("createTime"), []byte(torr.CreateTime))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("updateTime"), []byte(torr.UpdateTime))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("checkTime"), []byte(torr.CheckTime))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("magnet"), []byte(torr.Magnet))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("name"), []byte(torr.Name))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("originalname"), []byte(torr.Originalname))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("relased"), utils.I2B(torr.Relased))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("ffprobe_tryingdata"), utils.I2B(torr.FFProbeTryingdata))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("_sn"), []byte(torr.Sn))
	if err != nil {
		return err
	}

	err = torrB.Put([]byte("_so"), []byte(torr.So))
	if err != nil {
		return err
	}

	// Записываем Languages в отдельный бакет
	if len(torr.Languages) > 0 {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err = enc.Encode(torr.Languages)
		if err != nil {
			return err
		}
		err = torrB.Put([]byte("languages"), buf.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}

func readTorrent(parent *bolt.Bucket, uniqueID []byte) (*fdb.Torrent, error) {
	torrB := parent.Bucket(uniqueID)
	if torrB == nil {
		return nil, fmt.Errorf("bucket not found for unique ID: %s", uniqueID)
	}

	torr := &fdb.Torrent{}

	// Чтение основных полей
	sizeBytes := torrB.Get([]byte("size"))
	if sizeBytes != nil {
		torr.Size = utils.B2I(sizeBytes)
	}

	qualityBytes := torrB.Get([]byte("quality"))
	if qualityBytes != nil {
		torr.Quality = utils.B2I(qualityBytes)
	}

	videotypeBytes := torrB.Get([]byte("videotype"))
	if videotypeBytes != nil {
		torr.Videotype = string(videotypeBytes)
	}

	// Чтение Voices
	voicesBytes := torrB.Get([]byte("voices"))
	if voicesBytes != nil {
		buf := bytes.NewBuffer(voicesBytes)
		dec := gob.NewDecoder(buf)
		err := dec.Decode(&torr.Voices)
		if err != nil {
			return nil, err
		}
	}

	// Чтение Seasons
	seasonsB := torrB.Get([]byte("seasons"))
	if seasonsB != nil {
		buf := bytes.NewBuffer(seasonsB)
		dec := gob.NewDecoder(buf)
		err := dec.Decode(&torr.Seasons)
		if err != nil {
			return nil, err
		}
	}

	// Чтение остальных полей
	torr.TrackerName = string(torrB.Get([]byte("trackerName")))

	// Чтение Types
	typesB := torrB.Get([]byte("types"))
	if typesB != nil {
		buf := bytes.NewBuffer(typesB)
		dec := gob.NewDecoder(buf)
		err := dec.Decode(&torr.Types)
		if err != nil {
			return nil, err
		}
	}

	torr.URL = string(torrB.Get([]byte("url")))
	torr.Title = string(torrB.Get([]byte("title")))

	sidBytes := torrB.Get([]byte("sid"))
	if sidBytes != nil {
		torr.Sid = utils.B2I(sidBytes)
	}

	pirBytes := torrB.Get([]byte("pir"))
	if pirBytes != nil {
		torr.Pir = utils.B2I(pirBytes)
	}

	torr.SizeName = string(torrB.Get([]byte("sizeName")))
	torr.CreateTime = string(torrB.Get([]byte("createTime")))
	torr.UpdateTime = string(torrB.Get([]byte("updateTime")))
	torr.CheckTime = string(torrB.Get([]byte("checkTime")))
	torr.Magnet = string(torrB.Get([]byte("magnet")))
	torr.Name = string(torrB.Get([]byte("name")))
	torr.Originalname = string(torrB.Get([]byte("originalname")))

	releasedBytes := torrB.Get([]byte("relased"))
	if releasedBytes != nil {
		torr.Relased = utils.B2I(releasedBytes)
	}

	ffprobeDataBytes := torrB.Get([]byte("ffprobe_tryingdata"))
	if ffprobeDataBytes != nil {
		torr.FFProbeTryingdata = utils.B2I(ffprobeDataBytes)
	}

	torr.Sn = string(torrB.Get([]byte("_sn")))
	torr.So = string(torrB.Get([]byte("_so")))

	// Чтение Languages
	languagesB := torrB.Get([]byte("languages"))
	if languagesB != nil {
		buf := bytes.NewBuffer(languagesB)
		dec := gob.NewDecoder(buf)
		err := dec.Decode(&torr.Languages)
		if err != nil {
			return nil, err
		}
	}

	return torr, nil
}

func combineTorrents(torrents []*fdb.Torrent) *fdb.Torrent {
	if len(torrents) == 0 {
		return &fdb.Torrent{}
	}

	if len(torrents) == 1 {
		return torrents[0]
	}

	combined := torrents[0]

	for _, t := range torrents[1:] {
		if t.Size > combined.Size {
			combined.Size = t.Size
		}
		if t.Quality > combined.Quality {
			combined.Quality = t.Quality
		}
		if len(t.Videotype) > len(combined.Videotype) {
			combined.Videotype = t.Videotype
		}
		combined.Voices = uniqueStringArray(append(combined.Voices, t.Voices...))
		combined.Seasons = uniqueInt64Array(append(combined.Seasons, t.Seasons...))
		if len(t.TrackerName) > len(combined.TrackerName) {
			combined.TrackerName = t.TrackerName
		}
		combined.Types = uniqueStringArray(append(combined.Types, t.Types...))
		if len(t.URL) > len(combined.URL) {
			combined.URL = t.URL
		}
		if len(t.Title) > len(combined.Title) {
			combined.Title = t.Title
		}
		if t.Sid > combined.Sid {
			combined.Sid = t.Sid
		}
		if t.Pir > combined.Pir {
			combined.Pir = t.Pir
		}
		if len(t.SizeName) > len(combined.SizeName) {
			combined.SizeName = t.SizeName
		}
		if len(t.CreateTime) > len(combined.CreateTime) {
			combined.CreateTime = t.CreateTime
		}
		if len(t.Magnet) > len(combined.Magnet) {
			combined.Magnet = t.Magnet
		}
		if len(t.Name) > len(combined.Name) {
			combined.Name = t.Name
		}
		if len(t.Originalname) > len(combined.Originalname) {
			combined.Originalname = t.Originalname
		}
		if t.Relased > combined.Relased {
			combined.Relased = t.Relased
		}
		combined.Languages = uniqueStringArray(append(combined.Languages, t.Languages...))
	}

	return combined
}

// Функция для удаления дубликатов из массива строк
func uniqueStringArray(arr []string) []string {
	uniqueMap := make(map[string]struct{})
	uniqueArr := []string{}

	for _, str := range arr {
		if _, found := uniqueMap[str]; !found {
			uniqueMap[str] = struct{}{}
			uniqueArr = append(uniqueArr, str)
		}
	}
	return uniqueArr
}

// Функция для удаления дубликатов из массива int64
func uniqueInt64Array(arr []int64) []int64 {
	uniqueMap := make(map[int64]struct{})
	uniqueArr := []int64{}

	for _, val := range arr {
		if _, found := uniqueMap[val]; !found {
			uniqueMap[val] = struct{}{}
			uniqueArr = append(uniqueArr, val)
		}
	}
	return uniqueArr
}
