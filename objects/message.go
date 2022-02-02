package objects

const (
	MessageAttachmentTypePhoto       = "photo"
	MessageAttachmentTypeAudio       = "audio"
	MessageAttachmentTypeVideo       = "video"
	MessageAttachmentTypeDoc         = "doc"
	MessageAttachmentTypeLink        = "link"
	MessageAttachmentTypeMarket      = "market"
	MessageAttachmentTypeMarketAlbum = "market_album"
	MessageAttachmentTypeSticker     = "sticker"
)

// Message Объект, описывающий личное сообщение.
type Message struct {
	ID           int                `json:"id"`                  // Идентификатор сообщения.
	Date         int64              `json:"date"`                // Время отправки в Unix time.
	PeerID       int                `json:"peer_id"`             // Идентификатор назначения.
	FromID       int                `json:"from_id"`             // Идентификатор отправителя.
	Text         string             `json:"text"`                // Текст сообщения.
	RandomID     int                `json:"random_id,omitempty"` // Идентификатор, используемый при отправке сообщения. Возвращается только для исходящих сообщений.
	Ref          string             `json:"ref"`                 // Произвольный параметр для работы с источниками переходов.
	RefSource    string             `json:"ref_source"`          // Произвольный параметр для работы с источниками переходов.
	Attachments  MessageAttachments `json:"attachments"`         // Медиа вложения сообщения (фотографии, ссылки и т.п.).
	Important    bool               `json:"important"`           // True, если сообщение помечено как важное.
	Geo          MessageGeo         `json:"geo"`                 // Информация о местоположении.
	Payload      string             `json:"payload"`             // Сервисное поле для сообщений ботам (полезная нагрузка).
	Keyboard     Keyboard           `json:"keyboard"`            // Клавиатуры для ботов.
	FwdMessages  []*Message         `json:"fwd_messages"`        // Массив пересланных сообщений (если есть). Максимальное количество элементов — 100. Максимальная глубина вложенности для пересланных сообщений — 45, общее максимальное количество в цепочке с учетом вложенности — 500.
	ReplyMessage *Message           `json:"reply_message"`       // Сообщение, в ответ на которое отправлено текущее.
}

type MessageAttachments []MessageAttachment

// MessageAttachment Медиа вложения сообщения (фотографии, ссылки и т.п.).
type MessageAttachment struct {
	Photo       *Photo          `json:"photo,omitempty"`        // Фотография.
	Video       *Video          `json:"video,omitempty"`        // Видеозапись.
	Audio       *Audio          `json:"audio,omitempty"`        // Аудиозапись.
	Doc         *Doc            `json:"doc,omitempty"`          // Документ.
	Link        *Link           `json:"link,omitempty"`         // Ссылка.
	Market      *MarketItem     `json:"market,omitempty"`       // Товар.
	MarketAlbum *MarketAlbum    `json:"market_album,omitempty"` // Подборка товаров.
	Wall        *MessagePost    `json:"wall"`                   // Запись на стене. Обратите внимание, вместо поля owner_id возвращается to_id.
	WallReply   *MessageComment `json:"wall_reply"`             // Комментарий на стене.
	Sticker     *Sticker        `json:"sticker"`                // Стикер.
	Gift        *Gift           `json:"gift"`                   // Подарок.
}

// MessagePost Запись на стене
type MessagePost struct {
	Post
	ToID int `json:"to_id"`
}

// MessageComment Комментарий на стене.
type MessageComment struct {
	Comment
	PostID  int `json:"post_id"`  // Идентификатор записи, к которой оставлен комментарий.
	OwnerID int `json:"owner_id"` // Идентификатор владельца стены, на которой оставлен комментарий.
}

type MessageGeo struct {
	Geo
	Showmap int `json:"showmap"`
}
