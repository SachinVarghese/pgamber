# VAE Outlier Detection

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < create.sql
```

```sql
SELECT trainVAEOutlierDetector('individuals', ARRAY[0], 10);
```

```sql
SELECT *, isVAEOutlier(individuals) as outlier FROM individuals LIMIT 20;

SELECT *, isVAEOutlier(individuals) as outlier FROM individuals WHERE age > 55;

SELECT *, isVAEOutlier(individuals) as outlier FROM individuals WHERE isVAEOutlier(individuals) is TRUE AND age > 55;
```

```sql
SELECT dropVAEOutlierDetector('individuals');
```

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < drop.sql
```
