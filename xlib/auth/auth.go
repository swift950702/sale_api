package auth

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	uuid "github.com/satori/go.uuid"
)

// 签名算法, 随机, 不保存密钥, 每次都是随机的
// var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// var publicKey = &privateKey.PublicKey
// var hs = jwt.NewES256(
// 	jwt.ECDSAPublicKey(publicKey),
// 	jwt.ECDSAPrivateKey(privateKey),
// )
var hs = jwt.NewHS256([]byte("secret"))

// 记录登录信息的 JWT
type TokenPayload struct {
	jwt.Payload
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

// 签名
func Sign(username, password string) (string, error) {
	now := time.Now()
	pl := TokenPayload{
		Payload: jwt.Payload{
			Issuer:         "coolcat",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		Pwd:      password,
		Username: username,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

// 验证
func Verify(token []byte) (*TokenPayload, error) {
	pl := &TokenPayload{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}
