# Adding, Updating, and Deleting Data

## INSERT

```sql
INSERT INTO table (column1, column2)
           VALUES ('test', 45);
```

## UPDATE

```sql
UPDATE employee
SET status = 'Retired'
WHERE id = 1;
```

## DELETE

One tip to avoid accidents with DELETE is to first write a SELECT
and confirm that it returns only the rows you do want to delete, then
turn that into a DELETE statement.

```sql
DELETE FROM table
WHERE id = 1;
```

