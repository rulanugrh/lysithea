package middleware

import "github.com/rulanugrh/lysithea/internal/entity/web"

func CheckPermission(claims *jwtclaim) error {
	if claims.RoleID != 1 {
		return web.Forbidden("Sorry you not admin")
	} else if claims.RoleID != 2 {
		return web.Forbidden("Sorry you not owner")
	} else {
		return nil
	}
}
