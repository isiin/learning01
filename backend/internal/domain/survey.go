package domain

type Surveyor struct {
	ID         string
	Name       string
	OfficeID   string
	OfficeName string
}
type Surveyors []Surveyor

type SurveyorFilter struct {
	ID       string
	OfficeID string
}

type SurveyUseCase interface {
	GetSurveyors(filter SurveyorFilter) (Surveyors, error)
}

type SurveyRepository interface {
	GetSurveyors(filter SurveyorFilter) (Surveyors, error)
}
