package pg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/pg.v5"
	"gopkg.in/pg.v5/orm"
)

type HookTest struct {
	Id    int
	Value string

	afterQuery  int
	afterSelect int

	beforeInsert int
	afterInsert  int

	beforeUpdate int
	afterUpdate  int

	beforeDelete int
	afterDelete  int
}

func (t *HookTest) AfterQuery(db orm.DB) error {
	t.afterQuery++
	return nil
}

func (t *HookTest) AfterSelect(db orm.DB) error {
	t.afterSelect++
	return nil
}

func (t *HookTest) BeforeInsert(db orm.DB) error {
	t.beforeInsert++
	return nil
}

func (t *HookTest) AfterInsert(db orm.DB) error {
	t.afterInsert++
	return nil
}

func (t *HookTest) BeforeUpdate(db orm.DB) error {
	t.beforeUpdate++
	return nil
}

func (t *HookTest) AfterUpdate(db orm.DB) error {
	t.afterUpdate++
	return nil
}

func (t *HookTest) BeforeDelete(db orm.DB) error {
	t.beforeDelete++
	return nil
}

func (t *HookTest) AfterDelete(db orm.DB) error {
	t.afterDelete++
	return nil
}

var _ = Describe("HookTest", func() {
	var db *pg.DB

	BeforeEach(func() {
		db = pg.Connect(pgOptions())

		qs := []string{
			"CREATE TEMP TABLE hook_tests (id int, value text)",
			"INSERT INTO hook_tests VALUES (1, '')",
		}
		for _, q := range qs {
			_, err := db.Exec(q)
			Expect(err).NotTo(HaveOccurred())
		}
	})

	AfterEach(func() {
		Expect(db.Close()).NotTo(HaveOccurred())
	})

	It("calls AfterQuery for struct", func() {
		var hook HookTest
		_, err := db.QueryOne(&hook, "SELECT 1 AS id")
		Expect(err).NotTo(HaveOccurred())
		Expect(hook.afterQuery).To(Equal(1))
		Expect(hook.afterSelect).To(Equal(0))
	})

	It("calls AfterQuery and AfterSelect for struct model", func() {
		var hook HookTest
		err := db.Model(&hook).Select()
		Expect(err).NotTo(HaveOccurred())
		Expect(hook.afterQuery).To(Equal(1))
		Expect(hook.afterSelect).To(Equal(1))
	})

	It("calls AfterQuery for slice", func() {
		var hooks []HookTest
		_, err := db.Query(&hooks, "SELECT 1 AS id")
		Expect(err).NotTo(HaveOccurred())
		Expect(hooks).To(HaveLen(1))
		Expect(hooks[0].afterQuery).To(Equal(1))
		Expect(hooks[0].afterSelect).To(Equal(0))
	})

	It("calls AfterQuery and AfterSelect for slice model", func() {
		var hooks []HookTest
		err := db.Model(&hooks).Select()
		Expect(err).NotTo(HaveOccurred())
		Expect(hooks).To(HaveLen(1))
		Expect(hooks[0].afterQuery).To(Equal(1))
		Expect(hooks[0].afterSelect).To(Equal(1))
	})

	It("calls BeforeInsert and AfterInsert", func() {
		hook := &HookTest{
			Id:    1,
			Value: "value",
		}
		err := db.Insert(&hook)
		Expect(err).NotTo(HaveOccurred())
		Expect(hook.afterQuery).To(Equal(0))
		Expect(hook.beforeInsert).To(Equal(1))
		Expect(hook.afterInsert).To(Equal(1))
	})

	It("calls BeforeUpdate and AfterUpdate", func() {
		hook := &HookTest{
			Id: 1,
		}
		err := db.Update(&hook)
		Expect(err).NotTo(HaveOccurred())
		Expect(hook.afterQuery).To(Equal(0))
		Expect(hook.beforeUpdate).To(Equal(1))
		Expect(hook.afterUpdate).To(Equal(1))
	})

	It("calls BeforeDelete and AfterDelete", func() {
		hook := &HookTest{
			Id: 1,
		}
		err := db.Delete(&hook)
		Expect(err).NotTo(HaveOccurred())
		Expect(hook.afterQuery).To(Equal(0))
		Expect(hook.beforeDelete).To(Equal(1))
		Expect(hook.afterDelete).To(Equal(1))
	})
})
