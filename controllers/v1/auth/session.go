package controller

import (
	"bytes"
	"crypto/md5"
	"errors"
	"io"
	"time"
)

type Session struct {
	Sid     []byte
	Created time.Time
	Expiry  uint
}

var Sessions []Session

func (s *Session) GenSessionToken() {
	h := md5.New()
	s.Created = time.Now()

	cts, _ := s.Created.MarshalText()

	io.WriteString(h, "somerandomtext")
	io.WriteString(h, string(cts))

	s.Sid = h.Sum(nil)

	//  amongst the stupid md5 implementation, if you ever wondered
	//  why expiry is set to 1200secs, BECAUSE my college DECIDED TO
	//  GIVE ME THAT MANY SECONDS OF INTERNET ACCESS BEFORE THEY KICK OUT
	//  FUCKING HELL
	s.Expiry = 1200

	Sessions = append(Sessions, *s)
}

func MatchToken(token []byte) (bool, error) {

	for _, s := range Sessions {
		if bytes.Equal(s.Sid, token) {
			if uint(time.Now().Unix()-s.Created.Unix()) < s.Expiry {
				return true, nil
			} else {
				return false, errors.New("Session Expired")
			}
		}
	}
	return false, errors.New("Session Not Found")
}
