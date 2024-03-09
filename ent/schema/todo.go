package schema

import (
  "time"
  "entgo.io/ent"
  "entgo.io/ent/schema/field"
  "entgo.io/ent/dialect"
)

type Todo struct {
  ent.Schema
}

func (Todo) Fields() []ent.Field {
  return []ent.Field{
    field.String("title"),
    field.String("tag"),
    field.String("body").
      Optional().
      SchemaType(map[string]string{
        dialect.MySQL: "mediumtext",
      }),
    field.Bool("completed").
      Optional().
      GoType(Status(false)),
    field.Uint8("progress").
      Positive().
      Max(100).
      Default(0),
    field.Time("created_at").
      Default(time.Now),
    field.Time("updated_at").
      Default(time.Now),
  }
}

func (Todo) Edges() []ent.Edge {
  return nil
}

type (
	Int8    int8
	Status  bool
)

