package models

type Course struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	MaxScore      *int    `json:"maxScore"`
	MinScore      *int    `json:"minScore"`
	Description   string  `json:"description"`
	PreviewFile   File    `json:"previewFile"`
	EstimatedTime *string `json:"estimatedTime"`
	CreatedByUser User    `json:"createdByUser"`
}

type CourseResponse struct {
	Course Course `json:"course"`
}

type CoursesResponse struct {
	Courses []Course `json:"courses"`
}

type ListCoursesQuery struct {
	UserID string
}

func (q ListCoursesQuery) ToQueryParams() map[string]string {
	return map[string]string{"userId": q.UserID}
}

type CreateCourseRequest struct {
	Title           string  `json:"title"`
	MaxScore        *int    `json:"maxScore,omitempty"`
	MinScore        *int    `json:"minScore,omitempty"`
	Description     string  `json:"description"`
	EstimatedTime   *string `json:"estimatedTime,omitempty"`
	PreviewFileID   string  `json:"previewFileId"`
	CreatedByUserID string  `json:"createdByUserId"`
}

type UpdateCourseRequest struct {
	Title         *string `json:"title,omitempty"`
	MaxScore      *int    `json:"maxScore,omitempty"`
	MinScore      *int    `json:"minScore,omitempty"`
	Description   *string `json:"description,omitempty"`
	EstimatedTime *string `json:"estimatedTime,omitempty"`
}
