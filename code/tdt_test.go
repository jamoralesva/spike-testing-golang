package main

import "testing"
import "time"
import "fmt"

type TimeTestCase struct {
    gmt  string
    loc  string
    want string
}

//Helper function
func assertTime(t *testing.T, got, want string) {
    t.Helper() 
    if got != want {
        t.Errorf("got %s; want %s", got, want)
    }
}

func TestTime(t *testing.T) {
    testCases := []TimeTestCase {
        {"12:31", "Europe/Zuri", "13:31"},
        {"12:31", "America/New_York", "7:31"},
        {"08:08", "Australia/Sydney", "18:08"},
    }
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
            //Arrange
            loc, err := time.LoadLocation(tc.loc)
            if err != nil {
                t.Fatal("could not load location")
            }

            //Act
            gmt, _ := time.Parse("15:04", tc.gmt)
            got := gmt.In(loc).Format("15:04")
            
            //Assert
            assertTime(t, got, tc.want)
            
        })
    }
}
