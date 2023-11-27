package api

import (
	"math/rand"
	"regexp"
	"strconv"
)

func (rt *_router) GetUserFromLogin(login Login) (User, error) {
	dbUser, err := rt.db.GetDatabaseUser(login.LoginIntoDatabaseLogin())

	if err != nil {
		return UserDefault(), err
	}

	user := UserFromDatabaseUser(dbUser)

	return user, nil
}

func CheckAuthorization(user User, authRaw string) error {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

    tokenString := re.FindAllString(authRaw, -1)

	if len(tokenString) == 0 {
		return ErrUserUnauthorized
	}

    token, _ := strconv.Atoi(tokenString[0])

    if int(user.Id) != token {
		return ErrUserUnauthorized
	}

	return nil
}

func (rt *_router) GenerateRandomPhotoId() (uint64, error) {
	newId := rand.Uint64()

	for {
		validPhotoId, err := rt.db.CheckIfAvailablePhotoId(newId)

		if err != nil {
			return 0, err
		}

		if validPhotoId {
			break
		}

		newId = rand.Uint64()
	}

	return newId, nil
}