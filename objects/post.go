package objects

const (
	PostAttachmentTypePhoto       = "photo"
	PostAttachmentTypePostedPhoto = "posted_photo"
	PostAttachmentTypeAudio       = "audio"
	PostAttachmentTypeVideo       = "video"
	PostAttachmentTypeDoc         = "doc"
	PostAttachmentTypeGraffiti    = "graffiti"
	PostAttachmentTypeLink        = "link"
	PostAttachmentTypeNote        = "note"
	PostAttachmentTypeApp         = "app"
	PostAttachmentTypePoll        = "poll"
	PostAttachmentTypePage        = "page"
	PostAttachmentTypeAlbum       = "album"
	PostAttachmentTypePhotosList  = "photos_list"
	PostAttachmentTypeMarket      = "market"
	PostAttachmentTypeMarketAlbum = "market_album"
	PostAttachmentTypeSticker     = "sticker"
	PostAttachmentTypePrettyCards = "pretty_cards"
	PostAttachmentTypeEvent       = "event"
)

// Post Объект, описывающий запись на стене пользователя или сообщества
type Post struct {
	ID           int             `json:"id"`                     // Идентификатор записи.
	OwnerID      int             `json:"owner_id"`               // Идентификатор владельца стены, на которой размещена запись.
	FromID       int             `json:"from_id"`                // Идентификатор автора записи (от чьего имени опубликована запись).
	CreatedBy    int             `json:"created_by"`             // Идентификатор администратора, который опубликовал запись (возвращается только для сообществ при запросе с ключом доступа администратора). Возвращается в записях, опубликованных менее 24 часов назад.
	Date         int64           `json:"date"`                   // Время публикации записи в формате unix time.
	Text         string          `json:"text"`                   // Текст записи.
	ReplyOwnerID int             `json:"reply_owner_id"`         // Идентификатор владельца записи, в ответ на которую была оставлена текущая.
	ReplyPostID  int             `json:"reply_post_id"`          // Идентификатор записи, в ответ на которую была оставлена текущая.
	FriendsOnly  int             `json:"friends_only"`           // 1, если запись была создана с опцией «Только для друзей».
	Comments     PostComments    `json:"comments"`               // Информация о комментариях к записи.
	Copyright    PostCopyright   `json:"copyright"`              // Источник материала
	Likes        PostLikes       `json:"likes"`                  // Информация о лайках к записи.
	Reposts      PostReposts     `json:"reposts"`                // Информация о репостах записи («Рассказать друзьям»).
	Views        PostViews       `json:"views"`                  // Информация о просмотрах записи.
	PostType     string          `json:"post_type"`              // Тип записи, может принимать следующие значения: post, copy, reply, postpone, suggest.
	PostSource   PostSource      `json:"post_source,omitempty"`  // Информация о способе размещения записи.
	Attachments  PostAttachments `json:"attachments"`            // Медиа вложения записи (фотографии, ссылки и т.п.).
	Geo          PostGeo         `json:"geo"`                    // Информация о местоположении.
	SignerID     int             `json:"signer_id,omitempty"`    // Идентификатор автора, если запись была опубликована от имени сообщества и подписана пользователем.
	CopyHistory  []Post          `json:"copy_history,omitempty"` // Массив, содержащий историю репостов для записи. Возвращается только в том случае, если запись является репостом. Каждый из объектов массива, в свою очередь, является объектом-записью стандартного формата.
	CanPin       int             `json:"can_pin"`                // Информация о том, может ли текущий пользователь закрепить запись (1 — может, 0 — не может).
	CanDelete    int             `json:"can_delete"`             // Информация о том, может ли текущий пользователь удалить запись (1 — может, 0 — не может).
	CanEdit      int             `json:"can_edit"`               // Информация о том, может ли текущий пользователь редактировать запись (1 — может, 0 — не может).
	IsPinned     int             `json:"is_pinned"`              // Информация о том, что запись закреплена.
	MarkedAsAds  int             `json:"marked_as_ads"`          // Информация о том, содержит ли запись отметку «реклама» (1 — да, 0 — нет).
	IsFavorite   bool            `json:"is_favorite"`            // True, если объект добавлен в закладки у текущего пользователя.
	Donut        PostDonut       `json:"donut"`                  // Информация о записи VK Donut.
	PostponedID  int             `json:"postponed_id,omitempty"` // Идентификатор отложенной записи. Это поле возвращается тогда, когда запись стояла на таймере.
}

// PostComments Информация о комментариях к записи.
type PostComments struct {
	Count         int  `json:"count"`              // Количество комментариев.
	CanPost       int  `json:"can_post,omitempty"` // Информация о том, может ли текущий пользователь комментировать запись (1 — может, 0 — не может).
	GroupsCanPost bool `json:"groups_can_post"`    // Информация о том, могут ли сообщества комментировать запись.
	CanClose      int  `json:"can_close"`          // Может ли текущий пользователь закрыть комментарии к записи.
	CanOpen       int  `json:"can_open"`           // Может ли текущий пользователь открыть комментарии к записи.
}

// PostCopyright Источник материала.
type PostCopyright struct {
	ID   int    `json:"id"`
	Link string `json:"link,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// PostLikes Информация о лайках к записи.
type PostLikes struct {
	Count      int `json:"count"`                 // Число пользователей, которым понравилась запись.
	UserLikes  int `json:"user_likes,omitempty"`  // Наличие отметки «Мне нравится» от текущего пользователя (1 — есть, 0 — нет).
	CanLike    int `json:"can_like,omitempty"`    // Информация о том, может ли текущий пользователь поставить отметку «Мне нравится» (1 — может, 0 — не может).
	CanPublish int `json:"can_publish,omitempty"` // Информация о том, может ли текущий пользователь сделать репост записи (1 — может, 0 — не может).
}

// PostReposts Информация о репостах записи («Рассказать друзьям»).
type PostReposts struct {
	Count        int `json:"count"`         // Число пользователей, скопировавших запись.
	UserReposted int `json:"user_reposted"` // Наличие репоста от текущего пользователя (1 — есть, 0 — нет).
}

// PostViews Информация о просмотрах записи.
type PostViews struct {
	Count int `json:"count"` // Число просмотров записи.
}

// PostSource Информация о способе размещения записи.
type PostSource struct {
	Type     string `json:"type"`     // Тип источника. Возможные значения: vk — Запись создана через основной интерфейс сайта https://vk.com/; widget — Запись создана через виджет на стороннем сайте; api — Запись создана приложением через API; rss — Запись создана посредством импорта RSS-ленты со стороннего сайта; sms — Запись создана посредством отправки SMS-сообщения на специальный номер.
	Platform string `json:"platform"` // Название платформы, если оно доступно. Возможные значения: android; iphone; wphone.
	Data     string `json:"data"`     // Тип действия (только для Type = vk или widget). Возможные значения: profile_activity — Изменение статуса под именем пользователя (для Type = vk); profile_photo — Изменение профильной фотографии пользователя (для Type = vk); comments — Виджет комментариев (для Type = widget); like — Виджет «Мне нравится» (для Type = widget); poll — Виджет опросов (для Type = widget).
	Url      string `json:"url"`      // URL ресурса, с которого была опубликована запись.
}

type PostPhoto struct {
	Photo
	PostID int `json:"post_id"` // Идентификатор записи, в которую была загружена фотография.
}

type PostPostedPhoto struct {
	ID       int `json:"id"`        // Идентификатор фотографии.
	OwnerId  int `json:"owner_id"`  // Идентификатор владельца фотографии.
	Photo130 int `json:"photo_130"` // URL изображения для предпросмотра.
	Photo604 int `json:"photo_604"` // URL полноразмерного изображения.
}

type PostGraffiti struct {
	ID       int `json:"id"`        // Идентификатор граффити.
	OwnerId  int `json:"owner_id"`  // Идентификатор автора граффити.
	Photo130 int `json:"photo_130"` // URL изображения для предпросмотра.
	Photo604 int `json:"photo_604"` // URL полноразмерного изображения.
}

type PostApp struct {
	ID       int    `json:"id"`        // Идентификатор приложения.
	Name     string `json:"name"`      // Название приложения.
	Photo130 int    `json:"photo_130"` // URL изображения для предпросмотра.
	Photo604 int    `json:"photo_604"` // URL полноразмерного изображения.
}

type PostAlbum struct {
	ID          int    `json:"id"`          // Идентификатор альбома.
	Thumb       Photo  `json:"thumb"`       // Обложка альбома.
	OwnerID     int    `json:"owner_id"`    // Идентификатор владельца альбома.
	Title       string `json:"title"`       // Название альбома.
	Description string `json:"description"` // Описание альбома.
	Created     int64  `json:"created"`     // Дата создания альбома в формате Unix time.
	Updated     int64  `json:"updated"`     // Дата последнего обновления альбома в формате Unix time.
	Size        int    `json:"size"`        // Количество фотографий в альбоме.
}

// PrettyCard Карточка.
type PrettyCard struct {
	CardID   string  `json:"card_id"`   // Идентификатор карточки.
	LinkUrl  string  `json:"link_url"`  // URL карточки.
	Title    string  `json:"title"`     // Подпись карточки.
	Images   []Image `json:"images"`    // Изображения.
	Button   Button  `json:"button"`    // Объект кнопки.
	Price    string  `json:"price"`     // Цена.
	PriceOld string  `json:"price_old"` // Старая цена.
}

// PostEvent Встреча.
type PostEvent struct {
	ID           int    `json:"id"`            // Идентификатор встречи.
	Time         int64  `json:"time"`          // Время начала встречи в Unix time.
	MemberStatus int    `json:"member_status"` // Идёт ли текущий пользователь на встречу. Возможные значения: 1 — точно идёт; 2 — возможно пойдёт; 3 — не идёт.
	IsFavorite   bool   `json:"is_favorite"`   // Добавлена ли встреча в закладки.
	Address      string `json:"address"`       // Место проведения встречи.
	Text         string `json:"text"`          // Текст для отображения сниппета.
	ButtonText   string `json:"button_text"`   // Текст на кнопке сниппета.
	Friends      []int  `json:"friends"`       // Список идентификаторов друзей, которые также идут на мероприятие.
}

// PostAttachment Медиа вложение.
type PostAttachment struct {
	Type        string           `json:"type"`                   // Тип вложения (photo, note, audio и т.д.).
	Photo       *PostPhoto       `json:"photo,omitempty"`        // Фотография.
	PostedPhoto *PostPostedPhoto `json:"posted_photo,omitempty"` // Это устаревший тип вложения. Он может быть возвращен лишь для записей, созданных раньше 2013 года.
	Video       *Video           `json:"video,omitempty"`        // Видеозапись.
	Audio       *Audio           `json:"audio,omitempty"`        // Аудиозапись.
	Doc         *Doc             `json:"doc,omitempty"`          // Документ.
	Graffiti    *PostGraffiti    `json:"graffiti,omitempty"`     // Граффити.
	Link        *Link            `json:"link,omitempty"`         // Ссылка.
	Note        *Note            `json:"note,omitempty"`         // Заметка.
	App         *PostApp         `json:"app,omitempty"`          // Контент приложения.
	Poll        *Poll            `json:"poll,omitempty"`         // Опрос.
	Page        *WikiPage        `json:"page,omitempty"`         // Вики-страница (для версии API 5.20 и выше).
	Album       *PostAlbum       `json:"album,omitempty"`        // Альбом с фотографиями.
	PhotosList  []string         `json:"photos_list,omitempty"`  // Массив из строк, содержащих идентификаторы фотографий. Сами фотографии дублируются в виде приложенных объектов фотографий, однако этот список необходим в случае, если фотографий использовано больше максимального количества возвращаемых вложений (10).
	Market      *MarketItem      `json:"market,omitempty"`       // Товар.
	MarketAlbum *MarketAlbum     `json:"market_album,omitempty"` // Подборка товаров.
	Sticker     *Sticker         `json:"sticker,omitempty"`      // Стикер.
	PrettyCards []PrettyCard     `json:"pretty_cards,omitempty"` // Массив элементов-карточек.
	Event       *PostEvent       `json:"event,omitempty"`        // Встреча.
}

// PostGeo Информация о местоположении.
type PostGeo struct {
	Type        string `json:"type"`            // Тип места.
	Coordinates string `json:"coordinates"`     // Координаты места.
	Place       Place  `json:"place,omitempty"` // Описание места (если оно добавлено).
}

// PostDonut Информация о записи VK Donut.
type PostDonut struct {
	IsDonut            bool   `json:"is_donut"`              // Запись доступна только платным подписчикам VK Donut.
	PaidDuration       int    `json:"paid_duration"`         // Время, в течение которого запись будет доступна только платным подписчикам VK Donut.
	Placeholder        string `json:"placeholder"`           // Заглушка для пользователей, которые не оформили подписку VK Donut. Отображается вместо содержимого записи.
	CanPublishFreeCopy bool   `json:"can_publish_free_copy"` // Можно ли открыть запись для всех пользователей, а не только подписчиков VK Donut.
	EditMode           string `json:"edit_mode"`             // Информация о том, какие значения VK Donut можно изменить в записи. Возможные значения: all — всю информацию о VK Donut; duration — время, в течение которого запись будет доступна только платным подписчикам VK Donut.
}

type PostAttachments []PostAttachment

func (attachments *PostAttachments) GetAllPhotos() (photos []*PostPhoto) {
	for _, attachment := range *attachments {
		if attachment.IsPhoto() {
			photos = append(photos, attachment.Photo)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllPostedPhotos() (photos []*PostPostedPhoto) {
	for _, attachment := range *attachments {
		if attachment.IsPostedPhoto() {
			photos = append(photos, attachment.PostedPhoto)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllVideos() (videos []*Video) {
	for _, attachment := range *attachments {
		if attachment.IsVideo() {
			videos = append(videos, attachment.Video)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllAudios() (audios []*Audio) {
	for _, attachment := range *attachments {
		if attachment.IsAudio() {
			audios = append(audios, attachment.Audio)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllDocs() (docs []*Doc) {
	for _, attachment := range *attachments {
		if attachment.IsDoc() {
			docs = append(docs, attachment.Doc)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllGraffities() (graffities []*PostGraffiti) {
	for _, attachment := range *attachments {
		if attachment.IsGraffiti() {
			graffities = append(graffities, attachment.Graffiti)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllLinks() (links []*Link) {
	for _, attachment := range *attachments {
		if attachment.IsLink() {
			links = append(links, attachment.Link)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllNotes() (notes []*Note) {
	for _, attachment := range *attachments {
		if attachment.IsNote() {
			notes = append(notes, attachment.Note)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllApps() (apps []*PostApp) {
	for _, attachment := range *attachments {
		if attachment.IsApp() {
			apps = append(apps, attachment.App)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllPolls() (polls []*Poll) {
	for _, attachment := range *attachments {
		if attachment.IsPoll() {
			polls = append(polls, attachment.Poll)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllPages() (pages []*WikiPage) {
	for _, attachment := range *attachments {
		if attachment.IsPage() {
			pages = append(pages, attachment.Page)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllAlbums() (albums []*PostAlbum) {
	for _, attachment := range *attachments {
		if attachment.IsAlbum() {
			albums = append(albums, attachment.Album)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllPhotosLists() (photosLists [][]string) {
	for _, attachment := range *attachments {
		if attachment.IsPhotosList() {
			photosLists = append(photosLists, attachment.PhotosList)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllMarkets() (markets []*MarketItem) {
	for _, attachment := range *attachments {
		if attachment.IsMarket() {
			markets = append(markets, attachment.Market)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllMarketAlbums() (marketAlbums []*MarketAlbum) {
	for _, attachment := range *attachments {
		if attachment.IsMarketAlbum() {
			marketAlbums = append(marketAlbums, attachment.MarketAlbum)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllStickers() (stickers []*Sticker) {
	for _, attachment := range *attachments {
		if attachment.IsSticker() {
			stickers = append(stickers, attachment.Sticker)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllPrettyCards() (prettyCards [][]PrettyCard) {
	for _, attachment := range *attachments {
		if attachment.IsPrettyCards() {
			prettyCards = append(prettyCards, attachment.PrettyCards)
		}
	}
	return
}

func (attachments *PostAttachments) GetAllEvents() (events []*PostEvent) {
	for _, attachment := range *attachments {
		if attachment.IsEvent() {
			events = append(events, attachment.Event)
		}
	}
	return
}

func (attachment *PostAttachment) IsPhoto() bool {
	return attachment.Type == PostAttachmentTypePhoto && attachment.Photo != nil
}

func (attachment *PostAttachment) IsPostedPhoto() bool {
	return attachment.Type == PostAttachmentTypePostedPhoto && attachment.PostedPhoto != nil
}

func (attachment *PostAttachment) IsVideo() bool {
	return attachment.Type == PostAttachmentTypeVideo && attachment.Video != nil
}

func (attachment *PostAttachment) IsAudio() bool {
	return attachment.Type == PostAttachmentTypeAudio && attachment.Audio != nil
}

func (attachment *PostAttachment) IsDoc() bool {
	return attachment.Type == PostAttachmentTypeDoc && attachment.Doc != nil
}

func (attachment *PostAttachment) IsGraffiti() bool {
	return attachment.Type == PostAttachmentTypeGraffiti && attachment.Graffiti != nil
}

func (attachment *PostAttachment) IsLink() bool {
	return attachment.Type == PostAttachmentTypeLink && attachment.Link != nil
}

func (attachment *PostAttachment) IsNote() bool {
	return attachment.Type == PostAttachmentTypeNote && attachment.Note != nil
}

func (attachment *PostAttachment) IsApp() bool {
	return attachment.Type == PostAttachmentTypeApp && attachment.App != nil
}

func (attachment *PostAttachment) IsPoll() bool {
	return attachment.Type == PostAttachmentTypePoll && attachment.Poll != nil
}

func (attachment *PostAttachment) IsPage() bool {
	return attachment.Type == PostAttachmentTypePage && attachment.Page != nil
}

func (attachment *PostAttachment) IsAlbum() bool {
	return attachment.Type == PostAttachmentTypeAlbum && attachment.Album != nil
}

func (attachment *PostAttachment) IsPhotosList() bool {
	return attachment.Type == PostAttachmentTypePhotosList && attachment.PhotosList != nil
}

func (attachment *PostAttachment) IsMarket() bool {
	return attachment.Type == PostAttachmentTypeMarket && attachment.Market != nil
}

func (attachment *PostAttachment) IsMarketAlbum() bool {
	return attachment.Type == PostAttachmentTypeMarketAlbum && attachment.MarketAlbum != nil
}

func (attachment *PostAttachment) IsSticker() bool {
	return attachment.Type == PostAttachmentTypeSticker && attachment.Sticker != nil
}

func (attachment *PostAttachment) IsPrettyCards() bool {
	return attachment.Type == PostAttachmentTypePrettyCards && attachment.PrettyCards != nil
}

func (attachment *PostAttachment) IsEvent() bool {
	return attachment.Type == PostAttachmentTypeEvent && attachment.Event != nil
}

func (attachment *PostAttachment) GetPhoto() *PostPhoto {
	if attachment.IsPhoto() {
		return attachment.Photo
	}
	return nil
}

func (attachment *PostAttachment) GetPostedPhoto() *PostPostedPhoto {
	if attachment.IsPostedPhoto() {
		return attachment.PostedPhoto
	}
	return nil
}

func (attachment *PostAttachment) GetVideo() *Video {
	if attachment.IsVideo() {
		return attachment.Video
	}
	return nil
}

func (attachment *PostAttachment) GetAudio() *Audio {
	if attachment.IsAudio() {
		return attachment.Audio
	}
	return nil
}

func (attachment *PostAttachment) GetDoc() *Doc {
	if attachment.IsDoc() {
		return attachment.Doc
	}
	return nil
}

func (attachment *PostAttachment) GetGraffiti() *PostGraffiti {
	if attachment.IsGraffiti() {
		return attachment.Graffiti
	}
	return nil
}

func (attachment *PostAttachment) GetLink() *Link {
	if attachment.IsLink() {
		return attachment.Link
	}
	return nil
}

func (attachment *PostAttachment) GetNote() *Note {
	if attachment.IsNote() {
		return attachment.Note
	}
	return nil
}

func (attachment *PostAttachment) GetApp() *PostApp {
	if attachment.IsApp() {
		return attachment.App
	}
	return nil
}

func (attachment *PostAttachment) GetPoll() *Poll {
	if attachment.IsPoll() {
		return attachment.Poll
	}
	return nil
}

func (attachment *PostAttachment) GetPage() *WikiPage {
	if attachment.IsPage() {
		return attachment.Page
	}
	return nil
}

func (attachment *PostAttachment) GetAlbum() *PostAlbum {
	if attachment.IsAlbum() {
		return attachment.Album
	}
	return nil
}

func (attachment *PostAttachment) GetPhotosList() []string {
	if attachment.IsPhotosList() {
		return attachment.PhotosList
	}
	return nil
}

func (attachment *PostAttachment) GetMarket() *MarketItem {
	if attachment.IsMarket() {
		return attachment.Market
	}
	return nil
}

func (attachment *PostAttachment) GetMarketAlbum() *MarketAlbum {
	if attachment.IsMarketAlbum() {
		return attachment.MarketAlbum
	}
	return nil
}

func (attachment *PostAttachment) GetSticker() *Sticker {
	if attachment.IsSticker() {
		return attachment.Sticker
	}
	return nil
}

func (attachment *PostAttachment) GetPrettyCards() []PrettyCard {
	if attachment.IsPrettyCards() {
		return attachment.PrettyCards
	}
	return nil
}

func (attachment *PostAttachment) GetEvent() *PostEvent {
	if attachment.IsEvent() {
		return attachment.Event
	}
	return nil
}
