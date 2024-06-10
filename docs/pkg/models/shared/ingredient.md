# Ingredient


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Name`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | The name of the ingredient.                                                        | Sugar Syrup                                                                        |
| `ProductCode`                                                                      | **string*                                                                          | :heavy_minus_sign:                                                                 | The product code of the ingredient, only available when authenticated.             | AC-A2DF3                                                                           |
| `Stock`                                                                            | **int64*                                                                           | :heavy_minus_sign:                                                                 | The number of units of the ingredient in stock, only available when authenticated. | 10                                                                                 |
| `Type`                                                                             | [shared.IngredientType](../../../pkg/models/shared/ingredienttype.md)              | :heavy_check_mark:                                                                 | The type of ingredient.                                                            |                                                                                    |