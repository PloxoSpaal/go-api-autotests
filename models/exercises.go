package models

type Exercise struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	CourseId      int     `json:"courseId"`
	MaxScore      *int    `json:"maxScore"`
	MinScore      *int    `json:"minScore"`
	OrderIndex    int     `json:"orderIndex"`
	Description   string  `json:"description"`
	EstimatedTime *string `json:"estimatedTime"`
}

type ListExercisesQuery struct {
	CourseId string
}

func (q ListExercisesQuery) ToQueryParams() map[string]string {
	return map[string]string{"courseId": q.CourseId}
}

type ExercisesResponse struct {
	Exercises []Exercise `json:"exercises"`
}

type CreateExerciseRequest struct {
	Title         string  `json:"title"`
	CourseId      string  `json:"courseId"`
	MaxScore      *int    `json:"maxScore,omitempty"`
	MinScore      *int    `json:"minScore,omitempty"`
	OrderIndex    *int    `json:"orderIndex,omitempty"`
	Description   string  `json:"description"`
	EstimatedTime *string `json:"estimatedTime,omitempty"`
}

type ExerciseResponse struct {
	Exercise Exercise `json:"exercise"`
}

type UpdateExerciseRequest struct {
	Title         *string `json:"title,omitempty"`
	MaxScore      *int    `json:"maxScore,omitempty"`
	MinScore      *int    `json:"minScore,omitempty"`
	OrderIndex    *int    `json:"orderIndex,omitempty"`
	Description   *string `json:"description,omitempty"`
	EstimatedTime *string `json:"estimatedTime,omitempty"`
}
