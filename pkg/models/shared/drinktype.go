// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DrinkType - The type of drink.
type DrinkType string

const (
	DrinkTypeCocktail     DrinkType = "cocktail"
	DrinkTypeNonAlcoholic DrinkType = "non-alcoholic"
	DrinkTypeBeer         DrinkType = "beer"
	DrinkTypeWine         DrinkType = "wine"
	DrinkTypeSpirit       DrinkType = "spirit"
	DrinkTypeOther        DrinkType = "other"
)

func (e DrinkType) ToPointer() *DrinkType {
	return &e
}
func (e *DrinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cocktail":
		fallthrough
	case "non-alcoholic":
		fallthrough
	case "beer":
		fallthrough
	case "wine":
		fallthrough
	case "spirit":
		fallthrough
	case "other":
		*e = DrinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DrinkType: %v", v)
	}
}
