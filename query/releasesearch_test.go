package query

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"text/template"

	. "github.com/smartystreets/goconvey/convey"
)

var validators = NewReleaseQueryParamValidator()

func TestLimit(t *testing.T) {
	Convey("given a limit validator, and a set of limits as strings", t, func() {
		validator := validators["limit"]
		limits := []struct {
			given   string
			exValue int
			exError error
		}{
			{given: "XXX", exValue: 0, exError: errors.New("limit search parameter provided with non numeric characters")},
			{given: "-1", exValue: 0, exError: errors.New("limit search parameter provided with negative value")},
			{given: "5001", exValue: 0, exError: fmt.Errorf("limit search parameter provided with a value that is too high")},
			{given: "0", exValue: 0, exError: nil},
			{given: "5000", exValue: 5000, exError: nil},
		}

		for _, ls := range limits {
			v, e := validator(ls.given)

			So(v, ShouldEqual, ls.exValue)
			So(e, ShouldResemble, ls.exError)
		}
	})
}

func TestOffset(t *testing.T) {
	Convey("given an offset validator, and a set of offsets as strings", t, func() {
		validator := validators["offset"]
		offsets := []struct {
			given   string
			exValue int
			exError error
		}{
			{given: "XXX", exValue: 0, exError: errors.New("offset search parameter provided with non numeric characters")},
			{given: "-1", exValue: 0, exError: errors.New("offset search parameter provided with negative value")},
			{given: "0", exValue: 0, exError: nil},
			{given: "1", exValue: 1, exError: nil},
			{given: "15000", exValue: 15000, exError: nil},
		}

		for _, ls := range offsets {
			v, e := validator(ls.given)

			So(v, ShouldEqual, ls.exValue)
			So(e, ShouldResemble, ls.exError)
		}
	})
}

func TestDates(t *testing.T) {
	Convey("given a date validator, and a set of erroneous date strings", t, func() {
		validator := validators["date"]
		badDates := []string{"XXX", "Jan 21", "30/12/2021", "2021-13-31", "2021-12-32", "2021-02-29", "2300-12-31"}

		Convey("errors are generated, and zero values returned on validation", func() {
			for _, ds := range badDates {
				v, e := validator(ds)

				So(v, ShouldBeNil)
				So(e, ShouldNotBeNil)
			}
		})

		Convey("but a good date string is validated without error, and the appropriate Date returned", func() {
			date := "2022-12-31"
			v, e := validator(date)

			So(v, ShouldResemble, MustParseDate(date))
			So(e, ShouldBeNil)
		})
	})
}

func TestSort(t *testing.T) {
	Convey("given a sort validator, and a set of erroneous sort string options", t, func() {
		validator := validators["sort"]
		badSortOptions := []string{"dont sort", "sort-by-date", "date-ascending"}

		Convey("errors are generated, and zero values returned on validation", func() {
			for _, ds := range badSortOptions {
				v, e := validator(ds)

				So(v, ShouldBeNil)
				So(e, ShouldNotBeNil)
			}
		})

		Convey("but a good sort option string is validated without error, and the appropriate Sort option returned", func() {
			goodSortOptions := []struct {
				given   string
				exValue Sort
			}{
				{given: "release_date_asc", exValue: RelDateAsc},
				{given: "release_date_desc", exValue: RelDateDesc},
				{given: "title_asc", exValue: TitleAsc},
				{given: "title_desc", exValue: TitleDesc},
			}

			for _, gso := range goodSortOptions {
				v, e := validator(gso.given)

				So(v, ShouldEqual, gso.exValue)
				So(e, ShouldBeNil)
			}
		})
	})
}

func TestBuildSearchReleaseQuery(t *testing.T) {
	Convey("Should return InternalError for invalid template", t, func() {
		qb := createReleaseQueryBuilderForTemplate("dummy{{.Moo}}")
		query, err := qb.BuildSearchQuery(context.Background(), ReleaseSearchRequest{})

		So(err, ShouldNotBeNil)
		So(query, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "creation of search from template failed")
	})

	Convey("Should include all search parameters in elastic search query", t, func() {
		qb := createReleaseQueryBuilderForTemplate("Term={{.Term}};" +
			"From={{.From}};" +
			"Size={{.Size}};" +
			"SortBy={{.SortBy.ESString}};" +
			"ReleasedAfter={{.ReleasedAfter.ESString}};" +
			"ReleasedBefore={{.ReleasedBefore.ESString}};" +
			"Upcoming={{.Upcoming}};" +
			"Published={{.Published}};" +
			"Highlight={{.Highlight}};" +
			"Now={{.Now}};" +
			"NowES={{.Now.ESString}}")

		query, err := qb.BuildSearchQuery(context.Background(), ReleaseSearchRequest{
			Term:           "query+term",
			From:           0,
			Size:           25,
			SortBy:         TitleAsc,
			ReleasedAfter:  Date{},
			ReleasedBefore: MustParseDate("2020-12-31"),
			Upcoming:       true,
			Published:      false,
			Highlight:      true,
			Now:            MustParseDate("2001-01-01"),
		})

		So(err, ShouldBeNil)
		So(query, ShouldNotBeNil)
		queryString := string(query)
		So(queryString, ShouldContainSubstring, "Term=query+term")
		So(queryString, ShouldContainSubstring, "From=0")
		So(queryString, ShouldContainSubstring, "Size=25")
		So(queryString, ShouldContainSubstring, `SortBy={"description.title": "asc"}`)
		So(queryString, ShouldContainSubstring, "ReleasedAfter=null")
		So(queryString, ShouldContainSubstring, `ReleasedBefore="2020-12-31"`)
		So(queryString, ShouldContainSubstring, "Upcoming=true")
		So(queryString, ShouldContainSubstring, "Published=false")
		So(queryString, ShouldContainSubstring, "Highlight=true")
		So(queryString, ShouldContainSubstring, `Now=2001-01-01`)
		So(queryString, ShouldContainSubstring, `NowES="2001-01-01"`)
	})
}

func createReleaseQueryBuilderForTemplate(rawTemplate string) *ReleaseBuilder {
	temp, err := template.New("search.tmpl").Parse(rawTemplate)
	So(err, ShouldBeNil)
	return &ReleaseBuilder{
		searchTemplates: temp,
	}
}