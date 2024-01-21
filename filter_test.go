package kyte

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("without options", func(t *testing.T) {
		filter := Filter()
		if filter == nil {
			t.Error("Filter should not be nil")
		}

		if filter.kyte.source != nil {
			t.Error("Filter.kyte should be nil")
		}

		if filter.kyte.checkField != false {
			t.Error("Filter.kyte.checkField should be false")
		}
	})

	t.Run("with options", func(t *testing.T) {
		type Temp struct {
			Name string `bson:"name"`
		}

		filter := Filter(Source(&Temp{}), ValidateField(true))
		if filter == nil {
			t.Error("Filter should not be nil")
		}

		if filter.kyte.source == nil {
			t.Error("Filter.kyte should not be nil")
		}

		if filter.kyte.checkField != true {
			t.Error("Filter.kyte.checkField should be true")
		}
	})

	t.Run("with validate field is false", func(t *testing.T) {
		filter := Filter(ValidateField(false))
		if filter == nil {
			t.Error("Filter should not be nil")
		}

		if filter.kyte.source != nil {
			t.Error("Filter.kyte should be nil")
		}

		if filter.kyte.checkField != false {
			t.Error("Filter.kyte.checkField should be false")
		}
	})

}

func TestFilter_Equal(t *testing.T) {
	t.Parallel()

	t.Run("without source", func(t *testing.T) {
		q, err := Filter().Equal("name", "kyte").Build()
		if err != nil {
			t.Errorf("Filter.Equal should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.Equal should not return nil")
		}

		if q[0].Key != "name" {
			t.Errorf("Filter.Equal should return key name, got %v", q[0].Key)
		}

		if q[0].Value.(bson.M)["$eq"] != "kyte" {
			t.Errorf("Filter.Equal should return value map[$eq:kyte], got %v", q[0].Value)
		}
	})

	t.Run("with source", func(t *testing.T) {
		type Temp struct {
			Name string `bson:"name"`
		}
		var temp Temp
		q, err := Filter(Source(&temp)).Equal(&temp.Name, "kyte").Build()
		if err != nil {
			t.Errorf("Filter.Equal should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.Equal should not return nil")
		}

		if q[0].Key != "name" {
			t.Errorf("Filter.Equal should return key name, got %v", q[0].Key)
		}

		if q[0].Value.(bson.M)["$eq"] != "kyte" {
			t.Errorf("Filter.Equal should return value map[$eq:kyte], got %v", q[0].Value)
		}
	})

	t.Run("multiple", func(t *testing.T) {
		type Temp struct {
			Name    string `bson:"name"`
			Surname string `bson:"surname"`
			Age     int    `bson:"age"`
		}

		var temp Temp
		q, err := Filter(Source(&temp)).
			Equal(&temp.Name, "Joe").
			Equal(&temp.Surname, "Doe").
			Equal(&temp.Age, 10).
			Build()

		if err != nil {
			t.Errorf("Filter.Equal should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.Equal should not return nil")
		}

		for _, v := range q {
			if v.Key == "name" {
				if v.Value.(bson.M)["$eq"] != "Joe" {
					t.Errorf("Filter.Equal should return value map[$eq:Joe], got %v", v.Value)
				}
			}

			if v.Key == "surname" {
				if v.Value.(bson.M)["$eq"] != "Doe" {
					t.Errorf("Filter.Equal should return value map[$eq:Doe], got %v", v.Value)
				}
			}

			if v.Key == "age" {
				if v.Value.(bson.M)["$eq"] != 10 {
					t.Errorf("Filter.Equal should return value map[$eq:10], got %v", v.Value)
				}
			}
		}

	})
}

func TestFilter_NotEqual(t *testing.T) {
	t.Parallel()

	t.Run("without source", func(t *testing.T) {
		q, err := Filter().NotEqual("name", "kyte").Build()
		if err != nil {
			t.Errorf("Filter.NotEqual should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.NotEqual should not return nil")
		}

		if q[0].Key != "name" {
			t.Errorf("Filter.NotEqual should return key name, got %v", q[0].Key)
		}

		if q[0].Value.(bson.M)["$ne"] != "kyte" {
			t.Errorf("Filter.NotEqual should return value map[$ne:kyte], got %v", q[0].Value)
		}
	})

	t.Run("with source", func(t *testing.T) {
		type Temp struct {
			Name string `bson:"name"`
		}
		var temp Temp
		q, err := Filter(Source(&temp)).NotEqual(&temp.Name, "kyte").Build()
		if err != nil {
			t.Errorf("Filter.NotEqual should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.NotEqual should not return nil")
		}

		if q[0].Key != "name" {
			t.Errorf("Filter.NotEqual should return key name, got %v", q[0].Key)
		}

		if q[0].Value.(bson.M)["$ne"] != "kyte" {
			t.Errorf("Filter.NotEqual should return value map[$ne:kyte], got %v", q[0].Value)
		}
	})

	t.Run("multiple", func(t *testing.T) {
		type Temp struct {
			Name    string `bson:"name"`
			Surname string `bson:"surname"`
			Age     int    `bson:"age"`
		}

		var temp Temp
		q, err := Filter(Source(&temp)).
			NotEqual(&temp.Name, "Joe").
			NotEqual(&temp.Surname, "Doe").
			NotEqual(&temp.Age, 10).
			Build()

		if err != nil {
			t.Errorf("Filter.NotEqual should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.NotEqual should not return nil")
		}

		for _, v := range q {
			if v.Key == "name" {
				if v.Value.(bson.M)["$ne"] != "Joe" {
					t.Errorf("Filter.NotEqual should return value map[$ne:Joe], got %v", v.Value)
				}
			}

			if v.Key == "surname" {
				if v.Value.(bson.M)["$ne"] != "Doe" {
					t.Errorf("Filter.NotEqual should return value map[$ne:Doe], got %v", v.Value)
				}
			}

			if v.Key == "age" {
				if v.Value.(bson.M)["$ne"] != 10 {
					t.Errorf("Filter.NotEqual should return value map[$ne:10], got %v", v.Value)
				}
			}
		}
	})
}

func TestFilter_GreaterThan(t *testing.T) {
	t.Parallel()

	t.Run("without source", func(t *testing.T) {
		q, err := Filter().GreaterThan("age", 10).Build()
		if err != nil {
			t.Errorf("Filter.GreaterThan should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.GreaterThan should not return nil")
		}

		if q[0].Key != "age" {
			t.Errorf("Filter.GreaterThan should return key age, got %v", q[0].Key)
		}

		if q[0].Value.(bson.M)["$gt"] != 10 {
			t.Errorf("Filter.GreaterThan should return value map[$gt:10], got %v", q[0].Value)
		}
	})

	t.Run("with source", func(t *testing.T) {
		type Temp struct {
			Age int `bson:"age"`
		}
		var temp Temp
		q, err := Filter(Source(&temp)).GreaterThan(&temp.Age, 10).Build()
		if err != nil {
			t.Errorf("Filter.GreaterThan should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.GreaterThan should not return nil")
		}

		if q[0].Key != "age" {
			t.Errorf("Filter.GreaterThan should return key age, got %v", q[0].Key)
		}

		if q[0].Value.(bson.M)["$gt"] != 10 {
			t.Errorf("Filter.GreaterThan should return value map[$gt:10], got %v", q[0].Value)
		}
	})

	t.Run("multiple", func(t *testing.T) {
		type Temp struct {
			Name string `bson:"name"`
			Age  int    `bson:"age"`
		}

		var temp Temp
		q, err := Filter(Source(&temp)).
			GreaterThan(&temp.Name, "Joe").
			GreaterThan(&temp.Age, 10).
			Build()

		if err != nil {
			t.Errorf("Filter.GreaterThan should not return error: %v", err)
		}

		if q == nil {
			t.Error("Filter.GreaterThan should not return nil")
		}

		for _, v := range q {
			if v.Key == "name" {
				if v.Value.(bson.M)["$gt"] != "Joe" {
					t.Errorf("Filter.GreaterThan should return value map[$gt:Joe], got %v", v.Value)
				}
			}

			if v.Key == "age" {
				if v.Value.(bson.M)["$gt"] != 10 {
					t.Errorf("Filter.GreaterThan should return value map[$gt:10], got %v", v.Value)
				}
			}
		}
	})
}