package orm

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CreateTableTest struct {
	Id      int
	Int8    int8
	Uint8   uint8
	Int16   int16
	Uint16  uint16
	Int32   int32
	Uint32  uint32
	Int64   int64
	Uint64  uint64
	Float32 float32
	Float64 float64
	String  string
	Varchar string `sql:",type:varchar(500)"`
	Time    time.Time
}

type CreateTableWithoutPKTest struct {
	String string
}

var _ = Describe("CreateTable", func() {
	It("creates new table", func() {
		b, err := createTableQuery{model: CreateTableTest{}}.AppendQuery(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(b)).To(Equal(`CREATE TABLE "create_table_tests" (id bigserial, int8 smallint, uint8 smallint, int16 smallint, uint16 integer, int32 integer, uint32 bigint, int64 bigint, uint64 decimal, float32 real, float64 double precision, string text, varchar varchar(500), time timestamptz, PRIMARY KEY (id))`))
	})

	It("creates new table without primary key", func() {
		b, err := createTableQuery{model: CreateTableWithoutPKTest{}}.AppendQuery(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(b)).To(Equal(`CREATE TABLE "create_table_without_pk_tests" (string text)`))
	})
})
