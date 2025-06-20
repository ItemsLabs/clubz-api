// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schema

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PowerupAction is an object representing the database table.
type PowerupAction struct {
	ID        int `boil:"id" json:"id" toml:"id" yaml:"id"`
	Ordering  int `boil:"ordering" json:"ordering" toml:"ordering" yaml:"ordering"`
	ActionID  int `boil:"action_id" json:"action_id" toml:"action_id" yaml:"action_id"`
	PowerupID int `boil:"powerup_id" json:"powerup_id" toml:"powerup_id" yaml:"powerup_id"`

	R *powerupActionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L powerupActionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PowerupActionColumns = struct {
	ID        string
	Ordering  string
	ActionID  string
	PowerupID string
}{
	ID:        "id",
	Ordering:  "ordering",
	ActionID:  "action_id",
	PowerupID: "powerup_id",
}

var PowerupActionTableColumns = struct {
	ID        string
	Ordering  string
	ActionID  string
	PowerupID string
}{
	ID:        "powerup_actions.id",
	Ordering:  "powerup_actions.ordering",
	ActionID:  "powerup_actions.action_id",
	PowerupID: "powerup_actions.powerup_id",
}

// Generated where

var PowerupActionWhere = struct {
	ID        whereHelperint
	Ordering  whereHelperint
	ActionID  whereHelperint
	PowerupID whereHelperint
}{
	ID:        whereHelperint{field: "\"powerup_actions\".\"id\""},
	Ordering:  whereHelperint{field: "\"powerup_actions\".\"ordering\""},
	ActionID:  whereHelperint{field: "\"powerup_actions\".\"action_id\""},
	PowerupID: whereHelperint{field: "\"powerup_actions\".\"powerup_id\""},
}

// PowerupActionRels is where relationship names are stored.
var PowerupActionRels = struct {
	Action  string
	Powerup string
}{
	Action:  "Action",
	Powerup: "Powerup",
}

// powerupActionR is where relationships are stored.
type powerupActionR struct {
	Action  *Action  `boil:"Action" json:"Action" toml:"Action" yaml:"Action"`
	Powerup *Powerup `boil:"Powerup" json:"Powerup" toml:"Powerup" yaml:"Powerup"`
}

// NewStruct creates a new relationship struct
func (*powerupActionR) NewStruct() *powerupActionR {
	return &powerupActionR{}
}

func (r *powerupActionR) GetAction() *Action {
	if r == nil {
		return nil
	}
	return r.Action
}

func (r *powerupActionR) GetPowerup() *Powerup {
	if r == nil {
		return nil
	}
	return r.Powerup
}

// powerupActionL is where Load methods for each relationship are stored.
type powerupActionL struct{}

var (
	powerupActionAllColumns            = []string{"id", "ordering", "action_id", "powerup_id"}
	powerupActionColumnsWithoutDefault = []string{"action_id", "powerup_id"}
	powerupActionColumnsWithDefault    = []string{"id", "ordering"}
	powerupActionPrimaryKeyColumns     = []string{"id"}
	powerupActionGeneratedColumns      = []string{}
)

type (
	// PowerupActionSlice is an alias for a slice of pointers to PowerupAction.
	// This should almost always be used instead of []PowerupAction.
	PowerupActionSlice []*PowerupAction

	powerupActionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	powerupActionType                 = reflect.TypeOf(&PowerupAction{})
	powerupActionMapping              = queries.MakeStructMapping(powerupActionType)
	powerupActionPrimaryKeyMapping, _ = queries.BindMapping(powerupActionType, powerupActionMapping, powerupActionPrimaryKeyColumns)
	powerupActionInsertCacheMut       sync.RWMutex
	powerupActionInsertCache          = make(map[string]insertCache)
	powerupActionUpdateCacheMut       sync.RWMutex
	powerupActionUpdateCache          = make(map[string]updateCache)
	powerupActionUpsertCacheMut       sync.RWMutex
	powerupActionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single powerupAction record from the query.
func (q powerupActionQuery) One(exec boil.Executor) (*PowerupAction, error) {
	o := &PowerupAction{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for powerup_actions")
	}

	return o, nil
}

// All returns all PowerupAction records from the query.
func (q powerupActionQuery) All(exec boil.Executor) (PowerupActionSlice, error) {
	var o []*PowerupAction

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to PowerupAction slice")
	}

	return o, nil
}

// Count returns the count of all PowerupAction records in the query.
func (q powerupActionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count powerup_actions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q powerupActionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if powerup_actions exists")
	}

	return count > 0, nil
}

// Action pointed to by the foreign key.
func (o *PowerupAction) Action(mods ...qm.QueryMod) actionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ActionID),
	}

	queryMods = append(queryMods, mods...)

	return Actions(queryMods...)
}

// Powerup pointed to by the foreign key.
func (o *PowerupAction) Powerup(mods ...qm.QueryMod) powerupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PowerupID),
	}

	queryMods = append(queryMods, mods...)

	return Powerups(queryMods...)
}

// LoadAction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (powerupActionL) LoadAction(e boil.Executor, singular bool, maybePowerupAction interface{}, mods queries.Applicator) error {
	var slice []*PowerupAction
	var object *PowerupAction

	if singular {
		var ok bool
		object, ok = maybePowerupAction.(*PowerupAction)
		if !ok {
			object = new(PowerupAction)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePowerupAction)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePowerupAction))
			}
		}
	} else {
		s, ok := maybePowerupAction.(*[]*PowerupAction)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePowerupAction)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePowerupAction))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &powerupActionR{}
		}
		args[object.ActionID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &powerupActionR{}
			}

			args[obj.ActionID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`actions`),
		qm.WhereIn(`actions.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Action")
	}

	var resultSlice []*Action
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Action")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for actions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for actions")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Action = foreign
		if foreign.R == nil {
			foreign.R = &actionR{}
		}
		foreign.R.PowerupActions = append(foreign.R.PowerupActions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ActionID == foreign.ID {
				local.R.Action = foreign
				if foreign.R == nil {
					foreign.R = &actionR{}
				}
				foreign.R.PowerupActions = append(foreign.R.PowerupActions, local)
				break
			}
		}
	}

	return nil
}

// LoadPowerup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (powerupActionL) LoadPowerup(e boil.Executor, singular bool, maybePowerupAction interface{}, mods queries.Applicator) error {
	var slice []*PowerupAction
	var object *PowerupAction

	if singular {
		var ok bool
		object, ok = maybePowerupAction.(*PowerupAction)
		if !ok {
			object = new(PowerupAction)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePowerupAction)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePowerupAction))
			}
		}
	} else {
		s, ok := maybePowerupAction.(*[]*PowerupAction)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePowerupAction)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePowerupAction))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &powerupActionR{}
		}
		args[object.PowerupID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &powerupActionR{}
			}

			args[obj.PowerupID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`powerups`),
		qm.WhereIn(`powerups.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Powerup")
	}

	var resultSlice []*Powerup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Powerup")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for powerups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for powerups")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Powerup = foreign
		if foreign.R == nil {
			foreign.R = &powerupR{}
		}
		foreign.R.PowerupActions = append(foreign.R.PowerupActions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PowerupID == foreign.ID {
				local.R.Powerup = foreign
				if foreign.R == nil {
					foreign.R = &powerupR{}
				}
				foreign.R.PowerupActions = append(foreign.R.PowerupActions, local)
				break
			}
		}
	}

	return nil
}

// SetAction of the powerupAction to the related item.
// Sets o.R.Action to related.
// Adds o to related.R.PowerupActions.
func (o *PowerupAction) SetAction(exec boil.Executor, insert bool, related *Action) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"powerup_actions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"action_id"}),
		strmangle.WhereClause("\"", "\"", 2, powerupActionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ActionID = related.ID
	if o.R == nil {
		o.R = &powerupActionR{
			Action: related,
		}
	} else {
		o.R.Action = related
	}

	if related.R == nil {
		related.R = &actionR{
			PowerupActions: PowerupActionSlice{o},
		}
	} else {
		related.R.PowerupActions = append(related.R.PowerupActions, o)
	}

	return nil
}

// SetPowerup of the powerupAction to the related item.
// Sets o.R.Powerup to related.
// Adds o to related.R.PowerupActions.
func (o *PowerupAction) SetPowerup(exec boil.Executor, insert bool, related *Powerup) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"powerup_actions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"powerup_id"}),
		strmangle.WhereClause("\"", "\"", 2, powerupActionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PowerupID = related.ID
	if o.R == nil {
		o.R = &powerupActionR{
			Powerup: related,
		}
	} else {
		o.R.Powerup = related
	}

	if related.R == nil {
		related.R = &powerupR{
			PowerupActions: PowerupActionSlice{o},
		}
	} else {
		related.R.PowerupActions = append(related.R.PowerupActions, o)
	}

	return nil
}

// PowerupActions retrieves all the records using an executor.
func PowerupActions(mods ...qm.QueryMod) powerupActionQuery {
	mods = append(mods, qm.From("\"powerup_actions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"powerup_actions\".*"})
	}

	return powerupActionQuery{q}
}

// FindPowerupAction retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPowerupAction(exec boil.Executor, iD int, selectCols ...string) (*PowerupAction, error) {
	powerupActionObj := &PowerupAction{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"powerup_actions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, powerupActionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from powerup_actions")
	}

	return powerupActionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PowerupAction) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no powerup_actions provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(powerupActionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	powerupActionInsertCacheMut.RLock()
	cache, cached := powerupActionInsertCache[key]
	powerupActionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			powerupActionAllColumns,
			powerupActionColumnsWithDefault,
			powerupActionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(powerupActionType, powerupActionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(powerupActionType, powerupActionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"powerup_actions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"powerup_actions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "schema: unable to insert into powerup_actions")
	}

	if !cached {
		powerupActionInsertCacheMut.Lock()
		powerupActionInsertCache[key] = cache
		powerupActionInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PowerupAction.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PowerupAction) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	powerupActionUpdateCacheMut.RLock()
	cache, cached := powerupActionUpdateCache[key]
	powerupActionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			powerupActionAllColumns,
			powerupActionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update powerup_actions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"powerup_actions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, powerupActionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(powerupActionType, powerupActionMapping, append(wl, powerupActionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update powerup_actions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for powerup_actions")
	}

	if !cached {
		powerupActionUpdateCacheMut.Lock()
		powerupActionUpdateCache[key] = cache
		powerupActionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q powerupActionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for powerup_actions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for powerup_actions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PowerupActionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("schema: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powerupActionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"powerup_actions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, powerupActionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in powerupAction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all powerupAction")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PowerupAction) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("schema: no powerup_actions provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(powerupActionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	powerupActionUpsertCacheMut.RLock()
	cache, cached := powerupActionUpsertCache[key]
	powerupActionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			powerupActionAllColumns,
			powerupActionColumnsWithDefault,
			powerupActionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			powerupActionAllColumns,
			powerupActionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert powerup_actions, could not build update column list")
		}

		ret := strmangle.SetComplement(powerupActionAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(powerupActionPrimaryKeyColumns) == 0 {
				return errors.New("schema: unable to upsert powerup_actions, could not build conflict column list")
			}

			conflict = make([]string, len(powerupActionPrimaryKeyColumns))
			copy(conflict, powerupActionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"powerup_actions\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(powerupActionType, powerupActionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(powerupActionType, powerupActionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "schema: unable to upsert powerup_actions")
	}

	if !cached {
		powerupActionUpsertCacheMut.Lock()
		powerupActionUpsertCache[key] = cache
		powerupActionUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PowerupAction record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PowerupAction) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no PowerupAction provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), powerupActionPrimaryKeyMapping)
	sql := "DELETE FROM \"powerup_actions\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from powerup_actions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for powerup_actions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q powerupActionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no powerupActionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from powerup_actions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for powerup_actions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PowerupActionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powerupActionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"powerup_actions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, powerupActionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from powerupAction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for powerup_actions")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PowerupAction) Reload(exec boil.Executor) error {
	ret, err := FindPowerupAction(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PowerupActionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PowerupActionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powerupActionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"powerup_actions\".* FROM \"powerup_actions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, powerupActionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in PowerupActionSlice")
	}

	*o = slice

	return nil
}

// PowerupActionExists checks if the PowerupAction row exists.
func PowerupActionExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"powerup_actions\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if powerup_actions exists")
	}

	return exists, nil
}

// Exists checks if the PowerupAction row exists.
func (o *PowerupAction) Exists(exec boil.Executor) (bool, error) {
	return PowerupActionExists(exec, o.ID)
}
