# Ordering, Grouping, Aggregate Functions

## ORDER BY

The sequence of columns matter.

```sql
ORDER BY Name DESC;
``` 

## Aggregate functions

### COUNT

```sql
SELECT COUNT(*)
FROM Employees
WHERE Salary < 5000;
```

### SUM

```sql
SELECT SUM(minutes_watched)
FROM user_session;
```

### MAX

```sql
SELECT MAX(ListPrice)
FROM Product;
```

### MIN

```sql
SELECT MIN()
```

### AVG

```sql
SELECT AVG()
```

## Summarizing Data with Group By

### GROUP BY

When you use GROUP BY you also use an aggregate function.
When you use GROUP BY you also use an aggregate function.

```sql
SELECT COUNT(*), Category,
FROM Product
GROUP BY Category;
```

|Count|Category|
|----:|:-------|
|  640|Books   |
|  152|Tools   |

GROUP BY is intended to summarize.

## Using GROUP BY and HAVING

In SQL below, WHERE happens before GROUP BY, and after COUNT.

```sql
SELECT (*), Category
FROM Product
WHERE Price > 100
GROUP BY Category;
```

The HAVING keyword can be used together with GROUP BY.

```sql
SELECT COUNT(*), Category
FROM Product
WHERE Status <> 'Discontinues'
GROUP BY Category
HAVING COUNT(*) > 100;
```

## DISTINCT



