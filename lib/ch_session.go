package lib

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Mr-YongXuan/chainx/include"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Meta struct{
	Data   interface{}
	Expire int64
	token  string
}

type Sessions struct {
	DefaultExpire  int64
	ActiveSessions map [string]Meta
	TokenMap       map [string]string
	protect        sync.Mutex
}

func NewSessions(defaultExpire int64) *Sessions {
	s := &Sessions{}
	s.ActiveSessions = make(map [string]Meta)
	s.TokenMap = make(map [string]string)
	rand.Seed(time.Now().UnixNano())
	// seconds
	s.DefaultExpire = defaultExpire
	return s
}

/* Register add user information into sessions table */
func (s *Sessions) Register (res *include.ChResponse, identifier string, meta interface{}, expires int64) (token string, ok bool) {
	if identifier != "" {
		// sum token
		buffer := bytes.Buffer{}
		buffer.WriteString(identifier)
		buffer.WriteString(strconv.Itoa(time.Now().Nanosecond()))
		buffer.WriteString(strconv.Itoa(rand.Intn(90000)))
		token = fmt.Sprintf("%x", sha1.Sum(buffer.Bytes()))
		var currentExpires int64
		if expires == 0 {
			currentExpires = time.Now().Unix() + s.DefaultExpire
		} else {
			currentExpires = time.Now().Unix() + expires
		}

		s.protect.Lock()
		//register to sessions table
		s.ActiveSessions[identifier] = struct {
			Data   interface{}
			Expire int64
			token  string
		}{Data: meta, Expire: currentExpires, token: token}
		//register to token map
		s.TokenMap[token] = identifier
		cookieExpires := time.Unix(currentExpires, 0).Format("Mon, 02 Jan 2006 03:04:05 GMT")
		res.Headers["Set-Cookie"] = fmt.Sprintf("ChainxSession=%s; expires=%s", token, cookieExpires)
		s.protect.Unlock()
		ok = true

	} else {
		ok = false
	}
	return
}

/* TokenAvailable if token valid and not expired then return true */
func (s * Sessions) TokenAvailable (token string) bool {
	s.protect.Lock()
	defer s.protect.Unlock()
	if val, ok := s.TokenMap[token]; ok {
		if session, ok := s.ActiveSessions[val]; ok {
			if time.Now().Unix() < session.Expire {
				return true
			}
		}
	}
	return false
}

/* TokenAvailableWithIdentifier <- as you see */
func (s * Sessions) TokenAvailableWithIdentifier (identifier string) bool {
	s.protect.Lock()
	defer s.protect.Unlock()
	if session, ok := s.ActiveSessions[identifier]; ok {
		if time.Now().Unix() < session.Expire {
			return true
		}
	}
	return false
}

/* GetMetaWithToken <- as you see */
func (s *Sessions) GetMetaWithToken(token string) (data Meta, identifier string, err error) {
	if s.TokenAvailable(token) {
		s.protect.Lock()
		identifier = s.TokenMap[token]
		data = s.ActiveSessions[identifier]
		s.protect.Unlock()
		return
	}

	err = errors.New("token invalid or expired")
	return
}

func (s *Sessions) GetTokenWithRequest(req *include.ChRequest) (data Meta, identifier string, err error) {
	// parse request cookie
	if cookies := req.GetHeader("Cookie"); cookies != "" {
		for _, cookie := range strings.Split(cookies, "; ") {
			if chainxSession := strings.Split(cookie, "="); chainxSession[0] == "ChainxSession" && len(chainxSession) == 2 {
				data, identifier, err = s.GetMetaWithToken(chainxSession[1])
				return
			}
		}
	}

	err = errors.New("token invalid or expired")
	return
}

/* GetIdentifierWithToken <- as you see */
func (s *Sessions) GetIdentifierWithToken(token string) (identifier string, err error) {
	if s.TokenAvailable(token) {
		s.protect.Lock()
		identifier = s.TokenMap[token]
		s.protect.Unlock()
		return
	}

	err = errors.New("token invalid or expired")
	return
}

/* GetMetaWithIdentifier <- as you see */
func (s *Sessions) GetMetaWithIdentifier(identifier string) (data Meta, err error) {
	if s.TokenAvailableWithIdentifier(identifier) {
		s.protect.Lock()
		data = s.ActiveSessions[identifier]
		s.protect.Unlock()
		return
	}
	err = errors.New("token invalid or expired")
	return
}

/* GetTokenWithIdentifier <- as you see */
func (s *Sessions) GetTokenWithIdentifier(identifier string) (token string, err error) {
	if s.TokenAvailableWithIdentifier(identifier) {
		s.protect.Lock()
		token = s.ActiveSessions[identifier].token
		s.protect.Unlock()
		return
	}

	err = errors.New("token invalid or expired")
	return
}
