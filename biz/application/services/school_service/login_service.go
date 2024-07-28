package school_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/onebillion-0/user_sdk/biz/constants"
	"github.com/onebillion-0/user_sdk/biz/domain/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginService struct {
	Student repositories.MemberRepository
}

type Claims struct {
	Uid  int64  `json:"uid"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("your_secret_key")

func NewLoginService(stu repositories.MemberRepository) *LoginService {
	return &LoginService{
		Student: stu,
	}
}

func (s *LoginService) Login(ctx context.Context, uid int64, password string) (string, error) {
	member, err := s.Student.FindByID(ctx, uid)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return "", constants.ERROR_INVALID_USERNAME_OR_PASSWORD
	}
	if err != nil {
		fmt.Println("find admin fail, error:", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password)); err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Uid:  member.Uid,
		Role: string(member.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (s *LoginService) ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
