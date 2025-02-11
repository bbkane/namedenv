# vw_var_expanded

## Description

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE VIEW vw_var_expanded AS
SELECT
    var_id,
    env_id,
    (SELECT name FROM env WHERE env_id = var.env_id) AS env_name,
    name,
    value,
    comment,
    create_time,
    update_time
FROM var
```

</details>

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| var_id | INTEGER |  | true |  |  |  |
| env_id | INTEGER |  | true |  |  |  |
| env_name | TEXT |  | true |  |  |  |
| name | TEXT |  | true |  |  |  |
| value | TEXT |  | true |  |  |  |
| comment | TEXT |  | true |  |  |  |
| create_time | TEXT |  | true |  |  |  |
| update_time | TEXT |  | true |  |  |  |

## Referenced Tables

| Name | Columns | Comment | Type |
| ---- | ------- | ------- | ---- |
| [env](env.md) | 5 |  | table |
| [var](var.md) | 7 |  | table |

## Relations

![er](vw_var_expanded.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
