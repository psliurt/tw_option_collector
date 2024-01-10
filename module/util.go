package module

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func strToInt(str string) int {
	if strings.Trim(strings.Trim(str, " "), "\n") == "-" {
		return 0
	}
	str = strings.ReplaceAll(str, ",", "")
	rst, err := strconv.Atoi(str)
	handleError(err)
	return rst
}

func strToDecimal(str string) decimal.Decimal {
	if strings.Trim(strings.Trim(str, " "), "\n") == "-" {
		return decimal.NewFromInt(0)
	}
	str = strings.ReplaceAll(str, ",", "")
	rst, err := decimal.NewFromString(str)
	handleError(err)
	return rst
}

func strToIntData(str string, inbracket bool) int {
	if strings.Trim(strings.Trim(str, " "), "\n") == "-" {
		return 0
	}
	indexOfLeftBracket := strings.Index(str, "(")
	if !inbracket {
		return strToInt(strings.Trim(strings.Trim(str[0:indexOfLeftBracket], ""), "\n"))
	}

	indexOfRightBracket := strings.Index(str, ")")
	return strToInt(strings.Trim(strings.Trim(str[indexOfLeftBracket+1:indexOfRightBracket], ""), "\n"))

}

func strToDecimalData(str string, inbracket bool) decimal.Decimal {
	if strings.Trim(strings.Trim(str, " "), "\n") == "-" {
		return decimal.NewFromInt(0)
	}
	indexOfFirstPercentSign := strings.Index(str, "%")
	indexOfLeftBracket := strings.Index(str, "(")
	if !inbracket {
		return strToDecimal(strings.Trim(strings.Trim(str[0:indexOfFirstPercentSign], ""), "\n"))
	}

	indexOfRightBracket := strings.Index(str, ")")
	inBracketStr := str[indexOfLeftBracket+1 : indexOfRightBracket]
	indexOfSecondPercentSign := strings.Index(inBracketStr, "%")
	return strToDecimal(strings.Trim(strings.Trim(inBracketStr[0:indexOfSecondPercentSign], ""), "\n"))
}

func deleteSpace(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, " ", ""), "\n", "")
}
