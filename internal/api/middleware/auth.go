package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/config"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

type JWTKeySet struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

type AuthMidd struct {
	Env *config.Env
	keySet *JWTKeySet
}

func NewAuthMidd(env *config.Env) *AuthMidd {
	return &AuthMidd{
		Env:   env,
		keySet: &JWTKeySet{},
	}
}

// FetchPublicKeys faz a requisição para o endpoint de chaves públicas do Keycloak.
func (am *AuthMidd) FetchPublicKeys() error {
	jwksURI := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/certs", am.Env.KeycloakURL, am.Env.KeycloakRealm)
	resp, err := http.Get(jwksURI)
	if err != nil {
		return fmt.Errorf("falha na requisição para o JWKS URI: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status de resposta inesperado do JWKS URI: %s", resp.Status)
	}

	am.keySet = &JWTKeySet{}
	if err := json.NewDecoder(resp.Body).Decode(am.keySet); err != nil {
		return fmt.Errorf("falha ao decodificar a resposta JWKS: %w", err)
	}

	fmt.Println("Chaves públicas do Keycloak obtidas com sucesso!")
	return nil
}

func (am *AuthMidd) ValidateApiKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.Request.Header.Get("X-API-Key")
		if apiKey != am.Env.ApiKey {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid API key"})
			return
		}
		ctx.Next()
	}
}

// getJWKForKeyID busca a chave correta no JWKS usando o ID da chave do JWT.
func (am *AuthMidd) GetJWKForKeyID(keyID string) (jose.JSONWebKey, error) {
	for _, key := range am.keySet.Keys {
		if key.KeyID == keyID {
			return key, nil
		}
	}
	return jose.JSONWebKey{}, fmt.Errorf("chave com ID '%s' não encontrada", keyID)
}

// AuthMiddleware é o middleware que verifica o token JWT.
func (am *AuthMidd) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cabeçalho de Autorização não fornecido"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		
		// 1. Analisa o token para obter o cabeçalho e o ID da chave.
		parsedToken, err := jwt.ParseSigned(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token JWT inválido"})
			return
		}
		// -
		// --- CÓDIGO CORRIGIDO AQUI ---
		// Extrai o 'kid' diretamente do campo KeyID do cabeçalho
		keyID := parsedToken.Headers[0].KeyID
		if keyID == "" {
			fmt.Printf("Erro: a chave 'kid' não foi encontrada no cabeçalho do token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "ID da chave não encontrado no cabeçalho do token"})
			return
		}


		// 2. Busca a chave pública correspondente ao ID.
		jwk, err := am.GetJWKForKeyID(keyID)
		if err != nil {
			fmt.Printf("Erro: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Chave de validação não encontrada"})
			return
		}

		// 3. Valida o token usando a chave pública.
		claims := jwt.Claims{}
		err = parsedToken.Claims(jwk.Key, &claims)
		if err != nil {
			fmt.Printf("Falha na validação do token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de acesso inválido"})
			return
		}

		// Valida as claims do token (exp, iss, etc.).
		now := time.Now()
		err = claims.Validate(jwt.Expected{
			Issuer: fmt.Sprintf("%s/realms/%s", am.Env.KeycloakURL, am.Env.KeycloakRealm),
			Time:   now,
		})
		if err != nil {
			fmt.Printf("Falha na validação das claims do token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de acesso expirado ou inválido"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}