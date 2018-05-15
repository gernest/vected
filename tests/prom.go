package tests

import "github.com/gernest/mad"

func TestBefore() mad.Test {
	return mad.Describe("Before",
		mad.It("be called before the testcase", func(t mad.T) {
			before := 500
			ts := mad.Describe("TestT_Before",
				mad.Before(func() {
					before = 200
				}),
				mad.It("must set before value", func(rs mad.T) {
				}),
			)
			mad.Exec(ts)
			if before != 200 {
				t.Errorf("expected %v got %v", 200, before)
			}
		}),
	)
}

func TestAfter() mad.Test {
	return mad.Describe("After",
		mad.It("should be called after the testcase", func(t mad.T) {
			after := 500
			ts := mad.Describe("TestAfter",
				mad.After(func() {
					after = 200
				}),
				mad.It("must set before value", func(rs mad.T) {
					after = 0
				}),
			)
			mad.Exec(ts)
			if after != 200 {
				t.Errorf("expected %v got %v", 200, after)
			}
		}),
	)
}