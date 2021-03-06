package dal

import (
	"time"
	"PocketMusic/dal/model"
)

func GetLikeStuatus(Uid string,Mid uint)(bool,error)  {
	count:=0
	//println("mid:",Mid)
	err:=db.Table("likes").Where("mid = ? and status = 0 and uid =?",Mid,Uid).Count(&count).Error
	//println("count:",count)
	if count>0{
		return true,err
	}
	return false,err
}

func AddLike(Uid string,Mid uint)(bool,error)  {
	count:=0
	err:=db.Table("likes").Where("mid =? and uid =?",Mid,Uid).Count(&count).Error
	if err!=nil{
		return false,err
	}else if count>0{
		_,err:=db.DB().Exec("update likes set status = 0 where mid =? and Uid = ?",Mid,Uid)
		if err!=nil{
			return false,err
		}
	}else{
		_,err:=db.DB().Exec("insert into likes(uid,mid,status,update_at,create_at) values(?,?,?,?,?)",Uid,Mid,0,time.Now().Local(),time.Now().Local())
		if err !=nil{
			return false,err
		}
	}
	return true,err
}
func DeleteLike(Uid string,Mid uint)(bool,error) {
	_,err:=db.DB().Exec("update likes set status = 1 where mid = ? and uid=?",Mid,Uid)
	if err!=nil {
		return false, err
	}
	return true,err
}

func GetMusic(Mid uint,music *model.MusicInfo)  (error) {
	//println("mid:",Mid)
	count:=0
	err:=db.Table("songs").Where("mid =?",Mid).Count(&count).Error
	if err!=nil{
		return err
	}
	if count>0 {
		row := db.DB().QueryRow("select mid,mname,singer,lrc,url from songs where mid =?", Mid)
		err := row.Scan(&music.Mid, &music.Mname, &music.Singer, &music.Lrc, &music.Source)
		//println("music:", music.Singer)
		if err != nil {
			return err
		}
	}
	return err
}