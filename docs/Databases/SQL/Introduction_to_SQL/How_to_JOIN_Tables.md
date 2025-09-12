# How to Join tables

One-to-Many Relationship is the most common one.

This is an INNER JOIN (only results with exact match in both tables):
If there are rows with NULL they will not be returned.

```sql
SELECT FirstName, Employee.DepartmentId,
       Department.Name
FROM Employee
JOIN Department
  ON Employee.DepartmentId = Department.Id;
```

OUTER JOIN returns rows in table A that don't match table B, and the other results as in INNER JOIN:

```sql
...
FROM Employee OUTER JOIN Department
```

`LEFT OUTER JOIN` makes the table on the left have priority.
The word OUTER can be left out.

`FULL OUTER JOIN` is the most inclusive.

## NATURAL JOIN


```SQL
FROM Order NATURAL JOIN Customer;
```

## Table Alias

```sql
FROM orders o INNER JOIN customers c
```



