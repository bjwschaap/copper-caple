package orm_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/pg.v5/orm"
)

type A struct {
	Id int
}

func (A) Method() int {
	return 10
}

type B struct {
	A
}

var _ = Describe("Model embedding", func() {
	var strct reflect.Value
	var table *orm.Table

	BeforeEach(func() {
		strct = reflect.ValueOf(B{A: A{Id: 1}})
		table = orm.Tables.Get(strct.Type())
	})

	It("has fields", func() {
		Expect(table.Fields).To(HaveLen(1))
		Expect(table.FieldsMap).To(HaveLen(1))

		id, ok := table.FieldsMap["id"]
		Expect(ok).To(BeTrue())
		Expect(id.GoName).To(Equal("Id"))
		Expect(id.SQLName).To(Equal("id"))
		Expect(string(id.ColName)).To(Equal(`"id"`))
		Expect(id.Has(orm.PrimaryKeyFlag)).To(BeTrue())
		Expect(string(id.AppendValue(nil, strct, 1))).To(Equal("1"))

		Expect(table.PKs).To(HaveLen(1))
		Expect(table.PKs[0]).To(Equal(id))
	})

	It("has methods", func() {
		Expect(table.Methods).To(HaveLen(1))

		m, ok := table.Methods["Method"]
		Expect(ok).To(BeTrue())
		Expect(m.Index).To(Equal(0))
		Expect(string(m.AppendValue(nil, strct, 1))).To(Equal("10"))
	})
})

type C struct {
	Name int `sql:",pk"`
	Id   int
	UUID int
}

var _ = Describe("primary key annotation", func() {
	var table *orm.Table

	BeforeEach(func() {
		strct := reflect.ValueOf(C{})
		table = orm.Tables.Get(strct.Type())
	})

	It("has precedence over auto-detection", func() {
		Expect(table.PKs).To(HaveLen(1))
		Expect(table.PKs[0].GoName).To(Equal("Name"))
	})
})

type D struct {
	UUID int
}

var _ = Describe("uuid field", func() {
	var table *orm.Table

	BeforeEach(func() {
		strct := reflect.ValueOf(D{})
		table = orm.Tables.Get(strct.Type())
	})

	It("is detected as primary key", func() {
		Expect(table.PKs).To(HaveLen(1))
		Expect(table.PKs[0].GoName).To(Equal("UUID"))
	})
})

type E struct {
	Id          int
	StructField struct {
		Foo string
		Bar string
	}
}

var _ = Describe("struct field", func() {
	var table *orm.Table

	BeforeEach(func() {
		strct := reflect.ValueOf(E{})
		table = orm.Tables.Get(strct.Type())
	})

	It("is present in the list", func() {
		Expect(table.Fields).To(HaveLen(2))

		_, ok := table.FieldsMap["struct_field"]
		Expect(ok).To(BeTrue())
	})
})
