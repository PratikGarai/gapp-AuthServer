package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorErrorParser(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
			fmt.Println(err)
		}
		errs[f.StructField()] = err
	}
	return errs
}

func HandleJsonParsingError(c *gin.Context, err error) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		c.AbortWithStatusJSON(http.StatusBadRequest, ValidatorErrorParser(verr))
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error parsing JSON Body")
	}
}
