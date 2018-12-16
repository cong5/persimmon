webpackJsonp([27],{

/***/ 395:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_comments_vue__ = __webpack_require__(418);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_comments_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_comments_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_comments_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_comments_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_4fac2150_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_comments_vue__ = __webpack_require__(485);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_4fac2150_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_comments_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_4fac2150_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_comments_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(483)
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
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_comments_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_4fac2150_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_comments_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/plugins/comments.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-4fac2150", Component.options)
  } else {
    hotAPI.reload("data-v-4fac2150", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 418:
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
            listData: [],
            categorys: [],
            currentPage: 1,
            total: 0,
            pageSize: 20,
            myForm: {
                id: 0,
                name: '',
                url: '',
                email: '',
                markdown: ''
            },
            myRules: {
                name: [{ required: true, type: 'string', message: '请填写链接名称', trigger: 'blur' }],
                email: [{ required: true, type: 'string', message: '请填写邮箱', trigger: 'blur' }],
                markdown: [{ required: true, type: 'string', message: '请填写评论内容', trigger: 'blur' }]
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
            _util2.default.ajax.get('/backend/comments', {
                params: {
                    rows: that.pageSize,
                    page: this.currentPage
                }
            }).then(function (response) {
                var res = response.data;
                if (res.status === 200) {
                    that.listData = _util2.default.emailToMd5(res.list.data);
                    that.total = res.list.total;
                    that.currentPage = res.list.current_page;
                    that.listLoading = false;
                    that.$forceUpdate();
                } else {
                    that.$Notice.warning({
                        message: '数据获取失败',
                        type: 'error',
                        duration: 3 * 1000
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
            this.getData();
        },
        spam: function spam(row, status) {
            var id = row.id;
            var that = this;
            _util2.default.ajax.post('/backend/comments/spam', _util2.default.stringify({
                id: id,
                status: status
            })).then(function (response) {
                var res = response.data;
                that.$Notice.warning({
                    title: res.status === 200 ? '操作成功' : '操作失败',
                    desc: ''
                });
                that.$nextTick(function () {
                    that.getData();
                });
            }).catch(function (error) {
                console.log(error);
            });
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
            _util2.default.ajax.get('/backend/comments/' + row.id).then(function (response) {
                var res = response.data;
                if (res.status === 200) {
                    that.myForm = res.item;
                } else {
                    that.$Notice.warning({
                        message: '数据获取失败',
                        type: 'error'
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
                            message: '请选择需要删除的数据',
                            type: 'warning'
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    var ids = _util2.default.getIdByArr(that.checkedAll);
                    if (ids.length <= 0) {
                        that.$Notice.warning({
                            message: '请选择需要删除的数据',
                            type: 'warning'
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
                    _util2.default.ajax.post('/backend/comments/destroy', _util2.default.stringify(idsParam)).then(function (response) {
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
        postEditor: function postEditor(id) {
            var path = '/posts/edit/' + id;
            this.$router.push({ path: path });
        },
        submitMyForm: function submitMyForm() {
            var that = this;
            that.$refs['myForm'].validate(function (valid) {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }

                if (that.myForm.id > 0) {
                    _util2.default.ajax.post('/backend/comments/update', _util2.default.stringify(that.myForm)).then(function (response) {
                        var res = response.data;
                        that.$Notice.open({
                            title: res.status == 200 ? '编辑成功' : '编辑失败',
                            desc: ''
                        });
                        if (res.status == 200) {
                            that.closeForm('myForm');
                            that.getData();
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                } else {
                    _util2.default.ajax.post('/backend/links', that.myForm).then(function (response) {
                        var res = response.data;
                        if (res.status == 200) {
                            that.closeForm('myForm');
                            that.getData();
                        }
                        that.$Notice.warning({
                            title: res.status == 200 ? '新增成功' : '新增失败',
                            desc: ''
                        });
                    }).catch(function (error) {
                        if (error.response) {
                            if (error.response.status == 422) {
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
                url: '',
                logo: '',
                group: ''
            };
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

/***/ 483:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(484);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(22)("9de40b4a", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-4fac2150\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./comments.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-4fac2150\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./comments.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 484:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(21)(false);
// imports


// module
exports.push([module.i, "\n.links {\n  text-decoration: none;\n  color: #000;\n}\n.comment-avatar img {\n  border-radius: 50%;\n}\n.myForm {\n  width: 100% !important;\n}\n.myp-bold {\n  font-weight: bold;\n}\n.common-list-head {\n  border-bottom: 1px #eee solid;\n  padding-bottom: 5px;\n}\n.comment-list-body {\n  margin-top: 10px;\n}\n.comment-item {\n  border-bottom: 1px #eee dashed;\n  padding: 10px 0;\n}\n.comment-item-author {\n  padding-right: 20px;\n}\n.comment-item-avatar img {\n  float: left;\n  margin-right: 10px;\n  margin-top: 1px;\n}\n.comment-item-meta {\n  overflow: hidden;\n}\n.comment-item-content {\n  margin: .6em 0;\n}\n", ""]);

// exports


/***/ }),

/***/ 485:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
var render = function render() {
  var _vm = this;
  var _h = _vm.$createElement;
  var _c = _vm._self._c || _h;
  return _c("div", { staticClass: "pit-content" }, [_c("div", { staticClass: "data-list" }, [_c("div", { staticClass: "common-item" }, [_c("div", { staticClass: "common-list-head" }, [_c("Row", { attrs: { type: "flex" } }, [_c("Col", { attrs: { span: "4", order: "1" } }, [_vm._v("\n                    作者\n                    ")]), _vm._v(" "), _c("Col", { attrs: { span: "12", order: "2" } }, [_vm._v("\n                    评论\n                    ")]), _vm._v(" "), _c("Col", { attrs: { span: "4", order: "3" } }, [_vm._v("\n                    回复于\n                    ")]), _vm._v(" "), _c("Col", { attrs: { span: "4", order: "4" } }, [_vm._v("\n                    日期\n                    ")])], 1)], 1), _vm._v(" "), _c("div", { staticClass: "comment-list-body" }, _vm._l(_vm.listData, function (row) {
    return _c("Row", {
      key: row.id,
      staticClass: "comment-item",
      attrs: { type: "flex" }
    }, [_c("Col", { attrs: { span: "4", order: "1" } }, [_c("div", { staticClass: "comment-item-author" }, [_c("div", { staticClass: "comment-item-avatar" }, [_c("strong", [_c("img", {
      attrs: {
        src: "https://cn.gravatar.com/avatar/" + row.md5 + "?d=identicon&s=32",
        alt: ""
      }
    }), _vm._v(_vm._s(row.name) + "\n                            ")])]), _vm._v(" "), _c("div", { staticClass: "comment-item-meta" }, [_c("a", { attrs: { href: row.email, target: "_blank" } }, [_vm._v(_vm._s(row.email))]), _vm._v(" "), _c("br"), _vm._v(" "), _c("a", { attrs: { href: row.url, target: "_blank" } }, [_vm._v(_vm._s(row.url))])])])]), _vm._v(" "), _c("Col", { attrs: { span: "12", order: "2" } }, [_c("p", {
      staticClass: "comment-item-content",
      domProps: { innerHTML: _vm._s(row.content) }
    }), _vm._v(" "), _c("p", [_c("a", {
      attrs: { href: "javascript:void(0);" },
      on: {
        click: function click($event) {
          _vm.handleEdit(row);
        }
      }
    }, [_vm._v("编辑")]), _vm._v(" |\n                        "), row.status == 1 ? [_c("a", {
      attrs: { href: "javascript:void(0);" },
      on: {
        click: function click($event) {
          _vm.spam(row, 3);
        }
      }
    }, [_vm._v("垃圾评论")]), _vm._v(" |\n                        ")] : row.status == 2 ? [_c("a", {
      attrs: { href: "javascript:void(0);" },
      on: {
        click: function click($event) {
          _vm.spam(row, 1);
        }
      }
    }, [_vm._v("审核")]), _vm._v(" |\n                        ")] : row.status == 3 ? [_c("a", {
      attrs: { href: "javascript:void(0);" },
      on: {
        click: function click($event) {
          _vm.spam(row, 2);
        }
      }
    }, [_vm._v("不是垃圾评论")]), _vm._v(" |\n                        ")] : _vm._e(), _vm._v(" "), _c("a", {
      attrs: { href: "javascript:void(0);" },
      on: {
        click: function click($event) {
          _vm.handleDistory("one", row);
        }
      }
    }, [_vm._v("删除")])], 2)]), _vm._v(" "), _c("Col", { attrs: { span: "4", order: "3" } }, [_c("div", { staticClass: "comment-item-title" }, [_c("a", {
      staticClass: "myp-bold",
      attrs: {
        href: "javascript:void(0)",
        target: "_blank"
      },
      on: {
        click: function click($event) {
          _vm.postEditor(row.posts_id);
        }
      }
    }, [_vm._v(_vm._s(row.title))]), _vm._v(" "), _c("br"), _vm._v(" "), _c("a", {
      attrs: { href: "/post/" + row.slug, target: "_blank" }
    }, [_vm._v("查看文章")])])]), _vm._v(" "), _c("Col", { attrs: { span: "4", order: "4" } }, [_c("div", { staticClass: "comment-item-date" }, [_vm._v("\n                        " + _vm._s(_vm._f("moment")(row.created_at * 1000, "YYYY-MM-DD A h:mm")) + "\n                    ")])])], 1);
  }))])]), _vm._v(" "), _c("div", { staticStyle: { margin: "10px", overflow: "hidden" } }, [_c("div", { staticStyle: { float: "right" } }, [_c("Page", {
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
    attrs: {
      rules: _vm.myRules,
      "label-width": 100,
      model: _vm.myForm
    }
  }, [_c("FormItem", { attrs: { label: "名称", prop: "name" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.name,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "name", $$v);
      },
      expression: "myForm.name"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "链接" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.url,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "url", $$v);
      },
      expression: "myForm.url"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "E-Mail", prop: "email" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.email,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "email", $$v);
      },
      expression: "myForm.email"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "内容", prop: "markdown" } }, [_c("Input", {
    attrs: {
      type: "textarea",
      autosize: "",
      "auto-complete": "off"
    },
    model: {
      value: _vm.myForm.markdown,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "markdown", $$v);
      },
      expression: "myForm.markdown"
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
    require("vue-hot-reload-api").rerender("data-v-4fac2150", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=27.chunk.js.map