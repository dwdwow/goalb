package labels

type LabelName string
type LabelID int

type Label struct {
	LabelName LabelName `json:"labelName" bson:"labelName"`
	LabelID   LabelID   `json:"labelId" bson:"labelId"`
}

type AddressLabels struct {
	Address string  `json:"address" bson:"address"`
	Labels  []Label `json:"labels" bson:"labels"`
}
