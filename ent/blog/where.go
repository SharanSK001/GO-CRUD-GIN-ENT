// Code generated by ent, DO NOT EDIT.

package blog

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldName, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldTitle, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldEmail, v))
}

// Phonenumber applies equality check predicate on the "Phonenumber" field. It's identical to PhonenumberEQ.
func Phonenumber(v int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldPhonenumber, v))
}

// Blogcontent applies equality check predicate on the "Blogcontent" field. It's identical to BlogcontentEQ.
func Blogcontent(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldBlogcontent, v))
}

// Githublink applies equality check predicate on the "githublink" field. It's identical to GithublinkEQ.
func Githublink(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldGithublink, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldName, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldTitle, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Blog {
	return predicate.Blog(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Blog {
	return predicate.Blog(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldEmail, v))
}

// PhonenumberEQ applies the EQ predicate on the "Phonenumber" field.
func PhonenumberEQ(v int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldPhonenumber, v))
}

// PhonenumberNEQ applies the NEQ predicate on the "Phonenumber" field.
func PhonenumberNEQ(v int) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldPhonenumber, v))
}

// PhonenumberIn applies the In predicate on the "Phonenumber" field.
func PhonenumberIn(vs ...int) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldPhonenumber, vs...))
}

// PhonenumberNotIn applies the NotIn predicate on the "Phonenumber" field.
func PhonenumberNotIn(vs ...int) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldPhonenumber, vs...))
}

// PhonenumberGT applies the GT predicate on the "Phonenumber" field.
func PhonenumberGT(v int) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldPhonenumber, v))
}

// PhonenumberGTE applies the GTE predicate on the "Phonenumber" field.
func PhonenumberGTE(v int) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldPhonenumber, v))
}

// PhonenumberLT applies the LT predicate on the "Phonenumber" field.
func PhonenumberLT(v int) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldPhonenumber, v))
}

// PhonenumberLTE applies the LTE predicate on the "Phonenumber" field.
func PhonenumberLTE(v int) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldPhonenumber, v))
}

// PhonenumberIsNil applies the IsNil predicate on the "Phonenumber" field.
func PhonenumberIsNil() predicate.Blog {
	return predicate.Blog(sql.FieldIsNull(FieldPhonenumber))
}

// PhonenumberNotNil applies the NotNil predicate on the "Phonenumber" field.
func PhonenumberNotNil() predicate.Blog {
	return predicate.Blog(sql.FieldNotNull(FieldPhonenumber))
}

// BlogcontentEQ applies the EQ predicate on the "Blogcontent" field.
func BlogcontentEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldBlogcontent, v))
}

// BlogcontentNEQ applies the NEQ predicate on the "Blogcontent" field.
func BlogcontentNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldBlogcontent, v))
}

// BlogcontentIn applies the In predicate on the "Blogcontent" field.
func BlogcontentIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldBlogcontent, vs...))
}

// BlogcontentNotIn applies the NotIn predicate on the "Blogcontent" field.
func BlogcontentNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldBlogcontent, vs...))
}

// BlogcontentGT applies the GT predicate on the "Blogcontent" field.
func BlogcontentGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldBlogcontent, v))
}

// BlogcontentGTE applies the GTE predicate on the "Blogcontent" field.
func BlogcontentGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldBlogcontent, v))
}

// BlogcontentLT applies the LT predicate on the "Blogcontent" field.
func BlogcontentLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldBlogcontent, v))
}

// BlogcontentLTE applies the LTE predicate on the "Blogcontent" field.
func BlogcontentLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldBlogcontent, v))
}

// BlogcontentContains applies the Contains predicate on the "Blogcontent" field.
func BlogcontentContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldBlogcontent, v))
}

// BlogcontentHasPrefix applies the HasPrefix predicate on the "Blogcontent" field.
func BlogcontentHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldBlogcontent, v))
}

// BlogcontentHasSuffix applies the HasSuffix predicate on the "Blogcontent" field.
func BlogcontentHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldBlogcontent, v))
}

// BlogcontentEqualFold applies the EqualFold predicate on the "Blogcontent" field.
func BlogcontentEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldBlogcontent, v))
}

// BlogcontentContainsFold applies the ContainsFold predicate on the "Blogcontent" field.
func BlogcontentContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldBlogcontent, v))
}

// GithublinkEQ applies the EQ predicate on the "githublink" field.
func GithublinkEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldGithublink, v))
}

// GithublinkNEQ applies the NEQ predicate on the "githublink" field.
func GithublinkNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldGithublink, v))
}

// GithublinkIn applies the In predicate on the "githublink" field.
func GithublinkIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldGithublink, vs...))
}

// GithublinkNotIn applies the NotIn predicate on the "githublink" field.
func GithublinkNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldGithublink, vs...))
}

// GithublinkGT applies the GT predicate on the "githublink" field.
func GithublinkGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldGithublink, v))
}

// GithublinkGTE applies the GTE predicate on the "githublink" field.
func GithublinkGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldGithublink, v))
}

// GithublinkLT applies the LT predicate on the "githublink" field.
func GithublinkLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldGithublink, v))
}

// GithublinkLTE applies the LTE predicate on the "githublink" field.
func GithublinkLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldGithublink, v))
}

// GithublinkContains applies the Contains predicate on the "githublink" field.
func GithublinkContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldGithublink, v))
}

// GithublinkHasPrefix applies the HasPrefix predicate on the "githublink" field.
func GithublinkHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldGithublink, v))
}

// GithublinkHasSuffix applies the HasSuffix predicate on the "githublink" field.
func GithublinkHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldGithublink, v))
}

// GithublinkIsNil applies the IsNil predicate on the "githublink" field.
func GithublinkIsNil() predicate.Blog {
	return predicate.Blog(sql.FieldIsNull(FieldGithublink))
}

// GithublinkNotNil applies the NotNil predicate on the "githublink" field.
func GithublinkNotNil() predicate.Blog {
	return predicate.Blog(sql.FieldNotNull(FieldGithublink))
}

// GithublinkEqualFold applies the EqualFold predicate on the "githublink" field.
func GithublinkEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldGithublink, v))
}

// GithublinkContainsFold applies the ContainsFold predicate on the "githublink" field.
func GithublinkContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldGithublink, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		p(s.Not())
	})
}
