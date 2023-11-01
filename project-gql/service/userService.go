package service

import (
	"errors"
	"project-gql/graph/model"
	custommodel "project-gql/models"
	"project-gql/repository"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r repository.Users
	c repository.Company
}

func NewService(r repository.Users, c repository.Company) (*Service, error) {
	if r == nil {
		return nil, errors.New("db connection not given")
	}

	return &Service{r: r, c: c}, nil

}

func (s *Service) UserSignup(nu model.User) (*model.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return &model.User{}, errors.New("hashing password failed")
	}
	user := custommodel.User{UserName: nu.UserName, Email: nu.Email, PasswordHash: string(hashedPass)}

	cu, err := s.r.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return &model.User{}, errors.New("user creation failed")
	}
	user1 := model.User{UserName: cu.UserName, Email: cu.Email, PasswordHash: cu.PasswordHash}
	return &user1, nil

}
func (s *Service) Userlogin(l model.UserLogin) (*model.User, error) {
	fu, err := s.r.FetchUserByEmail(l.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return nil, errors.New("user login failed")
	}
	err = bcrypt.CompareHashAndPassword([]byte(fu.PasswordHash), []byte(l.Password))
	if err != nil {
		log.Error().Err(err).Msg("password of user incorrect")
		return nil, errors.New("user login failed")
	}
	// c := jwt.RegisteredClaims{
	// 	Issuer:    "service project",
	// 	Subject:   "2",io0ujhnmvc65:f
	// 	Audience:  jwt.ClaimStrings{"users"},
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	// 	IssuedAt:  jwt.NewNumericDate(time.Now()),
	// }
	//fmt.Println(c)
	ff := model.User{UserName: fu.UserName, Email: fu.Email}
	return &ff, nil

}
