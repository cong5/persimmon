webpackJsonp([23],{

/***/ 390:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_categorys_vue__ = __webpack_require__(412);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_categorys_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_categorys_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_categorys_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_categorys_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_053f9fb4_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_categorys_vue__ = __webpack_require__(467);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_053f9fb4_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_categorys_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_053f9fb4_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_categorys_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(465)
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
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_categorys_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_053f9fb4_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_categorys_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/posts/categorys.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-053f9fb4", Component.options)
  } else {
    hotAPI.reload("data-v-053f9fb4", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 412:
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
        var _this = this;

        return {
            listData: [],
            categorys: [],
            myForm: {
                id: 0,
                name: '',
                slug: '',
                description: '',
                pid: 0
            },
            myRules: {
                name: [{ required: true, type: "string", title: '请填写分类名称', trigger: 'blur' }],
                slug: [{ required: true, type: "string", title: '请填写分类别名', trigger: 'blur' }, { pattern: /^[a-zA-Z0-9_-]+$/, title: '只允许英文或者拼音,横杠(-),下划线(_)', trigger: 'blur' }]
            },
            editFormVisible: false,
            editFormLoading: false,
            listLoading: false,
            myFormTitle: '编辑',
            checkedAll: [],
            tableColumns: [{
                type: 'selection',
                width: 60,
                align: 'center'
            }, {
                title: '分类名称',
                key: 'name'
            }, {
                title: '分类别名',
                key: 'slug'
            }, {
                title: '日期',
                key: 'created_at',
                render: function render(h, params) {
                    var time = _util2.default.timeFormat(params.row.created_at);
                    return h('span', time);
                }
            }, {
                title: '操作',
                key: 'action',
                width: 150,
                align: 'center',
                render: function render(h, params) {
                    return h('div', [h('Button', {
                        props: {
                            type: 'primary',
                            size: 'small'
                        },
                        style: {
                            marginRight: '5px'
                        },
                        on: {
                            click: function click() {
                                _this.handleEdit(params.row);
                            }
                        }
                    }, '编辑'), h('Button', {
                        props: {
                            type: 'error',
                            size: 'small'
                        },
                        on: {
                            click: function click() {
                                _this.handleDistory('one', params.row);
                            }
                        }
                    }, '删除')]);
                }
            }]
        };
    },

    methods: {
        getData: function getData() {
            var that = this;
            that.listLoading = true;
            _util2.default.ajax.get('/backend/categories', {
                params: {
                    page: 1,
                    rows: 20
                }
            }).then(function (response) {
                var res = response.data;
                if (res.status === 200) {
                    that.listData = res.list.data;
                    that.listLoading = false;
                } else {
                    that.$Notice.warning({
                        title: '数据获取失败',
                        duration: 3 * 1000
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        handleCreate: function handleCreate() {
            var that = this;
            that.myFormTitle = '新增';
            that.myForm.id = 0;
            that.editFormVisible = true;
            that.setTopCategorys();
        },
        handleEdit: function handleEdit(row) {
            var that = this;
            that.setTopCategorys();
            that.editFormLoading = true;
            that.myFormTitle = '编辑';
            that.editFormVisible = true;
            _util2.default.ajax.get('/backend/categories/' + row.id).then(function (response) {
                var res = response.data;
                if (res !== false) {
                    that.myForm = res.item;
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
        handleDistory: function handleDistory(type, row) {
            var that = this,
                idsParam = {};
            switch (type) {
                case 'one':
                    if (parseInt(row.id) <= 0) {
                        that.$Notice.warning({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    var ids = _util2.default.getIdByArr(that.checkedAll);
                    if (ids.length <= 0) {
                        that.$Notice.warning({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: ids };
                    break;
                default:
                    break;
            }

            that.$Modal.confirm({
                title: '确认删除选中的记录吗?',
                content: '<p>您确认删除选中的记录吗?</p>',
                onOk: function onOk() {
                    that.listLoading = true;
                    _util2.default.ajax.post('/backend/categories/destroy', _util2.default.stringify(idsParam)).then(function (response) {
                        that.listLoading = false;
                        var res = response.data;
                        that.$Notice.warning({
                            title: res.status === 200 ? '删除成功' : '删除失败',
                            desc: ''
                        });
                        if (type === 'one') {
                            _util2.default.removeByValue(that.listData, row.id);
                        } else {
                            for (var index in that.checkedAll) {
                                _util2.default.removeByValue(that.listData, that.checkedAll[index].id);
                            }
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                },
                onCancel: function onCancel() {
                    that.listLoading = false;
                }
            });
        },
        submitMyForm: function submitMyForm() {
            var that = this;
            that.$refs['myForm'].validate(function (valid) {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }

                if (that.myForm.id > 0) {
                    _util2.default.ajax.put('/backend/categories/update', that.myForm).then(function (response) {
                        var res = response.data;
                        that.$Notice.open({
                            title: res.status === 200 ? '编辑成功' : '编辑失败',
                            desc: ''
                        });
                        if (res.status === 200) {
                            that.closeForm('myForm');
                            that.getData();
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                } else {
                    _util2.default.ajax.post('/backend/categories/store', that.myForm).then(function (response) {
                        var res = response.data;
                        if (res.status === 200) {
                            that.closeForm('myForm');
                        }
                        that.$Notice.open({
                            title: res.status === 200 ? '新增成功' : '新增失败',
                            desc: ''
                        });
                        that.getData();
                    }).catch(function (error) {
                        if (error.response) {
                            if (error.response.status === 501) {
                                for (var index in error.response.data) {
                                    that.$Notice.warning({
                                        title: '警告',
                                        desc: error.response.data[index][0]
                                    });
                                }
                            }
                        } else {
                            console.log(error);
                        }
                    });
                }
            });
        },
        closeForm: function closeForm() {
            this.editFormVisible = false;
            this.$refs['myForm'].resetFields();
            this.myForm = {
                id: 0,
                name: '',
                slug: '',
                description: '',
                pid: 0
            };
            console.log('closeForm');
        },
        setTopCategorys: function setTopCategorys() {
            var categorys = this.listData.concat();
            categorys.splice(0, 0, { id: 0, name: '顶级分类', hidden: true, pid: 0 });
            this.categorys = categorys;
        },
        selectRow: function selectRow(row) {
            this.checkedAll = row;
        },
        selectAll: function selectAll(selection) {
            this.checkedAll = selection;
        },
        translateWords: function translateWords(event) {
            var that = this;
            var query = that.myForm.name;
            if (query.match(/\w+/g) != null) {
                that.myForm.slug = query;
            }
            if (query == null || query === '') {
                return false;
            }
            _util2.default.ajax.get('/backend/utils/fanyi/' + query).then(function (response) {
                var res = response.data;
                if (res.status === 200) {
                    that.myForm.slug = res.item;
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    },
    mounted: function mounted() {
        this.getData();
    }
};

/***/ }),

/***/ 465:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(466);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(22)("638392aa", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-053f9fb4\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./categorys.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-053f9fb4\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./categorys.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 466:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(21)(false);
// imports


// module
exports.push([module.i, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", ""]);

// exports


/***/ }),

/***/ 467:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
var render = function render() {
  var _vm = this;
  var _h = _vm.$createElement;
  var _c = _vm._self._c || _h;
  return _c("div", { staticClass: "pit-content" }, [_c("div", { staticClass: "pit-action-btn" }, [_c("Button", {
    attrs: { type: "primary", icon: "plus" },
    on: { click: _vm.handleCreate }
  }, [_vm._v("新增")]), _vm._v(" "), _c("Button", {
    attrs: { type: "primary", icon: "delete" },
    on: {
      click: function click($event) {
        _vm.handleDistory("multi", {});
      }
    }
  }, [_vm._v("删除")])], 1), _vm._v(" "), _c("div", { staticClass: "data-list" }, [_c("Table", {
    attrs: {
      loading: _vm.listLoading,
      columns: _vm.tableColumns,
      data: _vm.listData,
      stripe: ""
    },
    on: { "on-select": _vm.selectRow, "on-select-all": _vm.selectAll }
  })], 1), _vm._v(" "), _c("Modal", {
    attrs: { title: _vm.myFormTitle },
    on: { "on-ok": _vm.submitMyForm, "on-cancel": _vm.closeForm },
    model: {
      value: _vm.editFormVisible,
      callback: function callback($$v) {
        _vm.editFormVisible = $$v;
      },
      expression: "editFormVisible"
    }
  }, [_c("div", { staticClass: "pit-dialog-edit-form" }, [_c("Form", {
    ref: "myForm",
    staticClass: "myForm",
    staticStyle: { width: "90%" },
    attrs: {
      rules: _vm.myRules,
      model: _vm.myForm,
      "label-position": "right",
      "label-width": 100
    }
  }, [_c("FormItem", { attrs: { label: "分类名称", prop: "name" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    on: {
      "on-blur": _vm.translateWords,
      "on-enter": _vm.translateWords
    },
    model: {
      value: _vm.myForm.name,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "name", $$v);
      },
      expression: "myForm.name"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "分类别名", prop: "slug" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.slug,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "slug", $$v);
      },
      expression: "myForm.slug"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "分类描述" } }, [_c("Input", {
    attrs: { type: "textarea" },
    model: {
      value: _vm.myForm.description,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "description", $$v);
      },
      expression: "myForm.description"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "父分类" } }, [_c("Select", {
    attrs: { placeholder: "请选择父分类" },
    model: {
      value: _vm.myForm.pid,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "pid", $$v);
      },
      expression: "myForm.pid"
    }
  }, _vm._l(_vm.categorys, function (item) {
    return _c("Option", { key: item.id, attrs: { value: item.id } }, [_vm._v(_vm._s(item.name) + "\n                        ")]);
  }))], 1), _vm._v(" "), _vm.myForm.id ? _c("FormItem", [_c("Input", {
    staticStyle: { display: "none" },
    model: {
      value: _vm.myForm.id,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "id", $$v);
      },
      expression: "myForm.id"
    }
  })], 1) : _vm._e()], 1), _vm._v(" "), _vm.editFormLoading ? _c("Spin", { attrs: { size: "large", fix: "" } }) : _vm._e()], 1)])], 1);
};
var staticRenderFns = [];
render._withStripped = true;
var esExports = { render: render, staticRenderFns: staticRenderFns };
exports.default = esExports;

if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-hot-reload-api").rerender("data-v-053f9fb4", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=23.chunk.js.map