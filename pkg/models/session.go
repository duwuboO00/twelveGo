package models

import (
	"database/sql"
	"time"

	"twelveGo/internal/database"

	"github.com/google/uuid"
)

// Session 代表用戶登入狀態
type Session struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	SessionID    string
	LastActivity time.Time
	CreatedAt    time.Time
}

// CreateSession 新增 session 記錄
func CreateSession(userID uuid.UUID, sessionID, ip, userAgent string) error {
	_, err := database.DB.Exec(
		"INSERT INTO user_sessions (user_id, session_id, ip_address, user_agent, status, last_activity) VALUES ($1,$2,$3,$4,true,now())",
		userID, sessionID, ip, userAgent,
	)
	return err
}

// GetSession 依 session_id 取得 Session
func GetSession(sessionID string) (*Session, error) {
	row := database.DB.QueryRow(`SELECT id, user_id, session_id, last_activity, created_at FROM user_sessions WHERE session_id=$1 AND status=true`, sessionID)
	s := &Session{}
	if err := row.Scan(&s.ID, &s.UserID, &s.SessionID, &s.LastActivity, &s.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return s, nil
}

// UpdateSessionActivity 更新最後活動時間
func UpdateSessionActivity(sessionID string) error {
	_, err := database.DB.Exec("UPDATE user_sessions SET last_activity=now() WHERE session_id=$1", sessionID)
	return err
}

// DeleteSession 刪除指定 session
func DeleteSession(sessionID string) error {
	_, err := database.DB.Exec("DELETE FROM user_sessions WHERE session_id=$1", sessionID)
	return err
}

// CleanupExpiredSessions 移除超過30天未活動的session
func CleanupExpiredSessions() error {
	_, err := database.DB.Exec("DELETE FROM user_sessions WHERE last_activity < now() - interval '30 days'")
	return err
}
