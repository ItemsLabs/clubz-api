// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schema

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAuthPermissions(t *testing.T) {
	t.Parallel()

	query := AuthPermissions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthPermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AuthPermissions().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthPermissionSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthPermissionExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AuthPermission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthPermissionExists to return true, but got false.")
	}
}

func testAuthPermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authPermissionFound, err := FindAuthPermission(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authPermissionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthPermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AuthPermissions().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AuthPermissions().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthPermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermissionOne := &AuthPermission{}
	authPermissionTwo := &AuthPermission{}
	if err = randomize.Struct(seed, authPermissionOne, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authPermissionTwo, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = authPermissionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authPermissionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthPermissions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthPermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authPermissionOne := &AuthPermission{}
	authPermissionTwo := &AuthPermission{}
	if err = randomize.Struct(seed, authPermissionOne, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authPermissionTwo, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = authPermissionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authPermissionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testAuthPermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthPermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(authPermissionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthPermissionToManyPermissionAuthGroupPermissions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c AuthGroupPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authGroupPermissionDBTypes, false, authGroupPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authGroupPermissionDBTypes, false, authGroupPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PermissionID = a.ID
	c.PermissionID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PermissionAuthGroupPermissions().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PermissionID == b.PermissionID {
			bFound = true
		}
		if v.PermissionID == c.PermissionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthPermissionSlice{&a}
	if err = a.L.LoadPermissionAuthGroupPermissions(tx, false, (*[]*AuthPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthGroupPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PermissionAuthGroupPermissions = nil
	if err = a.L.LoadPermissionAuthGroupPermissions(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthGroupPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthPermissionToManyPermissionAuthUserUserPermissions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c AuthUserUserPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PermissionID = a.ID
	c.PermissionID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PermissionAuthUserUserPermissions().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PermissionID == b.PermissionID {
			bFound = true
		}
		if v.PermissionID == c.PermissionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthPermissionSlice{&a}
	if err = a.L.LoadPermissionAuthUserUserPermissions(tx, false, (*[]*AuthPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthUserUserPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PermissionAuthUserUserPermissions = nil
	if err = a.L.LoadPermissionAuthUserUserPermissions(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthUserUserPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthPermissionToManyAddOpPermissionAuthGroupPermissions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c, d, e AuthGroupPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthGroupPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authGroupPermissionDBTypes, false, strmangle.SetComplement(authGroupPermissionPrimaryKeyColumns, authGroupPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthGroupPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPermissionAuthGroupPermissions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PermissionID {
			t.Error("foreign key was wrong value", a.ID, first.PermissionID)
		}
		if a.ID != second.PermissionID {
			t.Error("foreign key was wrong value", a.ID, second.PermissionID)
		}

		if first.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PermissionAuthGroupPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PermissionAuthGroupPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PermissionAuthGroupPermissions().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthPermissionToManyAddOpPermissionAuthUserUserPermissions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c, d, e AuthUserUserPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserUserPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserUserPermissionDBTypes, false, strmangle.SetComplement(authUserUserPermissionPrimaryKeyColumns, authUserUserPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthUserUserPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPermissionAuthUserUserPermissions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PermissionID {
			t.Error("foreign key was wrong value", a.ID, first.PermissionID)
		}
		if a.ID != second.PermissionID {
			t.Error("foreign key was wrong value", a.ID, second.PermissionID)
		}

		if first.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PermissionAuthUserUserPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PermissionAuthUserUserPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PermissionAuthUserUserPermissions().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthPermissionToOneDjangoContentTypeUsingContentType(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local AuthPermission
	var foreign DjangoContentType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ContentTypeID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContentType().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AuthPermissionSlice{&local}
	if err = local.L.LoadContentType(tx, false, (*[]*AuthPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ContentType == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ContentType = nil
	if err = local.L.LoadContentType(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ContentType == nil {
		t.Error("struct should have been eager loaded")
	}

}

func testAuthPermissionToOneSetOpDjangoContentTypeUsingContentType(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c DjangoContentType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DjangoContentType{&b, &c} {
		err = a.SetContentType(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ContentType != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContentTypeAuthPermissions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ContentTypeID != x.ID {
			t.Error("foreign key was wrong value", a.ContentTypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ContentTypeID))
		reflect.Indirect(reflect.ValueOf(&a.ContentTypeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContentTypeID != x.ID {
			t.Error("foreign key was wrong value", a.ContentTypeID, x.ID)
		}
	}
}

func testAuthPermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthPermissionSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthPermissions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authPermissionDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `ContentTypeID`: `integer`, `Codename`: `character varying`}
	_                     = bytes.MinRead
)

func testAuthPermissionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authPermissionAllColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthPermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authPermissionAllColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authPermissionAllColumns, authPermissionPrimaryKeyColumns) {
		fields = authPermissionAllColumns
	} else {
		fields = strmangle.SetComplement(
			authPermissionAllColumns,
			authPermissionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AuthPermissionSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthPermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(authPermissionAllColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AuthPermission{}
	if err = randomize.Struct(seed, &o, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthPermission: %s", err)
	}

	count, err := AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authPermissionDBTypes, false, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthPermission: %s", err)
	}

	count, err = AuthPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
