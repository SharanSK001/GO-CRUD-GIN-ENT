// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/blog"
	"entdemo/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogUpdate is the builder for updating Blog entities.
type BlogUpdate struct {
	config
	hooks    []Hook
	mutation *BlogMutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (bu *BlogUpdate) Where(ps ...predicate.Blog) *BlogUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetName sets the "name" field.
func (bu *BlogUpdate) SetName(s string) *BlogUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BlogUpdate) SetTitle(s string) *BlogUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetEmail sets the "email" field.
func (bu *BlogUpdate) SetEmail(s string) *BlogUpdate {
	bu.mutation.SetEmail(s)
	return bu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (bu *BlogUpdate) SetNillableEmail(s *string) *BlogUpdate {
	if s != nil {
		bu.SetEmail(*s)
	}
	return bu
}

// ClearEmail clears the value of the "email" field.
func (bu *BlogUpdate) ClearEmail() *BlogUpdate {
	bu.mutation.ClearEmail()
	return bu
}

// SetPhonenumber sets the "Phonenumber" field.
func (bu *BlogUpdate) SetPhonenumber(i int) *BlogUpdate {
	bu.mutation.ResetPhonenumber()
	bu.mutation.SetPhonenumber(i)
	return bu
}

// SetNillablePhonenumber sets the "Phonenumber" field if the given value is not nil.
func (bu *BlogUpdate) SetNillablePhonenumber(i *int) *BlogUpdate {
	if i != nil {
		bu.SetPhonenumber(*i)
	}
	return bu
}

// AddPhonenumber adds i to the "Phonenumber" field.
func (bu *BlogUpdate) AddPhonenumber(i int) *BlogUpdate {
	bu.mutation.AddPhonenumber(i)
	return bu
}

// ClearPhonenumber clears the value of the "Phonenumber" field.
func (bu *BlogUpdate) ClearPhonenumber() *BlogUpdate {
	bu.mutation.ClearPhonenumber()
	return bu
}

// SetBlogcontent sets the "Blogcontent" field.
func (bu *BlogUpdate) SetBlogcontent(s string) *BlogUpdate {
	bu.mutation.SetBlogcontent(s)
	return bu
}

// SetGithublink sets the "githublink" field.
func (bu *BlogUpdate) SetGithublink(s string) *BlogUpdate {
	bu.mutation.SetGithublink(s)
	return bu
}

// SetNillableGithublink sets the "githublink" field if the given value is not nil.
func (bu *BlogUpdate) SetNillableGithublink(s *string) *BlogUpdate {
	if s != nil {
		bu.SetGithublink(*s)
	}
	return bu
}

// ClearGithublink clears the value of the "githublink" field.
func (bu *BlogUpdate) ClearGithublink() *BlogUpdate {
	bu.mutation.ClearGithublink()
	return bu
}

// Mutation returns the BlogMutation object of the builder.
func (bu *BlogUpdate) Mutation() *BlogMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BlogMutation](ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlogUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlogUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlogUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlogUpdate) check() error {
	if v, ok := bu.mutation.Name(); ok {
		if err := blog.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Blog.name": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Blogcontent(); ok {
		if err := blog.BlogcontentValidator(v); err != nil {
			return &ValidationError{Name: "Blogcontent", err: fmt.Errorf(`ent: validator failed for field "Blog.Blogcontent": %w`, err)}
		}
	}
	return nil
}

func (bu *BlogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blog.Table,
			Columns: blog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blog.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(blog.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Email(); ok {
		_spec.SetField(blog.FieldEmail, field.TypeString, value)
	}
	if bu.mutation.EmailCleared() {
		_spec.ClearField(blog.FieldEmail, field.TypeString)
	}
	if value, ok := bu.mutation.Phonenumber(); ok {
		_spec.SetField(blog.FieldPhonenumber, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedPhonenumber(); ok {
		_spec.AddField(blog.FieldPhonenumber, field.TypeInt, value)
	}
	if bu.mutation.PhonenumberCleared() {
		_spec.ClearField(blog.FieldPhonenumber, field.TypeInt)
	}
	if value, ok := bu.mutation.Blogcontent(); ok {
		_spec.SetField(blog.FieldBlogcontent, field.TypeString, value)
	}
	if value, ok := bu.mutation.Githublink(); ok {
		_spec.SetField(blog.FieldGithublink, field.TypeString, value)
	}
	if bu.mutation.GithublinkCleared() {
		_spec.ClearField(blog.FieldGithublink, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlogUpdateOne is the builder for updating a single Blog entity.
type BlogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogMutation
}

// SetName sets the "name" field.
func (buo *BlogUpdateOne) SetName(s string) *BlogUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetTitle sets the "title" field.
func (buo *BlogUpdateOne) SetTitle(s string) *BlogUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetEmail sets the "email" field.
func (buo *BlogUpdateOne) SetEmail(s string) *BlogUpdateOne {
	buo.mutation.SetEmail(s)
	return buo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (buo *BlogUpdateOne) SetNillableEmail(s *string) *BlogUpdateOne {
	if s != nil {
		buo.SetEmail(*s)
	}
	return buo
}

// ClearEmail clears the value of the "email" field.
func (buo *BlogUpdateOne) ClearEmail() *BlogUpdateOne {
	buo.mutation.ClearEmail()
	return buo
}

// SetPhonenumber sets the "Phonenumber" field.
func (buo *BlogUpdateOne) SetPhonenumber(i int) *BlogUpdateOne {
	buo.mutation.ResetPhonenumber()
	buo.mutation.SetPhonenumber(i)
	return buo
}

// SetNillablePhonenumber sets the "Phonenumber" field if the given value is not nil.
func (buo *BlogUpdateOne) SetNillablePhonenumber(i *int) *BlogUpdateOne {
	if i != nil {
		buo.SetPhonenumber(*i)
	}
	return buo
}

// AddPhonenumber adds i to the "Phonenumber" field.
func (buo *BlogUpdateOne) AddPhonenumber(i int) *BlogUpdateOne {
	buo.mutation.AddPhonenumber(i)
	return buo
}

// ClearPhonenumber clears the value of the "Phonenumber" field.
func (buo *BlogUpdateOne) ClearPhonenumber() *BlogUpdateOne {
	buo.mutation.ClearPhonenumber()
	return buo
}

// SetBlogcontent sets the "Blogcontent" field.
func (buo *BlogUpdateOne) SetBlogcontent(s string) *BlogUpdateOne {
	buo.mutation.SetBlogcontent(s)
	return buo
}

// SetGithublink sets the "githublink" field.
func (buo *BlogUpdateOne) SetGithublink(s string) *BlogUpdateOne {
	buo.mutation.SetGithublink(s)
	return buo
}

// SetNillableGithublink sets the "githublink" field if the given value is not nil.
func (buo *BlogUpdateOne) SetNillableGithublink(s *string) *BlogUpdateOne {
	if s != nil {
		buo.SetGithublink(*s)
	}
	return buo
}

// ClearGithublink clears the value of the "githublink" field.
func (buo *BlogUpdateOne) ClearGithublink() *BlogUpdateOne {
	buo.mutation.ClearGithublink()
	return buo
}

// Mutation returns the BlogMutation object of the builder.
func (buo *BlogUpdateOne) Mutation() *BlogMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlogUpdateOne) Select(field string, fields ...string) *BlogUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blog entity.
func (buo *BlogUpdateOne) Save(ctx context.Context) (*Blog, error) {
	return withHooks[*Blog, BlogMutation](ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlogUpdateOne) SaveX(ctx context.Context) *Blog {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlogUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlogUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlogUpdateOne) check() error {
	if v, ok := buo.mutation.Name(); ok {
		if err := blog.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Blog.name": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Blogcontent(); ok {
		if err := blog.BlogcontentValidator(v); err != nil {
			return &ValidationError{Name: "Blogcontent", err: fmt.Errorf(`ent: validator failed for field "Blog.Blogcontent": %w`, err)}
		}
	}
	return nil
}

func (buo *BlogUpdateOne) sqlSave(ctx context.Context) (_node *Blog, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blog.Table,
			Columns: blog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blog.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blog.FieldID)
		for _, f := range fields {
			if !blog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(blog.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Email(); ok {
		_spec.SetField(blog.FieldEmail, field.TypeString, value)
	}
	if buo.mutation.EmailCleared() {
		_spec.ClearField(blog.FieldEmail, field.TypeString)
	}
	if value, ok := buo.mutation.Phonenumber(); ok {
		_spec.SetField(blog.FieldPhonenumber, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedPhonenumber(); ok {
		_spec.AddField(blog.FieldPhonenumber, field.TypeInt, value)
	}
	if buo.mutation.PhonenumberCleared() {
		_spec.ClearField(blog.FieldPhonenumber, field.TypeInt)
	}
	if value, ok := buo.mutation.Blogcontent(); ok {
		_spec.SetField(blog.FieldBlogcontent, field.TypeString, value)
	}
	if value, ok := buo.mutation.Githublink(); ok {
		_spec.SetField(blog.FieldGithublink, field.TypeString, value)
	}
	if buo.mutation.GithublinkCleared() {
		_spec.ClearField(blog.FieldGithublink, field.TypeString)
	}
	_node = &Blog{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
