package auth

import (
	"time"
	"errors"	
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/LNMMusic/goapi/config"
	"github.com/LNMMusic/goapi/handlers/db"
)

// NOT THE SAME AS MIDDLEWARE [JWT TOKEN MANAGE AUTH DATA]
type Claims map[string]interface{}
type Token struct {
	Id				uuid.UUID
	Expiration		int64				// unix
	ExpirationD		int64				// minutes (used only to set redis ttl)
	Claims			jwt.MapClaims
	Signing			string
}
func (t *Token) SetUp(exp int64, claims *Claims) {
	// exp in minutes
	t.Id 		 = uuid.New()
	t.Expiration = time.Now().Add(time.Minute * time.Duration(exp)).Unix()
	t.ExpirationD= exp

	t.Claims	 = jwt.MapClaims{
		"id":	t.Id,
		"exp" : t.Expiration,
		"expD": t.ExpirationD,
	};	for k, v := range *claims {t.Claims[k] = v}
}
func (t *Token) Sign(signKey string) error {
	var err error
	tc := jwt.NewWithClaims(jwt.SigningMethodHS256, t.Claims)
	t.Signing, err = tc.SignedString([]byte(signKey))
	if err != nil {
		return err
	}
	return nil
}
func (t *Token) SignRedis() error {
	// create session
	var userSessions UserSessions
	var session = UserSession {
		TokenID:		t.Id,
		TokenExp:		t.Expiration,
		TokenSign:		t.Signing,
		Active:	true,
	}
	key := t.Claims["user_id"].(uuid.UUID).String()


	// GET VALUES [sync userSessions]
	if err := JWTRedisGet(key, &userSessions); err == nil {
		userSessions.sync()

		if !userSessions.availability() {
			return errors.New("User Exceed Limit of Available Sessions [5]!")
		}
	}

	// add session
	userSessions = userSessions.addSession(session)

	
	// SET VALUES [userSessions]
	return JWTRedisSet(key, userSessions, t.Expiration)
}

type JWT struct {
	AccessToken		Token
	RefreshToken	Token
}
func CreateJWTToken(user_id uuid.UUID, user_isPremium bool, user_isAdmin bool) (*JWT, error) {
	var jwt = &JWT{}

	// MetaData
	jwt.AccessToken.SetUp(15, &Claims{
		"authorized":		true,
		"user_id":			user_id,
		"user_isPremium":	user_isPremium,
		"user_isAdmin":		user_isAdmin,
	})
	jwt.RefreshToken.SetUp(60*24*7, &Claims{
		"user_id":		user_id,
	})

	// Sign Up
	if err := jwt.AccessToken.Sign(config.EnvGet("SECRET_KEY")); err != nil {return nil, err}
	if err := jwt.RefreshToken.Sign(config.EnvGet("SECRET_KEY_REFRESH")); err!=nil {return nil, err}

	// Sign Up [Redis]
	if err := jwt.AccessToken.SignRedis(); err != nil {return nil, err}
	// if err := jwt.RefreshToken.SignRedis(); err != nil {return nil, err}

	return jwt, nil
}





// __________________________________________________________________________________________
// REDIS Handler [id -> token uuid]
type UserSession struct {
	TokenID		uuid.UUID
	TokenExp	int64				// unix
	TokenSign	string
	Active		bool
}
type UserSessions [5]UserSession
func (u *UserSessions) sync() {
	var now = time.Now().Unix()

	// drop expired tokens (set to nil)
	for ix, session := range u {
		if session.TokenExp <= now {
			u[ix] = UserSession{}
		}
	}

	// sort tokens by longer exp time
	limit := len(u)
	for i := 0; i < limit-1; i++ {
		var maxIx int
		for j := i+1; j < limit; j++ {
			if u[i].TokenExp < u[j].TokenExp {
				maxIx = j
			}
		}
		// switch
		var maxExpSession = u[maxIx]
		u[maxIx] = u[0]
		u[0] = maxExpSession
	}
}
func (u *UserSessions) availability() bool {
	// check no nil vals
	for _, session := range u {
		if !session.Active {
			return true
		}
	}
	return false
}
func (u UserSessions) addSession(session UserSession) UserSessions {
	// set new session at index 0
	var updatedSessions UserSessions
	updatedSessions[0] = session

	for ix, session := range u {
		if session.Active {
			updatedSessions[ix+1] = session
		} else {
			break
		}
	}
	return updatedSessions
}
func (u *UserSessions) deleteSession(tokenID string) {
	for ix, session := range u {
		if session.TokenID.String() == tokenID {
			u[ix] = UserSession{}
			break
		}
	}
}
// validation
func (u *UserSessions) tokenValid(tokenID string) bool {
	for _, session := range u {
		if session.TokenID.String() == tokenID {
			return true
		}
	}
	return false
}


// READ
func JWTRedisGet(key string, value interface{}) error {
	// Get JWT in Redis
	data, err := db.Redis.Db.Get(db.Redis.Ctx, key).Result(); if err != nil {
		return err
	}

	// Parse
	return json.Unmarshal([]byte(data), value)
}
// WRITE
func JWTRedisSet(key string, value interface{}, expiration int64) error {
	// Parse
	bytes, err := json.Marshal(value); if err != nil {return err}

	exp := time.Duration(expiration)*time.Minute
	return db.Redis.Db.Set(db.Redis.Ctx, key, string(bytes),exp).Err()
}



// ENDPOINTS Handlers
func JWTRedisHandlerValidate(key string, tokenID string) error {
	var userSessions UserSessions
	// get sessions
	if err := JWTRedisGet(key, &userSessions); err != nil {
		return err
	}

	if !userSessions.tokenValid(tokenID) {
		return errors.New("Token Has Expired!")
	}
	return nil
}
func JWTRedisHandlerGetSessions(key string) (UserSessions, error) {
	var userSessions UserSessions
	
	// Get
	if err := JWTRedisGet(key, &userSessions); err != nil {
		return userSessions, err
	}

	// Set [after sync]
	userSessions.sync()
	if err := JWTRedisSet(key, userSessions, 15); err != nil {
		return userSessions, err
	}

	return userSessions, nil
}
func JWTRedisHandlerDropSession(key string, tokenID string) error {
	var userSessions UserSessions

	// Get Sessions
	if err := JWTRedisGet(key, &userSessions); err != nil {
		return err
	}
	userSessions.deleteSession(tokenID)

	// Set Sessions
	return JWTRedisSet(key, userSessions, 15)
}
func JWTRedisHandlerDropSessions(key string) error {
	_, err := db.Redis.Db.Del(db.Redis.Ctx, key).Result()
	return err
}