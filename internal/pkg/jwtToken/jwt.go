// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package jwtToken

import (
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenExpireDuration = 7 * 24 * time.Hour
)

var jwtSecret = []byte("kdd452-934sg4-l4d4q6")

type Claims struct {
	ID       int64  `json:"id"`
	Source   string `json:"source"`
	Sign     string `json:"sign"`
	Business string `json:"business"`
	RoleId   int64  `json:"role_id"`
	jwt.StandardClaims
}

func Generate2Token(id int64, source string, sign string, roleId int64) (accessToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(AccessTokenExpireDuration)
	claims := Claims{
		ID:     id,
		Source: source,
		Sign:   sign,
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "test-user",
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", errors.Wrap(err, "failed to get accessToken")
	}

	return accessToken, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims == nil {
		return nil, errors.Wrap(err, "failed to parse token")
	}

	var claims *Claims
	var ok bool
	if claims, ok = tokenClaims.Claims.(*Claims); !ok || !tokenClaims.Valid {
		return nil, errors.Wrap(err, "failed to valid token")

	}

	if claims.ExpiresAt > time.Now().Unix() {
		return nil, errors.Wrap(err, "token is expires time")
	}

	return claims, nil

}
