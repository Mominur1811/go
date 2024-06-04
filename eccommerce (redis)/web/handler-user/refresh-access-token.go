package handleruser

import (
	"net/http"
)

func RefreshAccessToken(w http.ResponseWriter, r *http.Request) {

	//token := r.Header.Get("Authorization")

	/*
		claims, err := GetPayloadData(token)
		if err != nil {
			slog.Error("Failed to claims data from jwt_token", logger.Extra(map[string]any{
				"error":   err.Error(),
				"payload": claims,
			}))
			utils.SendError(w, http.StatusBadRequest, err.Error())
			return
		}

		newAccessToken, err := auth.GetAccessToken(claims, 2)
		if err != nil {
			slog.Error("Failed to generate jwt_token", logger.Extra(map[string]any{
				"error":   err.Error(),
				"payload": newAccessToken,
			}))
			utils.SendError(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.SendData(w, http.StatusOK, newAccessToken)
	*/

}
