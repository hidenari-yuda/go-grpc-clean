package repository

import (
	"fmt"
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utility"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type UserRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewUserRepositoryImpl(ex SQLExecuter) usecase.UserRepository {
	return &UserRepositoryImpl{
		Name:     "UserRepository",
		executer: ex,
	}
}

func (r *UserRepositoryImpl) SignUp(param *entity.SignUpParam) error {
	_, err := r.executer.Exec(
		"SignUp",
		`INSERT INTO users (
			uuid,
			firebase_id,
			name, 
			email, 
			password,
			created_at,
			updated_at
			) VALUES (
				?,
				?,
				?, 
				?,
				?,
				?,
				?
		)`,
		utility.CreateUUID(),
		"",
		"ゲスト",
		param.Email,
		param.Password,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// func (r *UserRepositoryImpl) SignIn(param *entity.SignInParam) (user *entity.User, err error) {
// 	err = r.executer.Get(
// 		"SignIn",
// 		user,
// 		"SELECT * FROM users WHERE email = ? AND password = ?",
// 		param.Email, param.Password)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

func (r *UserRepositoryImpl) GetByFirebaseId(firebaseId string) (*entity.User, error) {
	var (
		user entity.User
	)
	err := r.executer.Get(
		"GetByFirebaseId",
		&user,
		"SELECT * FROM users WHERE firebase_id = ?",
		firebaseId,
	)

	if err != nil {
		err = fmt.Errorf("failed to get user by firebase id: %w", err)
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
