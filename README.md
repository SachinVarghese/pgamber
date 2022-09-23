# pgAmber

## Data Observability for PostgreSQL using Alibi-Detect

Monitoring the data quality on databases can be crucial for many real-world applications. The number of downstream applications being built using data extracted from SQL-based data warehouses and ML feature stores is unprecedented. A good set of observability tools at the data extraction layer could have a significant positive impact on optimizing such data pipelines and processes.

This project hosts a set of stored functions for data observability on the Postgres database. These procedures are built using [alibi-detect package](https://docs.seldon.io/projects/alibi-detect/en/stable/index.html) with [python procedural language](https://www.postgresql.org/docs/current/plpython.html) to observe any data abnormalities. These detect functions when paired with clever SQL queries can help users foresee any downstream data quality related issues. The project name **pgAmber** is pronounced as 'paigamber'-> a synonym for clairvoyant or prophet. The current version is limited to detecting outlier records in data tables.

### Outlier Detection

- VAE Outlier Detection - [Quickstart](./procs/outlier-detection/vae/README.md)
