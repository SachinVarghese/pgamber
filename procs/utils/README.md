# Utility Functions :orange_circle:

## Alibi-Detect Version Check

1. Setup

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < version.sql
```

2. Run Query

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber -c "SELECT get_alibi_detect_version();"
```

```sql
SELECT get_alibi_detect_version();
```

3. Cleanup

```bash
PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d pgamber < drop.sql
```
