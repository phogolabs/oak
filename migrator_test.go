package gom_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"
	"github.com/svett/gom"
)

var _ = Describe("Migrator", func() {
	var migrator *gom.Migrator

	BeforeEach(func() {
		dir, err := ioutil.TempDir("", "gom")
		Expect(err).To(BeNil())

		gateway, err := gom.Open("sqlite3", "/tmp/gom.db")
		Expect(err).To(BeNil())

		migrator = &gom.Migrator{
			Dir:     dir,
			Gateway: gateway,
		}
	})

	AfterEach(func() {
		Expect(migrator.Gateway.Close()).To(Succeed())
	})

	Describe("Setup", func() {
		It("setups the project successfully", func() {
			Expect(migrator.Setup()).To(Succeed())

			directories := []string{
				filepath.Join(migrator.Dir, "/database/migration"),
				filepath.Join(migrator.Dir, "/database/statement"),
			}

			for _, dir := range directories {
				info, err := os.Stat(dir)
				Expect(err).To(Succeed())
				Expect(info.IsDir()).To(BeTrue())
			}

			path := fmt.Sprintf("/database/migration/%s_setup.sql", gom.MinTime.Format(gom.DateTimeFormat))
			path = filepath.Join(migrator.Dir, path)
			Expect(path).To(BeARegularFile())
			Expect(path).To(HaveSuffix("00060524000000_setup.sql"))
		})

		Context("when the project has already been configured", func() {
			JustBeforeEach(func() {
				Expect(migrator.Setup()).To(Succeed())
			})

			It("returns an error", func() {
				Expect(migrator.Setup()).To(MatchError("The project has already been configured"))
			})
		})
	})

	Describe("Create", func() {
		It("creates a migration successfully", func() {
			Expect(migrator.Setup()).To(Succeed())

			path, err := migrator.Create("test")
			Expect(err).To(BeNil())
			Expect(path).To(BeARegularFile())
			Expect(path).To(HaveSuffix("_test.sql"))
		})

		Context("when the project is not configured", func() {
			It("returns an error", func() {
				path, err := migrator.Create("test")
				Expect(err).To(MatchError("The project has not been configured"))
				Expect(path).To(BeEmpty())
			})
		})
	})
})
