package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("email").Optional(),
		field.Int("Phonenumber").Optional(),
		field.String("Blogcontent").NotEmpty(),
		field.String("githublink").Optional(),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return nil
}
