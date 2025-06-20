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

// MatchReward is an object representing the database table.
type MatchReward struct {
	ID           int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Position     null.Int `boil:"position" json:"position,omitempty" toml:"position" yaml:"position,omitempty"`
	Amount       float64  `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	MatchID      string   `boil:"match_id" json:"match_id" toml:"match_id" yaml:"match_id"`
	MaxPosition  null.Int `boil:"max_position" json:"max_position,omitempty" toml:"max_position" yaml:"max_position,omitempty"`
	MinPosition  int      `boil:"min_position" json:"min_position" toml:"min_position" yaml:"min_position"`
	Balls        int      `boil:"balls" json:"balls" toml:"balls" yaml:"balls"`
	Event        int      `boil:"event" json:"event" toml:"event" yaml:"event"`
	Game         float64  `boil:"game" json:"game" toml:"game" yaml:"game"`
	Lapt         float64  `boil:"lapt" json:"lapt" toml:"lapt" yaml:"lapt"`
	Shirts       int      `boil:"shirts" json:"shirts" toml:"shirts" yaml:"shirts"`
	SignedBalls  int      `boil:"signed_balls" json:"signed_balls" toml:"signed_balls" yaml:"signed_balls"`
	SignedShirts int      `boil:"signed_shirts" json:"signed_shirts" toml:"signed_shirts" yaml:"signed_shirts"`
	KickoffPack1 int      `boil:"kickoff_pack_1" json:"kickoff_pack_1" toml:"kickoff_pack_1" yaml:"kickoff_pack_1"`
	KickoffPack2 int      `boil:"kickoff_pack_2" json:"kickoff_pack_2" toml:"kickoff_pack_2" yaml:"kickoff_pack_2"`
	KickoffPack3 int      `boil:"kickoff_pack_3" json:"kickoff_pack_3" toml:"kickoff_pack_3" yaml:"kickoff_pack_3"`
	SeasonPack1  int      `boil:"season_pack_1" json:"season_pack_1" toml:"season_pack_1" yaml:"season_pack_1"`
	SeasonPack2  int      `boil:"season_pack_2" json:"season_pack_2" toml:"season_pack_2" yaml:"season_pack_2"`
	SeasonPack3  int      `boil:"season_pack_3" json:"season_pack_3" toml:"season_pack_3" yaml:"season_pack_3"`

	R *matchRewardR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L matchRewardL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MatchRewardColumns = struct {
	ID           string
	Position     string
	Amount       string
	MatchID      string
	MaxPosition  string
	MinPosition  string
	Balls        string
	Event        string
	Game         string
	Lapt         string
	Shirts       string
	SignedBalls  string
	SignedShirts string
	KickoffPack1 string
	KickoffPack2 string
	KickoffPack3 string
	SeasonPack1  string
	SeasonPack2  string
	SeasonPack3  string
}{
	ID:           "id",
	Position:     "position",
	Amount:       "amount",
	MatchID:      "match_id",
	MaxPosition:  "max_position",
	MinPosition:  "min_position",
	Balls:        "balls",
	Event:        "event",
	Game:         "game",
	Lapt:         "lapt",
	Shirts:       "shirts",
	SignedBalls:  "signed_balls",
	SignedShirts: "signed_shirts",
	KickoffPack1: "kickoff_pack_1",
	KickoffPack2: "kickoff_pack_2",
	KickoffPack3: "kickoff_pack_3",
	SeasonPack1:  "season_pack_1",
	SeasonPack2:  "season_pack_2",
	SeasonPack3:  "season_pack_3",
}

var MatchRewardTableColumns = struct {
	ID           string
	Position     string
	Amount       string
	MatchID      string
	MaxPosition  string
	MinPosition  string
	Balls        string
	Event        string
	Game         string
	Lapt         string
	Shirts       string
	SignedBalls  string
	SignedShirts string
	KickoffPack1 string
	KickoffPack2 string
	KickoffPack3 string
	SeasonPack1  string
	SeasonPack2  string
	SeasonPack3  string
}{
	ID:           "match_rewards.id",
	Position:     "match_rewards.position",
	Amount:       "match_rewards.amount",
	MatchID:      "match_rewards.match_id",
	MaxPosition:  "match_rewards.max_position",
	MinPosition:  "match_rewards.min_position",
	Balls:        "match_rewards.balls",
	Event:        "match_rewards.event",
	Game:         "match_rewards.game",
	Lapt:         "match_rewards.lapt",
	Shirts:       "match_rewards.shirts",
	SignedBalls:  "match_rewards.signed_balls",
	SignedShirts: "match_rewards.signed_shirts",
	KickoffPack1: "match_rewards.kickoff_pack_1",
	KickoffPack2: "match_rewards.kickoff_pack_2",
	KickoffPack3: "match_rewards.kickoff_pack_3",
	SeasonPack1:  "match_rewards.season_pack_1",
	SeasonPack2:  "match_rewards.season_pack_2",
	SeasonPack3:  "match_rewards.season_pack_3",
}

// Generated where

var MatchRewardWhere = struct {
	ID           whereHelperint
	Position     whereHelpernull_Int
	Amount       whereHelperfloat64
	MatchID      whereHelperstring
	MaxPosition  whereHelpernull_Int
	MinPosition  whereHelperint
	Balls        whereHelperint
	Event        whereHelperint
	Game         whereHelperfloat64
	Lapt         whereHelperfloat64
	Shirts       whereHelperint
	SignedBalls  whereHelperint
	SignedShirts whereHelperint
	KickoffPack1 whereHelperint
	KickoffPack2 whereHelperint
	KickoffPack3 whereHelperint
	SeasonPack1  whereHelperint
	SeasonPack2  whereHelperint
	SeasonPack3  whereHelperint
}{
	ID:           whereHelperint{field: "\"match_rewards\".\"id\""},
	Position:     whereHelpernull_Int{field: "\"match_rewards\".\"position\""},
	Amount:       whereHelperfloat64{field: "\"match_rewards\".\"amount\""},
	MatchID:      whereHelperstring{field: "\"match_rewards\".\"match_id\""},
	MaxPosition:  whereHelpernull_Int{field: "\"match_rewards\".\"max_position\""},
	MinPosition:  whereHelperint{field: "\"match_rewards\".\"min_position\""},
	Balls:        whereHelperint{field: "\"match_rewards\".\"balls\""},
	Event:        whereHelperint{field: "\"match_rewards\".\"event\""},
	Game:         whereHelperfloat64{field: "\"match_rewards\".\"game\""},
	Lapt:         whereHelperfloat64{field: "\"match_rewards\".\"lapt\""},
	Shirts:       whereHelperint{field: "\"match_rewards\".\"shirts\""},
	SignedBalls:  whereHelperint{field: "\"match_rewards\".\"signed_balls\""},
	SignedShirts: whereHelperint{field: "\"match_rewards\".\"signed_shirts\""},
	KickoffPack1: whereHelperint{field: "\"match_rewards\".\"kickoff_pack_1\""},
	KickoffPack2: whereHelperint{field: "\"match_rewards\".\"kickoff_pack_2\""},
	KickoffPack3: whereHelperint{field: "\"match_rewards\".\"kickoff_pack_3\""},
	SeasonPack1:  whereHelperint{field: "\"match_rewards\".\"season_pack_1\""},
	SeasonPack2:  whereHelperint{field: "\"match_rewards\".\"season_pack_2\""},
	SeasonPack3:  whereHelperint{field: "\"match_rewards\".\"season_pack_3\""},
}

// MatchRewardRels is where relationship names are stored.
var MatchRewardRels = struct {
	Match string
}{
	Match: "Match",
}

// matchRewardR is where relationships are stored.
type matchRewardR struct {
	Match *Match `boil:"Match" json:"Match" toml:"Match" yaml:"Match"`
}

// NewStruct creates a new relationship struct
func (*matchRewardR) NewStruct() *matchRewardR {
	return &matchRewardR{}
}

func (r *matchRewardR) GetMatch() *Match {
	if r == nil {
		return nil
	}
	return r.Match
}

// matchRewardL is where Load methods for each relationship are stored.
type matchRewardL struct{}

var (
	matchRewardAllColumns            = []string{"id", "position", "amount", "match_id", "max_position", "min_position", "balls", "event", "game", "lapt", "shirts", "signed_balls", "signed_shirts", "kickoff_pack_1", "kickoff_pack_2", "kickoff_pack_3", "season_pack_1", "season_pack_2", "season_pack_3"}
	matchRewardColumnsWithoutDefault = []string{"amount", "match_id", "min_position", "balls", "event", "game", "lapt", "shirts", "signed_balls", "signed_shirts", "kickoff_pack_1", "kickoff_pack_2", "kickoff_pack_3", "season_pack_1", "season_pack_2", "season_pack_3"}
	matchRewardColumnsWithDefault    = []string{"id", "position", "max_position"}
	matchRewardPrimaryKeyColumns     = []string{"id"}
	matchRewardGeneratedColumns      = []string{}
)

type (
	// MatchRewardSlice is an alias for a slice of pointers to MatchReward.
	// This should almost always be used instead of []MatchReward.
	MatchRewardSlice []*MatchReward

	matchRewardQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	matchRewardType                 = reflect.TypeOf(&MatchReward{})
	matchRewardMapping              = queries.MakeStructMapping(matchRewardType)
	matchRewardPrimaryKeyMapping, _ = queries.BindMapping(matchRewardType, matchRewardMapping, matchRewardPrimaryKeyColumns)
	matchRewardInsertCacheMut       sync.RWMutex
	matchRewardInsertCache          = make(map[string]insertCache)
	matchRewardUpdateCacheMut       sync.RWMutex
	matchRewardUpdateCache          = make(map[string]updateCache)
	matchRewardUpsertCacheMut       sync.RWMutex
	matchRewardUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single matchReward record from the query.
func (q matchRewardQuery) One(exec boil.Executor) (*MatchReward, error) {
	o := &MatchReward{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for match_rewards")
	}

	return o, nil
}

// All returns all MatchReward records from the query.
func (q matchRewardQuery) All(exec boil.Executor) (MatchRewardSlice, error) {
	var o []*MatchReward

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to MatchReward slice")
	}

	return o, nil
}

// Count returns the count of all MatchReward records in the query.
func (q matchRewardQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count match_rewards rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q matchRewardQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if match_rewards exists")
	}

	return count > 0, nil
}

// Match pointed to by the foreign key.
func (o *MatchReward) Match(mods ...qm.QueryMod) matchQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MatchID),
	}

	queryMods = append(queryMods, mods...)

	return Matches(queryMods...)
}

// LoadMatch allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (matchRewardL) LoadMatch(e boil.Executor, singular bool, maybeMatchReward interface{}, mods queries.Applicator) error {
	var slice []*MatchReward
	var object *MatchReward

	if singular {
		var ok bool
		object, ok = maybeMatchReward.(*MatchReward)
		if !ok {
			object = new(MatchReward)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMatchReward)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMatchReward))
			}
		}
	} else {
		s, ok := maybeMatchReward.(*[]*MatchReward)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMatchReward)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMatchReward))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &matchRewardR{}
		}
		args[object.MatchID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &matchRewardR{}
			}

			args[obj.MatchID] = struct{}{}

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
		qm.From(`matches`),
		qm.WhereIn(`matches.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Match")
	}

	var resultSlice []*Match
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Match")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for matches")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for matches")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Match = foreign
		if foreign.R == nil {
			foreign.R = &matchR{}
		}
		foreign.R.MatchRewards = append(foreign.R.MatchRewards, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MatchID == foreign.ID {
				local.R.Match = foreign
				if foreign.R == nil {
					foreign.R = &matchR{}
				}
				foreign.R.MatchRewards = append(foreign.R.MatchRewards, local)
				break
			}
		}
	}

	return nil
}

// SetMatch of the matchReward to the related item.
// Sets o.R.Match to related.
// Adds o to related.R.MatchRewards.
func (o *MatchReward) SetMatch(exec boil.Executor, insert bool, related *Match) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"match_rewards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"match_id"}),
		strmangle.WhereClause("\"", "\"", 2, matchRewardPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MatchID = related.ID
	if o.R == nil {
		o.R = &matchRewardR{
			Match: related,
		}
	} else {
		o.R.Match = related
	}

	if related.R == nil {
		related.R = &matchR{
			MatchRewards: MatchRewardSlice{o},
		}
	} else {
		related.R.MatchRewards = append(related.R.MatchRewards, o)
	}

	return nil
}

// MatchRewards retrieves all the records using an executor.
func MatchRewards(mods ...qm.QueryMod) matchRewardQuery {
	mods = append(mods, qm.From("\"match_rewards\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"match_rewards\".*"})
	}

	return matchRewardQuery{q}
}

// FindMatchReward retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMatchReward(exec boil.Executor, iD int, selectCols ...string) (*MatchReward, error) {
	matchRewardObj := &MatchReward{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"match_rewards\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, matchRewardObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from match_rewards")
	}

	return matchRewardObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MatchReward) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no match_rewards provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(matchRewardColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	matchRewardInsertCacheMut.RLock()
	cache, cached := matchRewardInsertCache[key]
	matchRewardInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			matchRewardAllColumns,
			matchRewardColumnsWithDefault,
			matchRewardColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(matchRewardType, matchRewardMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(matchRewardType, matchRewardMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"match_rewards\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"match_rewards\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "schema: unable to insert into match_rewards")
	}

	if !cached {
		matchRewardInsertCacheMut.Lock()
		matchRewardInsertCache[key] = cache
		matchRewardInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the MatchReward.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MatchReward) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	matchRewardUpdateCacheMut.RLock()
	cache, cached := matchRewardUpdateCache[key]
	matchRewardUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			matchRewardAllColumns,
			matchRewardPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update match_rewards, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"match_rewards\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, matchRewardPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(matchRewardType, matchRewardMapping, append(wl, matchRewardPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schema: unable to update match_rewards row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for match_rewards")
	}

	if !cached {
		matchRewardUpdateCacheMut.Lock()
		matchRewardUpdateCache[key] = cache
		matchRewardUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q matchRewardQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for match_rewards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for match_rewards")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MatchRewardSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), matchRewardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"match_rewards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, matchRewardPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in matchReward slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all matchReward")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MatchReward) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("schema: no match_rewards provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(matchRewardColumnsWithDefault, o)

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

	matchRewardUpsertCacheMut.RLock()
	cache, cached := matchRewardUpsertCache[key]
	matchRewardUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			matchRewardAllColumns,
			matchRewardColumnsWithDefault,
			matchRewardColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			matchRewardAllColumns,
			matchRewardPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert match_rewards, could not build update column list")
		}

		ret := strmangle.SetComplement(matchRewardAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(matchRewardPrimaryKeyColumns) == 0 {
				return errors.New("schema: unable to upsert match_rewards, could not build conflict column list")
			}

			conflict = make([]string, len(matchRewardPrimaryKeyColumns))
			copy(conflict, matchRewardPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"match_rewards\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(matchRewardType, matchRewardMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(matchRewardType, matchRewardMapping, ret)
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
		return errors.Wrap(err, "schema: unable to upsert match_rewards")
	}

	if !cached {
		matchRewardUpsertCacheMut.Lock()
		matchRewardUpsertCache[key] = cache
		matchRewardUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single MatchReward record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MatchReward) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no MatchReward provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), matchRewardPrimaryKeyMapping)
	sql := "DELETE FROM \"match_rewards\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from match_rewards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for match_rewards")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q matchRewardQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no matchRewardQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from match_rewards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for match_rewards")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MatchRewardSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), matchRewardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"match_rewards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, matchRewardPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from matchReward slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for match_rewards")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MatchReward) Reload(exec boil.Executor) error {
	ret, err := FindMatchReward(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MatchRewardSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MatchRewardSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), matchRewardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"match_rewards\".* FROM \"match_rewards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, matchRewardPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in MatchRewardSlice")
	}

	*o = slice

	return nil
}

// MatchRewardExists checks if the MatchReward row exists.
func MatchRewardExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"match_rewards\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if match_rewards exists")
	}

	return exists, nil
}

// Exists checks if the MatchReward row exists.
func (o *MatchReward) Exists(exec boil.Executor) (bool, error) {
	return MatchRewardExists(exec, o.ID)
}
