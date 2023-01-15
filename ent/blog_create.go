// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/blog"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogCreate is the builder for creating a Blog entity.
type BlogCreate struct {
	config
	mutation *BlogMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BlogCreate) SetName(s string) *BlogCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetTitle sets the "title" field.
func (bc *BlogCreate) SetTitle(s string) *BlogCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetEmail sets the "email" field.
func (bc *BlogCreate) SetEmail(s string) *BlogCreate {
	bc.mutation.SetEmail(s)
	return bc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (bc *BlogCreate) SetNillableEmail(s *string) *BlogCreate {
	if s != nil {
		bc.SetEmail(*s)
	}
	return bc
}

// SetPhonenumber sets the "Phonenumber" field.
func (bc *BlogCreate) SetPhonenumber(i int) *BlogCreate {
	bc.mutation.SetPhonenumber(i)
	return bc
}

// SetNillablePhonenumber sets the "Phonenumber" field if the given value is not nil.
func (bc *BlogCreate) SetNillablePhonenumber(i *int) *BlogCreate {
	if i != nil {
		bc.SetPhonenumber(*i)
	}
	return bc
}

// SetBlogcontent sets the "Blogcontent" field.
func (bc *BlogCreate) SetBlogcontent(s string) *BlogCreate {
	bc.mutation.SetBlogcontent(s)
	return bc
}

// SetGithublink sets the "githublink" field.
func (bc *BlogCreate) SetGithublink(s string) *BlogCreate {
	bc.mutation.SetGithublink(s)
	return bc
}

// SetNillableGithublink sets the "githublink" field if the given value is not nil.
func (bc *BlogCreate) SetNillableGithublink(s *string) *BlogCreate {
	if s != nil {
		bc.SetGithublink(*s)
	}
	return bc
}

// Mutation returns the BlogMutation object of the builder.
func (bc *BlogCreate) Mutation() *BlogMutation {
	return bc.mutation
}

// Save creates the Blog in the database.
func (bc *BlogCreate) Save(ctx context.Context) (*Blog, error) {
	return withHooks[*Blog, BlogMutation](ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlogCreate) SaveX(ctx context.Context) *Blog {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlogCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlogCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlogCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Blog.name"`)}
	}
	if v, ok := bc.mutation.Name(); ok {
		if err := blog.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Blog.name": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Blog.title"`)}
	}
	if v, ok := bc.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Blogcontent(); !ok {
		return &ValidationError{Name: "Blogcontent", err: errors.New(`ent: missing required field "Blog.Blogcontent"`)}
	}
	if v, ok := bc.mutation.Blogcontent(); ok {
		if err := blog.BlogcontentValidator(v); err != nil {
			return &ValidationError{Name: "Blogcontent", err: fmt.Errorf(`ent: validator failed for field "Blog.Blogcontent": %w`, err)}
		}
	}
	return nil
}

func (bc *BlogCreate) sqlSave(ctx context.Context) (*Blog, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlogCreate) createSpec() (*Blog, *sqlgraph.CreateSpec) {
	var (
		_node = &Blog{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blog.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Name(); ok {
		_spec.SetField(blog.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Email(); ok {
		_spec.SetField(blog.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := bc.mutation.Phonenumber(); ok {
		_spec.SetField(blog.FieldPhonenumber, field.TypeInt, value)
		_node.Phonenumber = value
	}
	if value, ok := bc.mutation.Blogcontent(); ok {
		_spec.SetField(blog.FieldBlogcontent, field.TypeString, value)
		_node.Blogcontent = value
	}
	if value, ok := bc.mutation.Githublink(); ok {
		_spec.SetField(blog.FieldGithublink, field.TypeString, value)
		_node.Githublink = value
	}
	return _node, _spec
}

// BlogCreateBulk is the builder for creating many Blog entities in bulk.
type BlogCreateBulk struct {
	config
	builders []*BlogCreate
}

// Save creates the Blog entities in the database.
func (bcb *BlogCreateBulk) Save(ctx context.Context) ([]*Blog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blog, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlogCreateBulk) SaveX(ctx context.Context) []*Blog {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlogCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlogCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}