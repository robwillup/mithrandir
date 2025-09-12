# Writing SQL Queries

## Like

Bear in mind that this can be inefficient.
Depending where the % is, the search can be very slow.
For example, % in the begining will make the DBMS search all the text.

Wild card multi character: %

```sql
SELECT *
FROM Vehicles
WHERE Name LIKE '%Genesis%';
```

Wild card single character: _

```sql
SELECT * 
FROM Vehicles
WHERE Name LIKE 'Genesis G_0';
```

## NULL

In SQL you don't use `= NULL`, you use `IS NULL`.
Or `IS NOT NULL`.

## DISTINCT

```sql
SELECT DISTINCT country
FROM customer;
```

## Here's the secret

> Before writing any SQL query, describe the results you want, in conversational language, in a single sentence.

SQL is a declarative language.


