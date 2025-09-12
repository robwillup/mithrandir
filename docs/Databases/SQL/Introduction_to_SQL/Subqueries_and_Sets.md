# Subqueries and Sets

## Creating a subquery

```sql
SELECT product_name, list_price
FROM product
WHERE list_price > (SELECT AVG(list_price) FROM product);
```

## Using column names with AS

```sql
SELECT nm AS name,
       desc AS 'Description of item'
```

## Combining Data with UNION and UNION ALL

Must be fetching same number of columns and compatible data types.
By default it discards duplicate values.
But if you want them to avoid performance issues you can use `UNION ALL`.

```sql
SELECT email FROM curent_customers
UNION
SELECT email FROM past_customers;
```

## UNION vs. JOIN

Use JOIN when the tables are different entities with a relationship.
Use UNION when the tables have similar data we want to append.


