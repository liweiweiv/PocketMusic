package dal

import (
	"PocketMusic/dal/model"
	"time"
)

func AddSongToList(Mid uint, Lid uint) (bool, error) {
	count := 0
	err := db.Table("list_songs").Where("mid =?", Mid).Count(&count).Error
	if err != nil {
		return false, err
	} else if count == 0 {
		_, err := db.DB().Exec("insert into list_songs(lid,mid,status,update_at,create_at) values(?,?,?,?,?,?)", Lid, Mid, 1, time.Now().Local(), time.Now().Local())
		if err != nil {
			return false, err
		}
	}
	return true, err
}

func DeleteSongFromList(Mid uint, Lid uint) (bool, error) {
	_, err := db.DB().Exec("update list_songs set status = 1 where lid = ? and mid = ?", Lid, Mid)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetListSongs(Lid uint) ([]map[string]interface{}, error) {
	whereParams := make(map[string]interface{})
	whereParams["lid"] = Lid
	whereParams["status"] = 0
	condition := CombineCondition(whereParams)
	var listsongs []*model.List_songs
	err := db.Where(condition).Find(&listsongs).Error
	result := make([]map[string]interface{}, 0)
	for key := range listsongs {
		res := make(map[string]interface{})
		res["id"] = listsongs[key].Id
		res["mid"] = listsongs[key].Mid
		result = append(result, res)
	}
	return result, err
}
