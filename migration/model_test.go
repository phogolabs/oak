package migration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/svett/gom/migration"
)

var _ = Describe("Model", func() {
	Describe("Item", func() {
		It("returns the filename correctly", func() {
			item := &migration.Item{
				Id:          "id",
				Description: "schema",
			}

			Expect(item.Filename()).To(Equal("id_schema.sql"))
		})

		Describe("Parse", func() {
			It("parses the item successfully", func() {
				filename := "20060102150405_schema.sql"
				item, err := migration.Parse(filename)
				Expect(err).NotTo(HaveOccurred())
				Expect(item.Id).To(Equal("20060102150405"))
				Expect(item.Description).To(Equal("schema"))
			})

			Context("when the filename is hes longer description", func() {
				It("parses the item successfully", func() {
					filename := "20060102150405_my_schema_for_this_db.sql"
					item, err := migration.Parse(filename)
					Expect(err).NotTo(HaveOccurred())
					Expect(item.Id).To(Equal("20060102150405"))
					Expect(item.Description).To(Equal("my_schema_for_this_db"))
				})
			})

			Context("when the filename does not contain two parts", func() {
				It("returns an error", func() {
					filename := "schema.sql"
					item, err := migration.Parse(filename)
					Expect(err).To(MatchError("Migration 'schema.sql' has an invalid file name"))
					Expect(item).To(BeNil())
				})
			})

			Context("when the filename does not have timestamp in its name", func() {
				It("returns an error", func() {
					filename := "id_schema.sql"
					item, err := migration.Parse(filename)
					Expect(err).To(MatchError("Migration 'id_schema.sql' has an invalid file name"))
					Expect(item).To(BeNil())
				})
			})
		})
	})
})
