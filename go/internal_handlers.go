package main

import (
	"database/sql"
	"errors"
	"net/http"
)

// このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 有効でかつ使用されてない椅子を取得し
	// 未割り当てのライド情報(r.chair_id IS NULL)の中で、最も距離的に近い椅子を取得する
	type Match struct {
		RideID  string `db:"ride_id"`
		ChairID string `db:"chair_id"`
	}
	var match Match
	query := `
		SELECT r.id as ride_id, c.id as chair_id
		FROM rides r
		CROSS JOIN (
			SELECT cl1.* 
			FROM chair_locations cl1
			INNER JOIN (
				SELECT chair_id, MAX(created_at) as latest_at
				FROM chair_locations
				GROUP BY chair_id
			) cl2 ON cl1.chair_id = cl2.chair_id 
			AND cl1.created_at = cl2.latest_at
		) cl
		JOIN chairs c ON c.id = cl.chair_id
		WHERE 
			r.chair_id IS NULL
			AND c.is_active = TRUE
			AND c.id NOT IN (
				SELECT chair_id 
				FROM rides r2 
				WHERE r2.chair_id IS NOT NULL
					AND r2.id IN (
					SELECT ride_id 
					FROM (
						SELECT ride_id, COUNT(chair_sent_at) = 6 AS completed 
						FROM ride_statuses 
						GROUP BY ride_id
					) status_count 
				WHERE completed = FALSE
				)
			)
		ORDER BY ST_Distance_Sphere(r.pickup_location, cl.location) ASC
		LIMIT 1
	`

	if err := db.GetContext(ctx, &match, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := db.ExecContext(ctx, "UPDATE rides SET chair_id = ? WHERE id = ?", match.ChairID, match.RideID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
