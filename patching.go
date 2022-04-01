package yamlpatch

import (
	"log"
	"strings"
	"unicode/utf8"
	zlog "github.com/rs/zerolog/log"
)


func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func CreatePatch(patch string) (ops []byte) {

	var patchOp string
	var patchValue string

	zlog.Debug().Msg("Received patch string " + patch)

	patchPath := strings.Split(patch, "=")[0]

	if len(strings.Split(patch, "=")) == 2 {
		patchValue = strings.Split(patch, "=")[1]
	}

	if strings.HasPrefix(patchPath, "+") {
		patchOp = "add"
		patchPath = trimFirstRune(patchPath)
	} else if strings.HasPrefix(patchPath, "-") {
		patchOp = "remove"
		patchPath = trimFirstRune(patchPath)
	} else {
		patchOp = "replace"
	}

	patchPath = "/" + strings.Replace(patchPath, ".", "/", -1)

	ops = []byte(`---
- op: ` + patchOp + `
  path: ` + patchPath + `
  value: ` + patchValue)
	return
}

func PatchFile(patchLine string, yfile []byte) (bs []byte){
	for _, p := range strings.Split(patchLine, ";") {
		ops := CreatePatch(p)
		patch, err := DecodePatch(ops)
		if err != nil {
			log.Fatalf("decoding patch failed: %s", err)
		}

		bs, err = patch.Apply(yfile)
		if err != nil {
			log.Fatalf("applying patch failed: %s", err)
		}
		yfile = bs
	}
	return
}
