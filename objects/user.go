package objects

// User Объект содержит информацию о пользователе ВКонтакте.
type User struct {
	ID              int    `json:"id"`                // Идентификатор пользователя.
	FirstName       string `json:"first_name"`        // Имя.
	LastName        string `json:"last_name"`         // Фамилия.
	Deactivated     string `json:"deactivated"`       // Поле возвращается, если страница пользователя удалена или заблокирована, содержит значение deleted или banned. В этом случае опциональные поля не возвращаются.
	IsClosed        bool   `json:"is_closed"`         // Скрыт ли профиль пользователя настройками приватности.
	CanAccessClosed bool   `json:"can_access_closed"` // Может ли текущий пользователь видеть профиль при is_closed = 1 (например, он есть в друзьях).
	UserOptional
}

// UserOptional Опциональные поля.
type UserOptional struct {
	About                  string            `json:"about,omitempty"`                     // Содержимое поля «О себе» из профиля.
	Activities             string            `json:"activities,omitempty"`                // Содержимое поля «Деятельность» из профиля.
	BDate                  string            `json:"bdate,omitempty"`                     // Дата рождения. Возвращается в формате D.M.YYYY или D.M (если год рождения скрыт). Если дата рождения скрыта целиком, поле отсутствует в ответе.
	Blacklisted            int               `json:"blacklisted,omitempty"`               // Информация о том, находится ли текущий пользователь в черном списке. Возможные значения: 1 — находится; 0 — не находится.
	BlacklistedByMe        int               `json:"blacklisted_by_me,omitempty"`         // Информация о том, находится ли пользователь в черном списке у текущего пользователя. Возможные значения: 1 — находится; 0 — не находится.
	Books                  string            `json:"books,omitempty"`                     // Содержимое поля «Любимые книги» из профиля пользователя.
	CanPost                int               `json:"can_post,omitempty"`                  // Информация о том, может ли текущий пользователь оставлять записи на стене. Возможные значения: 1 — может; 0 — не может.
	CanSeeAllPosts         int               `json:"can_see_all_posts,omitempty"`         // Информация о том, может ли текущий пользователь видеть чужие записи на стене. Возможные значения: 1 — может; 0 — не может.
	CanSeeAudio            int               `json:"can_see_audio,omitempty"`             // Информация о том, может ли текущий пользователь видеть аудиозаписи. Возможные значения: 1 — может; 0 — не может.
	CanSendFriendRequest   int               `json:"can_send_friend_request,omitempty"`   // Информация о том, будет ли отправлено уведомление пользователю о заявке в друзья от текущего пользователя. Возможные значения: 1 — уведомление будет отправлено; 0 — уведомление не будет отправлено.
	CanWritePrivateMessage int               `json:"can_write_private_message,omitempty"` // Информация о том, может ли текущий пользователь отправить личное сообщение. Возможные значения: 1 — может; 0 — не может.
	Career                 UserCareer        `json:"career,omitempty"`                    // Информация о карьере пользователя.
	City                   UserCity          `json:"city,omitempty"`                      // Информация о городе, указанном на странице пользователя в разделе «Контакты».
	CommonCount            int               `json:"common_count,omitempty"`              // Количество общих друзей с текущим пользователем.
	Connections            map[string]string `json:"connections,omitempty"`               // Возвращает данные об указанных в профиле сервисах пользователя, таких как: skype, facebook, twitter, livejournal, instagram. Для каждого сервиса возвращается отдельное поле с типом string, содержащее никнейм пользователя. Например, "instagram": "username".
	Contacts               UserContacts      `json:"contacts,omitempty"`                  // Информация о телефонных номерах пользователя.
	Counters               UserCounters      `json:"counters,omitempty"`                  // Количество различных объектов у пользователя. Поле возвращается только в методе users.get при запросе информации об одном пользователе, с передачей пользовательского access_token.
	Country                UserCountry       `json:"country,omitempty"`                   // Информация о стране, указанной на странице пользователя в разделе «Контакты».
	CropPhoto              UserCropPhoto     `json:"crop_photo,omitempty"`                // Возвращает данные о точках, по которым вырезаны профильная и миниатюрная фотографии пользователя, при наличии.
	Domain                 string            `json:"domain,omitempty"`                    // Короткий адрес страницы.
	Education              UserEducation     `json:"education,omitempty"`                 // Информация о высшем учебном заведении пользователя.
	Exports                []string          `json:"exports,omitempty"`                   // Внешние сервисы, в которые настроен экспорт из ВК (twitter, facebook, livejournal, instagram).
	UserFirstName
	FollowersCount   int    `json:"followers_count,omitempty"`     // Количество подписчиков пользователя.
	FriendStatus     int    `json:"friend_status,omitempty"`       // Статус дружбы с пользователем. Возможные значения: 0 — не является другом; 1 — отправлена заявка/подписка пользователю; 2 — имеется входящая заявка/подписка от пользователя; 3 — является другом.
	Games            string `json:"games,omitempty"`               // Содержимое поля «Любимые игры» из профиля.
	HasMobile        int    `json:"has_mobile,omitempty"`          // Информация о том, известен ли номер мобильного телефона пользователя. Возвращаемые значения: 1 — известен, 0 — не известен.
	HasPhoto         int    `json:"has_photo,omitempty"`           // Информация о том, установил ли пользователь фотографию для профиля. Возвращаемые значения: 1 — установил, 0 — не установил.
	HomeTown         int    `json:"home_town,omitempty"`           // Название родного города.
	Interests        string `json:"interests,omitempty"`           // Содержимое поля «Интересы» из профиля.
	IsFavorite       int    `json:"is_favorite,omitempty"`         // Информация о том, есть ли пользователь в закладках у текущего пользователя. Возможные значения: 1 — есть; 0 — нет.
	IsFriend         int    `json:"is_friend,omitempty"`           // Информация о том, является ли пользователь другом текущего пользователя. Возможные значения: 1 — да; 0 — нет.
	IsHiddenFromFeed int    `json:"is_hidden_from_feed,omitempty"` // Информация о том, скрыт ли пользователь из ленты новостей текущего пользователя. Возможные значения: 1 — да; 0 — нет.
	IsNoIndex        int    `json:"is_no_index,omitempty"`         // Индексируется ли профиль поисковыми сайтами. Возможные значения: 1 — профиль скрыт от поисковых сайтов; 0 — профиль доступен поисковым сайтам. (В настройках приватности: https://vk.com/settings?act=privacy, в пункте «Кому в интернете видна моя страница», выбрано значение «Всем»).
	UserLastName
	LastSeen     UserLastSeen     `json:"last_seen,omitempty"`      // Время последнего посещения.
	Lists        string           `json:"lists,omitempty"`          // Разделенные запятой идентификаторы списков друзей, в которых состоит пользователь. Поле доступно только для метода friends.get.
	MaidenName   string           `json:"maiden_name,omitempty"`    // Девичья фамилия.
	Military     UserMilitary     `json:"military,omitempty"`       // Информация о военной службе пользователя.
	Movies       string           `json:"movies,omitempty"`         // Содержимое поля «Любимые фильмы» из профиля пользователя.
	Music        string           `json:"music,omitempty"`          // Содержимое поля «Любимая музыка» из профиля пользователя.
	Nickname     string           `json:"nickname,omitempty"`       // Никнейм (отчество) пользователя.
	Occupation   UserOccupation   `json:"occupation,omitempty"`     // Информация о текущем роде занятия пользователя.
	Online       int              `json:"online,omitempty"`         // Информация о том, находится ли пользователь сейчас на сайте.
	Personal     UserPersonal     `json:"personal,omitempty"`       // Информация о полях из раздела «Жизненная позиция».
	Photo50      string           `json:"photo_50,omitempty"`       // URL квадратной фотографии пользователя, имеющей ширину 50 пикселей.
	Photo100     string           `json:"photo_100,omitempty"`      // URL квадратной фотографии пользователя, имеющей ширину 100 пикселей.
	Photo200Orig string           `json:"photo_200_orig,omitempty"` // URL фотографии пользователя, имеющей ширину 200 пикселей.
	Photo200     string           `json:"photo_200,omitempty"`      // URL квадратной фотографии, имеющей ширину 200 пикселей.
	Photo400Orig string           `json:"photo_400_orig,omitempty"` // URL фотографии, имеющей ширину 400 пикселей.
	PhotoID      string           `json:"photo_id,omitempty"`       // Строковый идентификатор главной фотографии профиля пользователя в формате {user_id}_{photo_id}, например, 6492_192164258.
	PhotoMax     string           `json:"photo_max,omitempty"`      // URL квадратной фотографии с максимальной шириной. Может быть возвращена фотография, имеющая ширину как 200, так и 100 пикселей.
	PhotoMaxOrig string           `json:"photo_max_orig,omitempty"` // URL фотографии максимального размера. Может быть возвращена фотография, имеющая ширину как 400, так и 200 пикселей.
	Quotes       string           `json:"quotes,omitempty"`         // Любимые цитаты.
	Relatives    []UserRelative   `json:"relatives,omitempty"`      // Список родственников.
	Relation     int              `json:"relation,omitempty"`       // Семейное положение. Возможные значения: 1 — не женат/не замужем; 2 — есть друг/есть подруга; 3 — помолвлен/помолвлена; 4 — женат/замужем; 5 — всё сложно; 6 — в активном поиске; 7 — влюблён/влюблена; 8 — в гражданском браке; 0 — не указано.
	Schools      []UserSchool     `json:"schools,omitempty"`        // Список школ, в которых учился пользователь.
	ScreenName   string           `json:"screen_name,omitempty"`    // Короткое имя страницы.
	Sex          int              `json:"sex,omitempty"`            // Пол. Возможные значения: 1 — женский; 2 — мужской; 0 — пол не указан.
	Site         string           `json:"site,omitempty"`           // Адрес сайта, указанный в профиле.
	Status       string           `json:"status,omitempty"`         // Статус пользователя.
	Timezone     int64            `json:"timezone,omitempty"`       // Временная зона. Только при запросе информации о текущем пользователе.
	Trending     int              `json:"trending,omitempty"`       // Информация о том, есть ли на странице пользователя «огонёк».
	TV           string           `json:"tv,omitempty"`             // Любимые телешоу.
	Universities []UserUniversity `json:"universities,omitempty"`   // Список вузов, в которых учился пользователь.
	Verified     int              `json:"verified,omitempty"`       // Возвращается 1, если страница пользователя верифицирована, 0 — если нет.
	WallDefault  string           `json:"wall_default,omitempty"`   // Режим стены по умолчанию. Возможные значения: owner, all.
}

// UserFirstName Имя пользователя.
type UserFirstName struct {
	FirstNameNom string `json:"first_name_nom,omitempty"` // Имя в именительном падеже.
	FirstNameGen string `json:"first_name_gen,omitempty"` // Имя в родительном падеже.
	FirstNameDat string `json:"first_name_dat,omitempty"` // Имя в дательном падеже.
	FirstNameAcc string `json:"first_name_acc,omitempty"` // Имя в винительном падеже.
	FirstNameIns string `json:"first_name_ins,omitempty"` // Имя в творительном падеже.
	FirstNameAbl string `json:"first_name_abl,omitempty"` // Имя в предложном падеже.
}

// UserLastName Фамилия пользователя.
type UserLastName struct {
	LastNameNom string `json:"last_name_nom,omitempty"` // Фамилия в именительном падеже.
	LastNameGen string `json:"last_name_gen,omitempty"` // Фамилия в родительном падеже.
	LastNameDat string `json:"last_name_dat,omitempty"` // Фамилия в дательном падеже.
	LastNameAc  string `json:"last_name_ac,omitempty"`  // Фамилия в винительном падеже.
	LastNameIns string `json:"last_name_ins,omitempty"` // Фамилия в творительном падеже.
	LastNameAbl string `json:"last_name_abl,omitempty"` // Фамилия в предложном падеже.

}

// UserCareer Информация о карьере пользователя.
type UserCareer struct {
	GroupID   int    `json:"group_id,omitempty"`  // Идентификатор сообщества (если доступно, иначе Company).
	Company   string `json:"company,omitempty"`   // Название компании (если доступно, иначе GroupID).
	CountryID int    `json:"country_id"`          // Идентификатор страны.
	CityID    int    `json:"city_id,omitempty"`   // Идентификатор города (если доступно, иначе CityName)
	CityName  string `json:"city_name,omitempty"` // Название города (если доступно, иначе CityID).
	From      int    `json:"from"`                // Год начала работы.
	Until     int    `json:"until"`               // Год окончания работы.
	Position  string `json:"position"`            // Должность.
}

// UserCity Информация о городе, указанном на странице пользователя в разделе «Контакты».
type UserCity struct {
	ID    int    `json:"id"`    // Идентификатор города, который можно использовать для получения его названия с помощью метода database.getCitiesById;
	Title string `json:"title"` // Название города.
}

// UserContacts Информация о телефонных номерах пользователя.
type UserContacts struct {
	MobilePhone string `json:"mobile_phone"` // Номер мобильного телефона пользователя (только для Standalone-приложений).
	HomePhone   string `json:"home_phone"`   // Дополнительный номер телефона пользователя.
}

// UserCounters Количество различных объектов у пользователя.
type UserCounters struct {
	Albums        int `json:"albums"`         // Количество фотоальбомов.
	Videos        int `json:"videos"`         // Количество видеозаписей.
	Audios        int `json:"audios"`         // Количество аудиозаписей.
	Photos        int `json:"photos"`         // Количество фотографий.
	Notes         int `json:"notes"`          // Количество заметок.
	Friends       int `json:"friends"`        // Количество друзей.
	Groups        int `json:"groups"`         // Количество сообществ.
	OnlineFriends int `json:"online_friends"` // Количество друзей онлайн.
	MutualFriends int `json:"mutual_friends"` // Количество общих друзей.
	UserVideos    int `json:"user_videos"`    // Количество видеозаписей с пользователем.
	Followers     int `json:"followers"`      // Количество подписчиков.
	Pages         int `json:"pages"`          // Количество объектов в блоке «Интересные страницы».
}

// UserCountry Информация о стране, указанной на странице пользователя в разделе «Контакты».
type UserCountry struct {
	ID    int    `json:"id"`    // Идентификатор страны, который можно использовать для получения ее названия с помощью метода database.getCountriesById;
	Title string `json:"title"` // Название страны.
}

// UserCropPhoto Возвращает данные о точках, по которым вырезаны профильная и миниатюрная фотографии пользователя, при наличии.
type UserCropPhoto struct {
	Photo Photo `json:"photo"` // Объект photo фотографии пользователя, из которой вырезается главное фото профиля.
	Crop  Crop  `json:"crop"`  // Вырезанная фотография пользователя.
	Rect  Rect  `json:"rect"`  // Миниатюрная квадратная фотография, вырезанная из фотографии crop.
}

// UserEducation Информация о высшем учебном заведении пользователя.
type UserEducation struct {
	University     int    `json:"university"`      // Идентификатор университета.
	UniversityName string `json:"university_name"` // Название университета.
	Faculty        int    `json:"faculty"`         // Идентификатор факультета.
	FacultyName    string `json:"faculty_name"`    // Название факультета.
	Graduation     int    `json:"graduation"`      // Год окончания.
}

// UserLastSeen Время последнего посещения.
type UserLastSeen struct {
	Time     int64 `json:"time"`     // Время последнего посещения в формате Unix time.
	Platform int   `json:"platform"` // Тип платформы. Возможные значения: 1 — мобильная версия; 2 — приложение для iPhone; 3 — приложение для iPad; 4 — приложение для Android; 5 — приложение для Windows Phone; 6 — приложение для Windows 10; 7 — полная версия сайта.
}

// UserMilitary Информация о военной службе пользователя.
type UserMilitary struct {
	Unit      string `json:"unit"`       // Номер части.
	UnitID    int    `json:"unit_id"`    // Идентификатор части в базе данных.
	CountryID int    `json:"country_id"` // Идентификатор страны, в которой находится часть.
	From      int    `json:"from"`       // Год начала службы.
	Until     int    `json:"until"`      // Год окончания службы.
}

// UserOccupation Информация о текущем роде занятия пользователя.
type UserOccupation struct {
	Type string `json:"type"` // Тип. Возможные значения: work — работа; school — среднее образование; university — высшее образование.
	ID   int    `json:"id"`   // Идентификатор школы, вуза, сообщества компании (в которой пользователь работает).
	Name string `json:"name"` // Название школы, вуза или места работы.
}

// UserPersonal Информация о полях из раздела «Жизненная позиция».
type UserPersonal struct {
	Political  int      `json:"political"`   // Политические предпочтения. Возможные значения: 1 — коммунистические; 2 — социалистические; 3 — умеренные; 4 — либеральные; 5 — консервативные; 6 — монархические; 7 — ультраконсервативные; 8 — индифферентные; 9 — либертарианские.
	Langs      []string `json:"langs"`       // Языки.
	Religion   string   `json:"religion"`    // Мировоззрение.
	InspiredBy string   `json:"inspired_by"` // Источники вдохновения.
	PeopleMain int      `json:"people_main"` // Главное в людях. Возможные значения: 1 — ум и креативность; 2 — доброта и честность; 3 — красота и здоровье; 4 — власть и богатство; 5 — смелость и упорство; 6 — юмор и жизнелюбие.
	LifeMain   int      `json:"life_main"`   // Главное в жизни. Возможные значения: 1 — семья и дети; 2 — карьера и деньги; 3 — развлечения и отдых; 4 — наука и исследования; 5 — совершенствование мира; 6 — саморазвитие; 7 — красота и искусство; 8 — слава и влияние.
	Smoking    int      `json:"smoking"`     // Отношение к курению. Возможные значения: 1 — резко негативное; 2 — негативное; 3 — компромиссное; 4 — нейтральное; 5 — положительное.
	Alcohol    int      `json:"alcohol"`     // Отношение к алкоголю. Возможные значения: 1 — резко негативное; 2 — негативное; 3 — компромиссное; 4 — нейтральное; 5 — положительное.
}

// UserRelative Родственник пользователя.
type UserRelative struct {
	ID   int    `json:"id,omitempty"` // Идентификатор пользователя.
	Name string `json:"name"`         // Имя родственника (если родственник не является пользователем ВКонтакте, то предыдущее значение ID возвращено не будет).
	Type string `json:"type"`         // Тип родственной связи. Возможные значения: child — сын/дочь; sibling — брат/сестра; parent — отец/мать; grandparent — дедушка/бабушка; grandchild — внук/внучка.
}

// UserSchool Школа пользователя.
type UserSchool struct {
	ID            string `json:"id"`             // Идентификатор школы.
	Country       int    `json:"country"`        // Идентификатор страны, в которой расположена школа.
	City          int    `json:"city"`           // Идентификатор города, в котором расположена школа.
	Name          string `json:"name"`           // Наименование школы.
	YearFrom      int    `json:"year_from"`      // Год начала обучения.
	YearTo        int    `json:"year_to"`        // Год окончания обучения.
	YearGraduated int    `json:"year_graduated"` // Год выпуска.
	Class         string `json:"class"`          // Буква класса.
	Speciality    string `json:"speciality"`     // Специализация.
	Type          int    `json:"type"`           // Идентификатор типа.
	TypeStr       string `json:"type_str"`       // Название типа. Возможные значения: 0 — "школа"; 1 — "гимназия"; 2 — "лицей"; 3 — "школа-интернат"; 4 — "школа вечерняя"; 5 — "школа музыкальная"; 6 — "школа спортивная"; 7 — "школа художественная"; 8 — "колледж"; 9 — "профессиональный лицей"; 10 — "техникум"; 11 — "ПТУ"; 12 — "училище"; 13 — "школа искусств".
}

// UserUniversity Вуз пользователя.
type UserUniversity struct {
	ID              int    `json:"id"`               // Идентификатор университета.
	Country         int    `json:"country"`          // Идентификатор страны, в которой расположен университет.
	City            int    `json:"city"`             // Идентификатор города, в котором расположен университет.
	Name            string `json:"name"`             // Наименование университета.
	Faculty         int    `json:"faculty"`          // Идентификатор факультета.
	FacultyName     string `json:"faculty_name"`     // Наименование факультета.
	Chair           int    `json:"chair"`            // Идентификатор кафедры.
	ChairName       string `json:"chair_name"`       // Наименование кафедры.
	Graduation      int    `json:"graduation"`       // Год окончания обучения.
	EducationForm   string `json:"education_form"`   // Форма обучения.
	EducationStatus string `json:"education_status"` // Статус (например, «Выпускник (специалист)»).
}
