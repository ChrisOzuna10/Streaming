package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Cache es una instancia global de go-cache
var Cache *cache.Cache

// InitCache inicializa el caché
func InitCache() {
	// Crea un caché con 5 minutos de expiración y limpieza cada 10 minutos
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
