# VAE Outlier Detection

### Quickstart

1. Spin up test setup using `docker-compose` :hammer_and_wrench:

```bash
make -C ../../../ test_setup_up
```

2. Ingest test data with golang script :card_index_dividers:

```bash
make -C ../../../ ingest_test_data
```

3. Setup VAE outlier detector procs :magic_wand:

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < create.sql
```

4. Train a VAE outlier detector for the data in `individuals` table by providing exclude idex list for non numeric rows & ids. Further, provide anoher parameter for outlier percentage. :crystal_ball:

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber
```

```sql
SELECT trainVAEOutlierDetector('individuals', ARRAY[0,5], 10);
```

5. Run outlier detection queries using sql commands :smile:

```sql
SELECT *, isVAEOutlier(individuals) as outlier FROM individuals LIMIT 10;
```

| id  | age | workclass | education | marital_status | occupation | relationship | race | sex | capital_gain | capital_loss | hours_per_week | country | outlier |
| --- | --- | --------- | --------- | -------------- | ---------- | ------------ | ---- | --- | ------------ | ------------ | -------------- | ------- | ------- |
| 1   | 39  | 7         | 1         | 1              | 1          | 1            | 4    | 1   | 2174         | 0            | 40             | 9       | f       |
| 2   | 50  | 6         | 1         | 0              | 8          | 0            | 4    | 1   | 0            | 0            | 13             | 9       | f       |
| 3   | 38  | 4         | 4         | 2              | 2          | 1            | 4    | 1   | 0            | 0            | 40             | 9       | f       |
| 4   | 53  | 4         | 3         | 0              | 2          | 0            | 2    | 1   | 0            | 0            | 40             | 9       | f       |
| 5   | 28  | 4         | 1         | 0              | 5          | 5            | 2    | 0   | 0            | 0            | 40             | 6       | f       |
| 6   | 37  | 4         | 5         | 0              | 8          | 5            | 4    | 0   | 0            | 0            | 40             | 9       | f       |
| 7   | 49  | 4         | 3         | 2              | 7          | 1            | 2    | 0   | 0            | 0            | 16             | 5       | f       |
| 8   | 52  | 6         | 4         | 0              | 8          | 0            | 4    | 1   | 0            | 0            | 45             | 9       | f       |
| 9   | 31  | 4         | 5         | 1              | 5          | 1            | 4    | 0   | 14084        | 0            | 50             | 9       | t       |
| 10  | 42  | 4         | 1         | 0              | 8          | 0            | 4    | 1   | 5178         | 0            | 40             | 9       | f       |

(10 rows)

6. Outlier detection queries works with filters too :wink:

```sql
SELECT *, isVAEOutlier(individuals) as outlier FROM individuals WHERE age > 55 LIMIT 10;
```

| id  | age | workclass | education | marital_status | occupation | relationship | race | sex | capital_gain | capital_loss | hours_per_week | country | outlier |
| --- | --- | --------- | --------- | -------------- | ---------- | ------------ | ---- | --- | ------------ | ------------ | -------------- | ------- | ------- |
| 25  | 59  | 4         | 4         | 2              | 4          | 4            | 4    | 0   | 0            | 0            | 40             | 9       | f       |
| 26  | 56  | 2         | 1         | 0              | 4          | 0            | 4    | 1   | 0            | 0            | 40             | 9       | f       |
| 46  | 57  | 1         | 1         | 0              | 5          | 0            | 2    | 1   | 0            | 0            | 40             | 9       | f       |
| 75  | 79  | 4         | 4         | 0              | 5          | 2            | 4    | 1   | 0            | 0            | 20             | 9       | f       |
| 78  | 67  | 0         | 3         | 0              | 0          | 0            | 4    | 1   | 0            | 0            | 2              | 9       | t       |
| 84  | 59  | 4         | 4         | 0              | 6          | 0            | 4    | 1   | 0            | 0            | 48             | 9       | f       |
| 91  | 57  | 4         | 0         | 0              | 5          | 0            | 4    | 1   | 0            | 0            | 40             | 9       | f       |

(7 rows)

```sql
SELECT *, isVAEOutlier(individuals) as outlier FROM individuals WHERE isVAEOutlier(individuals) is TRUE AND age > 55 LIMIT 10;
```

| id  | age | workclass | education | marital_status | occupation | relationship | race | sex | capital_gain | capital_loss | hours_per_week | country | outlier |
| --- | --- | --------- | --------- | -------------- | ---------- | ------------ | ---- | --- | ------------ | ------------ | -------------- | ------- | ------- |
| 78  | 67  | 0         | 3         | 0              | 0          | 0            | 4    | 1   | 0            | 0            | 2              | 9       | t       |

(1 row)

7. Remove VAE outlier detector created for specific table, :file_cabinet:

```sql
SELECT dropVAEOutlierDetector('individuals');
```

8. Drop VAE outlier detector procs :axe:

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < drop.sql
```

9. Spin down test setup :gear:

```bash
make -C ../../../ test_setup_down
```

10. Delete persisted data by removing local `docker` volume :broom:

```bash
make -C ../../../ purge_test_volume
```
