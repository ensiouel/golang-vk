package vk

import (
	"net/url"

	"github.com/ensiouel/golang-vk/objects"
)

type Wall struct {
	Count int            `json:"count"`
	Items []objects.Post `json:"items"`
}

func (session *Session) WallGet(owner interface{}, offset, count int, params url.Values) (wall *Wall, err error) {
	request := session.Call("wall.get", params)

	switch owner.(type) {
	case int:
		request.Set("owner_id", owner)
	case string:
		request.Set("domain", owner)
	}

	request.Set("offset", offset)
	request.Set("count", count)

	err = request.Execute(&wall)

	return
}

func (session *Session) WallPost(ownerID int, message, attachments string) (postID int, err error) {
	request := session.Call("wall.post", nil).
		Set("owner_id", ownerID).
		Set("message", message).
		Set("attachments", attachments)

	m := map[string]int{}
	err = request.Execute(&m)

	return m["post_id"], err
}
