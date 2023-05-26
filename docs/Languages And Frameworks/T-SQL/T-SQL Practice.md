# Practicing T-SQL

## Distinct

Select values that have an even ID and exclude duplicates:

```sql
SELECT DISTINCT name FROM city WHERE city.id % 2 = 0;
```

## Different

Find the number of total entries in the CITY table and the number of unique entries:

```sql
SELECT (SELECT count(city) FROM station) - (SELECT count(DISTINCT city) FROM station);
```

## Smallest and Largest

Select the two cities with the smallest and largest names:

```sql
SELECT TOP 1 city, LEN(city) lengthOfCity FROM station ORDER BY lengthOfCity, city;
SELECT TOP 1 city, LEN(city) lengthOfCity FROM station ORDER BY lengthOfCity DESC, city;
```

