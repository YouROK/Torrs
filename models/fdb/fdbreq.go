package fdb

import (
	"crypto/md5"
)

type FDBRequest struct {
	Nextread    bool          `json:"nextread"`
	Countread   int64         `json:"countread"`
	Take        int64         `json:"take"`
	Collections []*Collection `json:"collections"`
}

type Collection struct {
	Key   string `json:"Key"`
	Value Value  `json:"Value"`
}

type Value struct {
	Time     string              `json:"time"`
	FileTime int64               `json:"fileTime"`
	Torrents map[string]*Torrent `json:"torrents"`
}

type Torrent struct {
	Size              int64    `json:"size"`
	Quality           int64    `json:"quality"`
	Videotype         string   `json:"videotype"`
	Voices            []string `json:"voices"`
	Seasons           []int64  `json:"seasons"`
	TrackerName       string   `json:"trackerName"`
	Types             []string `json:"types"`
	URL               string   `json:"url"`
	Title             string   `json:"title"`
	Sid               int64    `json:"sid"`
	Pir               int64    `json:"pir"`
	SizeName          string   `json:"sizeName"`
	CreateTime        string   `json:"createTime"`
	UpdateTime        string   `json:"updateTime"`
	CheckTime         string   `json:"checkTime"`
	Magnet            string   `json:"magnet"`
	Name              string   `json:"name"`
	Originalname      string   `json:"originalname,omitempty"`
	Relased           int64    `json:"relased"`
	FFProbeTryingdata int64    `json:"ffprobe_tryingdata"`
	Sn                string   `json:"_sn"`
	So                string   `json:"_so,omitempty"`
	Languages         []string `json:"languages,omitempty"`
}

func (t *Torrent) GetUnique() []byte {
	hash := md5.Sum([]byte(t.Title + t.Magnet + t.TrackerName))
	return hash[:]
}
