// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Houses", testHouses)
	t.Run("Prefs", testPrefs)
	t.Run("Rooms", testRooms)
}

func TestDelete(t *testing.T) {
	t.Run("Houses", testHousesDelete)
	t.Run("Prefs", testPrefsDelete)
	t.Run("Rooms", testRoomsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Houses", testHousesQueryDeleteAll)
	t.Run("Prefs", testPrefsQueryDeleteAll)
	t.Run("Rooms", testRoomsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Houses", testHousesSliceDeleteAll)
	t.Run("Prefs", testPrefsSliceDeleteAll)
	t.Run("Rooms", testRoomsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Houses", testHousesExists)
	t.Run("Prefs", testPrefsExists)
	t.Run("Rooms", testRoomsExists)
}

func TestFind(t *testing.T) {
	t.Run("Houses", testHousesFind)
	t.Run("Prefs", testPrefsFind)
	t.Run("Rooms", testRoomsFind)
}

func TestBind(t *testing.T) {
	t.Run("Houses", testHousesBind)
	t.Run("Prefs", testPrefsBind)
	t.Run("Rooms", testRoomsBind)
}

func TestOne(t *testing.T) {
	t.Run("Houses", testHousesOne)
	t.Run("Prefs", testPrefsOne)
	t.Run("Rooms", testRoomsOne)
}

func TestAll(t *testing.T) {
	t.Run("Houses", testHousesAll)
	t.Run("Prefs", testPrefsAll)
	t.Run("Rooms", testRoomsAll)
}

func TestCount(t *testing.T) {
	t.Run("Houses", testHousesCount)
	t.Run("Prefs", testPrefsCount)
	t.Run("Rooms", testRoomsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Houses", testHousesHooks)
	t.Run("Prefs", testPrefsHooks)
	t.Run("Rooms", testRoomsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Houses", testHousesInsert)
	t.Run("Houses", testHousesInsertWhitelist)
	t.Run("Prefs", testPrefsInsert)
	t.Run("Prefs", testPrefsInsertWhitelist)
	t.Run("Rooms", testRoomsInsert)
	t.Run("Rooms", testRoomsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Houses", testHousesReload)
	t.Run("Prefs", testPrefsReload)
	t.Run("Rooms", testRoomsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Houses", testHousesReloadAll)
	t.Run("Prefs", testPrefsReloadAll)
	t.Run("Rooms", testRoomsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Houses", testHousesSelect)
	t.Run("Prefs", testPrefsSelect)
	t.Run("Rooms", testRoomsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Houses", testHousesUpdate)
	t.Run("Prefs", testPrefsUpdate)
	t.Run("Rooms", testRoomsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Houses", testHousesSliceUpdateAll)
	t.Run("Prefs", testPrefsSliceUpdateAll)
	t.Run("Rooms", testRoomsSliceUpdateAll)
}
