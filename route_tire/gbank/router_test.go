package gbank

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *Router {
	r := NewRouter()
	r.AddRoute("GET", "/", nil)
	r.AddRoute("GET", "/hello/:name", nil)
	r.AddRoute("GET", "/hello/b/c", nil)
	r.AddRoute("GET", "/hi/:name", nil)
	r.AddRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	r := NewRouter()
	ok := reflect.DeepEqual(r.parsePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(r.parsePath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(r.parsePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/bank")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "bank" {
		t.Fatal("name should be equal to 'bank'")
	}

	fmt.Printf("matched Path: %s, params['name']: %s\n", n.pattern, ps["name"])

}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be file1.txt")
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be css/test.css")
	}

}

//func TestGetRoutes(t *testing.T) {
//	r := newTestRouter()
//	nodes := r.GetRoutes("GET")
//	for i, n := range nodes {
//		fmt.Println(i+1, n)
//	}
//
//	if len(nodes) != 5 {
//		t.Fatal("the number of routes shoule be 4")
//	}
//}
