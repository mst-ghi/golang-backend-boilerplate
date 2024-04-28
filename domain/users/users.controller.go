package users

import (
	"app/core"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	root    *core.Controller
	service UsersServiceInterface
}

func NewUsersController() *UsersController {
	return &UsersController{
		root:    core.GetController(),
		service: NewUsersService(),
	}
}

// @tags     Users
// @security Bearer
// @router   /api/v1/users [get]
// @summary  get list of users
// @accept   json
// @produce  json
// @Param    search query string false "search on user name"
// @Param    page query string false "pagination page_value, default 1"
// @Param    take query string false "pagination take_value, default 20"
// @success  200 {object} core.Response[UsersMetaResponseType]
func (ctrl *UsersController) FindAll(c *gin.Context) {
	search, page, take := core.PaginateQueries(c)
	users, meta := ctrl.service.FindAll(search, page, take)
	ctrl.root.Success(c, UsersMetaResponse(users, meta))
}

// @tags     Users
// @security Bearer
// @router   /api/v1/users/{id} [get]
// @summary  get user by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[UserResponseType]
// @param    id path string true "User ID"
func (ctrl *UsersController) FindOne(c *gin.Context) {
	user, err := ctrl.service.FindOne(c.Param("id"))

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, UserResponse(user))
}
