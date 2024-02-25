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
	Size              int64       `json:"size"`
	Quality           int64       `json:"quality"`
	Videotype         Videotype   `json:"videotype"`
	Voices            []string    `json:"voices"`
	Seasons           []int64     `json:"seasons"`
	TrackerName       TrackerName `json:"trackerName"`
	Types             []Type      `json:"types"`
	URL               string      `json:"url"`
	Title             string      `json:"title"`
	Sid               int64       `json:"sid"`
	Pir               int64       `json:"pir"`
	SizeName          string      `json:"sizeName"`
	CreateTime        string      `json:"createTime"`
	UpdateTime        string      `json:"updateTime"`
	CheckTime         string      `json:"checkTime"`
	Magnet            string      `json:"magnet"`
	Name              string      `json:"name"`
	Originalname      *string     `json:"originalname,omitempty"`
	Relased           int64       `json:"relased"`
	FFProbeTryingdata int64       `json:"ffprobe_tryingdata"`
	Sn                string      `json:"_sn"`
	So                *string     `json:"_so,omitempty"`
	FFProbe           []FFProbe   `json:"ffprobe,omitempty"`
	Languages         []Language  `json:"languages,omitempty"`
}

type FFProbe struct {
	Index         int64          `json:"index"`
	CodecName     *CodecName     `json:"codec_name,omitempty"`
	CodecLongName *CodecLongName `json:"codec_long_name,omitempty"`
	CodecType     CodecType      `json:"codec_type"`
	Width         *int64         `json:"width,omitempty"`
	Height        *int64         `json:"height,omitempty"`
	CodedWidth    *int64         `json:"coded_width,omitempty"`
	CodedHeight   *int64         `json:"coded_height,omitempty"`
	BitRate       *string        `json:"bit_rate,omitempty"`
	SampleFmt     *SampleFmt     `json:"sample_fmt,omitempty"`
	SampleRate    *string        `json:"sample_rate,omitempty"`
	Channels      *int64         `json:"channels,omitempty"`
	ChannelLayout *ChannelLayout `json:"channel_layout,omitempty"`
	Tags          *Tags          `json:"tags,omitempty"`
}

type Tags struct {
	Language *Language `json:"language,omitempty"`
	Title    *string   `json:"title,omitempty"`
	Duration *string   `json:"DURATION,omitempty"`
	Bps      *string   `json:"BPS,omitempty"`
}

type ChannelLayout string

const (
	Stereo    ChannelLayout = "stereo"
	The51     ChannelLayout = "5.1"
	The51Side ChannelLayout = "5.1(side)"
	The61     ChannelLayout = "6.1"
)

type CodecLongName string

const (
	AACAdvancedAudioCoding                 CodecLongName = "AAC (Advanced Audio Coding)"
	ASSAdvancedSSASubtitle                 CodecLongName = "ASS (Advanced SSA) subtitle"
	AtscA52AAC3                            CodecLongName = "ATSC A/52A (AC-3)"
	AtscA52BAC3EAC3                        CodecLongName = "ATSC A/52B (AC-3, E-AC-3)"
	BinaryData                             CodecLongName = "binary data"
	DCADTSCoherentAcoustics                CodecLongName = "DCA (DTS Coherent Acoustics)"
	DVDSubtitles                           CodecLongName = "DVD subtitles"
	H264AVCMPEG4AVCMPEG4Part10             CodecLongName = "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10"
	H265HEVCHighEfficiencyVideoCoding      CodecLongName = "H.265 / HEVC (High Efficiency Video Coding)"
	HDMVPresentationGraphicStreamSubtitles CodecLongName = "HDMV Presentation Graphic Stream subtitles"
	MOVText                                CodecLongName = "MOV text"
	MP2MPEGAudioLayer2                     CodecLongName = "MP2 (MPEG audio layer 2)"
	MP3MPEGAudioLayer3                     CodecLongName = "MP3 (MPEG audio layer 3)"
	MPEG2Video                             CodecLongName = "MPEG-2 video"
	MPEG4Part2                             CodecLongName = "MPEG-4 part 2"
	MotionJPEG                             CodecLongName = "Motion JPEG"
	PCMSigned16BitLittleEndian             CodecLongName = "PCM signed 16-bit little-endian"
	PNGPortableNetworkGraphicsImage        CodecLongName = "PNG (Portable Network Graphics) image"
	RawUTF8Text                            CodecLongName = "raw UTF-8 text"
	SmpteVc1                               CodecLongName = "SMPTE VC-1"
	SubRIPSubtitle                         CodecLongName = "SubRip subtitle"
	TrueTypeFont                           CodecLongName = "TrueType font"
)

type CodecName string

const (
	AAC              CodecName = "aac"
	Ac3              CodecName = "ac3"
	Ass              CodecName = "ass"
	BinData          CodecName = "bin_data"
	CodecNameMOVText CodecName = "mov_text"
	DVDSubtitle      CodecName = "dvd_subtitle"
	Dts              CodecName = "dts"
	Eac3             CodecName = "eac3"
	H264             CodecName = "h264"
	HdmvPgsSubtitle  CodecName = "hdmv_pgs_subtitle"
	Hevc             CodecName = "hevc"
	Mjpeg            CodecName = "mjpeg"
	Mp2              CodecName = "mp2"
	Mp3              CodecName = "mp3"
	Mpeg2Video       CodecName = "mpeg2video"
	Mpeg4            CodecName = "mpeg4"
	PCMS16LE         CodecName = "pcm_s16le"
	PNG              CodecName = "png"
	Subrip           CodecName = "subrip"
	TTF              CodecName = "ttf"
	Text             CodecName = "text"
	Vc1              CodecName = "vc1"
)

type CodecType string

const (
	Attachment CodecType = "attachment"
	Audio      CodecType = "audio"
	Data       CodecType = "data"
	Subtitle   CodecType = "subtitle"
	Video      CodecType = "video"
)

type SampleFmt string

const (
	Fltp SampleFmt = "fltp"
	S16  SampleFmt = "s16"
	S16P SampleFmt = "s16p"
	S32P SampleFmt = "s32p"
)

type Language string

const (
	Afr   Language = "afr"
	Ara   Language = "ara"
	Bul   Language = "bul"
	Chi   Language = "chi"
	Cze   Language = "cze"
	Dan   Language = "dan"
	Dut   Language = "dut"
	Empty Language = "???"
	Eng   Language = "eng"
	Ewe   Language = "ewe"
	Fin   Language = "fin"
	Fra   Language = "fra"
	Fre   Language = "fre"
	Ger   Language = "ger"
	Gre   Language = "gre"
	Heb   Language = "heb"
	Hun   Language = "hun"
	Ice   Language = "ice"
	Ind   Language = "ind"
	Ita   Language = "ita"
	Jpn   Language = "jpn"
	Kor   Language = "kor"
	Lav   Language = "lav"
	Nob   Language = "nob"
	Nor   Language = "nor"
	Per   Language = "per"
	Pol   Language = "pol"
	Por   Language = "por"
	Rum   Language = "rum"
	Rus   Language = "rus"
	SPA   Language = "spa"
	Scc   Language = "scc"
	Scr   Language = "scr"
	Slo   Language = "slo"
	Slv   Language = "slv"
	Srp   Language = "srp"
	Swe   Language = "swe"
	Tur   Language = "tur"
	Ukr   Language = "ukr"
	Und   Language = "und"
	Vie   Language = "vie"
)

type TrackerName string

const (
	Baibako   TrackerName = "baibako"
	Bitru     TrackerName = "bitru"
	Hdrezka   TrackerName = "hdrezka"
	Kinozal   TrackerName = "kinozal"
	Lostfilm  TrackerName = "lostfilm"
	Megapeer  TrackerName = "megapeer"
	Nnmclub   TrackerName = "nnmclub"
	Rutor     TrackerName = "rutor"
	Rutracker TrackerName = "rutracker"
	Toloka    TrackerName = "toloka"
	Torrentby TrackerName = "torrentby"
)

type Type string

const (
	Documovie  Type = "documovie"
	Docuserial Type = "docuserial"
	Movie      Type = "movie"
	Multfilm   Type = "multfilm"
	Multserial Type = "multserial"
	Serial     Type = "serial"
	Tvshow     Type = "tvshow"
)

type Videotype string

const (
	Hdr Videotype = "hdr"
	SDR Videotype = "sdr"
)
