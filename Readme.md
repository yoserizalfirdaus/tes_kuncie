This is a backend service that handle checkout system with promo. It is written in Go, using PostgreSQL as the database.  
To run the application and the database we require Docker. I have created a Makefile script to do unit test and then run the application and the database. The database will be initialized using files in `sqlfiles` directory. To do that we can execute this command:
```
make start
```

Once the application start, we can interact with the server through port `9876` of our localhost.  

## List of API
This service currently serve these usecase:
- Add to Cart
Add item to the cart with provided transaction id. Promotion will be evaluated and applied for the eligible products.
- Checkout
Checkout the cart with transaction id using a given payment method.

See [here](docs/apidocs.md) for the complete API docs.

## Code Architecture
The code is written using Clean Code Architecture. This application is separated into 4 main layers. The layers in order from outside to inside are infrastructure, interface adapter, usecase, entity. The dependency go from outward to inside, means that a layer can import inner layer, but must not import outside layer. 

## Database
Currently there 4 tables. The product table store product data. The promotion table store promotion data, with the product requirement for the promo stored in table product_promo_requirement. And the outcome of the promo stored in table promo_outcome.  
So to handle promo: "Each sale of a MacBook Pro comes with a free Raspberry Pi B", the promo data stored in promotion table like this:
| id          | name                       | description                |
| ----------- | -------------------------- | -------------------------- |
| 1           | buy_macbook_free_raspberry | Each sale of a MacBook...  |

In product_promo_requirement:
| promo_id    | product_sku | minimum_qty |
| ----------- | ----------- | ----------- |
| 1           | 43N23P      | 1           |
| 1           | 234234      | 1           |

And in the promo_outcome:
| promo_id    | product_sku | promotion_type | amount      | qty         |
| ----------- | ----------- | -------------- | ----------- | ----------- |
| 1           | 43N23P      | percentage     | 100         | 1           |

## GraphQL
To use graphql to interact with the service, we need to create another service that handle graphql request. The service then will forward the request to the checkout service based on the query or mutation. Then graphql server will return the response back to the client. The schema for the service can be seen [here](docs/schema.graphql).