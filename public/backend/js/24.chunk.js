webpackJsonp([24],{

/***/ 397:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_settings_vue__ = __webpack_require__(420);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_settings_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_settings_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_settings_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_settings_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_70e7c3f2_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_settings_vue__ = __webpack_require__(491);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_70e7c3f2_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_settings_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_70e7c3f2_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_settings_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(489)
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
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_settings_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_70e7c3f2_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_settings_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/plugins/settings.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-70e7c3f2", Component.options)
  } else {
    hotAPI.reload("data-v-70e7c3f2", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 420:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});

var _util = __webpack_require__(17);

var _util2 = _interopRequireDefault(_util);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

exports.default = {
    data: function data() {
        return {
            editFormLoading: false,
            options: [],
            myForm: {}
        };
    },
    created: function created() {
        this.getData();
    },

    methods: {
        getData: function getData() {
            var dataType = {
                1: 'text',
                2: 'textarea'
            };
            var that = this;
            that.editFormLoading = true;
            _util2.default.ajax.get('/backend/settings').then(function (response) {
                var res = response.data;
                if (res.status === 200) {
                    var data = res.list;
                    if (data.length > 0) {
                        for (var index in data) {
                            that.myForm[data[index].name] = data[index].value;
                        }

                        for (var j = 0; j < data.length; j++) {
                            data[j].data_type = dataType[data[j].data_type];
                        }

                        that.options = data;
                    }
                } else {
                    that.$Notice.warning({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
                that.editFormLoading = false;
            }).catch(function (error) {
                console.log(error);
                that.editFormLoading = false;
            });
        },

        submitMyForm: function submitMyForm() {
            var that = this;
            _util2.default.ajax.post('/backend/settings/update', _util2.default.stringify(that.myForm)).then(function (response) {
                var res = response.data;
                that.$Notice.warning({
                    title: res.status === 200 ? '更新成功' : '更新失败',
                    desc: ''
                });
            }).catch(function (error) {
                console.log(error);
            });
        }
    },
    watch: {},
    mounted: function mounted() {}
};

/***/ }),

/***/ 489:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(490);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(22)("793d90dd", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-70e7c3f2\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./settings.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-70e7c3f2\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./settings.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 490:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(21)(false);
// imports


// module
exports.push([module.i, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", ""]);

// exports


/***/ }),

/***/ 491:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
var render = function render() {
  var _vm = this;
  var _h = _vm.$createElement;
  var _c = _vm._self._c || _h;
  return _c("div", { staticClass: "pit-post-form", staticStyle: { width: "60%" } }, [_c("Form", {
    ref: "myForm",
    staticClass: "pit-common",
    attrs: { model: _vm.myForm, "label-width": 150 }
  }, [_vm._l(_vm.options, function (option) {
    return _c("FormItem", {
      key: option.id,
      attrs: { value: option.id, label: option.title }
    }, [_c("Input", {
      attrs: { type: option.data_type },
      model: {
        value: _vm.myForm[option.name],
        callback: function callback($$v) {
          _vm.$set(_vm.myForm, option.name, $$v);
        },
        expression: "myForm[option.name]"
      }
    })], 1);
  }), _vm._v(" "), _c("FormItem", [_c("Button", {
    on: {
      click: function click($event) {
        _vm.closeForm("myForm");
      }
    }
  }, [_vm._v("取 消")]), _vm._v(" "), _c("Button", {
    attrs: { type: "primary" },
    on: {
      click: function click($event) {
        _vm.submitMyForm("myForm");
      }
    }
  }, [_vm._v("确 定")])], 1)], 2), _vm._v(" "), _vm.editFormLoading ? _c("Spin", { attrs: { size: "large", fix: "" } }) : _vm._e()], 1);
};
var staticRenderFns = [];
render._withStripped = true;
var esExports = { render: render, staticRenderFns: staticRenderFns };
exports.default = esExports;

if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-hot-reload-api").rerender("data-v-70e7c3f2", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=24.chunk.js.map