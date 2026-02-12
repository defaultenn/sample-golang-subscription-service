package erroring

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"test_task/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func TranslateTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Field is required."
	case "email":
		return "Incorrect email address."
	case "gt":
		return fmt.Sprintf("Value must be greather than %d.", fe.Value())
	case "uuid4":
		return "Value must be in format of UUID4."
	default:
		return fe.Tag()
	}
}

func HandleConcreteError(ctx *gin.Context, err error) {
	switch err {
	// перечисление и обработка ошибок слоя usecase
	default:
		{
			zerolog.Ctx(ctx.Request.Context()).Error().Fields(
				map[string]any{
					"error":       err.Error(),
					"stack_trace": string(debug.Stack()),
				},
			).Msg("unprocessed error")

			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				&HTTPInternalServerError{
					Message: err.Error(),
					Code:    InternalServerCode,
				},
			)
		}
	}
}

func HandleDatabaseError(ctx *gin.Context, err error) {
	// pgError := &pgconn.PgError{}

	// if errors.As(err, &pgError) {
	// 	message, exist := UNIQUE_CONSTRAINTS_TRANSLATIONS[pgError.ConstraintName]

	// 	if exist {
	// 		c.AbortWithStatusJSON(
	// 			http.StatusBadRequest,
	// 			&HTTPBadRequest[any]{
	// 				Message: message,
	// 			},
	// 		)
	// 		return
	// 	}
	// }

	zerolog.Ctx(ctx.Request.Context()).Error().Fields(
		map[string]any{
			"error": err.Error(),
		},
	).Msg("database error")

	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		&HTTPInternalServerError{
			Code:    InternalServerCode,
			Message: constants.ErrInternalServerError.Error(),
		},
	)
}

func Handle(ctx *gin.Context, err error) {

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Status(404)
		return
	}

	switch err.(type) {
	case *pgconn.PgError:
		HandleDatabaseError(ctx, err)
	case validator.ValidationErrors:
		{
			httpError := HTTPRequestValidationError{
				Message: constants.ErrIncorrectInputData.Error(),
				Code:    RequestValidationCode,
			}

			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				httpError.Data = make([]ValidationErrorFieldDTO, len(ve))
				for i, fe := range ve {
					httpError.Data[i] = ValidationErrorFieldDTO{
						strings.ToLower(fe.Field()),
						TranslateTag(fe),
					}
				}
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, &httpError)
		}
	default:
		HandleConcreteError(ctx, err)
	}
}
