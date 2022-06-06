## API docs

Host: http://localhost:9876

### Add to cart

Add item to the cart with specific transaction id. If cart with the transaction id is not exist, it will create a new cart. Promotion will be applied if the products in cart meet promo requirements.  

URL: http://localhost:9876/cart/add-to-cart  
Method: POST  
Content-Type: application/json  
Request body:  
| field          | data type   | description |
| -------------- | ----------- | ----------- |
| transaction_id | string      | transaction id of the cart. Must be unique for each cart |
| product_sku    | string      | sku of the product          |
| qty            | int         | qty of the product          |

Request body sample:
```
{
    "transaction_id": "abc123",
    "product_sku": "A304SD",
    "qty": 1
}
```

Response body:
| field          | data type   | description |
| -------------- | ----------- | ----------- |
| status | string      | `success` or `failed` |
| message | string      | message that explain the status |
| cart    |       |           |
| - transaction_id | string         | transaction id of cart |
| - transaction_time | string         | transaction time of cart |
| items    |       |           |
| - sku | string         | sku of the item          |
| - name | string         | name of the item          |
| - qty | int         | quantity of the item          |
| - price | float         | price for each item          |
| - total_price | float         | total price for item of that sku |
| promos    |       |           |
| - id | int         | id of the promo          |
| - name | string         | name of the promo          |
| - total_amount | float         | total amount of the promo          |
| subtotal | float      | subtotal of the cart |
| discount | float      | total discount of the cart |
| total_amount | float      | total amount of the cart |

Response body sample:
```
{
    "status": "success",
    "message": "success",
    "cart": {
        "transaction_id": "a1",
        "transaction_time": "2022-06-04T13:19:06Z",
        "items": [
            {
                "sku": "A304SD",
                "name": "Alexa Speaker",
                "qty": 3,
                "price": 109.5,
                "total_price": 328.5
            }
        ],
        "promos": [
            {
                "id": 3,
                "name": "buy3_alexa_disc_10",
                "total_amount": 32.85
            }
        ],
        "subtotal": 328.5,
        "discount": 32.85,
        "total_amount": 295.65
    }
}
```

### Checkout

Checkout and pay the cart. Require transaction id of the cart and payment method.  

URL: http://localhost:9876/cart/checkout  
Method: POST  
Content-Type: application/json  
Request body:  
| field          | data type   | description |
| -------------- | ----------- | ----------- |
| transaction_id | string      | transaction id of the cart |
| payment_method | string      | method to pay the transaction |

Request body sample:
```
{
    "transaction_id": "abc123",
    "payment_method": "cash"
}
```

Response body:
| field          | data type   | description |
| -------------- | ----------- | ----------- |
| status | string      | `success` or `failed` |
| message | string      | message that explain the status |

Response body sample:
```
{
    "status": "success",
    "message": "success"
}
```