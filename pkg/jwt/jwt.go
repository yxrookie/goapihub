// Package jwt: cope the jwt certification
package jwt

import (
	"errors"
	
	"goapihub/pkg/app"
	"goapihub/pkg/config"
	"goapihub/pkg/logger"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
    ErrTokenMalformed         error = errors.New("请求令牌格式有误")
    ErrTokenInvalid           error = errors.New("请求令牌无效")
    ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
    ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// JWT object 
type JWT struct {
	// secret key: encrypt JWT, read the config information: app.key
	SignKey []byte

	// refresh Token's maximum expiration time
	MaxRefresh time.Duration
}

// JWTCustomClaims: custom payload
type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	ExpireAtTime int64 `json:"expire_time"`

	 // StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
    // JWT 规定了7个官方字段，提供使用:
    // - iss (issuer)：发布者
    // - sub (subject)：主题
    // - iat (Issued At)：生成签名的时间
    // - exp (expiration time)：签名过期时间
    // - aud (audience)：观众，相当于接受者
    // - nbf (Not Before)：生效时间
    // - jti (JWT ID)：编号
	jwtpkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// Authorization:Bearer xxxx
func(jwt *JWT) getTokenFromHeader(c *gin.Context)(string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// spilt space to get two parts
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// ParseToken: parse Token, middleware call
func (jwt *JWT) ParseToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	//1. parse user's Token
	token, err := jwt.parseTokenString(tokenString)

	//2. parse failure
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	//3. parse claims from token, 和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid

}

// parseTokenString: use jwtpkg.ParseWithClaims parse Token
func (jwt *JWT) parseTokenString(tokenString string)(*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token)(interface{}, error) {
		return jwt.SignKey, nil
	})
}



// RefreshToken: update Token
func (jwt *JWT) RefreshToken(c *gin.Context)(string, error) {
	 // 1. 从 Header 里获取 token
	 tokenString, parseErr := jwt.getTokenFromHeader(c)
	 if parseErr != nil {
		 return "", parseErr
	 }
 
	 // 2. 调用 jwt 库解析用户传参的 Token
	 token, err := jwt.parseTokenString(tokenString)
 
	 // 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	 if err != nil {
		 validationErr, ok := err.(*jwtpkg.ValidationError)
		 // 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		 if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			 return "", err
		 }
	 }

	 // 4.parse JWTCustomClaims data 
	 claims := token.Claims.(*JWTCustomClaims)

	 // 5. 检查是否过了-- 最大允许刷新时间
	 x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	 if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	 }
	 return "", ErrTokenExpiredMaxRefresh
}

// createToken: create Token, 内部使用 外部调用 IssueToken
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// use HS256 algorithm to generate token
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}


// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
    timenow := app.TimenowInTimezone()

    var expireTime int64
    if config.GetBool("app.debug") {
        expireTime = config.GetInt64("jwt.debug_expire_time")
    } else {
        expireTime = config.GetInt64("jwt.expire_time")
    }

    expire := time.Duration(expireTime) * time.Minute
    return timenow.Add(expire).Unix()
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) IssueToken(userID string, userName string) string {
	
    // 1. 构造用户 claims 信息(负荷)
    expireAtTime := jwt.expireAtTime()
    claims := JWTCustomClaims{
        userID,
        userName,
        expireAtTime,
        jwtpkg.StandardClaims{
            NotBefore: app.TimenowInTimezone().Unix(), // 签名生效时间
            IssuedAt:  app.TimenowInTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
            ExpiresAt: expireAtTime,                   // 签名过期时间
            Issuer:    config.GetString("app.name"),   // 签名颁发者
        },
    }
	

    // 2. 根据 claims 生成token对象
    token, err := jwt.createToken(claims)
    if err != nil {
        logger.LogIf(err)
        return ""
    }

    return token
}