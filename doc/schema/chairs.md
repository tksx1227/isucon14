# chairs

## Description

椅子情報テーブル

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `chairs` (
  `id` varchar(26) NOT NULL COMMENT '椅子ID',
  `owner_id` varchar(26) NOT NULL COMMENT 'オーナーID',
  `name` varchar(30) NOT NULL COMMENT '椅子の名前',
  `model` text NOT NULL COMMENT '椅子のモデル',
  `is_active` tinyint(1) NOT NULL COMMENT '配椅子受付中かどうか',
  `access_token` varchar(255) NOT NULL COMMENT 'アクセストークン',
  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '登録日時',
  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='椅子情報テーブル'
```

</details>

## Columns

| Name | Type | Default | Nullable | Extra Definition | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | ---------------- | -------- | ------- | ------- |
| id | varchar(26) |  | false |  |  |  | 椅子ID |
| owner_id | varchar(26) |  | false |  |  |  | オーナーID |
| name | varchar(30) |  | false |  |  |  | 椅子の名前 |
| model | text |  | false |  |  |  | 椅子のモデル |
| is_active | tinyint(1) |  | false |  |  |  | 配椅子受付中かどうか |
| access_token | varchar(255) |  | false |  |  |  | アクセストークン |
| created_at | datetime(6) | CURRENT_TIMESTAMP(6) | false | DEFAULT_GENERATED |  |  | 登録日時 |
| updated_at | datetime(6) | CURRENT_TIMESTAMP(6) | false | DEFAULT_GENERATED on update CURRENT_TIMESTAMP(6) |  |  | 更新日時 |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| PRIMARY | PRIMARY KEY (id) USING BTREE |

## Relations

![er](chairs.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
