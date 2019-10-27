// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	// start :=
	// afterTenSeconds := start.Add(time.Second * 10)



	// 2011-04-25を受け取って、2011-04-25 00:00:00 +0000形式にするのがparseメソッド。
	// 受け取った2011-04-25 00:00:00 +0000に1万秒を足して返却すればいい。
	return t.Add(time.Second * 1e9)
}
