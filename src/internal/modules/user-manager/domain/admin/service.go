package admin

type AdminService interface {
	GetByID(input *GetAdminInput) (*Admin, error)
	GetByEmail(input *GetAdminByEmailInput) (*Admin, error)
	Create(input *CreateAdminInput) (*CreateAdminOutput, error)
	Update(input *UpdateAdminInput) (*UpdateAdminOutput, error)
	Delete(input *DeleteAdminInput) (*DeleteAdminOutput, error)
}
