package objects

import "time"

type DocPreviewPhoto struct {
	Sizes []PhotoSize `json:"sizes"`
}

type DocPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type DocPreviewAudioMessage struct {
	Duration time.Duration `json:"duration"`
	Waveform []int         `json:"waveform"`
	LinkOGG  string        `json:"link_ogg"`
	LinkMP3  string        `json:"link_mp3"`
}

type DocPreview struct {
	Photo        DocPreviewPhoto        `json:"photo,omitempty"`
	Graffiti     DocPreviewGraffiti     `json:"graffiti,omitempty"`
	AudioMessage DocPreviewAudioMessage `json:"audio_message,omitempty"`
}

type Doc struct {
	ID      int64      `json:"id"`
	OwnerID int64      `json:"owner_id"`
	Title   string     `json:"title"`
	Size    int64      `json:"size"`
	Ext     string     `json:"ext"`
	Url     string     `json:"url"`
	Date    int64      `json:"date"`
	Type    int        `json:"type"`
	Preview DocPreview `json:"preview"`
}
