package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
	"wanben/myebms/config"
	"wanben/myebms/model"
	"wanben/myebms/query"
	"wanben/myebms/repository"
	"wanben/myebms/utils"
)

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	Login(mobile string, password string) error
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	EditLocked(id string) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func (srv *UserService) List(req *query.ListQuery) (users []*model.User, err error) {
	if req.PageSize < 1 {
		req.PageSize = config.PAGE_SIZE
	}
	return srv.Repo.List(req)
}
func (srv *UserService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *UserService) Get(user model.User) (*model.User, error) {
	return srv.Repo.Get(user)
}
func (srv *UserService) Exist(user model.User) *model.User {
	return srv.Repo.Exist(user)
}

func (srv *UserService) ExistByUserID(id string) *model.User {
	return srv.Repo.ExistByUserID(id)
}

func (srv *UserService) Login(mobile string, password string) error {
	result2 := srv.Repo.ExistByPassword(mobile, password)
	if result2 == nil {
		return errors.New("用户不存在或密码错误")
	}
	return nil
}

func (srv *UserService) Add(user model.User) (*model.User, error) {
	//根据手机号判断是否存在用户
	result := srv.Repo.ExistByMobile(user.Mobile)
	if result != nil {
		return nil, errors.New("用户已经存在")
	}
	user.UserId = uuid.NewV4().String()
	/*密码为空，密码默认为123456*/
	if user.Password == "" {
		user.Password = utils.Md5("123456")
	}
	/*对密码进行加密*/
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, errors.New("加密错误")
	}
	user.Password = string(hasedPassword)
	user.IsDeleted = false
	user.IsLocked = false
	return srv.Repo.Add(user)
}
func (srv *UserService) Edit(user model.User) (bool, error) {
	if user.UserId == "" {
		return false, fmt.Errorf("参数错误")
	}

	exist := srv.Repo.ExistByUserID(user.UserId)
	if exist == nil {
		return false, errors.New("参数错误")
	}
	exist.NickName = user.NickName
	exist.Mobile = user.Mobile
	exist.Address = user.Address
	exist.UpdateAt = user.UpdateAt
	return srv.Repo.Edit(*exist)
}

func (srv *UserService) EditLocked(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}

	user := srv.ExistByUserID(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	user.IsLocked = !user.IsLocked
	return srv.Repo.EditLocked(*user)
}

func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}

	user := srv.ExistByUserID(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	user.IsDeleted = !user.IsDeleted
	return srv.Repo.Delete(*user)
}
