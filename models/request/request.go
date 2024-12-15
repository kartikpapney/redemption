package requestModel

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 32),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Password string `json:"password"`
}

func (a RegisterRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a LoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type RefreshRequest struct {
	Token string `json:"token"`
}

func (a RefreshRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(
			&a.Token,
			validation.Required,
			validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
		),
	)
}

type VerifyRequest struct {
	Token string `form:"token" binding:"required"`
}

func (a VerifyRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(
			&a.Token,
			validation.Required,
			validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
		),
	)
}

type CreateProductRequest struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
}

func (a CreateProductRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Manufacturer, validation.Required),
	)
}

type PaginatedRequest struct {
	Limit int64 `form:"limit" json:"limit"`
	Page  int64 `form:"page" json:"page"`
}

func NewPaginatedRequest(limit int64, page int64) (*PaginatedRequest, error) {
	paginatedRequest := PaginatedRequest{
		Limit: int64(limit),
		Page:  int64(page),
	}

	err := validation.ValidateStruct(&paginatedRequest,
		validation.Field(&paginatedRequest.Limit, validation.Required, validation.Min(1), validation.Max(10)),
		validation.Field(&paginatedRequest.Page, validation.Required, validation.Min(1)),
	)

	if err != nil {
		return nil, err
	}

	return &paginatedRequest, nil
}

type PathIdRequest struct {
	Id primitive.ObjectID `form:"id"`
}

func NewPathIdRequest(id string) (*PathIdRequest, error) {
	result, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	pathIdRequest := &PathIdRequest{
		Id: result,
	}
	return pathIdRequest, nil
}

type RequestMetadata struct {
	UserId primitive.ObjectID `json:"id" bson:"_id"`
}

func NewRequestMetadata(userId primitive.ObjectID) (*RequestMetadata, error) {
	return &RequestMetadata{
		UserId: userId,
	}, nil
}
