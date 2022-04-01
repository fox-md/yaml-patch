package yamlpatch_test

import (
	"bytes"
	"testing"

	yamlpatch "github.com/fox-md/yaml-patch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateAddPatch(t *testing.T) {

	addPatchBytes := []byte(`---
- op: add
  path: /login
  value: mike`)

  	addPatchString := "+login=mike"

	if !bytes.Equal(addPatchBytes, yamlpatch.CreatePatch(addPatchString)) {
		t.Fatalf(`Expected and actual values do not match`)
	}

}

func TestCreateRemovePatch(t *testing.T) {

	removePatchBytes := []byte(`---
- op: remove
  path: /login
  value: `)

	removePatchString := "-login"

	removePatch := yamlpatch.CreatePatch(removePatchString)

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

	removePatch := yamlpatch.CreatePatch(removePatchString)

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

	yamlFileOutActual := yamlpatch.PatchFile(patchString, yamlFile)

	if !bytes.Equal(yamlFileOut, yamlFileOutActual) {
		t.Fatalf(`Expected %q and actual %q values do not match`, yamlFileOut, yamlFileOutActual)
	}

}


var _ = Describe("Patch", func() {

	Describe("PatchFile", func() {

		Context("When patching should be OK", func() {

			It("Patch user and replicas", func() {

				yamlFile := []byte(`---
user: bob
spec:
  replicas: 0
tier: backend`)
			
				yamlFileExpected := []byte(`spec:
  replicas: 5
tier: backend
user: fox
`)
			
				patchString := "spec.replicas=5;user=fox"
				
				yamlFileActual := yamlpatch.PatchFile(patchString, yamlFile)

				Expect(yamlFileExpected).To(Equal(yamlFileActual))
			})

		})


		Context("When patching should not be OK", func() {

			It("Patch user and replicas", func() {

				yamlFile := []byte(`---
user: bob
spec:
  replicas: 0
tier: backend`)
			
				yamlFileExpected := []byte(`spec:
  replicas: 5
tier: backend
user: fox
`)
			
				patchString := "spec.replicas=5;user=mike"
				
				yamlFileActual := yamlpatch.PatchFile(patchString, yamlFile)

				Expect(yamlFileExpected).NotTo(Equal(yamlFileActual))
			})
		})
	})

})