package requests

import (
	"testing"
)

func TestRootNeedsToReturnIndexByConvention(t *testing.T){

    expected := "./www/index.html"
    got := getPathBaseOnText("/")

    if got != expected {

        t.Errorf("Expected %s but got %s", expected, got)
    }
}

func TestOtherStringWillRemainTheSameButPointingToWWWFolder(t *testing.T){

    expected := "./www/hello.html"
    got := getPathBaseOnText("/hello.html")

    if got != expected {

        t.Errorf("Expected %s but got %s", expected, got)
    }
}

func TestBlankAddressShouldProvideDefaultPage(t *testing.T){

    expected := "./www/index.html"
    got := getPathBaseOnText("")

    if got != expected {

        t.Errorf("Expected %s but got %s", expected, got)
    }
}

func TestHttpVersionShouldAcceptVersion(t *testing.T){

    expected := float32(1.1)
    got, err := getHttpVersionBasedOnText("HTTP/1.1")

    if got != expected &&
        err == nil {

        t.Errorf("Expected %f but got %f", expected, got)
    }
}

func TestHttpVersionShouldReactWhenItIsNotValidValueToParse(t *testing.T){

    expected := float32(0.0)
    got, err := getHttpVersionBasedOnText("HTTP")

    if got != expected &&
        err != nil {

        t.Errorf("Expected %f but got %f", expected, got)
    }
}
