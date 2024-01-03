package apikey

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func AuthApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストからAPIキーを取得
		apiKey := c.Request().Header.Get("X-API-KEY")

		// 環境変数から期待するAPIキーを取得
		expectedApiKey := os.Getenv("SAMPLE_API_KEY")

		// APIキーが期待するものと一致するかチェック
		if apiKey != expectedApiKey {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or missing API key"})
		}

		// APIキーが正しい場合、次のハンドラに処理を渡す
		return next(c)
	}
}
