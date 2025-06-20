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

func testDjangoCeleryBeatClockedschedules(t *testing.T) {
	t.Parallel()

	query := DjangoCeleryBeatClockedschedules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDjangoCeleryBeatClockedschedulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
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

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoCeleryBeatClockedschedulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DjangoCeleryBeatClockedschedules().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoCeleryBeatClockedschedulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoCeleryBeatClockedscheduleSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoCeleryBeatClockedschedulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DjangoCeleryBeatClockedscheduleExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DjangoCeleryBeatClockedschedule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DjangoCeleryBeatClockedscheduleExists to return true, but got false.")
	}
}

func testDjangoCeleryBeatClockedschedulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	djangoCeleryBeatClockedscheduleFound, err := FindDjangoCeleryBeatClockedschedule(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if djangoCeleryBeatClockedscheduleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDjangoCeleryBeatClockedschedulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DjangoCeleryBeatClockedschedules().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testDjangoCeleryBeatClockedschedulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DjangoCeleryBeatClockedschedules().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDjangoCeleryBeatClockedschedulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	djangoCeleryBeatClockedscheduleOne := &DjangoCeleryBeatClockedschedule{}
	djangoCeleryBeatClockedscheduleTwo := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, djangoCeleryBeatClockedscheduleOne, djangoCeleryBeatClockedscheduleDBTypes, false, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoCeleryBeatClockedscheduleTwo, djangoCeleryBeatClockedscheduleDBTypes, false, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = djangoCeleryBeatClockedscheduleOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoCeleryBeatClockedscheduleTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoCeleryBeatClockedschedules().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDjangoCeleryBeatClockedschedulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	djangoCeleryBeatClockedscheduleOne := &DjangoCeleryBeatClockedschedule{}
	djangoCeleryBeatClockedscheduleTwo := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, djangoCeleryBeatClockedscheduleOne, djangoCeleryBeatClockedscheduleDBTypes, false, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoCeleryBeatClockedscheduleTwo, djangoCeleryBeatClockedscheduleDBTypes, false, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = djangoCeleryBeatClockedscheduleOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoCeleryBeatClockedscheduleTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDjangoCeleryBeatClockedschedulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoCeleryBeatClockedschedulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(djangoCeleryBeatClockedscheduleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoCeleryBeatClockedscheduleToManyClockedDjangoCeleryBeatPeriodictasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a DjangoCeleryBeatClockedschedule
	var b, c DjangoCeleryBeatPeriodictask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, djangoCeleryBeatPeriodictaskDBTypes, false, djangoCeleryBeatPeriodictaskColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, djangoCeleryBeatPeriodictaskDBTypes, false, djangoCeleryBeatPeriodictaskColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ClockedID, a.ID)
	queries.Assign(&c.ClockedID, a.ID)
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ClockedDjangoCeleryBeatPeriodictasks().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ClockedID, b.ClockedID) {
			bFound = true
		}
		if queries.Equal(v.ClockedID, c.ClockedID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DjangoCeleryBeatClockedscheduleSlice{&a}
	if err = a.L.LoadClockedDjangoCeleryBeatPeriodictasks(tx, false, (*[]*DjangoCeleryBeatClockedschedule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ClockedDjangoCeleryBeatPeriodictasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ClockedDjangoCeleryBeatPeriodictasks = nil
	if err = a.L.LoadClockedDjangoCeleryBeatPeriodictasks(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ClockedDjangoCeleryBeatPeriodictasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDjangoCeleryBeatClockedscheduleToManyAddOpClockedDjangoCeleryBeatPeriodictasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a DjangoCeleryBeatClockedschedule
	var b, c, d, e DjangoCeleryBeatPeriodictask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoCeleryBeatClockedscheduleDBTypes, false, strmangle.SetComplement(djangoCeleryBeatClockedschedulePrimaryKeyColumns, djangoCeleryBeatClockedscheduleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DjangoCeleryBeatPeriodictask{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, djangoCeleryBeatPeriodictaskDBTypes, false, strmangle.SetComplement(djangoCeleryBeatPeriodictaskPrimaryKeyColumns, djangoCeleryBeatPeriodictaskColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DjangoCeleryBeatPeriodictask{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddClockedDjangoCeleryBeatPeriodictasks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ClockedID) {
			t.Error("foreign key was wrong value", a.ID, first.ClockedID)
		}
		if !queries.Equal(a.ID, second.ClockedID) {
			t.Error("foreign key was wrong value", a.ID, second.ClockedID)
		}

		if first.R.Clocked != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Clocked != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ClockedDjangoCeleryBeatPeriodictasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ClockedDjangoCeleryBeatPeriodictasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ClockedDjangoCeleryBeatPeriodictasks().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDjangoCeleryBeatClockedscheduleToManySetOpClockedDjangoCeleryBeatPeriodictasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a DjangoCeleryBeatClockedschedule
	var b, c, d, e DjangoCeleryBeatPeriodictask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoCeleryBeatClockedscheduleDBTypes, false, strmangle.SetComplement(djangoCeleryBeatClockedschedulePrimaryKeyColumns, djangoCeleryBeatClockedscheduleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DjangoCeleryBeatPeriodictask{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, djangoCeleryBeatPeriodictaskDBTypes, false, strmangle.SetComplement(djangoCeleryBeatPeriodictaskPrimaryKeyColumns, djangoCeleryBeatPeriodictaskColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetClockedDjangoCeleryBeatPeriodictasks(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ClockedDjangoCeleryBeatPeriodictasks().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetClockedDjangoCeleryBeatPeriodictasks(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ClockedDjangoCeleryBeatPeriodictasks().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ClockedID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ClockedID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ClockedID) {
		t.Error("foreign key was wrong value", a.ID, d.ClockedID)
	}
	if !queries.Equal(a.ID, e.ClockedID) {
		t.Error("foreign key was wrong value", a.ID, e.ClockedID)
	}

	if b.R.Clocked != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Clocked != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Clocked != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Clocked != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ClockedDjangoCeleryBeatPeriodictasks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ClockedDjangoCeleryBeatPeriodictasks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDjangoCeleryBeatClockedscheduleToManyRemoveOpClockedDjangoCeleryBeatPeriodictasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a DjangoCeleryBeatClockedschedule
	var b, c, d, e DjangoCeleryBeatPeriodictask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoCeleryBeatClockedscheduleDBTypes, false, strmangle.SetComplement(djangoCeleryBeatClockedschedulePrimaryKeyColumns, djangoCeleryBeatClockedscheduleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DjangoCeleryBeatPeriodictask{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, djangoCeleryBeatPeriodictaskDBTypes, false, strmangle.SetComplement(djangoCeleryBeatPeriodictaskPrimaryKeyColumns, djangoCeleryBeatPeriodictaskColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddClockedDjangoCeleryBeatPeriodictasks(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ClockedDjangoCeleryBeatPeriodictasks().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveClockedDjangoCeleryBeatPeriodictasks(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ClockedDjangoCeleryBeatPeriodictasks().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ClockedID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ClockedID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Clocked != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Clocked != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Clocked != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Clocked != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ClockedDjangoCeleryBeatPeriodictasks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ClockedDjangoCeleryBeatPeriodictasks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ClockedDjangoCeleryBeatPeriodictasks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDjangoCeleryBeatClockedschedulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
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

func testDjangoCeleryBeatClockedschedulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoCeleryBeatClockedscheduleSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testDjangoCeleryBeatClockedschedulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoCeleryBeatClockedschedules().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	djangoCeleryBeatClockedscheduleDBTypes = map[string]string{`ID`: `integer`, `ClockedTime`: `timestamp with time zone`, `Enabled`: `boolean`}
	_                                      = bytes.MinRead
)

func testDjangoCeleryBeatClockedschedulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(djangoCeleryBeatClockedschedulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(djangoCeleryBeatClockedscheduleAllColumns) == len(djangoCeleryBeatClockedschedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedschedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDjangoCeleryBeatClockedschedulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(djangoCeleryBeatClockedscheduleAllColumns) == len(djangoCeleryBeatClockedschedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedscheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoCeleryBeatClockedscheduleDBTypes, true, djangoCeleryBeatClockedschedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(djangoCeleryBeatClockedscheduleAllColumns, djangoCeleryBeatClockedschedulePrimaryKeyColumns) {
		fields = djangoCeleryBeatClockedscheduleAllColumns
	} else {
		fields = strmangle.SetComplement(
			djangoCeleryBeatClockedscheduleAllColumns,
			djangoCeleryBeatClockedschedulePrimaryKeyColumns,
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

	slice := DjangoCeleryBeatClockedscheduleSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDjangoCeleryBeatClockedschedulesUpsert(t *testing.T) {
	t.Parallel()

	if len(djangoCeleryBeatClockedscheduleAllColumns) == len(djangoCeleryBeatClockedschedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DjangoCeleryBeatClockedschedule{}
	if err = randomize.Struct(seed, &o, djangoCeleryBeatClockedscheduleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoCeleryBeatClockedschedule: %s", err)
	}

	count, err := DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, djangoCeleryBeatClockedscheduleDBTypes, false, djangoCeleryBeatClockedschedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoCeleryBeatClockedschedule struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoCeleryBeatClockedschedule: %s", err)
	}

	count, err = DjangoCeleryBeatClockedschedules().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
