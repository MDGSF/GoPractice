package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/google/btree"
	"github.com/gorilla/mux"
)

var GlobalTree *btree.BTree
var GlobalTree2 *btree.BTree

func main() {
	flag.Parse()

	GlobalTree = btree.New(32)

	go func() {
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()

	r := mux.NewRouter()
	r.HandleFunc("/addnode", AddNode)
	r.HandleFunc("/clone", Clone)
	r.HandleFunc("/delete2", Delete)
	r.HandleFunc("/gc", GC)

	srv := http.Server{
		Addr:         "localhost:8888",
		WriteTimeout: time.Second * 15,
		Handler:      r,
	}

	glog.Infof("listen at %v", "localhost:8888")
	if err := srv.ListenAndServe(); err != nil {
		glog.Fatal(err)
	}
}

func AddNode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	glog.Info(r.Form, r.PostForm)

	tree := r.Form.Get("tree")

	numstr := r.Form.Get("num")
	num, _ := strconv.Atoi(numstr)

	glog.Infof("tree1 len = %v", GlobalTree.Len())
	if GlobalTree2 != nil {
		glog.Infof("tree2 len = %v", GlobalTree2.Len())
	}

	if tree == "tree1" {
		for i := 0; i < num; i++ {
			GlobalTree.ReplaceOrInsert(btree.Int(i))
		}
	} else {
		for i := 0; i < num; i++ {
			GlobalTree2.ReplaceOrInsert(btree.Int(i))
		}
	}

	glog.Infof("tree1 len = %v", GlobalTree.Len())
	if GlobalTree2 != nil {
		glog.Infof("tree2 len = %v", GlobalTree2.Len())
	}
}

func Clone(w http.ResponseWriter, r *http.Request) {
	glog.Infof("clone, tree len = %v", GlobalTree.Len())
	GlobalTree2 = GlobalTree.Clone()
	glog.Infof("clone, tree len = %v, tree2 len = %v", GlobalTree.Len(), GlobalTree2.Len())
}

func Delete(w http.ResponseWriter, r *http.Request) {
	glog.Info("delete tree2")
	GlobalTree2.Clear(false)
}

func GC(w http.ResponseWriter, r *http.Request) {
	glog.Info("start GC")
	runtime.GC()
}
