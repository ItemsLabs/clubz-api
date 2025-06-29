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

func testAuthUserUserPermissions(t *testing.T) {
	t.Parallel()

	query := AuthUserUserPermissions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthUserUserPermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
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

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserUserPermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AuthUserUserPermissions().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserUserPermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthUserUserPermissionSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserUserPermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthUserUserPermissionExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AuthUserUserPermission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthUserUserPermissionExists to return true, but got false.")
	}
}

func testAuthUserUserPermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authUserUserPermissionFound, err := FindAuthUserUserPermission(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authUserUserPermissionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthUserUserPermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AuthUserUserPermissions().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthUserUserPermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AuthUserUserPermissions().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthUserUserPermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserUserPermissionOne := &AuthUserUserPermission{}
	authUserUserPermissionTwo := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, authUserUserPermissionOne, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserUserPermissionTwo, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = authUserUserPermissionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authUserUserPermissionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserUserPermissions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthUserUserPermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authUserUserPermissionOne := &AuthUserUserPermission{}
	authUserUserPermissionTwo := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, authUserUserPermissionOne, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserUserPermissionTwo, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = authUserUserPermissionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authUserUserPermissionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testAuthUserUserPermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserUserPermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(authUserUserPermissionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserUserPermissionToOneAuthPermissionUsingPermission(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local AuthUserUserPermission
	var foreign AuthPermission

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PermissionID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Permission().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AuthUserUserPermissionSlice{&local}
	if err = local.L.LoadPermission(tx, false, (*[]*AuthUserUserPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Permission == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Permission = nil
	if err = local.L.LoadPermission(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Permission == nil {
		t.Error("struct should have been eager loaded")
	}

}

func testAuthUserUserPermissionToOneAuthUserUsingUser(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local AuthUserUserPermission
	var foreign AuthUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AuthUserUserPermissionSlice{&local}
	if err = local.L.LoadUser(tx, false, (*[]*AuthUserUserPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

}

func testAuthUserUserPermissionToOneSetOpAuthPermissionUsingPermission(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthUserUserPermission
	var b, c AuthPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserUserPermissionDBTypes, false, strmangle.SetComplement(authUserUserPermissionPrimaryKeyColumns, authUserUserPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthPermission{&b, &c} {
		err = a.SetPermission(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Permission != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PermissionAuthUserUserPermissions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PermissionID != x.ID {
			t.Error("foreign key was wrong value", a.PermissionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PermissionID))
		reflect.Indirect(reflect.ValueOf(&a.PermissionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PermissionID != x.ID {
			t.Error("foreign key was wrong value", a.PermissionID, x.ID)
		}
	}
}
func testAuthUserUserPermissionToOneSetOpAuthUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a AuthUserUserPermission
	var b, c AuthUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserUserPermissionDBTypes, false, strmangle.SetComplement(authUserUserPermissionPrimaryKeyColumns, authUserUserPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthUser{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserAuthUserUserPermissions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testAuthUserUserPermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
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

func testAuthUserUserPermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthUserUserPermissionSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserUserPermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserUserPermissions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authUserUserPermissionDBTypes = map[string]string{`ID`: `integer`, `UserID`: `integer`, `PermissionID`: `integer`}
	_                             = bytes.MinRead
)

func testAuthUserUserPermissionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authUserUserPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authUserUserPermissionAllColumns) == len(authUserUserPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthUserUserPermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authUserUserPermissionAllColumns) == len(authUserUserPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthUserUserPermission{}
	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authUserUserPermissionDBTypes, true, authUserUserPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authUserUserPermissionAllColumns, authUserUserPermissionPrimaryKeyColumns) {
		fields = authUserUserPermissionAllColumns
	} else {
		fields = strmangle.SetComplement(
			authUserUserPermissionAllColumns,
			authUserUserPermissionPrimaryKeyColumns,
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

	slice := AuthUserUserPermissionSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthUserUserPermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(authUserUserPermissionAllColumns) == len(authUserUserPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AuthUserUserPermission{}
	if err = randomize.Struct(seed, &o, authUserUserPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthUserUserPermission: %s", err)
	}

	count, err := AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authUserUserPermissionDBTypes, false, authUserUserPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserUserPermission struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthUserUserPermission: %s", err)
	}

	count, err = AuthUserUserPermissions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
