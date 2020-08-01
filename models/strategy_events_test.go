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

func testStrategyEvents(t *testing.T) {
	t.Parallel()

	query := StrategyEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStrategyEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
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

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStrategyEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StrategyEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStrategyEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StrategyEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStrategyEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StrategyEventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if StrategyEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StrategyEventExists to return true, but got false.")
	}
}

func testStrategyEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	strategyEventFound, err := FindStrategyEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if strategyEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStrategyEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StrategyEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStrategyEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StrategyEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStrategyEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	strategyEventOne := &StrategyEvent{}
	strategyEventTwo := &StrategyEvent{}
	if err = randomize.Struct(seed, strategyEventOne, strategyEventDBTypes, false, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, strategyEventTwo, strategyEventDBTypes, false, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = strategyEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = strategyEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StrategyEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStrategyEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	strategyEventOne := &StrategyEvent{}
	strategyEventTwo := &StrategyEvent{}
	if err = randomize.Struct(seed, strategyEventOne, strategyEventDBTypes, false, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, strategyEventTwo, strategyEventDBTypes, false, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = strategyEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = strategyEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func strategyEventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func strategyEventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StrategyEvent) error {
	*o = StrategyEvent{}
	return nil
}

func testStrategyEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StrategyEvent{}
	o := &StrategyEvent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, strategyEventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StrategyEvent object: %s", err)
	}

	AddStrategyEventHook(boil.BeforeInsertHook, strategyEventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	strategyEventBeforeInsertHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.AfterInsertHook, strategyEventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	strategyEventAfterInsertHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.AfterSelectHook, strategyEventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	strategyEventAfterSelectHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.BeforeUpdateHook, strategyEventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	strategyEventBeforeUpdateHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.AfterUpdateHook, strategyEventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	strategyEventAfterUpdateHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.BeforeDeleteHook, strategyEventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	strategyEventBeforeDeleteHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.AfterDeleteHook, strategyEventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	strategyEventAfterDeleteHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.BeforeUpsertHook, strategyEventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	strategyEventBeforeUpsertHooks = []StrategyEventHook{}

	AddStrategyEventHook(boil.AfterUpsertHook, strategyEventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	strategyEventAfterUpsertHooks = []StrategyEventHook{}
}

func testStrategyEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStrategyEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(strategyEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStrategyEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
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

func testStrategyEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StrategyEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStrategyEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StrategyEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	strategyEventDBTypes = map[string]string{`ID`: `character varying`, `StrategyID`: `character varying`, `StrategyAction`: `character varying`, `Amount`: `integer`, `EventOn`: `timestamp without time zone`, `CreatedOn`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testStrategyEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(strategyEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(strategyEventAllColumns) == len(strategyEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStrategyEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(strategyEventAllColumns) == len(strategyEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StrategyEvent{}
	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, strategyEventDBTypes, true, strategyEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(strategyEventAllColumns, strategyEventPrimaryKeyColumns) {
		fields = strategyEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			strategyEventAllColumns,
			strategyEventPrimaryKeyColumns,
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

	slice := StrategyEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStrategyEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(strategyEventAllColumns) == len(strategyEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StrategyEvent{}
	if err = randomize.Struct(seed, &o, strategyEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StrategyEvent: %s", err)
	}

	count, err := StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, strategyEventDBTypes, false, strategyEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StrategyEvent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StrategyEvent: %s", err)
	}

	count, err = StrategyEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}