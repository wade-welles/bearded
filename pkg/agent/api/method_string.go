// generated by stringer -type Method; DO NOT EDIT

package api

import "fmt"

const _Method_name = "PingConnectGetConfig"

var _Method_index = [...]uint8{0, 4, 11, 20}

func (i Method) String() string {
	if i < 0 || i+1 >= Method(len(_Method_index)) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return _Method_name[_Method_index[i]:_Method_index[i+1]]
}