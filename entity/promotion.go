package entity

type Promotion struct {
	ID                  int64
	Name                string
	Description         string
	Status              string
	ProductRequirements map[string]PromoProductRequirement
	PromoOutcome        map[string]PromoOutcome
}

type PromoProductRequirement struct {
	ID         int64
	PromoID    int64
	ProductSku string
	MinQty     int64
}

type PromoOutcome struct {
	ID            int64
	PromoID       int64
	ProductSku    string
	PromotionType string
	Amount        float64
	Qty           int64
}
