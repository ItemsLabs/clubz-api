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

// Banner is an object representing the database table.
type Banner struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	Image       string    `boil:"image" json:"image" toml:"image" yaml:"image"`
	Points      int       `boil:"points" json:"points" toml:"points" yaml:"points"`
	Type        string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	Status      string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *bannerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bannerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BannerColumns = struct {
	ID          string
	Name        string
	Description string
	Image       string
	Points      string
	Type        string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	Image:       "image",
	Points:      "points",
	Type:        "type",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var BannerTableColumns = struct {
	ID          string
	Name        string
	Description string
	Image       string
	Points      string
	Type        string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "banners.id",
	Name:        "banners.name",
	Description: "banners.description",
	Image:       "banners.image",
	Points:      "banners.points",
	Type:        "banners.type",
	Status:      "banners.status",
	CreatedAt:   "banners.created_at",
	UpdatedAt:   "banners.updated_at",
}

// Generated where

var BannerWhere = struct {
	ID          whereHelperint
	Name        whereHelperstring
	Description whereHelperstring
	Image       whereHelperstring
	Points      whereHelperint
	Type        whereHelperstring
	Status      whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint{field: "\"banners\".\"id\""},
	Name:        whereHelperstring{field: "\"banners\".\"name\""},
	Description: whereHelperstring{field: "\"banners\".\"description\""},
	Image:       whereHelperstring{field: "\"banners\".\"image\""},
	Points:      whereHelperint{field: "\"banners\".\"points\""},
	Type:        whereHelperstring{field: "\"banners\".\"type\""},
	Status:      whereHelperstring{field: "\"banners\".\"status\""},
	CreatedAt:   whereHelpertime_Time{field: "\"banners\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"banners\".\"updated_at\""},
}

// BannerRels is where relationship names are stored.
var BannerRels = struct {
	UserBanners string
}{
	UserBanners: "UserBanners",
}

// bannerR is where relationships are stored.
type bannerR struct {
	UserBanners UserBannerSlice `boil:"UserBanners" json:"UserBanners" toml:"UserBanners" yaml:"UserBanners"`
}

// NewStruct creates a new relationship struct
func (*bannerR) NewStruct() *bannerR {
	return &bannerR{}
}

func (r *bannerR) GetUserBanners() UserBannerSlice {
	if r == nil {
		return nil
	}
	return r.UserBanners
}

// bannerL is where Load methods for each relationship are stored.
type bannerL struct{}

var (
	bannerAllColumns            = []string{"id", "name", "description", "image", "points", "type", "status", "created_at", "updated_at"}
	bannerColumnsWithoutDefault = []string{"name", "description", "image", "points", "type", "status", "created_at", "updated_at"}
	bannerColumnsWithDefault    = []string{"id"}
	bannerPrimaryKeyColumns     = []string{"id"}
	bannerGeneratedColumns      = []string{}
)

type (
	// BannerSlice is an alias for a slice of pointers to Banner.
	// This should almost always be used instead of []Banner.
	BannerSlice []*Banner

	bannerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bannerType                 = reflect.TypeOf(&Banner{})
	bannerMapping              = queries.MakeStructMapping(bannerType)
	bannerPrimaryKeyMapping, _ = queries.BindMapping(bannerType, bannerMapping, bannerPrimaryKeyColumns)
	bannerInsertCacheMut       sync.RWMutex
	bannerInsertCache          = make(map[string]insertCache)
	bannerUpdateCacheMut       sync.RWMutex
	bannerUpdateCache          = make(map[string]updateCache)
	bannerUpsertCacheMut       sync.RWMutex
	bannerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single banner record from the query.
func (q bannerQuery) One(exec boil.Executor) (*Banner, error) {
	o := &Banner{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for banners")
	}

	return o, nil
}

// All returns all Banner records from the query.
func (q bannerQuery) All(exec boil.Executor) (BannerSlice, error) {
	var o []*Banner

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to Banner slice")
	}

	return o, nil
}

// Count returns the count of all Banner records in the query.
func (q bannerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count banners rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bannerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if banners exists")
	}

	return count > 0, nil
}

// UserBanners retrieves all the user_banner's UserBanners with an executor.
func (o *Banner) UserBanners(mods ...qm.QueryMod) userBannerQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_banners\".\"banner_id\"=?", o.ID),
	)

	return UserBanners(queryMods...)
}

// LoadUserBanners allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bannerL) LoadUserBanners(e boil.Executor, singular bool, maybeBanner interface{}, mods queries.Applicator) error {
	var slice []*Banner
	var object *Banner

	if singular {
		var ok bool
		object, ok = maybeBanner.(*Banner)
		if !ok {
			object = new(Banner)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBanner)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBanner))
			}
		}
	} else {
		s, ok := maybeBanner.(*[]*Banner)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBanner)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBanner))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &bannerR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bannerR{}
			}
			args[obj.ID] = struct{}{}
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
		qm.From(`user_banners`),
		qm.WhereIn(`user_banners.banner_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_banners")
	}

	var resultSlice []*UserBanner
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_banners")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_banners")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_banners")
	}

	if singular {
		object.R.UserBanners = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userBannerR{}
			}
			foreign.R.Banner = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BannerID {
				local.R.UserBanners = append(local.R.UserBanners, foreign)
				if foreign.R == nil {
					foreign.R = &userBannerR{}
				}
				foreign.R.Banner = local
				break
			}
		}
	}

	return nil
}

// AddUserBanners adds the given related objects to the existing relationships
// of the banner, optionally inserting them as new records.
// Appends related to o.R.UserBanners.
// Sets related.R.Banner appropriately.
func (o *Banner) AddUserBanners(exec boil.Executor, insert bool, related ...*UserBanner) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BannerID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_banners\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"banner_id"}),
				strmangle.WhereClause("\"", "\"", 2, userBannerPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BannerID = o.ID
		}
	}

	if o.R == nil {
		o.R = &bannerR{
			UserBanners: related,
		}
	} else {
		o.R.UserBanners = append(o.R.UserBanners, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userBannerR{
				Banner: o,
			}
		} else {
			rel.R.Banner = o
		}
	}
	return nil
}

// Banners retrieves all the records using an executor.
func Banners(mods ...qm.QueryMod) bannerQuery {
	mods = append(mods, qm.From("\"banners\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"banners\".*"})
	}

	return bannerQuery{q}
}

// FindBanner retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBanner(exec boil.Executor, iD int, selectCols ...string) (*Banner, error) {
	bannerObj := &Banner{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"banners\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, bannerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from banners")
	}

	return bannerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Banner) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no banners provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(bannerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bannerInsertCacheMut.RLock()
	cache, cached := bannerInsertCache[key]
	bannerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bannerAllColumns,
			bannerColumnsWithDefault,
			bannerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bannerType, bannerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bannerType, bannerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"banners\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"banners\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "schema: unable to insert into banners")
	}

	if !cached {
		bannerInsertCacheMut.Lock()
		bannerInsertCache[key] = cache
		bannerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Banner.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Banner) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	bannerUpdateCacheMut.RLock()
	cache, cached := bannerUpdateCache[key]
	bannerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bannerAllColumns,
			bannerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update banners, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"banners\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bannerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bannerType, bannerMapping, append(wl, bannerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schema: unable to update banners row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for banners")
	}

	if !cached {
		bannerUpdateCacheMut.Lock()
		bannerUpdateCache[key] = cache
		bannerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q bannerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for banners")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for banners")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BannerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bannerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"banners\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bannerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in banner slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all banner")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Banner) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("schema: no banners provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(bannerColumnsWithDefault, o)

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

	bannerUpsertCacheMut.RLock()
	cache, cached := bannerUpsertCache[key]
	bannerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			bannerAllColumns,
			bannerColumnsWithDefault,
			bannerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bannerAllColumns,
			bannerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert banners, could not build update column list")
		}

		ret := strmangle.SetComplement(bannerAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(bannerPrimaryKeyColumns) == 0 {
				return errors.New("schema: unable to upsert banners, could not build conflict column list")
			}

			conflict = make([]string, len(bannerPrimaryKeyColumns))
			copy(conflict, bannerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"banners\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(bannerType, bannerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bannerType, bannerMapping, ret)
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
		return errors.Wrap(err, "schema: unable to upsert banners")
	}

	if !cached {
		bannerUpsertCacheMut.Lock()
		bannerUpsertCache[key] = cache
		bannerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Banner record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Banner) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no Banner provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bannerPrimaryKeyMapping)
	sql := "DELETE FROM \"banners\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from banners")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for banners")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bannerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no bannerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from banners")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for banners")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BannerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bannerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"banners\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bannerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from banner slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for banners")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Banner) Reload(exec boil.Executor) error {
	ret, err := FindBanner(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BannerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BannerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bannerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"banners\".* FROM \"banners\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bannerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in BannerSlice")
	}

	*o = slice

	return nil
}

// BannerExists checks if the Banner row exists.
func BannerExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"banners\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if banners exists")
	}

	return exists, nil
}

// Exists checks if the Banner row exists.
func (o *Banner) Exists(exec boil.Executor) (bool, error) {
	return BannerExists(exec, o.ID)
}
