package yamlpatch

import (
	"bytes"
	"testing"
)

func TestCreateAddPatch(t *testing.T) {

	addPatchBytes := []byte(`---
- op: add
  path: /login
  value: mike`)

  	addPatchString := "+login=mike"

	if !bytes.Equal(addPatchBytes, CreatePatch(addPatchString)) {
		t.Fatalf(`Expected and actual values do not match`)
	}

}

func TestCreateRemovePatch(t *testing.T) {

	removePatchBytes := []byte(`---
- op: remove
  path: /login
  value: `)

	removePatchString := "-login"

	removePatch := CreatePatch(removePatchString)

	if !bytes.Equal(removePatchBytes, removePatch) {
		t.Fatalf(`Expected %q and actual %q values do not match`, removePatchBytes, removePatch)
	}
}

func TestCreateReplacePatch(t *testing.T) {

	removePatchBytes := []byte(`---
- op: replace
  path: /spec/replicas
  value: 5`)

	removePatchString := "spec.replicas=5"

	removePatch := CreatePatch(removePatchString)

	if !bytes.Equal(removePatchBytes, removePatch) {
		t.Fatalf(`Expected %q and actual %q values do not match`, removePatchBytes, removePatch)
	}
}

func TestFilePatch(t *testing.T) {

	yamlFile := []byte(`---
user: bob
spec:
  replicas: 0
tier: backend`)

	yamlFileOut := []byte(`spec:
  replicas: 5
tier: backend
user: fox
`)

	patchString := "spec.replicas=5;user=fox"

	yamlFileOutActual := PatchFile(patchString, yamlFile)

	if !bytes.Equal(yamlFileOut, yamlFileOutActual) {
		t.Fatalf(`Expected %q and actual %q values do not match`, yamlFileOut, yamlFileOutActual)
	}

}