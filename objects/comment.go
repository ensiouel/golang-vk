package objects

// Comment Объект, описывающий комментарии на стене
type Comment struct {
	ID             int             `json:"id"`                      // Идентификатор комментария.
	FromID         int             `json:"from_id"`                 // Идентификатор автора комментария.
	Date           int64           `json:"date"`                    // Дата создания комментария в формате Unix time.
	Text           string          `json:"text"`                    // Текст комментария.
	Donut          CommentDonut    `json:"donut"`                   // Информация о VK Donut.
	ReplyToUser    int             `json:"reply_to_user,omitempty"` // Идентификатор пользователя или сообщества, в ответ которому оставлен текущий комментарий (если применимо).
	ReplyToComment int             `json:"reply_to_comment"`        // Идентификатор комментария, в ответ на который оставлен текущий (если применимо).
	Attachments    PostAttachments `json:"attachments"`             // Медиа вложения комментария (фотографии, ссылки и т.п.).
	ParentsStack   []int           `json:"parents_stack"`           // Массив идентификаторов родительских комментариев.
	Thread         CommentThread   `json:"thread"`                  // Информация о вложенной ветке комментариев.
}

// CommentDonut Информация о VK Donut.
type CommentDonut struct {
	IsDon       bool   `json:"is_don"`      // Является ли комментатор подписчиком VK Donut.
	Placeholder string `json:"placeholder"` // Заглушка для пользователей, которые не оформили подписку VK Donut.
}

// CommentThread Информация о вложенной ветке комментариев.
type CommentThread struct {
	Count           int       `json:"count"`             // Количество комментариев в ветке.
	Items           []Comment `json:"items,omitempty"`   // Массив объектов комментариев к записи (только для метода wall.getComments).
	CanPost         bool      `json:"can_post"`          // Может ли текущий пользователь оставлять комментарии в этой ветке.
	ShowReplyButton bool      `json:"show_reply_button"` // Нужно ли отображать кнопку «ответить» в ветке.
	GroupsCanPost   bool      `json:"groups_can_post"`   // Могут ли сообщества оставлять комментарии в ветке.
}
