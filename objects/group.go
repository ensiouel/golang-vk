package objects

// Group Объект содержит информацию о сообществе ВКонтакте.
type Group struct {
	ID           int    `json:"id"`            // Идентификатор сообщества.
	Name         string `json:"name"`          // Название сообщества.
	ScreenName   string `json:"screen_name"`   // Короткий адрес, например, apiclub.
	IsClosed     int    `json:"is_closed"`     // Является ли сообщество закрытым. Возможные значения: 0 — открытое; 1 — закрытое; 2 — частное.
	Deactivated  string `json:"deactivated"`   // Возвращается в случае, если сообщество удалено или заблокировано. Возможные значения: deleted — сообщество удалено; banned — сообщество заблокировано.
	IsAdmin      int    `json:"is_admin"`      // Требуется scope = groups. Информация о том, является ли текущий пользователь руководителем. Возможные значения: 1 — является; 0 — не является.
	AdminLevel   int    `json:"admin_level"`   // Требуется scope = groups. Уровень полномочий текущего пользователя (если IsAdmin = 1):
	IsMember     int    `json:"is_member"`     // Требуется scope = groups. Информация о том, является ли текущий пользователь участником. Возможные значения: 1 — является; 0 — не является.
	IsAdvertiser int    `json:"is_advertiser"` // Требуется scope = groups. Информация о том, является ли текущий пользователь рекламодателем. Возможные значения: 1 — является; 0 — не является.
	InvitedBy    int    `json:"invited_by"`    // Требуется scope = groups. Идентификатор пользователя, который отправил приглашение в сообщество. Поле возвращается только для метода groups.getInvites.
	Type         string `json:"type"`          // Тип сообщества: group — группа; page — публичная страница; event — мероприятие.
	Photo50      string `json:"photo_50"`      // URL главной фотографии с размером 50x50px.
	Photo100     string `json:"photo_100"`     // URL главной фотографии с размером 100х100px.
	Photo200     string `json:"photo_200"`     // URL главной фотографии в максимальном размере.
	GroupOptional
}

// GroupOptional Опциональные поля.
type GroupOptional struct {
	Activity          string         `json:"activity,omitempty"`   // Строка тематики паблика. У групп возвращается строковое значение, открыта ли группа или нет, а у событий дата начала.
	Addresses         []GroupAddress `json:"addresses,omitempty"`  // Информация об адресах сообщества.
	AgeLimits         int            `json:"age_limits,omitempty"` // Возрастное ограничение. 1 — нет; 2 — 16+; 3 — 18+.
	BanInfo           GroupBanInfo   `json:"ban_info,omitempty"`   // Информация о занесении в черный список сообщества (поле возвращается только при запросе информации об одном сообществе).
	CanCreateTopic    int            `json:"can_create_topic"`     // Информация о том, может ли текущий пользователь создать новое обсуждение в группе. Возможные значения: 1 — может; 0 — не может.
	CanMessage        int            `json:"can_message"`          // Информация о том, может ли текущий пользователь написать сообщение сообществу. Возможные значения: 1 — может; 0 — не может.
	CanPost           int            `json:"can_post"`             // Информация о том, может ли текущий пользователь оставлять записи на стене сообщества. Возможные значения: 1 — может; 0 — не может.
	CanSeeAllPosts    int            `json:"can_see_all_posts"`    // Информация о том, разрешено ли видеть чужие записи на стене группы. Возможные значения: 1 — может; 0 — не может.
	CanUploadDoc      int            `json:"can_upload_doc"`       // Информация о том, может ли текущий пользователь загружать документы в группу. Возможные значения: 1 — может; 0 — не может.
	CanUploadVideo    int            `json:"can_upload_video"`     // Информация о том, может ли текущий пользователь загружать видеозаписи в группу. Возможные значения: 1 — может; 0 — не может.
	City              GroupCity      `json:"city"`                 // Город, указанный в информации о сообществе.
	Contacts          []GroupContact `json:"contacts"`             // Информация из блока контактов публичной страницы.
	Counters          GroupCounters  `json:"counters"`             // Объект, содержащий счётчики сообщества, может включать любой набор из следующих полей: photos, albums, audios, videos, topics, docs.
	Country           GroupCountry   `json:"country"`              // Страна, указанная в информации о сообществе.
	Cover             GroupCover     `json:"cover"`                // Обложка сообщества.
	CropPhoto         GroupCropPhoto `json:"crop_photo"`           // Возвращает данные о точках, по которым вырезаны профильная и миниатюрная фотографии сообщества.
	Description       string         `json:"description"`          // Текст описания сообщества.
	FixedPost         int            `json:"fixed_post"`           // Идентификатор закрепленной записи.
	HasPhoto          int            `json:"has_photo"`            // Информация о том, установлена ли у сообщества главная фотография. Возможные значения: 1 — установлена; 0 — не установлена.
	IsFavorite        int            `json:"is_favorite"`          // Информация о том, находится ли сообщество в закладках у текущего пользователя. Возможные значения: 1 — находится; 0 — не находится.
	IsHiddenFromFeed  int            `json:"is_hidden_from_feed"`  // Информация о том, скрыто ли сообщество из ленты новостей текущего пользователя. Возможные значения: 1 — скрыто; 0 — не скрыто.
	IsMessagesBlocked int            `json:"is_messages_blocked"`  // Информация о том, заблокированы ли сообщения от этого сообщества (для текущего пользователя).
	Links             []GroupLink    `json:"links"`                // Информация из блока ссылок сообщества.
	MainAlbumID       int            `json:"main_album_id"`        // Идентификатор основного фотоальбома.
	MainSection       int            `json:"main_section"`         // Информация о главной секции. Возможные значения: 0 — отсутствует; 1 — фотографии; 2 — обсуждения; 3 — аудиозаписи; 4 — видеозаписи; 5 — товары.
	Market            GroupMarket    `json:"market"`               // Информация о магазине.
	MemberStatus      int            `json:"member_status"`        // Статус участника текущего пользователя. Возможные значения: 0 — не является участником; 1 — является участником; 2 — не уверен, что посетит мероприятие; 3 — отклонил приглашение; 4 — запрос на вступление отправлен; 5 — приглашен.
	MembersCount      int            `json:"members_count"`        // Количество участников в сообществе.
	Place             GroupPlace     `json:"place"`                // Место, указанное в информации о сообществе.
	PublicDateLabel   string         `json:"public_date_label"`    // Возвращается для публичных страниц.
	Site              string         `json:"site"`                 // Адрес сайта из поля «веб-сайт» в описании сообщества.
	StartDate         interface{}    `json:"start_date"`           //
	FinishDate        interface{}    `json:"finish_date"`          //
	Status            string         `json:"status"`               // Статус сообщества.
	Trending          int            `json:"trending"`             // Информация о том, есть ли у сообщества «огонёк».
	Verified          int            `json:"verified"`             // Информация о том, верифицировано ли сообщество. Возможные значения: 1 — является; 0 — не является.
	Wall              int            `json:"wall"`                 // Стена. Возможные значения: 0 — выключена; 1 — открытая; 2 — ограниченная; 3 — закрытая.
	WikiPage          string         `json:"wiki_page"`            // Название главной вики-страницы.
}

// GroupAddress Информация об адресе сообщества.
type GroupAddress struct {
	IsEnabled     bool `json:"is_enabled"`      // Включен ли блок адресов в сообществе.
	MainAddressID int  `json:"main_address_id"` // Идентификатор основного адреса.
}

// GroupBanInfo Информация о занесении в черный список сообщества.
type GroupBanInfo struct {
	EndDate int64  `json:"end_date"` // Срок окончания блокировки в формате unix time.
	Comment string `json:"comment"`  // Комментарий к блокировке.
}

// GroupCity Город, указанный в информации о сообществе.
type GroupCity struct {
	ID    int    `json:"id"`    // Идентификатор города.
	Title string `json:"title"` // Название города.
}

// GroupContact Контакт публичной страницы.
type GroupContact struct {
	UserID int    `json:"user_id"` // Идентификатор пользователя.
	Desc   string `json:"desc"`    // Должность.
	Phone  string `json:"phone"`   // Номер телефона.
	Email  string `json:"email"`   // Адрес e-mail.
}

// GroupCounters Объект, содержащий счётчики сообщества.
type GroupCounters struct {
	Photos int `json:"photos,omitempty"`
	Albums int `json:"albums,omitempty"`
	Audios int `json:"audios,omitempty"`
	Videos int `json:"videos,omitempty"`
	Topics int `json:"topics,omitempty"`
	Docs   int `json:"docs,omitempty"`
}

// GroupCountry Страна, указанная в информации о сообществе.
type GroupCountry struct {
	ID    int    `json:"id"`    // Идентификатор страны.
	Title string `json:"title"` // Название страны.
}

// GroupCover Обложка сообщества.
type GroupCover struct {
	Enabled int     `json:"enabled"` // Информация о том, включена ли обложка (1 — да, 0 — нет).
	Images  []Image `json:"images"`  // Копии изображений обложки.
}

// GroupCropPhoto Возвращает данные о точках, по которым вырезаны профильная и миниатюрная фотографии пользователя, при наличии.
type GroupCropPhoto struct {
	Photo Photo `json:"photo"` // Объект photo фотографии пользователя, из которой вырезается главное фото сообщества.
	Crop  Crop  `json:"crop"`  // Вырезанная фотография сообщества.
	Rect  Rect  `json:"rect"`  // Миниатюрная квадратная фотография, вырезанная из фотографии Crop.
}

// GroupLink Блока ссылки сообщества.
type GroupLink struct {
	ID       int    `json:"id"`        // Идентификатор ссылки.
	Url      string `json:"url"`       // URL.
	Name     string `json:"name"`      // Название ссылки.
	Desc     string `json:"desc"`      // Описание ссылки.
	Photo50  string `json:"photo_50"`  // URL изображения-превью шириной 50px.
	Photo100 string `json:"photo_100"` // URL изображения-превью шириной 100px.
}

// GroupMarket Информация о магазине.
type GroupMarket struct {
	Enabled      int      `json:"enabled"`                 // Информация о том, включен ли блок товаров в сообществе. Возможные значения: 1 — включен; 0 — выключен. Если Enabled = 0, объект не содержит других полей.
	Type         string   `json:"type,omitempty"`          // Информация о типе магазина. Возвращается, если в группе включен раздел "Товары". Возможные значения: basic — базовые товары; advanced — расширенные товары.
	PriceMin     int      `json:"price_min,omitempty"`     // Минимальная цена товаров.
	PriceMax     int      `json:"price_max,omitempty"`     // Максимальная цена товаров.
	MainAlbumId  int      `json:"main_album_id,omitempty"` // Идентификатор главной подборки товаров.
	ContactID    int      `json:"contact_id,omitempty"`    // Идентификатор контактного лица для связи с продавцом. Возвращается отрицательное значение, если для связи с продавцом используются сообщения сообщества;
	Currency     Currency `json:"currency,omitempty"`      // Информация о валюте.
	CurrencyText string   `json:"currency_text,omitempty"` // Строковое обозначение.
}

// GroupPlace Место, указанное в информации о сообществе.
type GroupPlace struct {
	ID        int    `json:"id"`        // Идентификатор места.
	Title     int    `json:"title"`     // Название места.
	Latitude  int    `json:"latitude"`  // Географическая широта в градусах (от -90 до 90).
	Longitude int    `json:"longitude"` // Географическая долгота в градусах (от -180 до 180).
	Type      string `json:"type"`      // Тип места.
	Country   int    `json:"country"`   // Идентификатор страны.
	City      int    `json:"city"`      // Идентификатор города.
	Address   string `json:"address"`   // Адрес.
}
