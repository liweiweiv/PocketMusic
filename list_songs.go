package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
)

type GetListAndSongsReq struct {
	Lid uint `json:"lid" form:"lid"`
	Mid uint `json:"mid" form:"mid"`
}

//type ListSongs struct {
//	Mid       uint   `json:"mid"`
//	Name      string `json:"name"`
//	MusicList *uint  `json:"music_list`
//}

func HandleGetListSongs(c *gin.Context) {
	var req GetListAndSongsReq
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	result, err := dal.GetListSongs(req.Lid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", result)
}

func HandleAddSongToList(c *gin.Context) {
	var req GetListAndSongsReq
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	stauts, err := dal.AddSongToList(req.Lid, req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", stauts)
}

func HandleDeleteSongFromList(c *gin.Context) {
	var req GetListAndSongsReq
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	status, err := dal.DeleteSongFromList(req.Lid, req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", status)
}
