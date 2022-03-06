package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
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
	json := jsoniter.ConfigFastest
	buf := bufio.NewReader(r)

	var user User
	result := make(DomainStat)
	i := 0
	for {
		line, err := buf.ReadBytes('\n')
		if len(line) == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err = json.Unmarshal(line, &user); err != nil {
			return nil, err
		}

		countDomains(user, domain, result)
		i++
	}

	return result, nil
}

func countDomains(user User, domain string, result DomainStat) {
	if strings.HasSuffix(user.Email, "."+domain) {
		index := strings.LastIndexByte(user.Email, '@')
		if index > 0 {
			domain := user.Email[index+1:]
			result[strings.ToLower(domain)]++
		}
	}
}
