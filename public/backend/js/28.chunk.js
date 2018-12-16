webpackJsonp([28],{

/***/ 388:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_own_space_vue__ = __webpack_require__(404);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_own_space_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_own_space_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_own_space_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_own_space_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5b3915ac_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_own_space_vue__ = __webpack_require__(437);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5b3915ac_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_own_space_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5b3915ac_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_own_space_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(435)
}
var normalizeComponent = __webpack_require__(5)
/* script */


/* template */

/* template functional */
var __vue_template_functional__ = false
/* styles */
var __vue_styles__ = injectStyle
/* scopeId */
var __vue_scopeId__ = null
/* moduleIdentifier (server only) */
var __vue_module_identifier__ = null
var Component = normalizeComponent(
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_own_space_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5b3915ac_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_own_space_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/own-space/own-space.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-5b3915ac", Component.options)
  } else {
    hotAPI.reload("data-v-5b3915ac", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 404:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});

var _util = __webpack_require__(17);

var _util2 = _interopRequireDefault(_util);

var _jsCookie = __webpack_require__(11);

var _jsCookie2 = _interopRequireDefault(_jsCookie);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

exports.default = {
    data: function data() {
        return {
            userName: '',
            saveLoading: false,
            myForm: {
                old_password: '',
                password: '',
                password_confirmation: ''
            },
            myRules: {
                old_password: [{ required: true, type: "string", message: '请填写旧密码', trigger: 'blur' }],
                password: [{ required: true, type: "string", message: '请填写新密码', trigger: 'blur' }, { min: 8, max: 64, message: '长度在 8 到 64 个字符', trigger: 'blur' }],
                password_confirmation: [{ required: true, type: "string", message: '请再填写一次密码', trigger: 'blur' }, { min: 8, max: 64, message: '长度在 8 到 64 个字符', trigger: 'blur' }]
            }
        };
    },
    created: function created() {},

    methods: {
        submitMyForm: function submitMyForm() {
            var that = this;
            that.$refs['myForm'].validate(function (valid) {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }
                _util2.default.ajax.post('/backend/user', that.myForm).then(function (response) {
                    var res = response.data;
                    that.$Notice.open({
                        title: res.status == 200 ? '信息更新成功' : '信息更新失败',
                        desc: ''
                    });
                    if (res.status == 200) {
                        that.closeForm('myForm');
                    }
                    setTimeout(function () {
                        _jsCookie2.default.remove('user');
                        that.$router.push({ path: '/login' });
                    }, 2 * 1000);
                }).catch(function (error) {
                    if (error.response) {
                        if (error.response.status == 422) {
                            for (var index in error.response.data) {
                                that.$Notice.open({
                                    title: '警告',
                                    message: error.response.data[index][0]
                                });
                            }
                        }
                    } else {
                        console.log(error);
                    }
                });
            });
        },
        closeForm: function closeForm() {
            this.saveLoading = false;
            this.$refs['myForm'].resetFields();
            this.myForm = {
                old_password: '',
                password: '',
                password_confirmation: ''
            };
            console.log('closeForm');
        }
    },
    mounted: function mounted() {
        var user = _jsCookie2.default.get('user');
        var userInfo = JSON.parse(user);
        this.userName = userInfo.name;
    }
};

/***/ }),

/***/ 435:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(436);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(22)("29e1eb0d", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-5b3915ac\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./own-space.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-5b3915ac\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./own-space.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 436:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(21)(false);
// imports


// module
exports.push([module.i, "\n.own-space-btn-box {\n  margin-bottom: 10px;\n}\n.own-space-btn-box button {\n  padding-left: 0;\n}\n.own-space-btn-box button span {\n  color: #2D8CF0;\n  transition: all .2s;\n}\n.own-space-btn-box button span:hover {\n  color: #0C25F1;\n  transition: all .2s;\n}\n.own-space-tra {\n  width: 10px;\n  height: 10px;\n  transform: rotate(45deg);\n  position: absolute;\n  top: 50%;\n  margin-top: -6px;\n  left: -3px;\n  box-shadow: 0 0 2px 3px rgba(0, 0, 0, 0.1);\n  background-color: white;\n  z-index: 100;\n}\n.own-space-input-identifycode-con {\n  position: absolute;\n  width: 200px;\n  height: 100px;\n  right: -220px;\n  top: 50%;\n  margin-top: -50px;\n  border-radius: 4px;\n  box-shadow: 0 0 2px 3px rgba(0, 0, 0, 0.1);\n}\n", ""]);

// exports


/***/ }),

/***/ 437:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
var render = function render() {
  var _vm = this;
  var _h = _vm.$createElement;
  var _c = _vm._self._c || _h;
  return _c("div", [_c("Card", [_c("p", { attrs: { slot: "title" }, slot: "title" }, [_c("Icon", { attrs: { type: "person" } }), _vm._v("\n            个人信息\n        ")], 1), _vm._v(" "), _c("div", [_c("Form", {
    ref: "myForm",
    attrs: {
      model: _vm.myForm,
      "label-width": 100,
      "label-position": "right",
      rules: _vm.myRules
    }
  }, [_c("FormItem", { attrs: { label: "用户名：" } }, [_c("span", [_vm._v(_vm._s(_vm.userName))])]), _vm._v(" "), _c("FormItem", { attrs: { label: "旧密码：", prop: "old_password" } }, [_c("div", {
    staticStyle: { display: "inline-block", width: "300px" }
  }, [_c("Input", {
    attrs: { type: "password" },
    model: {
      value: _vm.myForm.old_password,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "old_password", $$v);
      },
      expression: "myForm.old_password"
    }
  }, [_c("Icon", {
    attrs: {
      slot: "prepend",
      type: "ios-locked-outline"
    },
    slot: "prepend"
  })], 1)], 1)]), _vm._v(" "), _c("FormItem", { attrs: { label: "新密码：", prop: "password" } }, [_c("div", {
    staticStyle: { display: "inline-block", width: "300px" }
  }, [_c("Input", {
    attrs: { type: "password" },
    model: {
      value: _vm.myForm.password,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "password", $$v);
      },
      expression: "myForm.password"
    }
  }, [_c("Icon", {
    attrs: {
      slot: "prepend",
      type: "ios-locked-outline"
    },
    slot: "prepend"
  })], 1)], 1)]), _vm._v(" "), _c("FormItem", {
    attrs: {
      label: "确认密码：",
      prop: "password_confirmation"
    }
  }, [_c("div", {
    staticStyle: { display: "inline-block", width: "300px" }
  }, [_c("Input", {
    attrs: { type: "password" },
    model: {
      value: _vm.myForm.password_confirmation,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "password_confirmation", $$v);
      },
      expression: "myForm.password_confirmation"
    }
  }, [_c("Icon", {
    attrs: {
      slot: "prepend",
      type: "ios-locked-outline"
    },
    slot: "prepend"
  })], 1)], 1)]), _vm._v(" "), _c("div", [_c("Button", {
    staticStyle: { width: "100px" },
    attrs: { type: "text" },
    on: { click: _vm.closeForm }
  }, [_vm._v("取消")]), _vm._v(" "), _c("Button", {
    staticStyle: { width: "100px" },
    attrs: { type: "primary", loading: _vm.saveLoading },
    on: { click: _vm.submitMyForm }
  }, [_vm._v("保存\n                    ")])], 1)], 1)], 1)])], 1);
};
var staticRenderFns = [];
render._withStripped = true;
var esExports = { render: render, staticRenderFns: staticRenderFns };
exports.default = esExports;

if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-hot-reload-api").rerender("data-v-5b3915ac", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=28.chunk.js.map