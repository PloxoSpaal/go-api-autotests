package metadata

const (
	EpicAPI   = "Go API autotests"
	LayerE2E  = "e2e"
	SuiteHTTP = "HTTP API"
	SuiteGRPC = "gRPC API"
)

const (
	FeatureUsers          = "Users"
	FeatureFiles          = "Files"
	FeatureCourses        = "Courses"
	FeatureExercises      = "Exercises"
	FeatureAuthentication = "Authentication"
)

const (
	StoryLogin          = "Login"
	StoryGetEntity      = "Get entity"
	StoryGetEntities    = "Get entities"
	StoryCreateEntity   = "Create entity"
	StoryUpdateEntity   = "Update entity"
	StoryDeleteEntity   = "Delete entity"
	StoryValidateEntity = "Validate entity"
)

const (
	TagAPI  = "API"
	TagHTTP = "HTTP"
	TagGRPC = "GRPC"

	TagSmoke      = "SMOKE"
	TagNegative   = "NEGATIVE"
	TagRegression = "REGRESSION"

	TagUsers          = "USERS"
	TagFiles          = "FILES"
	TagCourses        = "COURSES"
	TagExercises      = "EXERCISES"
	TagAuthentication = "AUTHENTICATION"
)
