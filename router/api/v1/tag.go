package v1

import (
	"github.com/gin-gonic/gin"

	"qiu/blog/pkg/util"
	"qiu/blog/service"

	"qiu/blog/pkg/setting"

	"net/http"

	"qiu/blog/pkg/e"

	gin_http "qiu/blog/pkg/http"
)

//获取多个文章标签
// @Summary 获取多个文章标签
// @Produce  json
// @Param page query int false "Page"
// @Success 200 {object}  gin_http.ResponseJSON
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	tagService := service.GetTagService()

	data := make(map[string]int)
	data["pageNum"] = util.GetPage(c)
	data["pageSize"] = setting.AppSetting.PageSize

	tags := tagService.Get(data)
	gin_http.Response(c, http.StatusOK, e.SUCCESS, tags)
}

// @Summary 新增标签
// @Produce  json
// @Param name formData string true "Name"
// @Param created_by formData string false "CreatedBy"
// @Param token header string true "token"
// @Success 200 {object} gin_http.ResponseJSON
// @Failure  400 {object} gin_http.ResponseJSON
// @Failure  10001 {object} gin_http.ResponseJSON
// @Failure  10006 {object} gin_http.ResponseJSON
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {

	tagService := service.GetTagService()
	httpCode, errCode := tagService.Bind(c)

	if errCode != e.SUCCESS {
		gin_http.Response(c, httpCode, errCode, nil)
		return
	}
	err := tagService.Valid()
	if err != nil {
		gin_http.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	state := tagService.ExistTagByName()
	if state {
		gin_http.Response(c, http.StatusBadRequest, e.ERROR_EXIST_TAG, nil)
		return
	}

	claims := tagService.GetClaimsFromToken(c)
	if claims == nil {
		gin_http.Response(c, http.StatusBadRequest, e.ERROR_AUTH, nil)
		return
	}

	created_by := tagService.GetCreatedBy()
	if created_by == "" {
		tagService.SetCreatedBy(claims.Username)
	} else {
		if created_by != claims.Username {
			gin_http.Response(c, http.StatusBadRequest, e.ERROR_AUTH, nil)
			return
		}
	}

	err = tagService.Add()
	if err != nil {
		gin_http.Response(c, http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	gin_http.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// @Summary 修改标签
// @Produce  json
// @Param id formData int true "Id"
// @Param name formData string true "Name"
// @Param modified_by formData string false "Modifiedby"
// @Param token header string true "token"
// @Success 200 {object} gin_http.ResponseJSON
// @Failure  400 {object} gin_http.ResponseJSON
// @Failure  10007 {object} gin_http.ResponseJSON
// @Router /api/v1/tags [put]
func EditTag(c *gin.Context) {
	tagService := service.GetTagService()
	httpCode, errCode := tagService.Bind(c)

	if errCode != e.SUCCESS {
		gin_http.Response(c, httpCode, errCode, nil)
		return
	}

	err := tagService.Valid()
	if err != nil {
		gin_http.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	claims := tagService.GetClaimsFromToken(c)
	if claims == nil {
		gin_http.Response(c, http.StatusBadRequest, e.ERROR_AUTH, nil)
		return
	}

	modified_by := tagService.GetModifiedBy()
	if modified_by == "" {
		tagService.SetModifiedBy(claims.Username)
	} else {
		if modified_by != claims.Username {
			gin_http.Response(c, http.StatusBadRequest, e.ERROR_AUTH, nil)
			return
		}
	}

	state := tagService.Update()
	if !state {
		gin_http.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	gin_http.Response(c, http.StatusOK, e.SUCCESS, nil)

}

// @Summary 删除文章标签
// @Produce  json
// @Param id formData int true "Id"
// @Param name formData string true "Name"
// @Param token header string true "token"
// @Success 200 {object} gin_http.ResponseJSON
// @Failure  400 {object} gin_http.ResponseJSON
// @Failure  10008 {object} gin_http.ResponseJSON
// @Router /api/v1/tags [delete]
func DeleteTag(c *gin.Context) {
	tagService := service.GetTagService()
	httpCode, errCode := tagService.Bind(c)

	if errCode != e.SUCCESS {
		gin_http.Response(c, httpCode, errCode, nil)
		return
	}
	err := tagService.Valid()
	if err != nil {
		gin_http.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if state := tagService.Delete(); !state {
		gin_http.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}
	gin_http.Response(c, http.StatusOK, e.SUCCESS, nil)
}
