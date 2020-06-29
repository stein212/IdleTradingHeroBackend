// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMacdStrategies(t *testing.T) {
	t.Parallel()

	query := MacdStrategies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMacdStrategiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMacdStrategiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MacdStrategies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMacdStrategiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MacdStrategySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMacdStrategiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MacdStrategyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MacdStrategy exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MacdStrategyExists to return true, but got false.")
	}
}

func testMacdStrategiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	macdStrategyFound, err := FindMacdStrategy(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if macdStrategyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMacdStrategiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MacdStrategies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMacdStrategiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MacdStrategies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMacdStrategiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	macdStrategyOne := &MacdStrategy{}
	macdStrategyTwo := &MacdStrategy{}
	if err = randomize.Struct(seed, macdStrategyOne, macdStrategyDBTypes, false, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}
	if err = randomize.Struct(seed, macdStrategyTwo, macdStrategyDBTypes, false, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = macdStrategyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = macdStrategyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MacdStrategies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMacdStrategiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	macdStrategyOne := &MacdStrategy{}
	macdStrategyTwo := &MacdStrategy{}
	if err = randomize.Struct(seed, macdStrategyOne, macdStrategyDBTypes, false, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}
	if err = randomize.Struct(seed, macdStrategyTwo, macdStrategyDBTypes, false, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = macdStrategyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = macdStrategyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func macdStrategyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func macdStrategyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MacdStrategy) error {
	*o = MacdStrategy{}
	return nil
}

func testMacdStrategiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MacdStrategy{}
	o := &MacdStrategy{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MacdStrategy object: %s", err)
	}

	AddMacdStrategyHook(boil.BeforeInsertHook, macdStrategyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	macdStrategyBeforeInsertHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.AfterInsertHook, macdStrategyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	macdStrategyAfterInsertHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.AfterSelectHook, macdStrategyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	macdStrategyAfterSelectHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.BeforeUpdateHook, macdStrategyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	macdStrategyBeforeUpdateHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.AfterUpdateHook, macdStrategyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	macdStrategyAfterUpdateHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.BeforeDeleteHook, macdStrategyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	macdStrategyBeforeDeleteHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.AfterDeleteHook, macdStrategyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	macdStrategyAfterDeleteHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.BeforeUpsertHook, macdStrategyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	macdStrategyBeforeUpsertHooks = []MacdStrategyHook{}

	AddMacdStrategyHook(boil.AfterUpsertHook, macdStrategyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	macdStrategyAfterUpsertHooks = []MacdStrategyHook{}
}

func testMacdStrategiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMacdStrategiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(macdStrategyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMacdStrategyToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MacdStrategy
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, macdStrategyDBTypes, false, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MacdStrategySlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*MacdStrategy)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMacdStrategyToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MacdStrategy
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, macdStrategyDBTypes, false, strmangle.SetComplement(macdStrategyPrimaryKeyColumns, macdStrategyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MacdStrategies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testMacdStrategiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMacdStrategiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MacdStrategySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMacdStrategiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MacdStrategies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	macdStrategyDBTypes = map[string]string{`ID`: `character varying`, `UserID`: `character varying`, `Name`: `character varying`, `Instrument`: `character varying`, `Granularity`: `character varying`, `Ema26`: `integer`, `Ema12`: `integer`, `Ema9`: `integer`, `Status`: `character varying`, `CreatedOn`: `timestamp without time zone`, `LastEditedOn`: `timestamp without time zone`}
	_                   = bytes.MinRead
)

func testMacdStrategiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(macdStrategyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(macdStrategyAllColumns) == len(macdStrategyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMacdStrategiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(macdStrategyAllColumns) == len(macdStrategyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MacdStrategy{}
	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, macdStrategyDBTypes, true, macdStrategyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(macdStrategyAllColumns, macdStrategyPrimaryKeyColumns) {
		fields = macdStrategyAllColumns
	} else {
		fields = strmangle.SetComplement(
			macdStrategyAllColumns,
			macdStrategyPrimaryKeyColumns,
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

	slice := MacdStrategySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMacdStrategiesUpsert(t *testing.T) {
	t.Parallel()

	if len(macdStrategyAllColumns) == len(macdStrategyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MacdStrategy{}
	if err = randomize.Struct(seed, &o, macdStrategyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MacdStrategy: %s", err)
	}

	count, err := MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, macdStrategyDBTypes, false, macdStrategyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MacdStrategy struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MacdStrategy: %s", err)
	}

	count, err = MacdStrategies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
