package models

type StudentModel struct {
	University     string `json:"university"`
	StudyYear      string `json:"study_year"`
	Speciality     string `json:"speciality"`
	SpecialityCode string `json:"speciality_code"`
}
