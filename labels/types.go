package labels

type Label string

type AddressLabels struct {
	Address string  `json:"address" bson:"address"`
	Labels  []Label `json:"labels" bson:"labels"`
}
