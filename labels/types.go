package labels

type Label string

type LabelGrade string

const (
	LabelGradeStrict LabelGrade = "strict"
	LabelGradeGuess  LabelGrade = "guess"
)

type AddressLabels struct {
	Address string     `json:"address" bson:"address"`
	Labels  []Label    `json:"labels" bson:"labels"`
	Grade   LabelGrade `json:"grade" bson:"grade"`
	Remark  string     `json:"remark" bson:"remark"`
}
