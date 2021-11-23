package spreadsheet

import (
	"context"
	"errors"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	ScopeSpreadsheetsWrite = "https://www.googleapis.com/auth/spreadsheets"
	ScopeSpreadsheetsRead  = "https://www.googleapis.com/auth/spreadsheets.readonly"

	InsertDataOptionInsert   = "INSERT_ROWS"
	InsertDataOptionOvewrite = "OVERWRITE"

	ValueInputOptionRaw         = "RAW"
	ValueInputOptionUserEntered = "USER_ENTERED"
	ValueInputOptionUnspecified = "INPUT_VALUE_OPTION_UNSPECIFIED"

	// A1 notation. See https://developers.google.com/sheets/api/guides/concepts
	A1NotationFullFirstRow = "1:1"
)

type Service struct {
	service       *sheets.Service
	spreadsheetId string
	sheetId       string
	context       context.Context
}

func New(keyFile string) (*Service, error) {
	service, err := sheets.NewService(context.Background(), option.WithCredentialsFile(keyFile))
	if err != nil {
		return nil, err
	}

	return &Service{
		service: service,
		context: context.Background(),
	}, nil
}

func (s *Service) copy() *Service {
	res := *s
	return &res
}

func (s *Service) WithSpreadsheet(spreadsheetId string) *Service {
	res := s.copy()
	res.spreadsheetId = spreadsheetId
	return res
}

func (s *Service) WithSheet(sheetId string) *Service {
	res := s.copy()
	res.sheetId = sheetId
	return res
}

func (s *Service) WithContext(context context.Context) *Service {
	res := s.copy()
	res.context = context
	return res
}

func (s *Service) getCurrentSpreadsheet() (*sheets.Spreadsheet, error) {
	return s.service.Spreadsheets.Get(s.spreadsheetId).Do()
}

func (s *Service) Check() error {
	if len(s.spreadsheetId) == 0 {
		return errors.New("no spreadsheet id")
	}

	res, err := s.getCurrentSpreadsheet()
	if err != nil {
		return err
	}

	if res.SpreadsheetId != s.spreadsheetId {
		return errors.New("got incorrect spreadsheet")
	}

	return nil
}

func (s *Service) InsertRow(values ...interface{}) error {
	res, err := s.service.Spreadsheets.Values.
		Append(
			s.spreadsheetId,
			s.sheetId+"!"+A1NotationFullFirstRow,
			&sheets.ValueRange{
				MajorDimension: "ROWS",
				Values:         [][]interface{}{values},
			},
		).
		//ResponseDateTimeRenderOption("FORMATTED_STRING").
		InsertDataOption(InsertDataOptionInsert).
		ValueInputOption(ValueInputOptionRaw).
		Context(s.context).
		Do()

	if err != nil {
		return err
	}

	if res.Updates.UpdatedRows != 1 {
		return errors.New("row wasn't inserted")
	}

	return nil
}

func (s *Service) GetUrl() (string, error) {
	res, err := s.getCurrentSpreadsheet()
	if err != nil {
		return "", err
	}

	return res.SpreadsheetUrl, nil
}
