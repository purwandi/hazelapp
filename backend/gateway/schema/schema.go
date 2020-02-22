package schema

// $ cd gateway
// $ go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...

import (
	"bytes"
)

// GetRootSchema is to get root graphql schema
func GetRootSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
