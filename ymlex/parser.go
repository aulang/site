package ymlex

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"regexp"
	"strings"
)

// LF 换行符
const LF = '\n'

// REGULAR 暂时是支持一个，不支持嵌套
// 格式如：${ENV_NAME: defaultValue}，defaultValue可选
const REGULAR = "\\$\\{([^:^{^}]+):?([^:^{^}]*)\\}"

var regex *regexp.Regexp = nil

// 初始化正则表达式
func init() {
	var err error

	regex, err = regexp.Compile(REGULAR)

	if err != nil {
		panic("正则表达式错误：" + REGULAR)
	}
}

func parseLine(line []byte) []byte {
	matches := regex.FindSubmatch(line)

	if len(matches) == 0 {
		return line
	}

	// matches[0] = ${ENV_NAME: defaultValue}
	// matches[1] = ENV_NAME
	envValue := os.Getenv(strings.TrimSpace(string(matches[1])))
	// matches[2] = defaultValue
	if envValue == "" && matches[2] != nil {
		envValue = strings.TrimSpace(string(matches[2]))
	}

	// 替换所有的${ENV_NAME: defaultValue}
	return regex.ReplaceAll(line, []byte(envValue))
}

func Unmarshal(file []byte, out interface{}) (err error) {
	inBuffer := bytes.NewBuffer(file)

	var outBuffer bytes.Buffer
	for {
		line, err := inBuffer.ReadBytes(LF)

		if err == io.EOF {
			break
		}

		outBuffer.Write(parseLine(line))
	}

	return yaml.Unmarshal(outBuffer.Bytes(), out)
}
