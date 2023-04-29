# Reading notes:
- SQL query size table, index:
  ```
  select pg_relation_size(oid), relname from pg_class order by pg_relation_size(oid) desc;
  ```