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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CompetitionPhase is an object representing the database table.
type CompetitionPhase struct {
	ID                   string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt            time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt            time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ImportID             null.String `boil:"import_id" json:"import_id,omitempty" toml:"import_id" yaml:"import_id,omitempty"`
	Name                 string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Enabled              bool        `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	CompetitionEditionID string      `boil:"competition_edition_id" json:"competition_edition_id" toml:"competition_edition_id" yaml:"competition_edition_id"`

	R *competitionPhaseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L competitionPhaseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompetitionPhaseColumns = struct {
	ID                   string
	CreatedAt            string
	UpdatedAt            string
	ImportID             string
	Name                 string
	Enabled              string
	CompetitionEditionID string
}{
	ID:                   "id",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	ImportID:             "import_id",
	Name:                 "name",
	Enabled:              "enabled",
	CompetitionEditionID: "competition_edition_id",
}

var CompetitionPhaseTableColumns = struct {
	ID                   string
	CreatedAt            string
	UpdatedAt            string
	ImportID             string
	Name                 string
	Enabled              string
	CompetitionEditionID string
}{
	ID:                   "competition_phases.id",
	CreatedAt:            "competition_phases.created_at",
	UpdatedAt:            "competition_phases.updated_at",
	ImportID:             "competition_phases.import_id",
	Name:                 "competition_phases.name",
	Enabled:              "competition_phases.enabled",
	CompetitionEditionID: "competition_phases.competition_edition_id",
}

// Generated where

var CompetitionPhaseWhere = struct {
	ID                   whereHelperstring
	CreatedAt            whereHelpertime_Time
	UpdatedAt            whereHelpertime_Time
	ImportID             whereHelpernull_String
	Name                 whereHelperstring
	Enabled              whereHelperbool
	CompetitionEditionID whereHelperstring
}{
	ID:                   whereHelperstring{field: "\"competition_phases\".\"id\""},
	CreatedAt:            whereHelpertime_Time{field: "\"competition_phases\".\"created_at\""},
	UpdatedAt:            whereHelpertime_Time{field: "\"competition_phases\".\"updated_at\""},
	ImportID:             whereHelpernull_String{field: "\"competition_phases\".\"import_id\""},
	Name:                 whereHelperstring{field: "\"competition_phases\".\"name\""},
	Enabled:              whereHelperbool{field: "\"competition_phases\".\"enabled\""},
	CompetitionEditionID: whereHelperstring{field: "\"competition_phases\".\"competition_edition_id\""},
}

// CompetitionPhaseRels is where relationship names are stored.
var CompetitionPhaseRels = struct {
	CompetitionEdition string
}{
	CompetitionEdition: "CompetitionEdition",
}

// competitionPhaseR is where relationships are stored.
type competitionPhaseR struct {
	CompetitionEdition *CompetitionEdition `boil:"CompetitionEdition" json:"CompetitionEdition" toml:"CompetitionEdition" yaml:"CompetitionEdition"`
}

// NewStruct creates a new relationship struct
func (*competitionPhaseR) NewStruct() *competitionPhaseR {
	return &competitionPhaseR{}
}

func (r *competitionPhaseR) GetCompetitionEdition() *CompetitionEdition {
	if r == nil {
		return nil
	}
	return r.CompetitionEdition
}

// competitionPhaseL is where Load methods for each relationship are stored.
type competitionPhaseL struct{}

var (
	competitionPhaseAllColumns            = []string{"id", "created_at", "updated_at", "import_id", "name", "enabled", "competition_edition_id"}
	competitionPhaseColumnsWithoutDefault = []string{"id", "created_at", "updated_at", "name", "enabled", "competition_edition_id"}
	competitionPhaseColumnsWithDefault    = []string{"import_id"}
	competitionPhasePrimaryKeyColumns     = []string{"id"}
	competitionPhaseGeneratedColumns      = []string{}
)

type (
	// CompetitionPhaseSlice is an alias for a slice of pointers to CompetitionPhase.
	// This should almost always be used instead of []CompetitionPhase.
	CompetitionPhaseSlice []*CompetitionPhase

	competitionPhaseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	competitionPhaseType                 = reflect.TypeOf(&CompetitionPhase{})
	competitionPhaseMapping              = queries.MakeStructMapping(competitionPhaseType)
	competitionPhasePrimaryKeyMapping, _ = queries.BindMapping(competitionPhaseType, competitionPhaseMapping, competitionPhasePrimaryKeyColumns)
	competitionPhaseInsertCacheMut       sync.RWMutex
	competitionPhaseInsertCache          = make(map[string]insertCache)
	competitionPhaseUpdateCacheMut       sync.RWMutex
	competitionPhaseUpdateCache          = make(map[string]updateCache)
	competitionPhaseUpsertCacheMut       sync.RWMutex
	competitionPhaseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single competitionPhase record from the query.
func (q competitionPhaseQuery) One(exec boil.Executor) (*CompetitionPhase, error) {
	o := &CompetitionPhase{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for competition_phases")
	}

	return o, nil
}

// All returns all CompetitionPhase records from the query.
func (q competitionPhaseQuery) All(exec boil.Executor) (CompetitionPhaseSlice, error) {
	var o []*CompetitionPhase

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to CompetitionPhase slice")
	}

	return o, nil
}

// Count returns the count of all CompetitionPhase records in the query.
func (q competitionPhaseQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count competition_phases rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q competitionPhaseQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if competition_phases exists")
	}

	return count > 0, nil
}

// CompetitionEdition pointed to by the foreign key.
func (o *CompetitionPhase) CompetitionEdition(mods ...qm.QueryMod) competitionEditionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CompetitionEditionID),
	}

	queryMods = append(queryMods, mods...)

	return CompetitionEditions(queryMods...)
}

// LoadCompetitionEdition allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (competitionPhaseL) LoadCompetitionEdition(e boil.Executor, singular bool, maybeCompetitionPhase interface{}, mods queries.Applicator) error {
	var slice []*CompetitionPhase
	var object *CompetitionPhase

	if singular {
		var ok bool
		object, ok = maybeCompetitionPhase.(*CompetitionPhase)
		if !ok {
			object = new(CompetitionPhase)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCompetitionPhase)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCompetitionPhase))
			}
		}
	} else {
		s, ok := maybeCompetitionPhase.(*[]*CompetitionPhase)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCompetitionPhase)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCompetitionPhase))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &competitionPhaseR{}
		}
		args[object.CompetitionEditionID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &competitionPhaseR{}
			}

			args[obj.CompetitionEditionID] = struct{}{}

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
		qm.From(`competition_editions`),
		qm.WhereIn(`competition_editions.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CompetitionEdition")
	}

	var resultSlice []*CompetitionEdition
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CompetitionEdition")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for competition_editions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for competition_editions")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CompetitionEdition = foreign
		if foreign.R == nil {
			foreign.R = &competitionEditionR{}
		}
		foreign.R.CompetitionPhases = append(foreign.R.CompetitionPhases, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CompetitionEditionID == foreign.ID {
				local.R.CompetitionEdition = foreign
				if foreign.R == nil {
					foreign.R = &competitionEditionR{}
				}
				foreign.R.CompetitionPhases = append(foreign.R.CompetitionPhases, local)
				break
			}
		}
	}

	return nil
}

// SetCompetitionEdition of the competitionPhase to the related item.
// Sets o.R.CompetitionEdition to related.
// Adds o to related.R.CompetitionPhases.
func (o *CompetitionPhase) SetCompetitionEdition(exec boil.Executor, insert bool, related *CompetitionEdition) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"competition_phases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"competition_edition_id"}),
		strmangle.WhereClause("\"", "\"", 2, competitionPhasePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CompetitionEditionID = related.ID
	if o.R == nil {
		o.R = &competitionPhaseR{
			CompetitionEdition: related,
		}
	} else {
		o.R.CompetitionEdition = related
	}

	if related.R == nil {
		related.R = &competitionEditionR{
			CompetitionPhases: CompetitionPhaseSlice{o},
		}
	} else {
		related.R.CompetitionPhases = append(related.R.CompetitionPhases, o)
	}

	return nil
}

// CompetitionPhases retrieves all the records using an executor.
func CompetitionPhases(mods ...qm.QueryMod) competitionPhaseQuery {
	mods = append(mods, qm.From("\"competition_phases\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"competition_phases\".*"})
	}

	return competitionPhaseQuery{q}
}

// FindCompetitionPhase retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompetitionPhase(exec boil.Executor, iD string, selectCols ...string) (*CompetitionPhase, error) {
	competitionPhaseObj := &CompetitionPhase{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"competition_phases\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, competitionPhaseObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from competition_phases")
	}

	return competitionPhaseObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompetitionPhase) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no competition_phases provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(competitionPhaseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	competitionPhaseInsertCacheMut.RLock()
	cache, cached := competitionPhaseInsertCache[key]
	competitionPhaseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			competitionPhaseAllColumns,
			competitionPhaseColumnsWithDefault,
			competitionPhaseColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(competitionPhaseType, competitionPhaseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(competitionPhaseType, competitionPhaseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"competition_phases\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"competition_phases\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "schema: unable to insert into competition_phases")
	}

	if !cached {
		competitionPhaseInsertCacheMut.Lock()
		competitionPhaseInsertCache[key] = cache
		competitionPhaseInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CompetitionPhase.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompetitionPhase) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	competitionPhaseUpdateCacheMut.RLock()
	cache, cached := competitionPhaseUpdateCache[key]
	competitionPhaseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			competitionPhaseAllColumns,
			competitionPhasePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update competition_phases, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"competition_phases\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, competitionPhasePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(competitionPhaseType, competitionPhaseMapping, append(wl, competitionPhasePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schema: unable to update competition_phases row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for competition_phases")
	}

	if !cached {
		competitionPhaseUpdateCacheMut.Lock()
		competitionPhaseUpdateCache[key] = cache
		competitionPhaseUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q competitionPhaseQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for competition_phases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for competition_phases")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompetitionPhaseSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitionPhasePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"competition_phases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, competitionPhasePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in competitionPhase slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all competitionPhase")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompetitionPhase) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("schema: no competition_phases provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(competitionPhaseColumnsWithDefault, o)

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

	competitionPhaseUpsertCacheMut.RLock()
	cache, cached := competitionPhaseUpsertCache[key]
	competitionPhaseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			competitionPhaseAllColumns,
			competitionPhaseColumnsWithDefault,
			competitionPhaseColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			competitionPhaseAllColumns,
			competitionPhasePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert competition_phases, could not build update column list")
		}

		ret := strmangle.SetComplement(competitionPhaseAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(competitionPhasePrimaryKeyColumns) == 0 {
				return errors.New("schema: unable to upsert competition_phases, could not build conflict column list")
			}

			conflict = make([]string, len(competitionPhasePrimaryKeyColumns))
			copy(conflict, competitionPhasePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"competition_phases\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(competitionPhaseType, competitionPhaseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(competitionPhaseType, competitionPhaseMapping, ret)
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
		return errors.Wrap(err, "schema: unable to upsert competition_phases")
	}

	if !cached {
		competitionPhaseUpsertCacheMut.Lock()
		competitionPhaseUpsertCache[key] = cache
		competitionPhaseUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CompetitionPhase record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompetitionPhase) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no CompetitionPhase provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), competitionPhasePrimaryKeyMapping)
	sql := "DELETE FROM \"competition_phases\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from competition_phases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for competition_phases")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q competitionPhaseQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no competitionPhaseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from competition_phases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for competition_phases")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompetitionPhaseSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitionPhasePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"competition_phases\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, competitionPhasePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from competitionPhase slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for competition_phases")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CompetitionPhase) Reload(exec boil.Executor) error {
	ret, err := FindCompetitionPhase(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompetitionPhaseSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompetitionPhaseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitionPhasePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"competition_phases\".* FROM \"competition_phases\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, competitionPhasePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in CompetitionPhaseSlice")
	}

	*o = slice

	return nil
}

// CompetitionPhaseExists checks if the CompetitionPhase row exists.
func CompetitionPhaseExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"competition_phases\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if competition_phases exists")
	}

	return exists, nil
}

// Exists checks if the CompetitionPhase row exists.
func (o *CompetitionPhase) Exists(exec boil.Executor) (bool, error) {
	return CompetitionPhaseExists(exec, o.ID)
}
