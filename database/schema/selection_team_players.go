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

// SelectionTeamPlayer is an object representing the database table.
type SelectionTeamPlayer struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	SelectionID  string      `boil:"selection_id" json:"selection_id" toml:"selection_id" yaml:"selection_id"`
	Position     null.String `boil:"position" json:"position,omitempty" toml:"position" yaml:"position,omitempty"`
	JerseyNumber null.Int    `boil:"jersey_number" json:"jersey_number,omitempty" toml:"jersey_number" yaml:"jersey_number,omitempty"`
	PlayerID     string      `boil:"player_id" json:"player_id" toml:"player_id" yaml:"player_id"`

	R *selectionTeamPlayerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L selectionTeamPlayerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SelectionTeamPlayerColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	SelectionID  string
	Position     string
	JerseyNumber string
	PlayerID     string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	SelectionID:  "selection_id",
	Position:     "position",
	JerseyNumber: "jersey_number",
	PlayerID:     "player_id",
}

var SelectionTeamPlayerTableColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	SelectionID  string
	Position     string
	JerseyNumber string
	PlayerID     string
}{
	ID:           "selection_team_players.id",
	CreatedAt:    "selection_team_players.created_at",
	UpdatedAt:    "selection_team_players.updated_at",
	SelectionID:  "selection_team_players.selection_id",
	Position:     "selection_team_players.position",
	JerseyNumber: "selection_team_players.jersey_number",
	PlayerID:     "selection_team_players.player_id",
}

// Generated where

var SelectionTeamPlayerWhere = struct {
	ID           whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
	SelectionID  whereHelperstring
	Position     whereHelpernull_String
	JerseyNumber whereHelpernull_Int
	PlayerID     whereHelperstring
}{
	ID:           whereHelperstring{field: "\"selection_team_players\".\"id\""},
	CreatedAt:    whereHelpertime_Time{field: "\"selection_team_players\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"selection_team_players\".\"updated_at\""},
	SelectionID:  whereHelperstring{field: "\"selection_team_players\".\"selection_id\""},
	Position:     whereHelpernull_String{field: "\"selection_team_players\".\"position\""},
	JerseyNumber: whereHelpernull_Int{field: "\"selection_team_players\".\"jersey_number\""},
	PlayerID:     whereHelperstring{field: "\"selection_team_players\".\"player_id\""},
}

// SelectionTeamPlayerRels is where relationship names are stored.
var SelectionTeamPlayerRels = struct {
	Player string
}{
	Player: "Player",
}

// selectionTeamPlayerR is where relationships are stored.
type selectionTeamPlayerR struct {
	Player *Player `boil:"Player" json:"Player" toml:"Player" yaml:"Player"`
}

// NewStruct creates a new relationship struct
func (*selectionTeamPlayerR) NewStruct() *selectionTeamPlayerR {
	return &selectionTeamPlayerR{}
}

func (r *selectionTeamPlayerR) GetPlayer() *Player {
	if r == nil {
		return nil
	}
	return r.Player
}

// selectionTeamPlayerL is where Load methods for each relationship are stored.
type selectionTeamPlayerL struct{}

var (
	selectionTeamPlayerAllColumns            = []string{"id", "created_at", "updated_at", "selection_id", "position", "jersey_number", "player_id"}
	selectionTeamPlayerColumnsWithoutDefault = []string{"id", "created_at", "updated_at", "selection_id", "player_id"}
	selectionTeamPlayerColumnsWithDefault    = []string{"position", "jersey_number"}
	selectionTeamPlayerPrimaryKeyColumns     = []string{"id"}
	selectionTeamPlayerGeneratedColumns      = []string{}
)

type (
	// SelectionTeamPlayerSlice is an alias for a slice of pointers to SelectionTeamPlayer.
	// This should almost always be used instead of []SelectionTeamPlayer.
	SelectionTeamPlayerSlice []*SelectionTeamPlayer

	selectionTeamPlayerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	selectionTeamPlayerType                 = reflect.TypeOf(&SelectionTeamPlayer{})
	selectionTeamPlayerMapping              = queries.MakeStructMapping(selectionTeamPlayerType)
	selectionTeamPlayerPrimaryKeyMapping, _ = queries.BindMapping(selectionTeamPlayerType, selectionTeamPlayerMapping, selectionTeamPlayerPrimaryKeyColumns)
	selectionTeamPlayerInsertCacheMut       sync.RWMutex
	selectionTeamPlayerInsertCache          = make(map[string]insertCache)
	selectionTeamPlayerUpdateCacheMut       sync.RWMutex
	selectionTeamPlayerUpdateCache          = make(map[string]updateCache)
	selectionTeamPlayerUpsertCacheMut       sync.RWMutex
	selectionTeamPlayerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single selectionTeamPlayer record from the query.
func (q selectionTeamPlayerQuery) One(exec boil.Executor) (*SelectionTeamPlayer, error) {
	o := &SelectionTeamPlayer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for selection_team_players")
	}

	return o, nil
}

// All returns all SelectionTeamPlayer records from the query.
func (q selectionTeamPlayerQuery) All(exec boil.Executor) (SelectionTeamPlayerSlice, error) {
	var o []*SelectionTeamPlayer

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to SelectionTeamPlayer slice")
	}

	return o, nil
}

// Count returns the count of all SelectionTeamPlayer records in the query.
func (q selectionTeamPlayerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count selection_team_players rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q selectionTeamPlayerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if selection_team_players exists")
	}

	return count > 0, nil
}

// Player pointed to by the foreign key.
func (o *SelectionTeamPlayer) Player(mods ...qm.QueryMod) playerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PlayerID),
	}

	queryMods = append(queryMods, mods...)

	return Players(queryMods...)
}

// LoadPlayer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (selectionTeamPlayerL) LoadPlayer(e boil.Executor, singular bool, maybeSelectionTeamPlayer interface{}, mods queries.Applicator) error {
	var slice []*SelectionTeamPlayer
	var object *SelectionTeamPlayer

	if singular {
		var ok bool
		object, ok = maybeSelectionTeamPlayer.(*SelectionTeamPlayer)
		if !ok {
			object = new(SelectionTeamPlayer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSelectionTeamPlayer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSelectionTeamPlayer))
			}
		}
	} else {
		s, ok := maybeSelectionTeamPlayer.(*[]*SelectionTeamPlayer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSelectionTeamPlayer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSelectionTeamPlayer))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &selectionTeamPlayerR{}
		}
		args[object.PlayerID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &selectionTeamPlayerR{}
			}

			args[obj.PlayerID] = struct{}{}

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
		qm.From(`players`),
		qm.WhereIn(`players.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Player")
	}

	var resultSlice []*Player
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Player")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for players")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for players")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Player = foreign
		if foreign.R == nil {
			foreign.R = &playerR{}
		}
		foreign.R.SelectionTeamPlayers = append(foreign.R.SelectionTeamPlayers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PlayerID == foreign.ID {
				local.R.Player = foreign
				if foreign.R == nil {
					foreign.R = &playerR{}
				}
				foreign.R.SelectionTeamPlayers = append(foreign.R.SelectionTeamPlayers, local)
				break
			}
		}
	}

	return nil
}

// SetPlayer of the selectionTeamPlayer to the related item.
// Sets o.R.Player to related.
// Adds o to related.R.SelectionTeamPlayers.
func (o *SelectionTeamPlayer) SetPlayer(exec boil.Executor, insert bool, related *Player) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"selection_team_players\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"player_id"}),
		strmangle.WhereClause("\"", "\"", 2, selectionTeamPlayerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PlayerID = related.ID
	if o.R == nil {
		o.R = &selectionTeamPlayerR{
			Player: related,
		}
	} else {
		o.R.Player = related
	}

	if related.R == nil {
		related.R = &playerR{
			SelectionTeamPlayers: SelectionTeamPlayerSlice{o},
		}
	} else {
		related.R.SelectionTeamPlayers = append(related.R.SelectionTeamPlayers, o)
	}

	return nil
}

// SelectionTeamPlayers retrieves all the records using an executor.
func SelectionTeamPlayers(mods ...qm.QueryMod) selectionTeamPlayerQuery {
	mods = append(mods, qm.From("\"selection_team_players\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"selection_team_players\".*"})
	}

	return selectionTeamPlayerQuery{q}
}

// FindSelectionTeamPlayer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSelectionTeamPlayer(exec boil.Executor, iD string, selectCols ...string) (*SelectionTeamPlayer, error) {
	selectionTeamPlayerObj := &SelectionTeamPlayer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"selection_team_players\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, selectionTeamPlayerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from selection_team_players")
	}

	return selectionTeamPlayerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SelectionTeamPlayer) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no selection_team_players provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(selectionTeamPlayerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	selectionTeamPlayerInsertCacheMut.RLock()
	cache, cached := selectionTeamPlayerInsertCache[key]
	selectionTeamPlayerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			selectionTeamPlayerAllColumns,
			selectionTeamPlayerColumnsWithDefault,
			selectionTeamPlayerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(selectionTeamPlayerType, selectionTeamPlayerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(selectionTeamPlayerType, selectionTeamPlayerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"selection_team_players\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"selection_team_players\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "schema: unable to insert into selection_team_players")
	}

	if !cached {
		selectionTeamPlayerInsertCacheMut.Lock()
		selectionTeamPlayerInsertCache[key] = cache
		selectionTeamPlayerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the SelectionTeamPlayer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SelectionTeamPlayer) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	selectionTeamPlayerUpdateCacheMut.RLock()
	cache, cached := selectionTeamPlayerUpdateCache[key]
	selectionTeamPlayerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			selectionTeamPlayerAllColumns,
			selectionTeamPlayerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update selection_team_players, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"selection_team_players\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, selectionTeamPlayerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(selectionTeamPlayerType, selectionTeamPlayerMapping, append(wl, selectionTeamPlayerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schema: unable to update selection_team_players row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for selection_team_players")
	}

	if !cached {
		selectionTeamPlayerUpdateCacheMut.Lock()
		selectionTeamPlayerUpdateCache[key] = cache
		selectionTeamPlayerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q selectionTeamPlayerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for selection_team_players")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for selection_team_players")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SelectionTeamPlayerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), selectionTeamPlayerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"selection_team_players\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, selectionTeamPlayerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in selectionTeamPlayer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all selectionTeamPlayer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SelectionTeamPlayer) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("schema: no selection_team_players provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(selectionTeamPlayerColumnsWithDefault, o)

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

	selectionTeamPlayerUpsertCacheMut.RLock()
	cache, cached := selectionTeamPlayerUpsertCache[key]
	selectionTeamPlayerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			selectionTeamPlayerAllColumns,
			selectionTeamPlayerColumnsWithDefault,
			selectionTeamPlayerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			selectionTeamPlayerAllColumns,
			selectionTeamPlayerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert selection_team_players, could not build update column list")
		}

		ret := strmangle.SetComplement(selectionTeamPlayerAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(selectionTeamPlayerPrimaryKeyColumns) == 0 {
				return errors.New("schema: unable to upsert selection_team_players, could not build conflict column list")
			}

			conflict = make([]string, len(selectionTeamPlayerPrimaryKeyColumns))
			copy(conflict, selectionTeamPlayerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"selection_team_players\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(selectionTeamPlayerType, selectionTeamPlayerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(selectionTeamPlayerType, selectionTeamPlayerMapping, ret)
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
		return errors.Wrap(err, "schema: unable to upsert selection_team_players")
	}

	if !cached {
		selectionTeamPlayerUpsertCacheMut.Lock()
		selectionTeamPlayerUpsertCache[key] = cache
		selectionTeamPlayerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single SelectionTeamPlayer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SelectionTeamPlayer) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no SelectionTeamPlayer provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), selectionTeamPlayerPrimaryKeyMapping)
	sql := "DELETE FROM \"selection_team_players\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from selection_team_players")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for selection_team_players")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q selectionTeamPlayerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no selectionTeamPlayerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from selection_team_players")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for selection_team_players")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SelectionTeamPlayerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), selectionTeamPlayerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"selection_team_players\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, selectionTeamPlayerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from selectionTeamPlayer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for selection_team_players")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SelectionTeamPlayer) Reload(exec boil.Executor) error {
	ret, err := FindSelectionTeamPlayer(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SelectionTeamPlayerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SelectionTeamPlayerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), selectionTeamPlayerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"selection_team_players\".* FROM \"selection_team_players\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, selectionTeamPlayerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in SelectionTeamPlayerSlice")
	}

	*o = slice

	return nil
}

// SelectionTeamPlayerExists checks if the SelectionTeamPlayer row exists.
func SelectionTeamPlayerExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"selection_team_players\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if selection_team_players exists")
	}

	return exists, nil
}

// Exists checks if the SelectionTeamPlayer row exists.
func (o *SelectionTeamPlayer) Exists(exec boil.Executor) (bool, error) {
	return SelectionTeamPlayerExists(exec, o.ID)
}
