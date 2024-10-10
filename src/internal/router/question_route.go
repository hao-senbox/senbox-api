package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sen-global-api/config"
	controller "sen-global-api/internal/controller"
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/usecase"
	"sen-global-api/internal/middleware"
	"sen-global-api/pkg/sheet"
	"time"
)

func setupQuestionRoutes(engine *gin.Engine, conn *gorm.DB, config config.AppConfig, userSpreadsheet *sheet.Spreadsheet, uploaderSpreadsheet *sheet.Spreadsheet) {

	sessionRepository := repository.SessionRepository{
		AuthorizeEncryptKey: config.AuthorizeEncryptKey,

		TokenExpireTimeInHour: time.Duration(config.TokenExpireDurationInHour),
	}
	secureMiddleware := middleware.SecuredMiddleware{SessionRepository: sessionRepository}
	userRepository := repository.UserRepository{DBConn: conn}
	questionRepository := repository.QuestionRepository{DBConn: conn}
	formRepo := &repository.FormRepository{DBConn: conn, DefaultRequestPageSize: config.DefaultRequestPageSize}
	ctx := context.Background()
	spreadSheet, _ := sheet.NewUserSpreadsheet(config, ctx)
	controller := controller.QuestionController{
		DBConn: conn,
		GetUserQuestionsUseCase: usecase.GetUserQuestionsUseCase{
			UserRepository: userRepository,
			DeviceQuestionRepository: repository.DeviceQuestionRepository{
				DBConn: conn,
			},
		},
		GetUserFromTokenUseCase: usecase.GetUserFromTokenUseCase{
			UserRepository:    userRepository,
			SessionRepository: sessionRepository,
		},
		GetQuestionByIdUseCase: usecase.GetQuestionByIdUseCase{
			QuestionRepository: questionRepository,
		},
		GetDeviceIdFromTokenUseCase: usecase.GetDeviceIdFromTokenUseCase{
			SessionRepository: &sessionRepository,
			DeviceRepository:  &repository.DeviceRepository{DBConn: conn, DefaultRequestPageSize: config.DefaultRequestPageSize, DefaultOutputSpreadsheetUrl: config.OutputSpreadsheetUrl},
		},
		GetQuestionByFormUseCase: usecase.GetQuestionsByFormUseCase{
			QuestionRepository:          &questionRepository,
			DeviceFormDatasetRepository: &repository.DeviceFormDatasetRepository{DBConn: conn},
			CodeCountingRepository:      repository.NewCodeCountingRepository(),
			DB:                          conn,
		},
		GetFormByIdUseCase: usecase.GetFormByIdUseCase{
			FormRepository: formRepo,
		},
		GetAllQuestionsUseCase: usecase.GetAllQuestionsUseCase{
			QuestionRepository: &questionRepository,
		},
		CreateFormUseCase: usecase.CreateFormUseCase{
			FormRepository:         formRepo,
			FormQuestionRepository: &repository.FormQuestionRepository{DBConn: conn},
		},
		GetRawQuestionFromSpreadsheetUseCase: usecase.GetRawQuestionFromSpreadsheetUseCase{
			SpreadsheetId:     config.Google.SpreadsheetId,
			SpreadsheetReader: spreadSheet.Reader,
		},
		SyncQuestionsUseCase: usecase.SyncQuestionsUseCase{
			QuestionRepository: &questionRepository,
		},
		GetButtonsQuestionDetailUseCase: usecase.GetButtonsQuestionDetailUseCase{
			QuestionRepository: &questionRepository,
			Reader:             spreadSheet.Reader,
		},
		GetShowPicsQuestionDetailUseCase: usecase.GetShowPicsQuestionDetailUseCase{
			QuestionRepository: &questionRepository,
		},
		FindDeviceFromRequestCase: usecase.FindDeviceFromRequestCase{
			DeviceRepository:  &repository.DeviceRepository{DBConn: conn, DefaultRequestPageSize: config.DefaultRequestPageSize, DefaultOutputSpreadsheetUrl: config.OutputSpreadsheetUrl},
			SessionRepository: &sessionRepository,
		},
	}

	form := engine.Group("v1/form", secureMiddleware.Secured())
	{
		form.GET("", controller.GetFormQRCode)
	}

	question := engine.Group("v1/question", secureMiddleware.Secured())
	{
		question.GET("/buttons", controller.GetButtonsQuestion)

		question.GET("/show-pics", controller.GetShowPicsQuestion)
	}
}
