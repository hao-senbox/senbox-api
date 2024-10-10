package usecase

import (
	"errors"
	"sen-global-api/config"
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/request"
	"sen-global-api/internal/domain/value"
	"sen-global-api/pkg/sheet"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MarkToDoAsDoneUseCase struct {
	*repository.ToDoRepository
	*sheet.Reader
	*sheet.Writer
	dbConn *gorm.DB
	cfg    config.AppConfig
}

func (c *MarkToDoAsDoneUseCase) Execute(device entity.SDevice, code string, index int, selectValue string) error {
	todoList, err := c.ToDoRepository.GetToDoListByQRCode(code, c.dbConn)
	if err != nil {
		return err
	}

	values, err := c.Reader.Get(sheet.ReadSpecificRangeParams{
		SpreadsheetId: todoList.SpreadsheetID,
		ReadRange:     todoList.SheetName + `!K12:K1000`,
	})

	if err != nil {
		log.Error(err)
		return err
	}

	var completedTask entity.Task = entity.Task{}
	updatedTasks := make([]entity.Task, 0)
	for _, task := range todoList.Tasks.Data.Tasks {
		if task.Index == index {
			completedTask = task
			task.Selected = selectValue
		}
		updatedTasks = append(updatedTasks, task)
	}
	todoList.Tasks.Data.Tasks = updatedTasks
	c.ToDoRepository.Save(c.dbConn, &todoList)

	log.Info("completedTask: ", completedTask)

	completedRowNo, err := findFirstRow(strconv.Itoa(index), values, 12)
	if err != nil {
		return err
	}

	completedData := make([][]interface{}, 0)
	completedData = append(completedData, []interface{}{selectValue})
	completedData = append(completedData, []interface{}{time.Now().Format("2006-01-02 15:04:05")})
	completedData = append(completedData, []interface{}{selectValue})
	completedData = append(completedData, []interface{}{device.PrimaryUserInfo})
	completedData = append(completedData, []interface{}{device.SecondaryUserInfo})
	completedData = append(completedData, []interface{}{device.TertiaryUserInfo})
	completedData = append(completedData, []interface{}{device.DeviceId})
	_, err = c.Writer.UpdateRange(sheet.WriteRangeParams{
		Range:     todoList.SheetName + "!P" + strconv.Itoa(completedRowNo) + ":W",
		Dimension: "COLUMNS",
		Rows:      completedData,
	}, todoList.SpreadsheetID)

	if err != nil {
		log.Error("Unable to write to sheet ToDo: ", todoList.SpreadsheetID, err)
		return err
	}

	//TODO: Write Todo History
	historyData := make([][]interface{}, 0)
	historyData = append(historyData, []interface{}{time.Now().Format("2006-01-02 15:04:05")})
	historyData = append(historyData, []interface{}{device.DeviceId})
	historyData = append(historyData, []interface{}{device.DeviceName})
	historyData = append(historyData, []interface{}{device.Note})
	historyData = append(historyData, []interface{}{device.PrimaryUserInfo})
	historyData = append(historyData, []interface{}{device.SecondaryUserInfo})
	historyData = append(historyData, []interface{}{device.TertiaryUserInfo})
	historyData = append(historyData, []interface{}{todoList.ID})
	historyData = append(historyData, []interface{}{todoList.SheetName})
	historyData = append(historyData, []interface{}{"https://docs.google.com/spreadsheets/d/" + todoList.SpreadsheetID})
	historyData = append(historyData, []interface{}{completedTask.Name})
	historyData = append(historyData, []interface{}{completedTask.DueDate})
	historyData = append(historyData, []interface{}{completedTask.Value})
	historyData = append(historyData, []interface{}{selectValue})

	_, err = c.Writer.WriteRanges(sheet.WriteRangeParams{
		Range:     todoList.HistorySheetName + "!K11",
		Dimension: "COLUMNS",
		Rows:      historyData,
	}, todoList.HistorySpreadsheetID)

	if err != nil {
		log.Error("Write Todo History Error: ", todoList.HistorySpreadsheetID, err)
	}

	return nil
}

func NewMarkToDoAsDoneUseCase(cfg config.AppConfig, dbConn *gorm.DB, reader *sheet.Reader, writer *sheet.Writer) *MarkToDoAsDoneUseCase {
	return &MarkToDoAsDoneUseCase{
		ToDoRepository: &repository.ToDoRepository{},
		Reader:         reader,
		Writer:         writer,
		dbConn:         dbConn,
		cfg:            cfg,
	}
}

func findFirstRow(id string, values [][]interface{}, startRow int) (int, error) {
	rowNo := 0
	for rowindex, row := range values {
		if len(row) > 0 {
			if row[0].(string) == id {
				return rowindex + startRow, nil
			}
		}
	}
	return rowNo, errors.New("Cannot determine row number for todo index: " + id)
}

func (c *MarkToDoAsDoneUseCase) LogTask(req request.LogTaskRequest, device entity.SDevice) error {
	todo, err := c.ToDoRepository.FindById(req.ToDoID, c.dbConn)
	if err != nil {
		return err
	}

	if todo == nil {
		return errors.New("todo not found")
	}

	if todo.Type != value.ToDoTypeCompose {
		return errors.New("invalid todo type")
	}

	if todo.HistorySpreadsheetID == "" || todo.HistorySheetName == "" {
		return errors.New("todo's history sheet was not set up")
	}

	historyData := make([][]interface{}, 0)
	historyData = append(historyData, []interface{}{time.Now().Format("2006-01-02 15:04:05")})
	historyData = append(historyData, []interface{}{device.DeviceId})
	historyData = append(historyData, []interface{}{device.DeviceName})
	historyData = append(historyData, []interface{}{device.Note})
	historyData = append(historyData, []interface{}{device.PrimaryUserInfo})
	historyData = append(historyData, []interface{}{device.SecondaryUserInfo})
	historyData = append(historyData, []interface{}{device.TertiaryUserInfo})
	historyData = append(historyData, []interface{}{todo.ID})
	historyData = append(historyData, []interface{}{todo.SheetName})
	historyData = append(historyData, []interface{}{"https://docs.google.com/spreadsheets/d/" + todo.SpreadsheetID})
	historyData = append(historyData, []interface{}{req.Name})
	historyData = append(historyData, []interface{}{req.DueDate})
	historyData = append(historyData, []interface{}{req.Value})
	historyData = append(historyData, []interface{}{""})

	switch req.LogType {
	case request.LogTaskType_Create:
		historyData = append(historyData, []interface{}{"created"})
	case request.LogTaskType_Update:
		historyData = append(historyData, []interface{}{"updated"})
	case request.LogTaskType_Deleted:
		historyData = append(historyData, []interface{}{"deleted"})
	}

	historyData = append(historyData, []interface{}{""})

	_, err = c.Writer.WriteRanges(sheet.WriteRangeParams{
		Range:     todo.HistorySheetName + "!K11",
		Dimension: "COLUMNS",
		Rows:      historyData,
	}, todo.HistorySpreadsheetID)

	return err
}
