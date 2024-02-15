package modules

type Example struct {
	Inputs []string `json:"input"`
	Output string   `json:"output"`
}

type Exercise struct {
	ID                   string `json:"id" bson:"_id,omitempty"`
	Name                 string `json:"name" bson:"name"`
	Description          string `json:"description" bson:"description"`
	Examples             []Example
	BasisOperationGO     string `json:"basisOperationGO" bson:"basisOperationGO"`
	BasisOperationNodeJS string `json:"basisOperationNodeJS" bson:"basisOperationNodeJS"`
}
