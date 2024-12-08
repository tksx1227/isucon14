-- 走行距離管理用のテーブルを作成
DROP TABLE IF EXISTS chair_mileage;
CREATE TABLE chair_mileage
(
  chair_id   VARCHAR(26) NULL COMMENT '割り当てられた椅子ID',
  distance   INTEGER NOT NULL COMMENT '総走行距離',
  updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '最終走行距離更新日時',
  UNIQUE (chair_id)
)
  COMMENT '椅子ごとの総走行距離テーブル';

-- chair_locationsのchair_idをユニークにする
ALTER TABLE chair_locations ADD CONSTRAINT unique_chair_id UNIQUE (chair_id);
