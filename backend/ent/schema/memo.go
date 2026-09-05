package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Memo holds the schema definition for the Memo entity.
type Memo struct {
	ent.Schema
}

func (Memo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "memos"},
	}
}

func (Memo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").
			Comment("用户ID"),
		field.String("title").
			MaxLen(200).
			NotEmpty().
			Comment("备忘录标题"),
		field.String("content").
			SchemaType(map[string]string{dialect.Postgres: "text"}).
			NotEmpty().
			Comment("备忘录内容（支持 Markdown）"),
		field.Bool("pinned").
			Default(false).
			Comment("是否置顶"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			SchemaType(map[string]string{dialect.Postgres: "timestamptz"}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{dialect.Postgres: "timestamptz"}),
	}
}

func (Memo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("pinned", "updated_at"),
	}
}
