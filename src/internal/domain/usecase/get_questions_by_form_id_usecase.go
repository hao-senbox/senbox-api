package usecase

import (
	"encoding/json"
	"net/http"
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/response"
	"sen-global-api/internal/domain/value"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GetQuestionsByFormUseCase struct {
	*repository.QuestionRepository
	*repository.CodeCountingRepository
	*gorm.DB
}

func (receiver *GetQuestionsByFormUseCase) GetQuestionByForm(form entity.SForm) (*response.QuestionListResponse, *response.FailedResponse) {
	questions, err := receiver.GetQuestionsByFormId(form.ID)

	if err != nil {
		return nil, &response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: "Failed to get questions",
		}
	}
	var result = make([]response.QuestionListData, 0)
	var rawQuestions = make([]response.QuestionListData, 0)
	for _, question := range questions {
		var att response.QuestionAttributes
		err = json.Unmarshal([]byte(question.Attributes), &att)
		if err != nil {
			continue
		}

		// Skip Send Notification
		// because they are not for mobile
		if question.QuestionType == value.GetStringValue(value.QuestionSendNotification) {
			continue
		}

		q := response.QuestionListData{
			QuestionId:     question.QuestionId,
			QuestionType:   strings.ToUpper(question.QuestionType),
			Question:       question.Question,
			Attributes:     att,
			Order:          question.Order,
			AnswerRequired: question.AnswerRequired,
			Enabled:        question.EnableOnMobile == value.QuestionForMobile_Enabled,
		}

		rawQuestions = append(rawQuestions, q)
	}

	for _, rawQuestion := range rawQuestions {
		qType, err := value.GetQuestionType(rawQuestion.QuestionType)
		if err != nil {
			continue
		}
		if !value.IsGeneralQuestionType(qType) {
			return nil, &response.FailedResponse{
				Code:  http.StatusBadRequest,
				Error: "Failed to get questions",
			}
		}

		// Check code counting & code generation
		switch qType {
		case value.QuestionCodeCounting:
			q, err := receiver.BuildCodeCountingQuestion(rawQuestion)
			if err != nil {
				return nil, &response.FailedResponse{
					Code:  555,
					Error: "Could not parsed user form data fo question: " + rawQuestion.QuestionId + " err: " + err.Error(),
				}
			}
			result = append(result, q)
		case value.QuestionRandomizer:
			q, err := receiver.BuildRandomizerQuestion(rawQuestion)
			if err != nil {
				return nil, &response.FailedResponse{
					Code:  555,
					Error: "Could not parsed user form data fo question: " + rawQuestion.QuestionId + " err: " + err.Error(),
				}
			}
			result = append(result, q)
		default:
			result = append(result, rawQuestion)
		}
	}

	return &response.QuestionListResponse{
		Data: response.QuestionListResponseData{
			QuestionListData: result,
			DecryptPassword:  form.Password,
			FormName:         form.Name,
		},
	}, nil
}

func (receiver *GetQuestionsByFormUseCase) GetQuestionsByForm(form entity.SForm) *response.QuestionListResponse {
	questions, err := receiver.GetQuestionsByFormId(form.ID)
	if err != nil {
		return nil
	}
	var result = make([]response.QuestionListData, 0)
	var rawQuestions = make([]response.QuestionListData, 0)
	for _, question := range questions {
		var att response.QuestionAttributes
		err = json.Unmarshal([]byte(question.Attributes), &att)
		if err != nil {
			continue
		}
		q := response.QuestionListData{
			QuestionId:     question.QuestionId,
			QuestionType:   strings.ToUpper(question.QuestionType),
			Question:       question.Question,
			Attributes:     att,
			Order:          question.Order,
			AnswerRequired: question.AnswerRequired,
			Enabled:        question.EnableOnMobile == value.QuestionForMobile_Enabled,
		}

		rawQuestions = append(rawQuestions, q)
	}

	for _, rawQuestion := range rawQuestions {
		qType, err := value.GetQuestionType(rawQuestion.QuestionType)
		if err != nil {
			continue
		}
		if value.IsGeneralQuestionType(qType) {
			result = append(result, rawQuestion)
		}
	}

	return &response.QuestionListResponse{
		Data: response.QuestionListResponseData{
			QuestionListData: result,
			DecryptPassword:  form.Password,
			FormName:         form.Name,
		},
	}
}

func (receiver *GetQuestionsByFormUseCase) BuildCodeCountingQuestion(question response.QuestionListData) (response.QuestionListData, error) {
	var att response.QuestionAttributes
	var attInJSONString string

	newCodeCountingValue, err := receiver.CreateForQuestionWithID(question.QuestionId, receiver.DB)
	if err != nil {
		log.Error(err)
		return response.QuestionListData{}, err
	}
	attInJSONString = `{"value": "` + newCodeCountingValue + `"}`

	err = json.Unmarshal([]byte(attInJSONString), &att)
	if err != nil {
		return response.QuestionListData{}, err
	}

	q := response.QuestionListData{
		QuestionId:     question.QuestionId,
		QuestionType:   strings.ToUpper(question.QuestionType),
		Question:       question.Question,
		Attributes:     att,
		Order:          question.Order,
		AnswerRequired: question.AnswerRequired,
		Enabled:        question.Enabled,
	}

	return q, nil
}

func (receiver *GetQuestionsByFormUseCase) BuildRandomizerQuestion(question response.QuestionListData) (response.QuestionListData, error) {
	var att response.QuestionAttributes
	var attInJSONString string

	code := att.Value + value.GetRandomString(8)
	attInJSONString = `{"value": "` + question.Attributes.Value + code + `"}`

	err := json.Unmarshal([]byte(attInJSONString), &att)
	if err != nil {
		return response.QuestionListData{}, err
	}

	q := response.QuestionListData{
		QuestionId:     question.QuestionId,
		QuestionType:   strings.ToUpper(question.QuestionType),
		Question:       question.Question,
		Attributes:     att,
		Order:          question.Order,
		AnswerRequired: question.AnswerRequired,
		Enabled:        question.Enabled,
	}

	return q, nil
}
