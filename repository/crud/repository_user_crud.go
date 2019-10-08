package crud

import (
	"blogos/models"

	"github.com/jinzhu/gorm"
)

type repositoryuserCRUD struct {
	db *gorm.DB
}

func NewRepositoryUserCRUD(db *gorm.DB) *repositoryuserCRUD {
	return &repositoryuserCRUD{db}
}

func (r *repositoryuserCRUD) Save(user models.User) (models.User,error){
	var err error
	done:=make(chan bool)
	go func (ch chan<- bool)  {
		defer close(done)
		if err=r.db.Debug().Model(&models.User{}).Create(&user).Error;err!=nil{
			done <- false
			return
		}
		done<-true
	}(done)
	if res:=<-done;res{
		return user,nil
	}else{
		return models.User{},err
	}
}


func (r *repositoryuserCRUD) FindAll() ([]models.User,error){
	var err error
	var users []models.User
	done:=make(chan bool)
	go func (ch chan<- bool)  {
		defer close(done)
		if err=r.db.Debug().Model(&models.User{}).Select("username,id").Limit(10).Find(&users).Error;err!=nil{
			done <- false
			return
		}
		done<-true
	}(done)
	if res:=<-done;res{
		return users,nil
	}else{
		return nil,err
	}
}

func (r *repositoryuserCRUD) FindById(id uint) (models.User,error){
	var err error
	var user models.User
	done:=make(chan bool)
	go func (ch chan<- bool)  {
		defer close(done)

		if err=r.db.Debug().Model(&models.User{}).Select("username").Take(&user).Error;err!=nil{
			done <- false
			return
		}
		done<-true
	}(done)
	if res:=<-done;res{
		return models.User{},nil
	}else{
		return models.User{},err
	}
}


func (r *repositoryuserCRUD) UpdateUser(id uint,user models.User) (int64,error){
	var err error
	done:=make(chan bool)
	go func (ch chan<- bool)  {
		defer close(done)
		if err=r.db.Debug().Model(&models.User{}).Where("id = ?",id).Take(&user).UpdateColumns(&user).Error;err!=nil{
			done <- false
			return
		}
		done<-true
	}(done)
	if res:=<-done;res{
		return 1,nil
	}else{
		return 0,err
	}
}

func (r *repositoryuserCRUD) DeleteUser(id uint) (int64,error){
	var err error
	done:=make(chan bool)
	go func (ch chan<- bool)  {
		defer close(done)
		if err=r.db.Debug().Model(&models.User{}).Where("id = ?",id).Take(&models.User{}).Delete(models.User{}).Error;err!=nil{
			done <- false
			return
		}
		done<-true
	}(done)
	if res:=<-done;res{
		return 1,nil
	}else{
		return 0,err
	}
}
