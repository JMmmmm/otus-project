package hw10programoptimization

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	err = nil
	var user User
	var json = jsoniter.ConfigFastest
	buf := bufio.NewReader(r)
	i := 0
	for {
		line, err := buf.ReadBytes('\n')
		if len(line) == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return result, err
		}

		if err = json.Unmarshal(line, &user); err != nil {
			return result, err
		}
		result[i] = user
		i++
	}

	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		if strings.HasSuffix(user.Email, "."+domain) {
			index := strings.LastIndexByte(user.Email, '@')
			if index > 0 {
				domain := user.Email[index+1:]
				result[strings.ToLower(domain)]++
			}
		}
	}
	return result, nil
}
