
# My Cool API
My API usually works as expected.

Table of Contents

1. [retrieves orders for given customer defined by customer ID](#order)

<a name="order"></a>

## order

| Specification | Value |
|-----|-----|
| Resource Path | /order |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api/ |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /orders/by-customer/\{customer_id\} | [GET](#GetResults) | retrieves orders for given customer defined by customer ID |



<a name="GetResults"></a>

#### API: /orders/by-customer/\{customer_id\} (GET)


retrieves orders for given customer defined by customer ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| customer_id | path | int | Customer ID | Yes |
| order_id | query | int | Retrieve order with given ID only |  |
| order_nr | query | string | Retrieve order with given number only |  |
| created_from | query | string | Date-time string, MySQL format. If specified, API will retrieve orders that were created starting from created_from |  |
| created_to | query | string | Date-time string, MySQL format. If specified, API will retrieve orders that were created before created_to |  |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Results](#github.com.angelospanag.council_tax_scraper.domain.Results) |  |




### Models

<a name="github.com.angelospanag.council_tax_scraper.domain.Result"></a>

#### Result

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| address | string |  |
| council_tax_band | string |  |
| reference_number | string |  |

<a name="github.com.angelospanag.council_tax_scraper.domain.Results"></a>

#### Results

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| results | array |  |


