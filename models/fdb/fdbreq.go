package fdb

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
	Size              int64     `json:"size"`
	Quality           int64     `json:"quality"`
	Videotype         string    `json:"videotype"`
	Voices            []string  `json:"voices"`
	Seasons           []int64   `json:"seasons"`
	TrackerName       string    `json:"trackerName"`
	Types             []string  `json:"types"`
	URL               string    `json:"url"`
	Title             string    `json:"title"`
	Sid               int64     `json:"sid"`
	Pir               int64     `json:"pir"`
	SizeName          string    `json:"sizeName"`
	CreateTime        string    `json:"createTime"`
	UpdateTime        string    `json:"updateTime"`
	CheckTime         string    `json:"checkTime"`
	Magnet            string    `json:"magnet"`
	Name              string    `json:"name"`
	Originalname      *string   `json:"originalname,omitempty"`
	Relased           int64     `json:"relased"`
	FFProbeTryingdata int64     `json:"ffprobe_tryingdata"`
	Sn                string    `json:"_sn"`
	So                *string   `json:"_so,omitempty"`
	FFProbe           []FFProbe `json:"ffprobe,omitempty"`
	Languages         []string  `json:"languages,omitempty"`
}

type FFProbe struct {
	Index         int64   `json:"index"`
	CodecName     *string `json:"codec_name,omitempty"`
	CodecLongName *string `json:"codec_long_name,omitempty"`
	CodecType     string  `json:"codec_type"`
	Width         *int64  `json:"width,omitempty"`
	Height        *int64  `json:"height,omitempty"`
	CodedWidth    *int64  `json:"coded_width,omitempty"`
	CodedHeight   *int64  `json:"coded_height,omitempty"`
	BitRate       *string `json:"bit_rate,omitempty"`
	SampleFmt     *string `json:"sample_fmt,omitempty"`
	SampleRate    *string `json:"sample_rate,omitempty"`
	Channels      *int64  `json:"channels,omitempty"`
	ChannelLayout *string `json:"channel_layout,omitempty"`
	Tags          *Tags   `json:"tags,omitempty"`
}

type Tags struct {
	Language *string `json:"language,omitempty"`
	Title    *string `json:"title,omitempty"`
	Duration *string `json:"DURATION,omitempty"`
	Bps      *string `json:"BPS,omitempty"`
}
