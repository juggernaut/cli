package app_files_test

import (
	. "cf/app_files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path"
	"path/filepath"
)

var _ = Describe("AppFiles", func() {
	fixturePath := filepath.Join("..", "..", "fixtures", "applications")

	Describe("AppFilesInDir", func() {
		It("all files have '/' path separators", func() {
			files, err := AppFilesInDir(fixturePath)
			Expect(err).ShouldNot(HaveOccurred())

			for _, afile := range files {
				Expect(afile.Path).Should(Equal(filepath.ToSlash(afile.Path)))
			}
		})

		It("excludes files based on the .cfignore file", func() {
			appPath := filepath.Join(fixturePath, "app-with-cfignore")
			files, err := AppFilesInDir(appPath)
			Expect(err).ShouldNot(HaveOccurred())

			paths := []string{}
			for _, file := range files {
				paths = append(paths, file.Path)
			}

			Expect(paths).To(Equal([]string{
				path.Join("dir1", "child-dir", "file3.txt"),
				path.Join("dir1", "file1.txt"),
			}))
		})
	})
})
