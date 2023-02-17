package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type UserInteractor interface {
	// Gest API
	// Create
	Create(user *entity.User) (*entity.User, error)

	// Update
	Update(user *entity.User) (*entity.User, error)

	// Get
	GetById(id uint) (*entity.User, error)
	SignIn(param *entity.SignInParam) (*entity.User, error)
	GetByFirebaseToken(token string) (*entity.User, error)

	// admin API
	GetAll() ([]*entity.User, error)
}

type UserInteractorImpl struct {
	firebase usecase.Firebase
	// userClient     usecase.UserClient
	userRepository usecase.UserRepository
}

func NewUserInteractorImpl(
	fb usecase.Firebase,
	// uc usecase.UserClient,
	uR usecase.UserRepository,
) UserInteractor {
	return &UserInteractorImpl{
		firebase: fb,
		// userClient:     uc,
		userRepository: uR,
	}
}

// Create
// type Createuser struct {
// 	Param *entity.User
// }

// type CreateOutput struct {
// 	Ok bool
// }

func (i *UserInteractorImpl) Create(user *entity.User) (*entity.User, error) {

	// ユーザー登録
	err := i.userRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update
func (i *UserInteractorImpl) Update(user *entity.User) (*entity.User, error) {

	// ユーザー登録
	err := i.userRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetById
func (i *UserInteractorImpl) GetById(id uint) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	// ユーザー登録
	user, err = i.userRepository.GetById(id)
	if err != nil {
		fmt.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// SignIn
func (i *UserInteractorImpl) SignIn(param *entity.SignInParam) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	firebaseId, err := i.firebase.VerifyIDToken(param.Token)
	if err != nil {
		return user, err
	}

	fmt.Println("exported firebaseToken is:", param.Token)
	fmt.Println("exported firebaseId is:", firebaseId)

	// ユーザー登録
	user, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return user, err
	}

	return user, nil

}

// GetByFirebaseToken
func (i *UserInteractorImpl) GetByFirebaseToken(token string) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	firebaseId, err := i.firebase.VerifyIDToken(token)
	if err != nil {
		return user, err
	}

	fmt.Println("exported firebaseId is:", firebaseId)

	user, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return user, err
	}

	fmt.Println("exported user is:", user)

	return user, nil
}

// GetAll
func (i *UserInteractorImpl) GetAll() ([]*entity.User, error) {
	var (
		users []*entity.User
		err   error
	)

	users, err = i.userRepository.GetAll()
	if err != nil {
		return users, err
	}

	// i.userClient.DetectTextFromImage()

	return users, nil
}
