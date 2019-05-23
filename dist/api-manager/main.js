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
		var e = new Error("Cannot find module '" + req + "'");
		e.code = 'MODULE_NOT_FOUND';
		throw e;
	});
}
webpackEmptyAsyncContext.keys = function() { return []; };
webpackEmptyAsyncContext.resolve = webpackEmptyAsyncContext;
module.exports = webpackEmptyAsyncContext;
webpackEmptyAsyncContext.id = "./src/$$_lazy_route_resource lazy recursive";

/***/ }),

/***/ "./src/app/api-list/api-list.component.html":
/*!**************************************************!*\
  !*** ./src/app/api-list/api-list.component.html ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<div fxLayout=\"row wrap\">\n  <div *ngFor=\"let API of APIs\" [hidden]=\"API.hide\"  fxFlex=\"100%\">\n    <mat-list>\n    <api-card [API]=\"API\"></api-card>\n    <mat-divider></mat-divider>\n    </mat-list>\n  </div>\n</div>\n"

/***/ }),

/***/ "./src/app/api-list/api-list.component.scss":
/*!**************************************************!*\
  !*** ./src/app/api-list/api-list.component.scss ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2FwaS1saXN0L2FwaS1saXN0LmNvbXBvbmVudC5zY3NzIn0= */"

/***/ }),

/***/ "./src/app/api-list/api-list.component.ts":
/*!************************************************!*\
  !*** ./src/app/api-list/api-list.component.ts ***!
  \************************************************/
/*! exports provided: ApiListComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "ApiListComponent", function() { return ApiListComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _api_proxy_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../api-proxy.service */ "./src/app/api-proxy.service.ts");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");




var ApiListComponent = /** @class */ (function () {
    function ApiListComponent(apiProxyService, dialog) {
        var _this = this;
        this.dialog = dialog;
        this.APIs = [];
        this.APIs$ = apiProxyService.getAPIProxies();
        this.APIs$.subscribe(function (res) { return _this.APIs = res; });
    }
    ApiListComponent.prototype.ngOnChanges = function (changes) {
        if (changes["search"] || changes["tenant"]) {
            for (var i = 0; i < this.APIs.length; ++i) {
                this.APIs[i].hide = false;
                if (!this.APIs[i].apiName.toUpperCase().includes(this.search) &&
                    !this.APIs[i].apiURL.toUpperCase().includes(this.search) &&
                    !this.APIs[i].apiTargetDestination.toUpperCase().includes(this.search) &&
                    !this.APIs[i].apiTargetEndpoint.toUpperCase().includes(this.search) || (this.tenant && this.tenant !== this.APIs[i].maxTenant())) {
                    this.APIs[i].hide = true;
                }
            }
        }
    };
    tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Input"])(),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:type", String)
    ], ApiListComponent.prototype, "search", void 0);
    tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Input"])(),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:type", String)
    ], ApiListComponent.prototype, "tenant", void 0);
    ApiListComponent = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'app-api-list',
            template: __webpack_require__(/*! ./api-list.component.html */ "./src/app/api-list/api-list.component.html"),
            styles: [__webpack_require__(/*! ./api-list.component.scss */ "./src/app/api-list/api-list.component.scss")]
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_api_proxy_service__WEBPACK_IMPORTED_MODULE_2__["APIProxyService"],
            _angular_material__WEBPACK_IMPORTED_MODULE_3__["MatDialog"]])
    ], ApiListComponent);
    return ApiListComponent;
}());



/***/ }),

/***/ "./src/app/api-proxy.service.ts":
/*!**************************************!*\
  !*** ./src/app/api-proxy.service.ts ***!
  \**************************************/
/*! exports provided: APIProxy, APIProxyService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "APIProxy", function() { return APIProxy; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "APIProxyService", function() { return APIProxyService; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common/http */ "./node_modules/@angular/common/fesm5/http.js");
/* harmony import */ var rxjs_operators__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! rxjs/operators */ "./node_modules/rxjs/_esm5/operators/index.js");
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! rxjs */ "./node_modules/rxjs/_esm5/index.js");





var APIProxy = /** @class */ (function () {
    function APIProxy(name, url, destination, endpoint, inQA, inProduction) {
        this.prodColor = "secondary";
        this.qaColor = "secondary";
        this.inDev = true;
        this.devColor = "primary";
        this.working = false;
        this.style = "display: none;";
        this.hide = false;
        this.apiName = name;
        this.apiURL = url;
        this.apiTargetDestination = destination;
        this.apiTargetEndpoint = endpoint;
        this.inQA = inQA;
        this.qaColor = this.inQA ? 'primary' : 'secondary';
        this.testedInDev = this.inQA ? true : false;
        this.inProduction = inProduction;
        this.prodColor = this.inProduction ? 'primary' : 'secondary';
        this.testedInQA = this.inProduction ? true : false;
        console.log(this.testedInDev);
    }
    APIProxy.prototype.toggleHide = function () {
        this.hide = !this.hide;
    };
    APIProxy.prototype.moveToProduction = function () {
        var _this = this;
        if (!this.inProduction) {
            if (this.working) {
                return;
            }
            this.working = true;
            setTimeout(function () {
                _this.inProduction = true;
                _this.prodColor = "primary";
                _this.working = false;
            }, 2000);
        }
    };
    APIProxy.prototype.maxTenant = function () {
        if (this.inProduction) {
            return "Production";
        }
        if (this.inQA) {
            return "QA";
        }
        return "Development";
    };
    APIProxy.prototype.setQA = function () {
        this.inQA = true;
        this.qaColor = "primary";
        this.testedInDev = true;
    };
    APIProxy.prototype.setProduction = function () {
        this.inProduction = true;
        this.prodColor = "primary";
        this.testedInQA = true;
    };
    APIProxy.prototype.moveToQA = function () {
        var _this = this;
        if (!this.inQA) {
            if (this.working) {
                return;
            }
            this.working = true;
            setTimeout(function () {
                _this.working = false;
                _this.qaColor = "primary";
                _this.inQA = true;
            }, 2000);
        }
    };
    APIProxy.prototype.setStyles = function () {
        var styles = {
            'display': this.working ? 'inline' : 'none',
            'text-align': 'center'
        };
        return styles;
    };
    return APIProxy;
}());

var APIProxyService = /** @class */ (function () {
    function APIProxyService(http) {
        this.http = http;
        this.APIM_HOST_MAP = {
            'dev': 'dfc3ccb1f',
            'qa': 'd8b3bfb89',
            'sandbox': 'd6c83d68e',
        };
        this.API_MAP = {};
        this.APIM_DEV_HOST = "https://produs2apiportalapimgmtpphx-dfc3ccb1f.us2.hana.ondemand.com/apiportal/api/1.0/";
        this.APIM_PROXY_PATH = "Management.svc/APIProxyEndPoints";
        this.APIM_TARGET_PATH = "Management.svc/APITargetEndPoints";
    }
    APIProxyService.prototype.getAPIProxies = function () {
        var APIEndpoints = this.http.get('/api/devapi/' + this.APIM_PROXY_PATH);
        var APITargets = this.http.get('/api/devapi/' + this.APIM_TARGET_PATH);
        var APIEndpointsQA = this.http.get('/api/qaapi/' + this.APIM_PROXY_PATH);
        var APIEndpointsProduction = this.http.get('/api/prodapi/' + this.APIM_PROXY_PATH);
        return Object(rxjs__WEBPACK_IMPORTED_MODULE_4__["forkJoin"])([APIEndpoints, APITargets, APIEndpointsQA, APIEndpointsProduction]).pipe(
        //return forkJoin([APIEndpoints, APITargets, APIEndpointsQA]).pipe(
        Object(rxjs_operators__WEBPACK_IMPORTED_MODULE_3__["map"])(function (results) {
            var result = [];
            var resp1 = results[0]['d']['results'];
            var resp2 = results[1]['d']['results'];
            var resp3 = results[2]['d']['results'];
            var resp4 = results[3]['d']['results'];
            var resultMap = {};
            for (var i in resp1) {
                resultMap[resp1[i]['FK_API_NAME']] =
                    new APIProxy(resp1[i]['FK_API_NAME'], resp1[i]['base_path'], "", "", false, false);
            }
            for (var _i = 0, resp2_1 = resp2; _i < resp2_1.length; _i++) {
                var api = resp2_1[_i];
                if (resultMap[api['FK_API_NAME']]) {
                    resultMap[api['FK_API_NAME']].apiTargetDestination =
                        (api['provider_id'] == 'NONE' ? 'URL' : api['provider_id']);
                    resultMap[api['FK_API_NAME']].apiTargetEndpoint =
                        (api['relativePath'] == null ? api['url'] : api['relativePath']);
                }
            }
            for (var _a = 0, _b = results[2]['d']['results']; _a < _b.length; _a++) {
                var api = _b[_a];
                if (resultMap[api['FK_API_NAME']]) {
                    resultMap[api['FK_API_NAME']].setQA();
                }
            }
            for (var _c = 0, _d = results[3]['d']['results']; _c < _d.length; _c++) {
                var api = _d[_c];
                if (resultMap[api['FK_API_NAME']]) {
                    resultMap[api['FK_API_NAME']].setProduction();
                }
            }
            return Object.values(resultMap);
        }));
    };
    APIProxyService = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Injectable"])({
            providedIn: 'root'
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_common_http__WEBPACK_IMPORTED_MODULE_2__["HttpClient"]])
    ], APIProxyService);
    return APIProxyService;
}());



/***/ }),

/***/ "./src/app/api-test.service.ts":
/*!*************************************!*\
  !*** ./src/app/api-test.service.ts ***!
  \*************************************/
/*! exports provided: ApiTestService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "ApiTestService", function() { return ApiTestService; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common/http */ "./node_modules/@angular/common/fesm5/http.js");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");




var ApiTestService = /** @class */ (function () {
    function ApiTestService(http, snackBar) {
        this.http = http;
        this.snackBar = snackBar;
    }
    ApiTestService.prototype.Test = function (API, tenant, token, metaDataPath) {
        return tslib__WEBPACK_IMPORTED_MODULE_0__["__awaiter"](this, void 0, void 0, function () {
            var payload, response;
            var _this = this;
            return tslib__WEBPACK_IMPORTED_MODULE_0__["__generator"](this, function (_a) {
                switch (_a.label) {
                    case 0:
                        payload = {};
                        payload['Token'] = token;
                        payload['Path'] = metaDataPath;
                        return [4 /*yield*/, this.http.post("/api/test/" + tenant + "/" + API.apiName, JSON.stringify(payload)).toPromise()
                                .then(function (data) {
                                return data;
                            })
                                .catch(function (error) {
                                _this.snackBar.open("Error! test failed with status code " + error.error, null, {
                                    duration: 3000,
                                    panelClass: ["failed-snackbar"]
                                });
                                return null;
                            })];
                    case 1:
                        response = _a.sent();
                        return [2 /*return*/, response];
                }
            });
        });
    };
    ApiTestService = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Injectable"])({
            providedIn: 'root'
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_common_http__WEBPACK_IMPORTED_MODULE_2__["HttpClient"],
            _angular_material__WEBPACK_IMPORTED_MODULE_3__["MatSnackBar"]])
    ], ApiTestService);
    return ApiTestService;
}());



/***/ }),

/***/ "./src/app/api/api.component.html":
/*!****************************************!*\
  !*** ./src/app/api/api.component.html ***!
  \****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<mat-list-item>\n  <h1 style=\"font-weight: bold;\"matLine>{{this.API.apiName}}</h1>\n  <h4 matLine class=\"demo-2\">URL: {{this.API.apiURL}} </h4>\n  <h4 matLine> Target: {{this.API.apiTargetDestination}}: {{this.API.apiTargetEndpoint}}</h4>\n  <div style=\"padding-top: 1rem; padding-bottom: 0.25rem\"matLine>\n    <button [color]=\"this.API.devColor\"\n            (click)=\"openDialog(this.API, 'DEV', !this.API.inDev)\"\n            style=\"display: inline; margin-right: 1rem;\"\n            mat-raised-button>\n      {{this.API.inDev ? \"Test in Development\" : \"Move to Development\"}}\n    </button>\n\n    <button (click)=\"openDialog(this.API, 'QA', !this.API.inQA)\"\n            [color]=\"this.API.qaColor\"\n            style=\"display: inline; margin-right: 1rem;\"\n            mat-raised-button>\n      {{this.API.inQA ? \"Test in QA\" : \"Move to QA\"}}\n    </button>\n\n    <button\n      (click)=\"this.API.inProduction ? null : openDialog(this.API, 'PROD', !this.API.inProd)\"\n      [color]=\"this.API.prodColor\"\n      [disabled]=\"!this.API.inQA\"\n      [disableRipple]=\"this.API.inProduction\"\n      style=\"display: inline; margin-right: 1rem;\"\n      mat-raised-button>\n\n      {{this.API.inProduction ? \"In Production\" : \"Move to Production\"}}\n    </button>\n    <mat-spinner\n      color=\"accent\"\n      [ngStyle]=\"this.API.setStyles()\"\n      [diameter]=\"28\"></mat-spinner>\n  </div>\n</mat-list-item>\n"

/***/ }),

/***/ "./src/app/api/api.component.scss":
/*!****************************************!*\
  !*** ./src/app/api/api.component.scss ***!
  \****************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2FwaS9hcGkuY29tcG9uZW50LnNjc3MifQ== */"

/***/ }),

/***/ "./src/app/api/api.component.ts":
/*!**************************************!*\
  !*** ./src/app/api/api.component.ts ***!
  \**************************************/
/*! exports provided: APIComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "APIComponent", function() { return APIComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _api_proxy_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../api-proxy.service */ "./src/app/api-proxy.service.ts");
/* harmony import */ var _test_dialog_test_dialog_component__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../test-dialog/test-dialog.component */ "./src/app/test-dialog/test-dialog.component.ts");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");





var APIComponent = /** @class */ (function () {
    function APIComponent(dialog) {
        this.dialog = dialog;
    }
    APIComponent.prototype.openDialog = function (API, environment, transport) {
        var _this = this;
        var dialogRef = this.dialog.open(_test_dialog_test_dialog_component__WEBPACK_IMPORTED_MODULE_3__["TestDialogComponent"], {
            width: '500px',
            data: { api: API, environment: environment, success: false, transport: transport }
        });
        dialogRef.afterClosed().subscribe(function (result) {
            console.log('The dialog was closed');
            if (result) {
                console.log(environment);
                if (environment === "QA") {
                    _this.API.testedInQA = true;
                    _this.API.inQA = true;
                    _this.API.qaColor = "primary";
                }
                else if (environment === "PROD") {
                    _this.API.inProduction = true;
                    _this.API.prodColor = "primary";
                }
            }
        });
    };
    tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Input"])(),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:type", _api_proxy_service__WEBPACK_IMPORTED_MODULE_2__["APIProxy"])
    ], APIComponent.prototype, "API", void 0);
    APIComponent = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'api-card',
            template: __webpack_require__(/*! ./api.component.html */ "./src/app/api/api.component.html"),
            styles: [__webpack_require__(/*! ./api.component.scss */ "./src/app/api/api.component.scss")]
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_material__WEBPACK_IMPORTED_MODULE_4__["MatDialog"]])
    ], APIComponent);
    return APIComponent;
}());



/***/ }),

/***/ "./src/app/apitransport.service.ts":
/*!*****************************************!*\
  !*** ./src/app/apitransport.service.ts ***!
  \*****************************************/
/*! exports provided: ApiTransportService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "ApiTransportService", function() { return ApiTransportService; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common/http */ "./node_modules/@angular/common/fesm5/http.js");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");




var ApiTransportService = /** @class */ (function () {
    function ApiTransportService(http, snackBar) {
        this.http = http;
        this.snackBar = snackBar;
    }
    ApiTransportService.prototype.Transport = function (API, tenant, token, metaDataPath, email, password) {
        return tslib__WEBPACK_IMPORTED_MODULE_0__["__awaiter"](this, void 0, void 0, function () {
            var payload, response;
            var _this = this;
            return tslib__WEBPACK_IMPORTED_MODULE_0__["__generator"](this, function (_a) {
                switch (_a.label) {
                    case 0:
                        payload = {};
                        payload['Token'] = token;
                        payload['Path'] = metaDataPath;
                        payload['Email'] = email;
                        payload['Password'] = password;
                        tenant = tenant === "PROD" ? "QA" : "DEV";
                        return [4 /*yield*/, this.http.post("/api/transport/" + tenant + "/" + API.apiName, JSON.stringify(payload)).toPromise()
                                .then(function (data) {
                                return data;
                            })
                                .catch(function (error) {
                                _this.snackBar.open("Error! Transport failed with status code " + error.error, null, {
                                    duration: 3000,
                                    panelClass: ["failed-snackbar"]
                                });
                                return null;
                            })];
                    case 1:
                        response = _a.sent();
                        return [2 /*return*/, response];
                }
            });
        });
    };
    ApiTransportService = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Injectable"])({
            providedIn: 'root'
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_common_http__WEBPACK_IMPORTED_MODULE_2__["HttpClient"],
            _angular_material__WEBPACK_IMPORTED_MODULE_3__["MatSnackBar"]])
    ], ApiTransportService);
    return ApiTransportService;
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
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/router */ "./node_modules/@angular/router/fesm5/router.js");



var routes = [];
var AppRoutingModule = /** @class */ (function () {
    function AppRoutingModule() {
    }
    AppRoutingModule = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["NgModule"])({
            imports: [_angular_router__WEBPACK_IMPORTED_MODULE_2__["RouterModule"].forRoot(routes)],
            exports: [_angular_router__WEBPACK_IMPORTED_MODULE_2__["RouterModule"]]
        })
    ], AppRoutingModule);
    return AppRoutingModule;
}());



/***/ }),

/***/ "./src/app/app.component.html":
/*!************************************!*\
  !*** ./src/app/app.component.html ***!
  \************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<app-main-nav></app-main-nav>\n"

/***/ }),

/***/ "./src/app/app.component.scss":
/*!************************************!*\
  !*** ./src/app/app.component.scss ***!
  \************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2FwcC5jb21wb25lbnQuc2NzcyJ9 */"

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
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");


var AppComponent = /** @class */ (function () {
    function AppComponent() {
        this.title = 'api-manager';
    }
    AppComponent = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'app-root',
            template: __webpack_require__(/*! ./app.component.html */ "./src/app/app.component.html"),
            styles: [__webpack_require__(/*! ./app.component.scss */ "./src/app/app.component.scss")]
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
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_platform_browser__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/platform-browser */ "./node_modules/@angular/platform-browser/fesm5/platform-browser.js");
/* harmony import */ var _angular_flex_layout__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/flex-layout */ "./node_modules/@angular/flex-layout/esm5/flex-layout.es5.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _app_routing_module__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./app-routing.module */ "./src/app/app-routing.module.ts");
/* harmony import */ var _app_component__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./app.component */ "./src/app/app.component.ts");
/* harmony import */ var _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/platform-browser/animations */ "./node_modules/@angular/platform-browser/fesm5/animations.js");
/* harmony import */ var _main_nav_main_nav_component__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ./main-nav/main-nav.component */ "./src/app/main-nav/main-nav.component.ts");
/* harmony import */ var _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @angular/cdk/layout */ "./node_modules/@angular/cdk/esm5/layout.es5.js");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @angular/forms */ "./node_modules/@angular/forms/fesm5/forms.js");
/* harmony import */ var _api_api_component__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ./api/api.component */ "./src/app/api/api.component.ts");
/* harmony import */ var _api_proxy_service__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ./api-proxy.service */ "./src/app/api-proxy.service.ts");
/* harmony import */ var _api_test_service__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ./api-test.service */ "./src/app/api-test.service.ts");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! @angular/common/http */ "./node_modules/@angular/common/fesm5/http.js");
/* harmony import */ var _test_dialog_test_dialog_component__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! ./test-dialog/test-dialog.component */ "./src/app/test-dialog/test-dialog.component.ts");
/* harmony import */ var _api_list_api_list_component__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ./api-list/api-list.component */ "./src/app/api-list/api-list.component.ts");
/* harmony import */ var _apitransport_service__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ./apitransport.service */ "./src/app/apitransport.service.ts");



















var AppModule = /** @class */ (function () {
    function AppModule() {
    }
    AppModule = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_3__["NgModule"])({
            declarations: [
                _app_component__WEBPACK_IMPORTED_MODULE_5__["AppComponent"],
                _main_nav_main_nav_component__WEBPACK_IMPORTED_MODULE_7__["MainNavComponent"],
                _api_api_component__WEBPACK_IMPORTED_MODULE_11__["APIComponent"],
                _test_dialog_test_dialog_component__WEBPACK_IMPORTED_MODULE_15__["TestDialogComponent"],
                _api_list_api_list_component__WEBPACK_IMPORTED_MODULE_16__["ApiListComponent"],
            ],
            imports: [
                _angular_platform_browser__WEBPACK_IMPORTED_MODULE_1__["BrowserModule"],
                _app_routing_module__WEBPACK_IMPORTED_MODULE_4__["AppRoutingModule"],
                _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_6__["BrowserAnimationsModule"],
                _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_8__["LayoutModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatToolbarModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatButtonModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatSidenavModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatIconModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatListModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatCardModule"],
                _angular_flex_layout__WEBPACK_IMPORTED_MODULE_2__["FlexLayoutModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatProgressSpinnerModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatFormFieldModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatInputModule"],
                _angular_forms__WEBPACK_IMPORTED_MODULE_10__["ReactiveFormsModule"],
                _angular_common_http__WEBPACK_IMPORTED_MODULE_14__["HttpClientModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatDialogModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatSelectModule"],
                _angular_forms__WEBPACK_IMPORTED_MODULE_10__["FormsModule"],
                _angular_material__WEBPACK_IMPORTED_MODULE_9__["MatSnackBarModule"],
                _angular_common_http__WEBPACK_IMPORTED_MODULE_14__["HttpClientXsrfModule"].withOptions({
                    cookieName: 'BIGipServerprodus2apiportalapimgmtpphx.us2.hana.ondemand.com	',
                }),
            ],
            providers: [_api_proxy_service__WEBPACK_IMPORTED_MODULE_12__["APIProxyService"], _api_test_service__WEBPACK_IMPORTED_MODULE_13__["ApiTestService"], _apitransport_service__WEBPACK_IMPORTED_MODULE_17__["ApiTransportService"]],
            bootstrap: [_app_component__WEBPACK_IMPORTED_MODULE_5__["AppComponent"]],
            entryComponents: [
                _test_dialog_test_dialog_component__WEBPACK_IMPORTED_MODULE_15__["TestDialogComponent"]
            ],
        })
    ], AppModule);
    return AppModule;
}());



/***/ }),

/***/ "./src/app/main-nav/main-nav.component.html":
/*!**************************************************!*\
  !*** ./src/app/main-nav/main-nav.component.html ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<mat-sidenav-container class=\"sidenav-container\" >\n  <!--\n  <mat-sidenav #drawer class=\"sidenav\" fixedInViewport=\"true\"\n      [attr.role]=\"(isHandset$ | async) ? 'dialog' : 'navigation'\"\n      [mode]=\"(isHandset$ | async) ? 'over' : 'side'\"\n      [opened]=\"!(isHandset$ | async)\">\n    <mat-toolbar>Menu</mat-toolbar>\n    <mat-nav-list>\n      <a mat-list-item href=\"#\">Link 1</a>\n      <a mat-list-item href=\"#\">Link 2</a>\n      <a mat-list-item href=\"#\">Link 3</a>\n    </mat-nav-list>\n  </mat-sidenav>\n-->\n  <mat-sidenav-content>\n\n    <mat-toolbar style=\"position: sticky; top: 0; z-index: 1; height: 100px;\"class=\"mat-elevation-z5\" color=\"primary\">\n\n      <!--\n      <button\n        type=\"button\"\n        aria-label=\"Toggle sidenav\"\n        mat-icon-button\n        (click)=\"drawer.toggle()\"\n        *ngIf=\"isHandset$ | async\">\n        <mat-icon aria-label=\"Side nav toggle icon\">menu</mat-icon>\n      </button> -->\n      <img class=\"img-fluid\" style=\"width: auto; height: 100%;\" src=\"assets/img/Nike-Swoosh.svg\">\n      <span>SAPAE APIs</span>\n      <div fxFlex=\"40\">\n      </div>\n      <div fxFlex>\n        <mat-form-field style=\"margin-right: 1rem;\">\n          <input  matInput placeholder=\"Search\" [formControl]=\"search\">\n        </mat-form-field>\n        <mat-form-field>\n          <mat-select [(value)]=\"tenant\" placeholder=\"Tenant\">\n            <mat-option>None</mat-option>\n            <mat-option value=\"Development\">Development</mat-option>\n            <mat-option value=\"QA\">QA</mat-option>\n            <mat-option value=\"Production\">Production</mat-option>\n          </mat-select>\n        </mat-form-field>\n      </div>\n\n    </mat-toolbar>\n\n    <app-api-list [search]=\"search.value.toUpperCase()\"\n                  [tenant]=\"tenant\"></app-api-list>\n\n  </mat-sidenav-content>\n</mat-sidenav-container>\n"

/***/ }),

/***/ "./src/app/main-nav/main-nav.component.scss":
/*!**************************************************!*\
  !*** ./src/app/main-nav/main-nav.component.scss ***!
  \**************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = ".sidenav-container {\n  height: 100%; }\n\n.sidenav {\n  width: 200px; }\n\n.sidenav .mat-toolbar {\n  background: inherit; }\n\n.mat-toolbar.mat-primary {\n  position: -webkit-sticky;\n  position: sticky;\n  top: 0;\n  z-index: 1; }\n\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9Vc2Vycy90YW5kcjkvRG9jdW1lbnRzL3RtcC9hbmd1bGFyL2FwaS1tYW5hZ2VyL3NyYy9hcHAvbWFpbi1uYXYvbWFpbi1uYXYuY29tcG9uZW50LnNjc3MiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUE7RUFDRSxZQUFZLEVBQUE7O0FBR2Q7RUFDRSxZQUFZLEVBQUE7O0FBR2Q7RUFDRSxtQkFBbUIsRUFBQTs7QUFHckI7RUFDRSx3QkFBZ0I7RUFBaEIsZ0JBQWdCO0VBQ2hCLE1BQU07RUFDTixVQUFVLEVBQUEiLCJmaWxlIjoic3JjL2FwcC9tYWluLW5hdi9tYWluLW5hdi5jb21wb25lbnQuc2NzcyIsInNvdXJjZXNDb250ZW50IjpbIi5zaWRlbmF2LWNvbnRhaW5lciB7XG4gIGhlaWdodDogMTAwJTtcbn1cblxuLnNpZGVuYXYge1xuICB3aWR0aDogMjAwcHg7XG59XG5cbi5zaWRlbmF2IC5tYXQtdG9vbGJhciB7XG4gIGJhY2tncm91bmQ6IGluaGVyaXQ7XG59XG5cbi5tYXQtdG9vbGJhci5tYXQtcHJpbWFyeSB7XG4gIHBvc2l0aW9uOiBzdGlja3k7XG4gIHRvcDogMDtcbiAgei1pbmRleDogMTtcbn1cbiJdfQ== */"

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
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_cdk_layout__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/cdk/layout */ "./node_modules/@angular/cdk/esm5/layout.es5.js");
/* harmony import */ var rxjs_operators__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! rxjs/operators */ "./node_modules/rxjs/_esm5/operators/index.js");
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/forms */ "./node_modules/@angular/forms/fesm5/forms.js");





var MainNavComponent = /** @class */ (function () {
    function MainNavComponent(breakpointObserver) {
        this.breakpointObserver = breakpointObserver;
        this.isHandset$ = this.breakpointObserver.observe(_angular_cdk_layout__WEBPACK_IMPORTED_MODULE_2__["Breakpoints"].Handset)
            .pipe(Object(rxjs_operators__WEBPACK_IMPORTED_MODULE_3__["map"])(function (result) { return result.matches; }));
        this.valueChange = new _angular_core__WEBPACK_IMPORTED_MODULE_1__["EventEmitter"]();
        this.search = new _angular_forms__WEBPACK_IMPORTED_MODULE_4__["FormControl"]('');
    }
    tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Output"])(),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:type", Object)
    ], MainNavComponent.prototype, "valueChange", void 0);
    MainNavComponent = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'app-main-nav',
            template: __webpack_require__(/*! ./main-nav.component.html */ "./src/app/main-nav/main-nav.component.html"),
            styles: [__webpack_require__(/*! ./main-nav.component.scss */ "./src/app/main-nav/main-nav.component.scss")]
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_cdk_layout__WEBPACK_IMPORTED_MODULE_2__["BreakpointObserver"]])
    ], MainNavComponent);
    return MainNavComponent;
}());



/***/ }),

/***/ "./src/app/test-dialog/test-dialog.component.html":
/*!********************************************************!*\
  !*** ./src/app/test-dialog/test-dialog.component.html ***!
  \********************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<div fxLayout=\"row wrap\" >\n  <div fxFlex=\"100\" style=\"text-align: center;\">\n    <h3 *ngIf=\"this.data.transport; else elseBlock\">Transporting {{data.api.apiName}} to {{this.data.environment}}</h3>\n    <ng-template #elseBlock> Testing {{data.api.apiName}} in {{this.data.environment}}</ng-template>\n  </div>\n\n  <div fxFlex=\"70%\" style=\"text-align: left;\">\n    <ng-container *ngFor=\"let test of tests\">\n      <h4>{{test.name}}</h4>\n    </ng-container>\n  </div>\n\n  <div fxFlex style=\"text-align: left;\">\n    <ng-container *ngFor=\"let test of tests;\">\n      <h4>\n        <mat-icon *ngIf=\"test.success !== null && !this.working && !test.success\"\n                  color=\"warn\"\n                  [inline]=\"true\"\n                  mat-list-icon>clear</mat-icon>\n        <mat-icon *ngIf=\"test.success !== null && !this.working && test.success\"\n                  style=\"color: green;\"\n                  [inline]=\"true\"\n                  mat-list-icon>done</mat-icon>\n      </h4>\n    </ng-container>\n  </div>\n\n  <div fxFlex=\"100\">\n    <ng-container *ngIf=\"this.data.transport\">\n      <mat-form-field style=\"width: 60%;\">\n        <input matInput type=\"email\" placeholder=\"Email\" [formControl]=\"email\">\n      </mat-form-field>\n      <mat-form-field style=\"width: 60%;\">\n        <input matInput type=\"password\" placeholder=\"Password\" [formControl]=\"password\">\n      </mat-form-field>\n    </ng-container>\n    <mat-form-field>\n      <input matInput [formControl]=\"metaDataPath\" placeholder=\"Path to Metadata\">\n      <span matPrefix>{{data.api.apiURL}}</span>\n    </mat-form-field>\n    <mat-form-field style=\"width: 60%; margin-right: 1rem;\">\n      <textarea matInput [formControl]=\"token\" placeholder=\"Input your bearer token\"></textarea>\n    </mat-form-field>\n\n    <button\n      color=\"primary\"\n      [disabled]=\"!this.token.valid\"\n      (click)=\"testService()\"\n      mat-raised-button>\n      <ng-container *ngIf=\"this.data.transport; else buttonElseBlock\">Transport</ng-container>\n      <ng-template #buttonElseBlock>Run Tests</ng-template>\n    </button>\n    <mat-spinner\n          *ngIf=\"this.working\"\n          color=\"accent\"\n          style=\"display: inline; margin-left: 1rem;\"\n          [diameter]=\"30\">\n    </mat-spinner>\n  </div>\n</div>\n"

/***/ }),

/***/ "./src/app/test-dialog/test-dialog.component.scss":
/*!********************************************************!*\
  !*** ./src/app/test-dialog/test-dialog.component.scss ***!
  \********************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL3Rlc3QtZGlhbG9nL3Rlc3QtZGlhbG9nLmNvbXBvbmVudC5zY3NzIn0= */"

/***/ }),

/***/ "./src/app/test-dialog/test-dialog.component.ts":
/*!******************************************************!*\
  !*** ./src/app/test-dialog/test-dialog.component.ts ***!
  \******************************************************/
/*! exports provided: TestDialogComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "TestDialogComponent", function() { return TestDialogComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_material__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/material */ "./node_modules/@angular/material/esm5/material.es5.js");
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/forms */ "./node_modules/@angular/forms/fesm5/forms.js");
/* harmony import */ var _api_test_service__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../api-test.service */ "./src/app/api-test.service.ts");
/* harmony import */ var _apitransport_service__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ../apitransport.service */ "./src/app/apitransport.service.ts");







var Test = /** @class */ (function () {
    function Test(name) {
        this.name = name;
        this.success = null;
    }
    return Test;
}());
;
var TestDialogComponent = /** @class */ (function () {
    function TestDialogComponent(dialogRef, data, apiTestService, apiTransportService, snackBar) {
        this.dialogRef = dialogRef;
        this.data = data;
        this.snackBar = snackBar;
        this.tests = [
            new Test("Client ID Authorized"),
            new Test("Rejects Incorrect Token"),
            new Test("Successfully Called")
        ];
        this.apiTestService = apiTestService;
        this.apiTransportService = apiTransportService;
        this.token = new _angular_forms__WEBPACK_IMPORTED_MODULE_3__["FormControl"]('', [_angular_forms__WEBPACK_IMPORTED_MODULE_3__["Validators"].required]);
        this.metaDataPath = new _angular_forms__WEBPACK_IMPORTED_MODULE_3__["FormControl"]('/$metadata', [_angular_forms__WEBPACK_IMPORTED_MODULE_3__["Validators"].required]);
        this.email = new _angular_forms__WEBPACK_IMPORTED_MODULE_3__["FormControl"]('example@nike.com', [_angular_forms__WEBPACK_IMPORTED_MODULE_3__["Validators"].required]);
        this.password = new _angular_forms__WEBPACK_IMPORTED_MODULE_3__["FormControl"]('', [_angular_forms__WEBPACK_IMPORTED_MODULE_3__["Validators"].required]);
        if (this.data.transport) {
            this.tests.push(new Test("Client ID Authorized in Target Tenant"));
            this.tests.push(new Test("Authorized to Transport"));
        }
        this.working = false;
    }
    TestDialogComponent.prototype.testService = function () {
        return tslib__WEBPACK_IMPORTED_MODULE_0__["__awaiter"](this, void 0, void 0, function () {
            var result, transport, _i, result_1, test;
            return tslib__WEBPACK_IMPORTED_MODULE_0__["__generator"](this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (this.token.invalid) {
                            return [2 /*return*/];
                        }
                        this.working = true;
                        result = [];
                        if (!!this.data.transport) return [3 /*break*/, 2];
                        return [4 /*yield*/, this.apiTestService.Test(this.data.api, this.data.environment, this.token.value, this.metaDataPath.value)];
                    case 1:
                        result = _a.sent();
                        return [3 /*break*/, 4];
                    case 2: return [4 /*yield*/, this.apiTransportService.Transport(this.data.api, this.data.environment, this.token.value, this.metaDataPath.value, this.email.value, this.password.value)];
                    case 3:
                        result = _a.sent();
                        _a.label = 4;
                    case 4:
                        this.working = false;
                        if (result === null) {
                            return [2 /*return*/];
                        }
                        transport = false;
                        for (_i = 0, result_1 = result; _i < result_1.length; _i++) {
                            test = result_1[_i];
                            if (test['Name'] === 'kvm authorization test') {
                                this.tests[0].success = test['Pass'];
                            }
                            if (test['Name'] === 'unauthorized client test') {
                                this.tests[1].success = test['Pass'];
                            }
                            if (test['Name'] === 'api authentication test') {
                                this.tests[2].success = test['Pass'];
                            }
                            if (test['Name'] === 'authorized in target tenant' && this.tests.length > 3) {
                                this.tests[3].success = test['Pass'];
                            }
                            if (test['Name'] === 'transport' && test['Pass'] === true) {
                                transport = true;
                            }
                            if (test['Name'] === 'ldap authentication test' && this.tests.length > 4) {
                                this.tests[4].success = test['Pass'];
                            }
                        }
                        if (transport && this.tests.length >= 4 && this.tests[0].success && this.tests[1].success && this.tests[2].success && this.tests[3].success) {
                            this.snackBar.open("Transport Succeeded!", null, {
                                duration: 3000,
                                panelClass: ["succeed-snackbar"]
                            });
                            this.dialogRef.close(true);
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    TestDialogComponent = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'app-test-dialog',
            template: __webpack_require__(/*! ./test-dialog.component.html */ "./src/app/test-dialog/test-dialog.component.html"),
            styles: [__webpack_require__(/*! ./test-dialog.component.scss */ "./src/app/test-dialog/test-dialog.component.scss")]
        }),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__param"](1, Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Inject"])(_angular_material__WEBPACK_IMPORTED_MODULE_2__["MAT_DIALOG_DATA"])),
        tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_material__WEBPACK_IMPORTED_MODULE_2__["MatDialogRef"], Object, _api_test_service__WEBPACK_IMPORTED_MODULE_4__["ApiTestService"],
            _apitransport_service__WEBPACK_IMPORTED_MODULE_5__["ApiTransportService"],
            _angular_material__WEBPACK_IMPORTED_MODULE_2__["MatSnackBar"]])
    ], TestDialogComponent);
    return TestDialogComponent;
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
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.
var environment = {
    production: false
};
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
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
/* harmony import */ var hammerjs__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! hammerjs */ "./node_modules/hammerjs/hammer.js");
/* harmony import */ var hammerjs__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(hammerjs__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm5/core.js");
/* harmony import */ var _angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/platform-browser-dynamic */ "./node_modules/@angular/platform-browser-dynamic/fesm5/platform-browser-dynamic.js");
/* harmony import */ var _app_app_module__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./app/app.module */ "./src/app/app.module.ts");
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./environments/environment */ "./src/environments/environment.ts");





if (_environments_environment__WEBPACK_IMPORTED_MODULE_4__["environment"].production) {
    Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["enableProdMode"])();
}
Object(_angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_2__["platformBrowserDynamic"])().bootstrapModule(_app_app_module__WEBPACK_IMPORTED_MODULE_3__["AppModule"])
    .catch(function (err) { return console.error(err); });


/***/ }),

/***/ 0:
/*!***************************!*\
  !*** multi ./src/main.ts ***!
  \***************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(/*! /Users/tandr9/Documents/tmp/angular/api-manager/src/main.ts */"./src/main.ts");


/***/ })

},[[0,"runtime","vendor"]]]);
//# sourceMappingURL=main.js.map