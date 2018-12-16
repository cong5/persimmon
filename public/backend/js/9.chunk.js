webpackJsonp([9],{

/***/ 396:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_options_vue__ = __webpack_require__(419);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_options_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_options_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_options_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_options_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5c60f7ea_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_options_vue__ = __webpack_require__(488);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5c60f7ea_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_options_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5c60f7ea_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_options_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(486)
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
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_options_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_5c60f7ea_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_options_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/plugins/options.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-5c60f7ea", Component.options)
  } else {
    hotAPI.reload("data-v-5c60f7ea", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 419:
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
            tableColumns: [{
                type: 'selection',
                width: 60,
                align: 'center'
            }, {
                title: '配置项说明',
                key: 'title'
            }, {
                title: '配置项名称',
                key: 'name'
            }, {
                title: '配置分组',
                key: 'group'
            }, {
                title: '配置项备注',
                key: 'remark'
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
                            icon: 'edit',
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
                    }, ''), h('Button', {
                        props: {
                            icon: 'trash-b',
                            size: 'small'
                        },
                        on: {
                            click: function click() {
                                _this.handleDistory('one', params.row);
                            }
                        }
                    }, '')]);
                }
            }],
            listData: [],
            currentPage: 1,
            total: 0,
            pageSize: 20,
            myForm: {
                id: 0,
                title: '',
                name: '',
                value: '',
                group: '',
                remark: ''
            },
            myRules: {
                title: [{ required: true, type: "string", message: '请填写配置项说明', trigger: 'blur' }],
                name: [{ required: true, type: "string", message: '请填写配置项名称', trigger: 'blur' }, { pattern: /^[a-zA-Z0-9_]+$/, message: '只允许英文或者拼音,横杠(-)', trigger: 'blur' }]
            },
            editFormVisible: false,
            editFormLoading: false,
            listLoading: false,
            myFormTitle: '编辑',
            checkedAll: [],
            sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
        };
    },

    methods: {
        getData: function getData() {
            var that = this;
            that.listLoading = true;
            _util2.default.ajax.get('/backend/options', {
                params: {
                    rows: this.pageSize
                }
            }).then(function (response) {
                var res = response.data;
                if (res.status == 200) {
                    that.listData = res.list.data;
                    that.total = res.list.total;
                    that.currentPage = res.list.current_page;
                    that.listLoading = false;
                } else {
                    that.$Notice.open({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        handleSizeChange: function handleSizeChange(val) {
            this.pageSize = val;
            this.getData();
        },
        handleCurrentChange: function handleCurrentChange(val) {
            this.currentPage = val;
        },

        handleCreate: function handleCreate() {
            var that = this;
            that.myFormTitle = '新增';
            that.myForm.id = 0;
            that.editFormVisible = true;
        },
        handleEdit: function handleEdit(row) {
            var that = this;
            that.editFormLoading = true;
            that.myFormTitle = '编辑';
            that.editFormVisible = true;
            _util2.default.ajax.get('/backend/options/' + row.id).then(function (response) {
                var res = response.data;
                if (res.status === 200) {
                    that.myForm = res.item;
                } else {
                    that.$Notice.open({
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
                        that.$Notice.open({
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
                        that.$Notice.open({
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
                    _util2.default.ajax.post('/backend/options/destroy', _util2.default.stringify(idsParam)).then(function (response) {
                        that.listLoading = false;
                        var res = response.data;
                        that.$Notice.open({
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
                    _util2.default.ajax.post('/backend/options/update', _util2.default.stringify(that.myForm)).then(function (response) {
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
                    _util2.default.ajax.post('/backend/options/store', _util2.default.stringify(that.myForm)).then(function (response) {
                        var res = response.data;
                        if (res.status === 200) {
                            that.closeForm('myForm');
                            that.getData();
                        }
                        that.$Notice.open({
                            title: res.status === 200 ? '新增成功' : '新增失败',
                            desc: ''
                        });
                    }).catch(function (error) {
                        if (error.response) {
                            if (error.response.status === 422) {
                                for (var index in error.response.data) {
                                    that.$notify({
                                        title: '警告',
                                        message: error.response.data[index][0]
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
                url: '',
                logo: '',
                group: ''
            };
            console.log('closeForm');
        },
        selectRow: function selectRow(row) {
            this.checkedAll = row;
        },
        selectAll: function selectAll(selection) {
            this.checkedAll = selection;
        }
    },
    watch: {},
    mounted: function mounted() {
        this.getData();
    }
};

/***/ }),

/***/ 486:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(487);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(21)("29854515", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-5c60f7ea\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./options.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-5c60f7ea\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./options.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 487:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(22)(false);
// imports


// module
exports.push([module.i, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", ""]);

// exports


/***/ }),

/***/ 488:
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
    attrs: { icon: "plus" },
    on: { click: _vm.handleCreate }
  }), _vm._v(" "), _c("Button", {
    attrs: { icon: "trash-b" },
    on: {
      click: function click($event) {
        _vm.handleDistory("multi", {});
      }
    }
  })], 1), _vm._v(" "), _c("div", { staticClass: "data-list" }, [_c("Table", {
    attrs: {
      loading: _vm.listLoading,
      columns: _vm.tableColumns,
      data: _vm.listData,
      stripe: ""
    },
    on: { "on-select": _vm.selectRow, "on-select-all": _vm.selectAll }
  })], 1), _vm._v(" "), _c("div", { staticStyle: { margin: "10px", overflow: "hidden" } }, [_c("div", { staticStyle: { float: "right" } }, [_c("Page", {
    staticClass: "myp-page",
    attrs: {
      total: _vm.total,
      "page-size": _vm.pageSize,
      current: _vm.currentPage,
      "page-size-opts": _vm.sizeOpts,
      "show-elevator": true,
      "show-sizer": true
    },
    on: {
      "on-page-size-change": _vm.handleSizeChange,
      "on-change": _vm.handleCurrentChange
    }
  })], 1)]), _vm._v(" "), _c("Modal", {
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
    staticStyle: { width: "80%" },
    attrs: {
      rules: _vm.myRules,
      "label-width": 100,
      model: _vm.myForm
    }
  }, [_c("FormItem", { attrs: { label: "配置项说明", prop: "title" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.title,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "title", $$v);
      },
      expression: "myForm.title"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "配置项名称", prop: "name" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.name,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "name", $$v);
      },
      expression: "myForm.name"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "配置项值" } }, [_c("Input", {
    attrs: {
      type: _vm.myForm.data_type,
      autosize: _vm.myForm.data_type == "textarea",
      "auto-complete": "off"
    },
    model: {
      value: _vm.myForm.value,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "value", $$v);
      },
      expression: "myForm.value"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "配置分组" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.group,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "group", $$v);
      },
      expression: "myForm.group"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "配置项备注" } }, [_c("Input", {
    attrs: {
      type: "textarea",
      autosize: { minRows: 3, maxRows: 5 },
      "auto-complete": "off"
    },
    model: {
      value: _vm.myForm.option_remark,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "option_remark", $$v);
      },
      expression: "myForm.option_remark"
    }
  })], 1), _vm._v(" "), _vm.myForm.id ? _c("FormItem", [_c("Input", {
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
    require("vue-hot-reload-api").rerender("data-v-5c60f7ea", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=9.chunk.js.map