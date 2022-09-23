# VAE Outlier Detection

The Variational Auto-Encoder (VAE) outlier detector is trained on a batch of **numerical columns** specified on a reference data table. The inbuilt encoder reduces the record to a specified latent dimension and from there VAE detector tries to reconstruct the input it receives. If the input data cannot be reconstructed well, the reconstruction error is high and the data can be flagged as an outlier. The reconstruction error is measured as the mean squared error (MSE) between the input and the reconstructed instance. Read more about this technique on the [`alibi-detect` documentation](https://docs.seldon.io/projects/alibi-detect/en/stable/od/methods/vae.html#Variational-Auto-Encoder).

## Quickstart :racing_car:

- :hammer_and_wrench: Spin up test setup using `docker-compose`. This step creates a docker container running a postgres instance with some additional dependencies.

```bash
make -C ../../../ test_setup_up
```

- :card_index_dividers: Ingest test dataset with person’s characteristics as `individuals` table using the following script.

```bash
make -C ../../../ ingest_test_data
```

- :magic_wand: Setup VAE outlier detector stored procedures. This is the secret sauce.

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < create.sql
```

- :crystal_ball: Train a VAE outlier detector for the data table by providing the following,

- name of the table to be used as reference data -> `individuals`
- an exclude index list for non numeric rows & ids -> `ARRAY[0 , 5]`
- outlier percentage in the dataset -> `10.5%`
- the latent dimension for the encoder.-> `2`

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber
```

```sql
SELECT createVAEOutlierDetector('individuals', ARRAY[0,5], 10.5, 2);
```

- :detective: This step creates an `isVAEOutlier` procedure to detect outlier records. Now, run outlier detection queries using sql commands :smile:

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

- :bar_chart: Outlier detection queries works with your awesome filters too. :wink:

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

- :file_cabinet: The VAE outlier detectors are created for specific table, when you want to re-train or remove the detector artifacts run,

```sql
SELECT dropVAEOutlierDetector('individuals');
```

- :axe: Drop VAE outlier detector procs to cleanup.

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < drop.sql
```

- :gear: Spin down test setup using `docker-compose`.

```bash
make -C ../../../ test_setup_down
```

- :broom: Delete persisted data by removing local `docker` volume.

```bash
make -C ../../../ purge_test_volume
```
