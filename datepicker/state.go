package datepicker

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	queryDataSeparator = ":"
)

type state struct {
	cmd   int
	param int
}

func (datePicker *DatePicker) encodeState(st state) string {
	var parts []string

	parts = append(parts, strconv.Itoa(st.cmd))
	parts = append(parts, strconv.Itoa(st.param))

	return datePicker.prefix + strings.Join(parts, queryDataSeparator)
}

func (datePicker *DatePicker) decodeState(queryData string) state {
	parts := strings.Split(strings.TrimPrefix(queryData, datePicker.prefix), queryDataSeparator)

	if len(parts) != 2 {
		panic(fmt.Errorf("invalid data format, expected 2 parts, got %d", len(parts)))
	}

	cmd, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Errorf("invalid command: %s", err))
	}

	param, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Errorf("invalid parameter: %s", err))
	}

	return state{
		cmd:   cmd,
		param: param,
	}
}
