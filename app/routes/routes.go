// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Help(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Help", args).URL
}


type tProducts struct {}
var Products tProducts


func (_ tProducts) Show(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Products.Show", args).URL
}

func (_ tProducts) Crud(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Products.Crud", args).URL
}

func (_ tProducts) GetProduct(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("Products.GetProduct", args).URL
}

func (_ tProducts) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Products.Delete", args).URL
}

func (_ tProducts) Pop(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Products.Pop", args).URL
}

func (_ tProducts) Insert(
		pro interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pro", pro)
	return revel.MainRouter.Reverse("Products.Insert", args).URL
}


type tController struct {}
var Controller tController



type tTxnController struct {}
var TxnController tTxnController


func (_ tTxnController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TxnController.Begin", args).URL
}

func (_ tTxnController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TxnController.Commit", args).URL
}

func (_ tTxnController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TxnController.Rollback", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


