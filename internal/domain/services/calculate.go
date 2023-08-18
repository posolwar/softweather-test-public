package services

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var ErrUnexpectedSymb = errors.New("unexpected symbol")

// Method for solving mathematical expressions
//
// Example: "2+2-3-5+1" -> (-3, nil)
func Calculate(expression string) (int, error) {
	var (
		amount         int
		operationCache rune            = '+'
		stringCache    strings.Builder // TODO: grow или замена на []string, отрезок от expression
		err            error
	)

	defer stringCache.Reset()

	for _, char := range expression {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			stringCache.WriteRune(char)

			continue
		case '+', '-':
			if stringCache.Len() == 0 {
				continue
			}

			amount, err = doOperation(&stringCache, amount, operationCache)
			if err != nil {
				return 0, err
			}

			stringCache.Reset()

			operationCache = char
		default:
			return 0, errors.Wrap(ErrUnexpectedSymb, "getted symb: "+string(char))
		}
	}

	// Для очистки кэша (работа с последней цифрой)
	amount, err = doOperation(&stringCache, amount, operationCache)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func doOperation(builder *strings.Builder, amount int, operation rune) (int, error) {
	number, err := strconv.Atoi(builder.String())
	if err != nil {
		return 0, errors.Wrap(err, "ошибка получения числа из строки")
	}

	switch operation {
	case '+':
		return amount + number, nil
	case '-':
		return amount - number, nil
	default:
		return 0, errors.Wrap(ErrUnexpectedSymb, "getted symb: "+string(operation))
	}
}
