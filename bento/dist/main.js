(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["main"],{

/***/ "./src/$$_lazy_route_resource lazy recursive":
/*!**********************************************************!*\
  !*** ./src/$$_lazy_route_resource lazy namespace object ***!
  \**********************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

function webpackEmptyAsyncContext(req) {
	// Here Promise.resolve().then() is used instead of new Promise() to prevent
	// uncaught exception popping up in devtools
	return Promise.resolve().then(function() {
		var e = new Error('Cannot find module "' + req + '".');
		e.code = 'MODULE_NOT_FOUND';
		throw e;
	});
}
webpackEmptyAsyncContext.keys = function() { return []; };
webpackEmptyAsyncContext.resolve = webpackEmptyAsyncContext;
module.exports = webpackEmptyAsyncContext;
webpackEmptyAsyncContext.id = "./src/$$_lazy_route_resource lazy recursive";

/***/ }),

/***/ "./src/app/amai.service.ts":
/*!*********************************!*\
  !*** ./src/app/amai.service.ts ***!
  \*********************************/
/*! exports provided: AmaiService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AmaiService", function() { return AmaiService; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/common/http */ "./node_modules/@angular/common/fesm5/http.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var AmaiService = /** @class */ (function () {
    //url : string = "";
    /*since i'm new at ts i don't exactly understatnd why
     * the constructor requires de http var on it
     */
    function AmaiService(http) {
        this.http = http;
        //local API url
        this.url = "http://localhost:9000";
    }
    /* this functions returns the array of object that is given by the api
     */
    AmaiService.prototype.getProductsObservable = function () {
        return this.http.get(this.url + "/show");
    };
    AmaiService.prototype.deleteProduct = function (id) {
        console.log("delete");
        this.http.delete(this.url + "/delete/" + id).subscribe();
    };
    AmaiService.prototype.createUser = function (name, age, gender, userEmail, aLevel, role) {
        console.log("create");
        this.http.post(this.url + "/insert", { UserName: name, UserAge: age, UserGender: gender, UserEmail: userEmail,
            UserAcademicLevel: aLevel, UserRole: role })
            .subscribe(function (res) { return Response; });
    };
    AmaiService = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Injectable"])(),
        __metadata("design:paramtypes", [_angular_common_http__WEBPACK_IMPORTED_MODULE_1__["HttpClient"]])
    ], AmaiService);
    return AmaiService;
}());



/***/ }),

/***/ "./src/app/app-routing.module.ts":
/*!***************************************!*\
  !*** ./src/app/app-routing.module.ts ***!
  \***************************************/
/*! exports provided: AppRoutingModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppRoutingModule", function() { return AppRoutingModule; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/router */ "./node_modules/@angular/router/fesm5/router.js");
/* harmony import */ var _users_table_users_table_component__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./users-table/users-table.component */ "./src/app/users-table/users-table.component.ts");
/* harmony import */ var _create_user_create_user_component__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./create-user/create-user.component */ "./src/app/create-user/create-user.component.ts");
/* harmony import */ var _panel_panel_component__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./panel/panel.component */ "./src/app/panel/panel.component.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};





// Pretty self explicatory we specify a route and it's component.
var routes = [
    { path: '', component: _users_table_users_table_component__WEBPACK_IMPORTED_MODULE_2__["UsersTableComponent"] },
    { path: 'tools', component: _create_user_create_user_component__WEBPACK_IMPORTED_MODULE_3__["CreateUserComponent"] },
    { path: 'panel', component: _panel_panel_component__WEBPACK_IMPORTED_MODULE_4__["PanelComponent"] }
];
var AppRoutingModule = /** @class */ (function () {
    function AppRoutingModule() {
    }
    AppRoutingModule = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["NgModule"])({
            imports: [_angular_router__WEBPACK_IMPORTED_MODULE_1__["RouterModule"].forRoot(routes)],
            exports: [_angular_router__WEBPACK_IMPORTED_MODULE_1__["RouterModule"]]
        })
    ], AppRoutingModule);
    return AppRoutingModule;
}());



/***/ }),

/***/ "./src/app/app.component.css":
/*!***********************************!*\
  !*** ./src/app/app.component.css ***!
  \***********************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "/* Application-wide Styles */\r\nh1 {\r\n    color: #369;\r\n    font-family: Arial, Helvetica, sans-serif;\r\n    font-size: 250%;\r\n}\r\nh2, h3 {\r\n    color: #444;\r\n    font-family: Arial, Helvetica, sans-serif;\r\n    font-weight: lighter;\r\n}\r\nbody {\r\n    height: 100%;\r\n    width: 100%;\r\n}\r\nbody, input[text], button {\r\n    color: #888;\r\n    font-family: Cambria, Georgia;\r\n}\r\n/* everywhere else */\r\n* {\r\n    font-family: Arial, Helvetica, sans-serif;\r\n}\r\n/*\r\nheader {\r\n    background-color: #1b1b1b;\r\n    color: white;\r\n    grid-area: head;\r\n}\r\n\r\nheader h1{\r\n    text-align: center;\r\n    padding: 8px;\r\n    margin: auto;\r\n    color: white;\r\n}\r\n*/\r\n/*\r\nnav {\r\n    background-color: #212121;\r\n    border-right: 5px solid #1b1b1b;\r\n    grid-area: nav;\r\n}\r\n\r\nnav a {\r\n    padding: 8px;\r\n    display: block;\r\n    color: white;\r\n    text-decoration: none;\r\n    background-color: #f44336;\r\n}\r\n*/\r\n/*\r\n.sidenav-container{\r\n    grid-area: nav;\r\n}\r\n*/\r\n/*\r\n.active{\r\n    background-color: red !important;\r\n}\r\n\r\nmain {\r\n    color: #eeeeee;\r\n    min-height: 100vh;\r\n    padding: 16px;\r\n    /* background-color: #212121; */\r\n/*\r\n    grid-area: main;\r\n}\r\n\r\n\r\nfooter {\r\n    color: #eeeeee;\r\n    background-color: #1b1b1b;\r\n    grid-column: foot;\r\n}\r\n\r\nfooter p{\r\n    text-align: center;\r\n    padding: 8px;\r\n    margin: auto;\r\n    color: white;\r\n}\r\n\r\n#page {\r\n    display: grid;\r\n    width: 100%;\r\n    display: grid;\r\n    grid-template: /*[header-left] \"head head head\" 11% [header-right]\r\n                 [main-left]   \"nav main main\" 79%  [main-right]\r\n                 /*[footer-left] \"nav main main\" 10% [footer-right]\r\n                 / 15% 1fr;\r\n}\r\n*/\r\n\r\n"

/***/ }),

/***/ "./src/app/app.component.html":
/*!************************************!*\
  !*** ./src/app/app.component.html ***!
  \************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = " <section id=\"page\">\n     <!--\n    <header> <h1> Test Page </h1> </header>\n   <!-- <nav>\n        <!-- RouterLinkActive it's a class from Angular than let us highlight the button when\n            it's selected, routerLinkActiveOptions it's used to ensure that only will happen on selected \n        <a [routerLinkActive]=\"'active'\" [routerLinkActiveOptions]=\"{exact:true}\" routerLink=\"/\">Home</a>\n        <a [routerLinkActive]=\"'active'\" [routerLinkActiveOptions]=\"{exact:true}\" class=\"\" routerLink=\"/crud\">crud</a>\n        <a [routerLinkActive]=\"'active'\" [routerLinkActiveOptions]=\"{exact:true}\" class=\"\" routerLink=\"/panel\">panel</a>\n   </nav> -->\n    <app-main-nav></app-main-nav>\n        <!-- This components returns the component specified in the app-routing module\n    <main>\n            in the variable const routes \n    <router-outlet></router-outlet>\n   <app-panel></app-panel>\n    </main>\n        -->\n\n     <!--\n    <footer> <p> Footer</p></footer>\n        -->\n</section>\n"

/***/ }),

/***/ "./src/app/app.component.ts":
/*!**********************************!*\
  !*** ./src/app/app.component.ts ***!
  \**********************************/
/*! exports provided: AppComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppComponent", function() { return AppComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};

var AppComponent = /** @class */ (function () {
    function AppComponent() {
        this.title = 'app';
    }
    AppComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-root',
            template: __webpack_require__(/*! ./app.component.html */ "./src/app/app.component.html"),
            styles: [__webpack_require__(/*! ./app.component.css */ "./src/app/app.component.css")]
        })
    ], AppComponent);
    return AppComponent;
}());



/***/ }),

/***/ "./src/app/app.module.ts":
/*!*******************************!*\
  !*** ./src/app/app.module.ts ***!
  \*******************************/
/*! exports provided: AppModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppModule", function() { return AppModule; });
/* harmony import */ var _angular_platform_browser__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/platform-browser */ "./node_modules/@angular/platform-browser/fesm5/platform-browser.js");
/* harmony import */ var _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/platform-browser/animations */ "./node_modules/@angular/platform-browser/fesm5/animations.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/forms */ "./node_modules/@angular/forms/fesm5/forms.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/common/http */ "./node_modules/@angular/common/fesm5/http.js");
/* harmony import */ var _app_component__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./app.component */ "./src/app/app.component.ts");
/* harmony import */ var _home_home_component__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ./home/home.component */ "./src/app/home/home.component.ts");
/* harmony import */ var _crud_crud_component__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ./crud/crud.component */ "./src/app/crud/crud.component.ts");
/* harmony import */ var _app_routing_module__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ./app-routing.module */ "./src/app/app-routing.module.ts");
/* harmony import */ var _amai_service__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ./amai.service */ "./src/app/amai.service.ts");
/* harmony import */ var _panel_panel_component__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ./panel/panel.component */ "./src/app/panel/panel.component.ts");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");
/* harmony import */ var _angular_material_form_field__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! @angular/material/form-field */ "./node_modules/@angular/material/esm5/form-field.es5.js");
/* harmony import */ var _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! @angular/cdk/layout */ "./node_modules/@angular/cdk/esm5/layout.es5.js");
/* harmony import */ var _main_nav_main_nav_component__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ./main-nav/main-nav.component */ "./src/app/main-nav/main-nav.component.ts");
/* harmony import */ var _users_table_users_table_component__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! ./users-table/users-table.component */ "./src/app/users-table/users-table.component.ts");
/* harmony import */ var _create_user_create_user_component__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ./create-user/create-user.component */ "./src/app/create-user/create-user.component.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};

















var AppModule = /** @class */ (function () {
    function AppModule() {
    }
    AppModule = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_2__["NgModule"])({
            declarations: [
                _app_component__WEBPACK_IMPORTED_MODULE_5__["AppComponent"],
                _home_home_component__WEBPACK_IMPORTED_MODULE_6__["HomeComponent"],
                _crud_crud_component__WEBPACK_IMPORTED_MODULE_7__["CrudComponent"],
                _panel_panel_component__WEBPACK_IMPORTED_MODULE_10__["PanelComponent"],
                _main_nav_main_nav_component__WEBPACK_IMPORTED_MODULE_14__["MainNavComponent"],
                _users_table_users_table_component__WEBPACK_IMPORTED_MODULE_15__["UsersTableComponent"],
                _create_user_create_user_component__WEBPACK_IMPORTED_MODULE_16__["CreateUserComponent"]
            ],
            imports: [
                _angular_platform_browser__WEBPACK_IMPORTED_MODULE_0__["BrowserModule"],
                _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_1__["BrowserAnimationsModule"],
                _angular_forms__WEBPACK_IMPORTED_MODULE_3__["FormsModule"],
                _app_routing_module__WEBPACK_IMPORTED_MODULE_8__["AppRoutingModule"],
                _angular_common_http__WEBPACK_IMPORTED_MODULE_4__["HttpClientModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatGridListModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatCardModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatMenuModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatIconModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatButtonModule"],
                _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_13__["LayoutModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatToolbarModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatSidenavModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatListModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatTableModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatPaginatorModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatSortModule"],
                _angular_material_form_field__WEBPACK_IMPORTED_MODULE_12__["MatFormFieldModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatInputModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatSelectModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_11__["MatOptionModule"]
            ],
            providers: [_amai_service__WEBPACK_IMPORTED_MODULE_9__["AmaiService"]],
            bootstrap: [_app_component__WEBPACK_IMPORTED_MODULE_5__["AppComponent"]]
        })
    ], AppModule);
    return AppModule;
}());



/***/ }),

/***/ "./src/app/create-user/create-user.component.css":
/*!*******************************************************!*\
  !*** ./src/app/create-user/create-user.component.css ***!
  \*******************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ".container{\r\n    margin: 20px;\r\n    display: -ms-grid;\r\n    display: grid;\r\n        grid-template-areas: \"card-create card-show\";\r\n    grid-gap: 20px;\r\n}\r\n\r\n@media only screen and (max-width: 800px) {\r\n    .container{\r\n        margin: 20px;\r\n        display: -ms-grid;\r\n        display: grid;\r\n            grid-template-areas: \"card-create\"\r\n                             \"card-show\";\r\n        grid-gap: 20px;\r\n    }\r\n    .card-create{\r\n        -ms-grid-row: 1;\r\n        -ms-grid-column: 1;\r\n    }\r\n    .card-show{\r\n        -ms-grid-row: 2;\r\n        -ms-grid-column: 1;\r\n    }\r\n}\r\n\r\n.form{\r\n    display: -ms-grid;\r\n    display: grid;\r\n}\r\n\r\n.card-create{\r\n    -ms-grid-row: 1;\r\n    -ms-grid-column: 1;\r\n    grid-area: card-create;\r\n    display: -ms-grid;\r\n    display: grid;\r\n    max-width: 400px;\r\n}\r\n\r\ninput::-webkit-outer-spin-button,\r\ninput::-webkit-inner-spin-button {\r\n    display: none;\r\n}\r\n\r\ninput{\r\n    -moz-appearance: textfield;\r\n}\r\n\r\n.mat-form-field{\r\n    margin: 7px 7px;\r\n\r\n}\r\n\r\n.card-show{\r\n    -ms-grid-row: 1;\r\n    -ms-grid-column: 2;\r\n    grid-area: card-show;\r\n    display: -ms-grid;\r\n    display: grid;\r\n    max-width: 400px;\r\n}\r\n\r\n#id-search{\r\n    text-align: right;\r\n    width: 25%;\r\n    position: absolute;\r\n    top: -20px;\r\n    right: 10px;\r\n}\r\n\r\n.btn{\r\n    margin: 20px 5px;\r\n}\r\n"

/***/ }),

/***/ "./src/app/create-user/create-user.component.html":
/*!********************************************************!*\
  !*** ./src/app/create-user/create-user.component.html ***!
  \********************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<div class=\"container\">\n<mat-card class=\"card-create\">\n  <mat-card-content>\n    <mat-card-title> Crear usuario </mat-card-title>\n<form action=\"\" method=\"POST\" accept-charset=\"utf-8\">\n    <div class=\"form\">\n        <!-- nombre -->\n      <mat-form-field id=\"name\" hintLabel=\"Max 40 characters\">\n        <input [(ngModel)]=\"user.UserName\" name=\"name\" matInput #input maxlength=\"40\" placeholder=\"Nombre\">\n        <mat-hint align=\"end\">{{input.value?.length || 0}}/40</mat-hint>\n      </mat-form-field>\n\n        <!-- genero -->\n      <mat-form-field>\n        <mat-select  [(ngModel)]=\"user.Gender\" name=\"gender\" id=\"gender\" placeholder=\"Genero\">\n          <mat-option value=\"Hombre\">Hombre</mat-option>\n         <mat-option value=\"Mujer\">Mujer</mat-option>\n        </mat-select>\n        <mat-hint align=\"end\">Selecciona tu genero ^</mat-hint>\n      </mat-form-field>\n\n        <!-- edad -->\n      <mat-form-field id=\"age\" hintLabel=\"\">\n        <input  [(ngModel)]=\"user.Age\" name=\"age\" matInput #input type=\"number\" min=\"0\" max=\"120\" placeholder=\"0\">\n        <mat-hint align=\"end\">{{input.value?.length || 0}}/120</mat-hint>\n      </mat-form-field>\n\n        <!-- Email -->\n      <mat-form-field id=\"email\" hintLabel=\"Max 40 characters\">\n        <input  [(ngModel)]=\"user.UserEmail\" name=\"email\" matInput #input maxlength=\"40\" type=\"email\" placeholder=\"Email@example.com\">\n        <mat-hint align=\"end\">{{input.value?.length || 0}}/40</mat-hint>\n      </mat-form-field>\n\n        <!-- Niverl academico -->\n      <mat-form-field>\n        <mat-select  [(ngModel)]=\"user.AcademicLevel\" name=\"academicLevel\" id=\"academicLevel\" placeholder=\"Nivel educativo\">\n          <mat-option value=\"Educacion Media\">Educacion media</mat-option>\n          <mat-option value=\"Bachillerato\">Bachillerato</mat-option>\n         <mat-option value=\"Universidad\">Universidad</mat-option>\n        </mat-select>\n        <mat-hint align=\"end\">Selecciona tu genero ^</mat-hint>\n      </mat-form-field>\n\n        <button class=\"btn\" mat-flat-button \n            (click)=\"createUser(user.UserName,user.Age,user.Gender,user.UserEmail, user.AcademicLevel,user.Role)\" \n            type=\"submit\">Crear</button>\n    </div> <!-- from div -->\n    </form>\n\n  </mat-card-content>\n</mat-card>\n\n<mat-card class=\"card-show\">\n    <mat-card-title> Buscar usuario \n      <mat-form-field id=\"id-search\" hintLabel=\"\">\n        <input matInput placeholder=\"id\">\n      </mat-form-field>\n    </mat-card-title>\n    <mat-label> Nombre: </mat-label>\n    <mat-label> Genero: </mat-label>\n    <mat-label> Edad: </mat-label>\n    <mat-label> Email: </mat-label>\n    <mat-label> Nivel Educativo: </mat-label>\n</mat-card>\n\n</div>\n"

/***/ }),

/***/ "./src/app/create-user/create-user.component.ts":
/*!******************************************************!*\
  !*** ./src/app/create-user/create-user.component.ts ***!
  \******************************************************/
/*! exports provided: CreateUserComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CreateUserComponent", function() { return CreateUserComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _amai_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../amai.service */ "./src/app/amai.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var CreateUserComponent = /** @class */ (function () {
    function CreateUserComponent(service) {
        this.service = service;
        this.user = {
            UserName: '',
            Age: 0,
            Gender: '',
            UserEmail: '',
            AcademicLevel: '',
            Role: 'Estudiante',
        };
    }
    /*
@Output()
created = new EventEmitter<string>();
     */
    CreateUserComponent.prototype.createUser = function (name, age, gender, userEmail, aLevel, role) {
        console.log("on create");
        console.log(name, age, gender, userEmail, aLevel, role);
        this.service.createUser(name, age, gender, userEmail, aLevel, role);
        //reset placeholder
        this.user = {
            UserName: '',
            Age: 0,
            Gender: '',
            UserEmail: '',
            AcademicLevel: '',
            Role: 'Estudiante',
        };
        console.log(this.user);
        //this.created.emit('created');
    };
    CreateUserComponent.prototype.ngOnInit = function () {
    };
    CreateUserComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-create-user',
            template: __webpack_require__(/*! ./create-user.component.html */ "./src/app/create-user/create-user.component.html"),
            styles: [__webpack_require__(/*! ./create-user.component.css */ "./src/app/create-user/create-user.component.css")]
        }),
        __metadata("design:paramtypes", [_amai_service__WEBPACK_IMPORTED_MODULE_1__["AmaiService"]])
    ], CreateUserComponent);
    return CreateUserComponent;
}());



/***/ }),

/***/ "./src/app/crud/crud.component.css":
/*!*****************************************!*\
  !*** ./src/app/crud/crud.component.css ***!
  \*****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ""

/***/ }),

/***/ "./src/app/crud/crud.component.html":
/*!******************************************!*\
  !*** ./src/app/crud/crud.component.html ***!
  \******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<h2> Users List </h2>\r\n<app-users-table>\r\n</app-users-table>\r\n\r\n<!--\r\n  <h1 class=\"mat-h1\">Users</h1>\r\n<ul *ngIf=\"users\"> <!-- Angular if, if items (object in home.component.ts) render this. -->\r\n<!--\r\n    <li *ngFor=\"let user of users\"> <!-- Angular for -->\r\n<!--\r\n  <h1 class=\"mat-h1\">Users</h1>\r\n      {{user.UserName}}\r\n      {{user.UserAge}}\r\n      {{user.UserEmail}}\r\n      {{user.ID}}\r\n      <form  action=\"\" [ngFormOptions]=\"{ updateOn: 'submit' }\" #myForm=\"ngForm\" method=\"POST\" accept-charset=\"utf-8\">\r\n          <button class=\"btn\" type=\"submit\" (click)=\"onSubmit(user.ID)\">delete </button>\r\n      </form>\r\n   </li>\r\n</ul>\r\n    <app-create-product (created)=\"refresh($event)\"></app-create-product>\r\n-->\r\n"

/***/ }),

/***/ "./src/app/crud/crud.component.ts":
/*!****************************************!*\
  !*** ./src/app/crud/crud.component.ts ***!
  \****************************************/
/*! exports provided: CrudComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CrudComponent", function() { return CrudComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _amai_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../amai.service */ "./src/app/amai.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var CrudComponent = /** @class */ (function () {
    function CrudComponent(service) {
        this.service = service;
    }
    CrudComponent.prototype.onSubmit = function (id) {
        console.log("on submit");
        this.service.deleteProduct(id);
        this.update();
    };
    CrudComponent.prototype.refresh = function ($event) {
        this.update();
    };
    CrudComponent.prototype.update = function () {
        var _this = this;
        this.service.getProductsObservable().subscribe(function (res) {
            _this.users = res;
        });
    };
    CrudComponent.prototype.ngOnInit = function () {
        this.update();
    };
    CrudComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-crud',
            template: __webpack_require__(/*! ./crud.component.html */ "./src/app/crud/crud.component.html"),
            styles: [__webpack_require__(/*! ./crud.component.css */ "./src/app/crud/crud.component.css")],
        }),
        __metadata("design:paramtypes", [_amai_service__WEBPACK_IMPORTED_MODULE_1__["AmaiService"]])
    ], CrudComponent);
    return CrudComponent;
}());



/***/ }),

/***/ "./src/app/home/home.component.css":
/*!*****************************************!*\
  !*** ./src/app/home/home.component.css ***!
  \*****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ""

/***/ }),

/***/ "./src/app/home/home.component.html":
/*!******************************************!*\
  !*** ./src/app/home/home.component.html ***!
  \******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<ul *ngIf=\"users\"> <!-- Angular if, if items (object in home.component.ts) render this. -->\r\n    <li *ngFor=\"let user of users\"> <!-- Angular for -->\r\n      {{user.UserName}}\r\n      {{user.UserAge}}\r\n      {{user.UserEmail}}\r\n   </li>\r\n</ul>\r\n<p>home<p>\r\n"

/***/ }),

/***/ "./src/app/home/home.component.ts":
/*!****************************************!*\
  !*** ./src/app/home/home.component.ts ***!
  \****************************************/
/*! exports provided: HomeComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "HomeComponent", function() { return HomeComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _amai_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../amai.service */ "./src/app/amai.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var HomeComponent = /** @class */ (function () {
    function HomeComponent(service) {
        this.service = service;
    }
    HomeComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.service.getProductsObservable().subscribe(function (res) {
            _this.users = res;
        });
    };
    HomeComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-home',
            template: __webpack_require__(/*! ./home.component.html */ "./src/app/home/home.component.html"),
            styles: [__webpack_require__(/*! ./home.component.css */ "./src/app/home/home.component.css")]
        }),
        __metadata("design:paramtypes", [_amai_service__WEBPACK_IMPORTED_MODULE_1__["AmaiService"]])
    ], HomeComponent);
    return HomeComponent;
}());



/***/ }),

/***/ "./src/app/main-nav/main-nav.component.css":
/*!*************************************************!*\
  !*** ./src/app/main-nav/main-nav.component.css ***!
  \*************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n.toolbar{\n    display: -ms-grid;\n    display: grid;\n}\n\n.content{\n    display: -ms-grid;\n    display: grid;\n    grid-gap: 75px;\n}\n\n.sidenav-container {\n  height: 100%;\n}\n\n.sidenav {\n  width: 200px;\n}\n\n.mat-toolbar.mat-primary {\n  position: -webkit-sticky;\n  position: sticky;\n  top: 0;\n}\n"

/***/ }),

/***/ "./src/app/main-nav/main-nav.component.html":
/*!**************************************************!*\
  !*** ./src/app/main-nav/main-nav.component.html ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<mat-sidenav-container class=\"sidenav-container\">\n  <mat-sidenav\n    #drawer\n    class=\"sidenav\"\n    fixedInViewport=\"true\"\n    [attr.role]=\"(isHandset$ | async) ? 'dialog' : 'navigation'\"\n    [mode]=\"(isHandset$ | async) ? 'over' : 'side'\"\n    [opened]=\"!(isHandset$ | async)\">\n    <mat-toolbar color=\"primary\">Menu</mat-toolbar>\n    <mat-nav-list>\n      <a mat-list-item routerLink=\"/\">Users</a>\n      <a mat-list-item routerLink=\"/tools\">Tools</a>\n      <a mat-list-item routerLink=\"/panel\">Place Holder</a>\n    </mat-nav-list>\n  </mat-sidenav>\n  <mat-sidenav-content>\n      <div class=\"toolbar\">\n    <mat-toolbar color=\"primary\" style=\"z-index:1; position:fixed\">\n      <button\n        type=\"button\"\n        aria-label=\"Toggle sidenav\"\n        mat-icon-button\n        (click)=\"drawer.toggle()\"\n        *ngIf=\"isHandset$ | async\">\n        <mat-icon aria-label=\"Side nav toggle icon\">menu</mat-icon>\n      </button>\n      <span class=\"title-bar\">Admin Module</span>\n    </mat-toolbar>\n      </div>\n    <!-- Add Content Here -->\n      <div class=\"content\">\n    <router-outlet></router-outlet>\n      </div>\n  </mat-sidenav-content>\n</mat-sidenav-container>\n"

/***/ }),

/***/ "./src/app/main-nav/main-nav.component.ts":
/*!************************************************!*\
  !*** ./src/app/main-nav/main-nav.component.ts ***!
  \************************************************/
/*! exports provided: MainNavComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "MainNavComponent", function() { return MainNavComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/cdk/layout */ "./node_modules/@angular/cdk/esm5/layout.es5.js");
/* harmony import */ var rxjs_operators__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! rxjs/operators */ "./node_modules/rxjs/_esm5/operators/index.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var MainNavComponent = /** @class */ (function () {
    function MainNavComponent(breakpointObserver) {
        this.breakpointObserver = breakpointObserver;
        this.isHandset$ = this.breakpointObserver.observe(_angular_cdk_layout__WEBPACK_IMPORTED_MODULE_1__["Breakpoints"].Handset)
            .pipe(Object(rxjs_operators__WEBPACK_IMPORTED_MODULE_2__["map"])(function (result) { return result.matches; }));
    }
    MainNavComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-main-nav',
            template: __webpack_require__(/*! ./main-nav.component.html */ "./src/app/main-nav/main-nav.component.html"),
            styles: [__webpack_require__(/*! ./main-nav.component.css */ "./src/app/main-nav/main-nav.component.css")]
        }),
        __metadata("design:paramtypes", [_angular_cdk_layout__WEBPACK_IMPORTED_MODULE_1__["BreakpointObserver"]])
    ], MainNavComponent);
    return MainNavComponent;
}());



/***/ }),

/***/ "./src/app/panel/panel.component.css":
/*!*******************************************!*\
  !*** ./src/app/panel/panel.component.css ***!
  \*******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ".grid-container {\n  margin: 20px;\n}\n\n.dashboard-card {\n  position: absolute;\n  top: 15px;\n  left: 15px;\n  right: 15px;\n  bottom: 15px;\n}\n\n.more-button {\n  position: absolute;\n  top: 5px;\n  right: 10px;\n}\n\n.dashboard-card-content {\n  text-align: center;\n}"

/***/ }),

/***/ "./src/app/panel/panel.component.html":
/*!********************************************!*\
  !*** ./src/app/panel/panel.component.html ***!
  \********************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<div class=\"grid-container\">\n  <h1 class=\"mat-h1\">Dashboard</h1>\n  <mat-grid-list cols=\"2\" rowHeight=\"350px\">\n    <mat-grid-tile *ngFor=\"let card of cards\" [colspan]=\"card.cols\" [rowspan]=\"card.rows\">\n      <mat-card class=\"dashboard-card\">\n        <mat-card-header>\n          <mat-card-title>\n            {{card.title}}\n            <button mat-icon-button class=\"more-button\" [matMenuTriggerFor]=\"menu\" aria-label=\"Toggle menu\">\n              <mat-icon>more_vert</mat-icon>\n            </button>\n            <mat-menu #menu=\"matMenu\" xPosition=\"before\">\n              <button mat-menu-item>Expand</button>\n              <button mat-menu-item>Remove</button>\n            </mat-menu>\n          </mat-card-title>\n        </mat-card-header>\n        <mat-card-content class=\"dashboard-card-content\">\n          <div>Card Content Here</div>\n        </mat-card-content>\n      </mat-card>\n    </mat-grid-tile>\n  </mat-grid-list>\n</div>\n"

/***/ }),

/***/ "./src/app/panel/panel.component.ts":
/*!******************************************!*\
  !*** ./src/app/panel/panel.component.ts ***!
  \******************************************/
/*! exports provided: PanelComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "PanelComponent", function() { return PanelComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/cdk/layout */ "./node_modules/@angular/cdk/esm5/layout.es5.js");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var PanelComponent = /** @class */ (function () {
    /*
  cards = this.breakpointObserver.observe(Breakpoints.Handset).pipe(
    map(({ matches }) => {
      if (matches) {
        return [
          { title: 'Card 1', cols: 1, rows: 1 },
          { title: 'Card 2', cols: 1, rows: 1 },
          { title: 'Card 3', cols: 1, rows: 1 },
          { title: 'Card 4', cols: 1, rows: 1 }
        ];
      }

      return [
        { title: 'Card 1', cols: 2, rows: 1 },
        { title: 'Card 2', cols: 1, rows: 1 },
        { title: 'Card 3', cols: 1, rows: 2 },
        { title: 'Card 4', cols: 1, rows: 1 }
      ];
    })
  );
     */
    function PanelComponent(breakpointObserver) {
        this.breakpointObserver = breakpointObserver;
        /** Based on the screen size, switch from standard to one column per row */
        this.cards = [
            { title: 'Card 1', cols: 1, rows: 1 },
            { title: 'Card 2', cols: 1, rows: 1 },
            { title: 'Card 3', cols: 1, rows: 1 },
            { title: 'Card 4', cols: 1, rows: 1 }
        ];
    }
    PanelComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-panel',
            template: __webpack_require__(/*! ./panel.component.html */ "./src/app/panel/panel.component.html"),
            styles: [__webpack_require__(/*! ./panel.component.css */ "./src/app/panel/panel.component.css")]
        }),
        __metadata("design:paramtypes", [_angular_cdk_layout__WEBPACK_IMPORTED_MODULE_1__["BreakpointObserver"]])
    ], PanelComponent);
    return PanelComponent;
}());



/***/ }),

/***/ "./src/app/users-table/users-table-datasource.ts":
/*!*******************************************************!*\
  !*** ./src/app/users-table/users-table-datasource.ts ***!
  \*******************************************************/
/*! exports provided: UsersTableDataSource */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "UsersTableDataSource", function() { return UsersTableDataSource; });
/* harmony import */ var _angular_cdk_collections__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/cdk/collections */ "./node_modules/@angular/cdk/esm5/collections.es5.js");
var __extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();

/*
export interface UsersTableItem {
  name: string;
  id: number;
}
 */
/*
const EXAMPLE_DATA: UsersTableItem[] = [
  {id: 1, name: 'Hydrogen'},
  {id: 2, name: 'Helium'},
  {id: 3, name: 'Lithium'},
  {id: 4, name: 'Beryllium'},
  {id: 5, name: 'Boron'},
  {id: 6, name: 'Carbon'},
  {id: 7, name: 'Nitrogen'},
  {id: 8, name: 'Oxygen'},
  {id: 9, name: 'Fluorine'},
  {id: 10, name: 'Neon'},
  {id: 11, name: 'Sodium'},
  {id: 12, name: 'Magnesium'},
  {id: 13, name: 'Aluminum'},
  {id: 14, name: 'Silicon'},
  {id: 15, name: 'Phosphorus'},
  {id: 16, name: 'Sulfur'},
  {id: 17, name: 'Chlorine'},
  {id: 18, name: 'Argon'},
  {id: 19, name: 'Potassium'},
  {id: 20, name: 'Calcium'},
];
 */
/**
 * Data source for the UsersTable view. This class should
 * encapsulate all logic for fetching and manipulating the displayed data
 * (including sorting, pagination, and filtering).
 */
var UsersTableDataSource = /** @class */ (function (_super) {
    __extends(UsersTableDataSource, _super);
    function UsersTableDataSource(paginator, sort, service) {
        var _this = _super.call(this) || this;
        _this.paginator = paginator;
        _this.sort = sort;
        _this.service = service;
        _this.usersList = [];
        return _this;
    }
    /**
     * Connect this data source to the table. The table will only update when
     * the returned stream emits new items.
     * @returns A stream of the items to be rendered.
     */
    UsersTableDataSource.prototype.connect = function () {
        var _this = this;
        // Combine everything that affects the rendered data into one update
        // stream for the data-table to consume.
        this.service.getProductsObservable().subscribe(function (res) {
            _this.usersList = res;
        });
        return this.service.getProductsObservable();
        /*
        const dataMutations = [
          observableOf(this.usersList),
          this.paginator.page,
          this.sort.sortChange
        ];
    
        // Set the paginators length
        this.paginator.length = this.usersList.length;
    
        return merge(...dataMutations).pipe(map(() => {
          return this.getPagedData(this.getSortedData([...this.usersList]));
        }));
         */
    };
    /**
     *  Called when the table is being destroyed. Use this function, to clean up
     * any open connections or free any held resources that were set up during connect.
     */
    UsersTableDataSource.prototype.disconnect = function () { };
    /**
     * Paginate the data (client-side). If you're using server-side pagination,
    }
  
     * this would be replaced by requesting the appropriate data from the server.
     */
    UsersTableDataSource.prototype.getPagedData = function (data) {
        var startIndex = this.paginator.pageIndex * this.paginator.pageSize;
        return data.splice(startIndex, this.paginator.pageSize);
    };
    /**
     * Sort the data (client-side). If you're using server-side sorting,
     * this would be replaced by requesting the appropriate data from the server.
     */
    UsersTableDataSource.prototype.getSortedData = function (data) {
        var _this = this;
        if (!this.sort.active || this.sort.direction === '') {
            return data;
        }
        return data.sort(function (a, b) {
            var isAsc = _this.sort.direction === 'asc';
            switch (_this.sort.active) {
                case 'name': return compare(a.UserName, b.UserName, isAsc);
                case 'email': return compare(+a.UserEmail, +b.UserEmail, isAsc);
                default: return 0;
            }
        });
    };
    return UsersTableDataSource;
}(_angular_cdk_collections__WEBPACK_IMPORTED_MODULE_0__["DataSource"]));

/** Simple sort comparator for example ID/Name columns (for client-side sorting). */
function compare(a, b, isAsc) {
    return (a < b ? -1 : 1) * (isAsc ? 1 : -1);
}


/***/ }),

/***/ "./src/app/users-table/users-table.component.css":
/*!*******************************************************!*\
  !*** ./src/app/users-table/users-table.component.css ***!
  \*******************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ".container{\r\n    margin: 20px;\r\n}\r\n"

/***/ }),

/***/ "./src/app/users-table/users-table.component.html":
/*!********************************************************!*\
  !*** ./src/app/users-table/users-table.component.html ***!
  \********************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<div class=\"container\">\n  <h1 class=\"mat-h1\">Usuarios</h1>\n<div class=\"mat-elevation-z2\">\n  <table mat-table #table [dataSource]=\"dataSource\" matSort style=\"width: 100%;\">\n\n    <!-- button -->\n    <ng-container matColumnDef=\"Options\">\n     <th mat-header-cell *matHeaderCellDef mat-sort-header> Options </th>\n      <td mat-cell *matCellDef=\"let row\" >\n          <button mat-flat-button> <mat-icon>delete</mat-icon> </button>\n      </td>\n    </ng-container>\n\n    <!-- Id Column -->\n    <ng-container matColumnDef=\"IdUser\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>ID</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.Id}}</td>\n    </ng-container>\n\n    <!-- Role Column -->\n    <ng-container matColumnDef=\"UserRole\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>Role</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.Role}}</td>\n    </ng-container>\n\n    <!-- Name Column -->\n    <ng-container matColumnDef=\"UserName\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>Name</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.UserName}}</td>\n    </ng-container>\n\n    <!-- Gender Column -->\n    <ng-container matColumnDef=\"UserGender\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>Gender</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.Gender}}</td>\n    </ng-container>\n\n    <!-- Email Column -->\n    <ng-container matColumnDef=\"UserEmail\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>Email</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.UserEmail}}</td>\n    </ng-container>\n\n    <!-- Academic Level Column -->\n    <ng-container matColumnDef=\"UserAcademicLevel\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>Academic Level</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.AcademicLevel}}</td>\n    </ng-container>\n\n    <!-- Rate Column -->\n    <ng-container matColumnDef=\"UserRate\">\n      <th mat-header-cell *matHeaderCellDef mat-sort-header>Rate</th>\n      <td mat-cell *matCellDef=\"let row\">{{row.Rate}}</td>\n    </ng-container>\n\n    <tr mat-header-row *matHeaderRowDef=\"displayedColumns\"></tr>\n    <tr mat-row *matRowDef=\"let row; columns: displayedColumns;\"></tr>\n  </table>\n\n  <mat-paginator #paginator\n    [length]=\"dataSource?.data?.length\"\n    [pageIndex]=\"0\"\n    [pageSize]=\"50\"\n    [pageSizeOptions]=\"[25, 50, 100, 250]\">\n  </mat-paginator>\n</div>\n</div>\n"

/***/ }),

/***/ "./src/app/users-table/users-table.component.ts":
/*!******************************************************!*\
  !*** ./src/app/users-table/users-table.component.ts ***!
  \******************************************************/
/*! exports provided: UsersTableComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "UsersTableComponent", function() { return UsersTableComponent; });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");
/* harmony import */ var _users_table_datasource__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./users-table-datasource */ "./src/app/users-table/users-table-datasource.ts");
/* harmony import */ var _amai_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../amai.service */ "./src/app/amai.service.ts");
var __decorate = (undefined && undefined.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (undefined && undefined.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};




var UsersTableComponent = /** @class */ (function () {
    function UsersTableComponent(service) {
        this.service = service;
        /** Columns displayed in the table. Columns IDs can be added, removed, or reordered. */
        this.displayedColumns = ['Options', 'IdUser', 'UserRole', 'UserName',
            'UserGender', 'UserEmail', 'UserAcademicLevel', 'UserRate'];
    }
    UsersTableComponent.prototype.ngOnInit = function () {
        this.dataSource = new _users_table_datasource__WEBPACK_IMPORTED_MODULE_2__["UsersTableDataSource"](this.paginator, this.sort, this.service);
    };
    __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["ViewChild"])(_angular_material__WEBPACK_IMPORTED_MODULE_1__["MatPaginator"]),
        __metadata("design:type", _angular_material__WEBPACK_IMPORTED_MODULE_1__["MatPaginator"])
    ], UsersTableComponent.prototype, "paginator", void 0);
    __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["ViewChild"])(_angular_material__WEBPACK_IMPORTED_MODULE_1__["MatSort"]),
        __metadata("design:type", _angular_material__WEBPACK_IMPORTED_MODULE_1__["MatSort"])
    ], UsersTableComponent.prototype, "sort", void 0);
    UsersTableComponent = __decorate([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["Component"])({
            selector: 'app-users-table',
            template: __webpack_require__(/*! ./users-table.component.html */ "./src/app/users-table/users-table.component.html"),
            styles: [__webpack_require__(/*! ./users-table.component.css */ "./src/app/users-table/users-table.component.css")]
        }),
        __metadata("design:paramtypes", [_amai_service__WEBPACK_IMPORTED_MODULE_3__["AmaiService"]])
    ], UsersTableComponent);
    return UsersTableComponent;
}());



/***/ }),

/***/ "./src/environments/environment.ts":
/*!*****************************************!*\
  !*** ./src/environments/environment.ts ***!
  \*****************************************/
/*! exports provided: environment */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "environment", function() { return environment; });
// This file can be replaced during build by using the `fileReplacements` array.
// `ng build ---prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.
var environment = {
    production: false
};
/*
 * In development mode, to ignore zone related error stack frames such as
 * `zone.run`, `zoneDelegate.invokeTask` for easier debugging, you can
 * import the following file, but please comment it out in production mode
 * because it will have performance impact when throw error
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.


/***/ }),

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/*! no exports provided */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/platform-browser-dynamic */ "./node_modules/@angular/platform-browser-dynamic/fesm5/platform-browser-dynamic.js");
/* harmony import */ var _app_app_module__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./app/app.module */ "./src/app/app.module.ts");
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./environments/environment */ "./src/environments/environment.ts");




if (_environments_environment__WEBPACK_IMPORTED_MODULE_3__["environment"].production) {
    Object(_angular_core__WEBPACK_IMPORTED_MODULE_0__["enableProdMode"])();
}
Object(_angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_1__["platformBrowserDynamic"])().bootstrapModule(_app_app_module__WEBPACK_IMPORTED_MODULE_2__["AppModule"])
    .catch(function (err) { return console.log(err); });


/***/ }),

/***/ 0:
/*!***************************!*\
  !*** multi ./src/main.ts ***!
  \***************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(/*! C:\Users\developer\go\src\github.com\pdmp\amai\bento\src\main.ts */"./src/main.ts");


/***/ })

},[[0,"runtime","vendor"]]]);
//# sourceMappingURL=main.js.map