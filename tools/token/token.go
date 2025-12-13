package token

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"sync"
	"time"
)

// SessionData, bir oturumda saklanacak verileri temsil eder.
type SessionData struct {
	UserID    int
	ExpiresAt time.Time
}

// Manager, tüm aktif oturumları bellekte yönetir.
// Üretim ortamında bu verileri Redis veya bir veritabanında saklamak daha ölçeklenebilir bir çözümdür.
type Manager struct {
	mu       sync.RWMutex
	sessions map[string]SessionData
}

// NewManager yeni bir oturum yöneticisi oluşturur.
func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]SessionData),
	}
}

// CreateSession yeni bir oturum oluşturur, saklar ve session token'ını döndürür.
func (m *Manager) CreateSession(userID int, duration time.Duration) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[token] = SessionData{
		UserID:    userID,
		ExpiresAt: time.Now().Add(duration),
	}

	return token, nil
}

// GetSessionData verilen token'a karşılık gelen oturum verisini döndürür.
func (m *Manager) GetSessionData(token string) (SessionData, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, exists := m.sessions[token]
	// Oturum yoksa veya süresi dolmuşsa geçersiz kabul et.
	if !exists || time.Now().After(data.ExpiresAt) {
		if exists {
			// Süresi dolmuş oturumu temizle.
			delete(m.sessions, token)
		}
		return SessionData{}, false
	}
	return data, true
}

// RevokeSession verilen token'ı silerek oturumu sonlandırır.
func (m *Manager) RevokeSession(token string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.sessions, token)
}

// generateToken kriptografik olarak güvenli, rastgele bir token oluşturur.
func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
