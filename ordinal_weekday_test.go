package recurrence

import (
	"testing"
	"time"
)

func TestOrdinalWeekday(t *testing.T) {
	r := YearRange(2006, time.UTC)

	// First
	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Sunday),
		"2006-01-01", "2006-02-05", "2006-03-05", "2006-04-02", "2006-05-07",
		"2006-06-04", "2006-07-02", "2006-08-06", "2006-09-03", "2006-10-01",
		"2006-11-05", "2006-12-03")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Monday),
		"2006-01-02", "2006-02-06", "2006-03-06", "2006-04-03", "2006-05-01",
		"2006-06-05", "2006-07-03", "2006-08-07", "2006-09-04", "2006-10-02",
		"2006-11-06", "2006-12-04")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Tuesday),
		"2006-01-03", "2006-02-07", "2006-03-07", "2006-04-04", "2006-05-02",
		"2006-06-06", "2006-07-04", "2006-08-01", "2006-09-05", "2006-10-03",
		"2006-11-07", "2006-12-05")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Wednesday),
		"2006-01-04", "2006-02-01", "2006-03-01", "2006-04-05", "2006-05-03",
		"2006-06-07", "2006-07-05", "2006-08-02", "2006-09-06", "2006-10-04",
		"2006-11-01", "2006-12-06")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Thursday),
		"2006-01-05", "2006-02-02", "2006-03-02", "2006-04-06", "2006-05-04",
		"2006-06-01", "2006-07-06", "2006-08-03", "2006-09-07", "2006-10-05",
		"2006-11-02", "2006-12-07")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Friday),
		"2006-01-06", "2006-02-03", "2006-03-03", "2006-04-07", "2006-05-05",
		"2006-06-02", "2006-07-07", "2006-08-04", "2006-09-01", "2006-10-06",
		"2006-11-03", "2006-12-01")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(First, Saturday),
		"2006-01-07", "2006-02-04", "2006-03-04", "2006-04-01", "2006-05-06",
		"2006-06-03", "2006-07-01", "2006-08-05", "2006-09-02", "2006-10-07",
		"2006-11-04", "2006-12-02")

	// Second
	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Sunday),
		"2006-01-08", "2006-02-12", "2006-03-12", "2006-04-09", "2006-05-14",
		"2006-06-11", "2006-07-09", "2006-08-13", "2006-09-10", "2006-10-08",
		"2006-11-12", "2006-12-10")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Monday),
		"2006-01-09", "2006-02-13", "2006-03-13", "2006-04-10", "2006-05-08",
		"2006-06-12", "2006-07-10", "2006-08-14", "2006-09-11", "2006-10-09",
		"2006-11-13", "2006-12-11")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Tuesday),
		"2006-01-10", "2006-02-14", "2006-03-14", "2006-04-11", "2006-05-09",
		"2006-06-13", "2006-07-11", "2006-08-08", "2006-09-12", "2006-10-10",
		"2006-11-14", "2006-12-12")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Wednesday),
		"2006-01-11", "2006-02-08", "2006-03-08", "2006-04-12", "2006-05-10",
		"2006-06-14", "2006-07-12", "2006-08-09", "2006-09-13", "2006-10-11",
		"2006-11-08", "2006-12-13")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Thursday),
		"2006-01-12", "2006-02-09", "2006-03-09", "2006-04-13", "2006-05-11",
		"2006-06-08", "2006-07-13", "2006-08-10", "2006-09-14", "2006-10-12",
		"2006-11-09", "2006-12-14")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Friday),
		"2006-01-13", "2006-02-10", "2006-03-10", "2006-04-14", "2006-05-12",
		"2006-06-09", "2006-07-14", "2006-08-11", "2006-09-08", "2006-10-13",
		"2006-11-10", "2006-12-08")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Second, Saturday),
		"2006-01-14", "2006-02-11", "2006-03-11", "2006-04-08", "2006-05-13",
		"2006-06-10", "2006-07-08", "2006-08-12", "2006-09-09", "2006-10-14",
		"2006-11-11", "2006-12-09")

	// Third
	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Sunday),
		"2006-01-15", "2006-02-19", "2006-03-19", "2006-04-16", "2006-05-21",
		"2006-06-18", "2006-07-16", "2006-08-20", "2006-09-17", "2006-10-15",
		"2006-11-19", "2006-12-17")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Monday),
		"2006-01-16", "2006-02-20", "2006-03-20", "2006-04-17", "2006-05-15",
		"2006-06-19", "2006-07-17", "2006-08-21", "2006-09-18", "2006-10-16",
		"2006-11-20", "2006-12-18")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Tuesday),
		"2006-01-17", "2006-02-21", "2006-03-21", "2006-04-18", "2006-05-16",
		"2006-06-20", "2006-07-18", "2006-08-15", "2006-09-19", "2006-10-17",
		"2006-11-21", "2006-12-19")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Wednesday),
		"2006-01-18", "2006-02-15", "2006-03-15", "2006-04-19", "2006-05-17",
		"2006-06-21", "2006-07-19", "2006-08-16", "2006-09-20", "2006-10-18",
		"2006-11-15", "2006-12-20")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Thursday),
		"2006-01-19", "2006-02-16", "2006-03-16", "2006-04-20", "2006-05-18",
		"2006-06-15", "2006-07-20", "2006-08-17", "2006-09-21", "2006-10-19",
		"2006-11-16", "2006-12-21")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Friday),
		"2006-01-20", "2006-02-17", "2006-03-17", "2006-04-21", "2006-05-19",
		"2006-06-16", "2006-07-21", "2006-08-18", "2006-09-15", "2006-10-20",
		"2006-11-17", "2006-12-15")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Third, Saturday),
		"2006-01-21", "2006-02-18", "2006-03-18", "2006-04-15", "2006-05-20",
		"2006-06-17", "2006-07-15", "2006-08-19", "2006-09-16", "2006-10-21",
		"2006-11-18", "2006-12-16")

	// Fourth
	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Sunday),
		"2006-01-22", "2006-02-26", "2006-03-26", "2006-04-23", "2006-05-28",
		"2006-06-25", "2006-07-23", "2006-08-27", "2006-09-24", "2006-10-22",
		"2006-11-26", "2006-12-24")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Monday),
		"2006-01-23", "2006-02-27", "2006-03-27", "2006-04-24", "2006-05-22",
		"2006-06-26", "2006-07-24", "2006-08-28", "2006-09-25", "2006-10-23",
		"2006-11-27", "2006-12-25")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Tuesday),
		"2006-01-24", "2006-02-28", "2006-03-28", "2006-04-25", "2006-05-23",
		"2006-06-27", "2006-07-25", "2006-08-22", "2006-09-26", "2006-10-24",
		"2006-11-28", "2006-12-26")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Wednesday),
		"2006-01-25", "2006-02-22", "2006-03-22", "2006-04-26", "2006-05-24",
		"2006-06-28", "2006-07-26", "2006-08-23", "2006-09-27", "2006-10-25",
		"2006-11-22", "2006-12-27")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Thursday),
		"2006-01-26", "2006-02-23", "2006-03-23", "2006-04-27", "2006-05-25",
		"2006-06-22", "2006-07-27", "2006-08-24", "2006-09-28", "2006-10-26",
		"2006-11-23", "2006-12-28")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Friday),
		"2006-01-27", "2006-02-24", "2006-03-24", "2006-04-28", "2006-05-26",
		"2006-06-23", "2006-07-28", "2006-08-25", "2006-09-22", "2006-10-27",
		"2006-11-24", "2006-12-22")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fourth, Saturday),
		"2006-01-28", "2006-02-25", "2006-03-25", "2006-04-22", "2006-05-27",
		"2006-06-24", "2006-07-22", "2006-08-26", "2006-09-23", "2006-10-28",
		"2006-11-25", "2006-12-23")

	// Fifth
	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Sunday),
		"2006-01-29", "2006-04-30", "2006-07-30", "2006-09-31", "2006-10-29",
		"2006-12-31")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Monday),
		"2006-01-30", "2006-04-31", "2006-05-29", "2006-07-31", "2006-10-30")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Tuesday),
		"2006-01-31", "2006-05-30", "2006-08-29", "2006-10-31")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Wednesday),
		"2006-02-29", "2006-03-29", "2006-05-31", "2006-08-30", "2006-11-29")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Thursday),
		"2006-02-30", "2006-03-30", "2006-06-29", "2006-08-31", "2006-11-30")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Friday),
		"2006-02-31", "2006-03-31", "2006-06-30", "2006-09-29", "2006-11-31", "2006-12-29")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Fifth, Saturday),
		"2006-04-29", "2006-06-31", "2006-07-29", "2006-09-30", "2006-12-30")

	// Last
	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Sunday),
		"2006-01-29", "2006-02-26", "2006-03-26", "2006-04-30", "2006-05-28",
		"2006-06-25", "2006-07-30", "2006-08-27", "2006-09-24", "2006-10-29",
		"2006-11-26", "2006-12-31")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Monday),
		"2006-01-30", "2006-02-27", "2006-03-27", "2006-04-24", "2006-05-29",
		"2006-06-26", "2006-07-31", "2006-08-28", "2006-09-25", "2006-10-30",
		"2006-11-27", "2006-12-25")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Tuesday),
		"2006-01-31", "2006-02-28", "2006-03-28", "2006-04-25", "2006-05-30",
		"2006-06-27", "2006-07-25", "2006-08-29", "2006-09-26", "2006-10-31",
		"2006-11-28", "2006-12-26")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Wednesday),
		"2006-01-25", "2006-02-22", "2006-03-29", "2006-04-26", "2006-05-31",
		"2006-06-28", "2006-07-26", "2006-08-30", "2006-09-27", "2006-10-25",
		"2006-11-29", "2006-12-27")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Thursday),
		"2006-01-26", "2006-02-23", "2006-03-30", "2006-04-27", "2006-05-25",
		"2006-06-29", "2006-07-27", "2006-08-31", "2006-09-28", "2006-10-26",
		"2006-11-30", "2006-12-28")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Friday),
		"2006-01-27", "2006-02-24", "2006-03-31", "2006-04-28", "2006-05-26",
		"2006-06-30", "2006-07-28", "2006-08-25", "2006-09-29", "2006-10-27",
		"2006-11-24", "2006-12-29")

	assertIsOnlyOccurring(t, r, OrdinalWeekday(Last, Saturday),
		"2006-01-28", "2006-02-25", "2006-03-25", "2006-04-29", "2006-05-27",
		"2006-06-24", "2006-07-29", "2006-08-26", "2006-09-30", "2006-10-28",
		"2006-11-25", "2006-12-30")
}
