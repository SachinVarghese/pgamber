# Utils

## Alibi-Detect Version Check

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < version.sql
```

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber -c "SELECT get_alibi_detect_version();"
```

```sql
SELECT get_alibi_detect_version();
```

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < drop.sql
```
