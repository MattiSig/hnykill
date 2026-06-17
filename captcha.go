package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// captchaSecret signs captcha tokens so the expected answer can travel in the
// (untrusted) form without being forgeable. It comes from CAPTCHA_SECRET, or a
// fresh random secret generated at startup — fine for a single-instance family
// app, where a restart simply invalidates any in-flight challenges.
var captchaSecret = loadCaptchaSecret()

// captchaTTL is how long a generated challenge stays valid.
const captchaTTL = 10 * time.Minute

func loadCaptchaSecret() []byte {
	if s := os.Getenv("CAPTCHA_SECRET"); s != "" {
		return []byte(s)
	}
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand should never fail; degrade to a time-seeded secret.
		return []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
	return b
}

// newCaptcha returns a human-readable arithmetic prompt and a signed, time-limited
// token that encodes the expected answer. The token is safe to embed in a hidden
// form field: it cannot be altered without invalidating the signature.
func newCaptcha() (prompt, token string) {
	a := mrand.Intn(9) + 1 // 1..9
	b := mrand.Intn(9) + 1 // 1..9
	answer := a + b
	prompt = fmt.Sprintf("What is %d + %d?", a, b)
	exp := time.Now().Add(captchaTTL).Unix()
	payload := fmt.Sprintf("%d|%d", exp, answer)
	sig := sign(payload)
	token = base64.RawURLEncoding.EncodeToString([]byte(payload)) + "." + sig
	return prompt, token
}

// verifyCaptcha reports whether userAnswer matches the answer encoded in token
// and the token is well-formed, correctly signed and not expired.
func verifyCaptcha(token, userAnswer string) bool {
	dot := strings.LastIndexByte(token, '.')
	if dot < 0 {
		return false
	}
	raw, sig := token[:dot], token[dot+1:]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(raw)
	if err != nil {
		return false
	}
	payload := string(payloadBytes)
	if !hmac.Equal([]byte(sig), []byte(sign(payload))) {
		return false
	}
	exp, answer, ok := splitPayload(payload)
	if !ok || time.Now().Unix() > exp {
		return false
	}
	got, err := strconv.Atoi(strings.TrimSpace(userAnswer))
	if err != nil {
		return false
	}
	return got == answer
}

func splitPayload(payload string) (exp int64, answer int, ok bool) {
	parts := strings.SplitN(payload, "|", 2)
	if len(parts) != 2 {
		return 0, 0, false
	}
	exp, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, false
	}
	answer, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, false
	}
	return exp, answer, true
}

func sign(payload string) string {
	mac := hmac.New(sha256.New, captchaSecret)
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
