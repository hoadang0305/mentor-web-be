package validation

import (
	"encoding/json"
	"github.com/hoadang0305/mentor-web-be/internal/domain/http_common"
	stringutils "github.com/hoadang0305/mentor-web-be/internal/utils/string_utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func BindJsonAndValidate(c *gin.Context, dest interface{}) error {
	err := c.ShouldBindJSON(&dest)

	if err != nil {
		checkErr(c, err)
	}
	return err
}

func checkErr(c *gin.Context, err error) {
	switch t := err.(type) {
	case *json.UnmarshalTypeError:
		httpErr := http_common.Error{
			Message: http_common.ErrorMessage.InvalidDataType, Code: http_common.ErrorResponseCode.InvalidDataType, Field: t.Field,
		}
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(httpErr))
		return
	case *json.SyntaxError:
		httpErr := http_common.Error{Message: err.Error(), Code: http_common.ErrorResponseCode.InvalidRequest}
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(httpErr))
		return
	case validator.ValidationErrors:
		httpErrs := handleValidationErrors(err)
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(httpErrs...))
		return
	default:
		httpErr := http_common.Error{Message: err.Error(), Code: http_common.ErrorResponseCode.InvalidRequest, Field: ""}
		c.JSON(http.StatusBadRequest, httpErr)
		return
	}
}

func handleValidationErrors(errs error) (httpErrs []http_common.Error) {
	for _, fieldErr := range errs.(validator.ValidationErrors) {
		errCodeWithStructNs := http_common.CustomValidationErrCode[strings.ToLower(fieldErr.StructNamespace())]
		field := stringutils.FirstLetterToLower(fieldErr.Field())
		if errCodeWithStructNs == "" {
			// handle builtin validation
			httpErrs = append(httpErrs, http_common.Error{
				Message: http_common.ErrorMessage.InvalidRequest, Code: http_common.ErrorResponseCode.InvalidRequest, Field: field,
			})
		} else {
			// handle custom validation
			httpErrs = append(httpErrs, http_common.Error{
				Message: http_common.ErrorMessage.InvalidRequest, Code: errCodeWithStructNs, Field: field,
			})
		}
	}
	return httpErrs
}
