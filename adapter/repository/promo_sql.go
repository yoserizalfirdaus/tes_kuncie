package repository

import "github.com/yoserizalfirdaus/tes_kuncie/entity"

type PromoRepository struct {
	sqldb SQLDatabase
}

type PromoRow struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type PromoReqRow struct {
	ID         int64  `db:"id"`
	PromoID    int64  `db:"promo_id"`
	ProductSku string `db:"product_sku"`
	MinimumQty int64  `db:"minimum_qty"`
}

type PromoOutcomeRow struct {
	ID            int64   `db:"id"`
	PromoID       int64   `db:"promo_id"`
	ProductSku    string  `db:"product_sku"`
	PromotionType string  `db:"promotion_type"`
	Amount        float64 `db:"amount"`
	Qty           int64   `db:"qty"`
}

func NewPromoRepo(db SQLDatabase) (PromoRepository, error) {
	return PromoRepository{db}, nil
}

func (r PromoRepository) GetPromoByProductSku(sku string) (entity.Promotion, error) {
	promoID := int64(0)
	err := r.sqldb.Get(&promoID,
		"SELECT promo_id FROM product_promo_requirement WHERE product_sku = $1", sku)
	if err != nil {
		return entity.Promotion{}, err
	}

	promoRow := PromoRow{}
	err = r.sqldb.Get(&promoRow,
		"SELECT id, name, description FROM promotion WHERE id = $1", promoID)
	if err != nil {
		return entity.Promotion{}, err
	}

	promoReq := []PromoReqRow{}
	err = r.sqldb.Select(&promoReq,
		"SELECT id, promo_id, product_sku, minimum_qty FROM product_promo_requirement WHERE promo_id = $1", promoID)
	if err != nil {
		return entity.Promotion{}, err
	}

	promoOutcome := []PromoOutcomeRow{}
	err = r.sqldb.Select(&promoOutcome,
		"SELECT id, promo_id, product_sku, promotion_type, amount, qty FROM promo_outcome WHERE promo_id = $1", promoID)
	if err != nil {
		return entity.Promotion{}, err
	}

	promotion := promoRow.toEntity()
	promotion.ProductRequirements = map[string]entity.PromoProductRequirement{}
	for _, v := range promoReq {
		promotion.ProductRequirements[v.ProductSku] = v.toEntity()
	}

	promotion.PromoOutcome = map[string]entity.PromoOutcome{}
	for _, v := range promoOutcome {
		promotion.PromoOutcome[v.ProductSku] = v.toEntity()
	}

	return promotion, err
}

func (r PromoRow) toEntity() entity.Promotion {
	return entity.Promotion{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func (r PromoReqRow) toEntity() entity.PromoProductRequirement {
	return entity.PromoProductRequirement{
		ID:         r.ID,
		PromoID:    r.PromoID,
		ProductSku: r.ProductSku,
		MinQty:     r.MinimumQty,
	}
}

func (r PromoOutcomeRow) toEntity() entity.PromoOutcome {
	return entity.PromoOutcome{
		ID:            r.ID,
		PromoID:       r.PromoID,
		ProductSku:    r.ProductSku,
		PromotionType: r.PromotionType,
		Amount:        r.Amount,
		Qty:           r.Qty,
	}
}
