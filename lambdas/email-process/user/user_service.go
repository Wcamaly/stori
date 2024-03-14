package user

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
)

type UserConfig struct {
	BaseUrl string
	Timeout int
	Cache   map[string]*UserData
	Mu      sync.Mutex
}

func NewUserService(baseUrl string, timeout int) *UserConfig {
	return &UserConfig{
		BaseUrl: baseUrl,
		Timeout: timeout,
		Cache:   make(map[string]*UserData),
	}
}

func (u *UserConfig) GetUser(email string) *UserData {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	if userData, found := u.Cache[email]; found {
		return userData
	}

	userData, err := u.GetUserByEmail(email)
	if err != nil {
		return nil
	}
	u.Cache[email] = userData
	return userData
}

func (u *UserConfig) GetUserByEmail(email string) (*UserData, error) {
	data, err := http.Get(u.BaseUrl + "/user?email=" + url.QueryEscape(email))
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	var res *UserData
	err = json.NewDecoder(data.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserConfig) GetUsers() ([]*UserData, error) {
	res := make([]*UserData, 0)
	for _, userData := range u.Cache {
		res = append(res, userData)
	}
	return res, nil
}
