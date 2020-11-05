# gp-pg-issue-1773


```bash
git clone https://github.com/bearsh/gp-pg-issue-1773.git
cd gp-pg-issue-1773
go build
./gp-pg-issue-1773 DB_URL
```

## Output in my test

```bash
./gp-pg-issue-1773 postgres://postgres:postgres@localhost:5432/gp-pg-issue-1773?sslmode=disable
>>> create table 'table' in db 'gp-pg-issue-1773' on 'DB<Addr="localhost:5432">'
DROP TABLE IF EXISTS public."table" CASCADE;
CREATE TABLE public."table" (
        id serial NOT NULL,
        data bytea,
        CONSTRAINT table_pk PRIMARY KEY (id)
);
>>> insert t1: '&{{} 0 [1 2 3 4]}' into 'table'
INSERT INTO "table" ("id", "data") VALUES (DEFAULT, '{1,2,3,4}') RETURNING "id"
>>> t1 after inserting: '&{{} 1 [1 2 3 4]}'
>>> select the just inserted row
SELECT "t"."id", "t"."data" FROM "table" AS "t" WHERE "t"."id" = 1
>>> error: got '\\', wanted '{'
```

Output from pgcli:

```
gp-pg-issue-1773> SELECT * FROM "table"
+------+----------------------+
| id   | data                 |
|------+----------------------|
| 1    | \x7b312c322c332c347d |
+------+----------------------+
SELECT 1
Time: 0.025s
```
