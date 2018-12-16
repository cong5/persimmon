<style lang="less">

</style>

<template>
    <div class="pit-content" style="margin: 20px;">
        <div class="data-list">
            <Tag type="dot" v-for="item in listData" :key="item.id" fade closable @click="handleEdit(row)" @on-close="handleDistory('one', item)">{{item.name}}</Tag>
            <Button icon="ios-plus-empty" type="dashed" size="small" @click="handleCreate">添加标签</Button>
        </div>

        <div style="margin: 10px;overflow: hidden">
            <div style="float: right;">
                <Page :total="total"
                      @on-page-size-change="handleSizeChange"
                      @on-change="handleCurrentChange"
                      :page-size="pageSize"
                      :current="currentPage"
                      :page-size-opts="sizeOpts"
                      :show-elevator="true"
                      :show-sizer="true"
                      class="myp-page">
                </Page>
            </div>
        </div>

        <Modal :title="myFormTitle" v-model="editFormVisible" @on-ok="submitMyForm" @on-cancel="closeForm">
            <div class="pit-dialog-edit-form">
                <Form ref="myForm" :rules="myRules" class="myForm" :label-width="100" :model="myForm">
                    <FormItem label="标签名称" prop="name">
                        <Input v-model="myForm.name" auto-complete="off" @on-blur="translateWords" @on-enter="translateWords"></Input>
                    </FormItem>
                    <FormItem label="标签别名" prop="slug">
                        <Input v-model="myForm.slug" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem v-if="myForm.id">
                        <Input v-model="myForm.id" style="display: none;"></Input>
                    </FormItem>
                </Form>
                <Spin size="large" fix v-if="editFormLoading"></Spin>
            </div>
        </Modal>

    </div>
</template>

<script>
    import util from '../../libs/util';

    export default{
        data(){
            return {
                listData: [],
                categorys: [],
                currentPage: 1,
                total: 0,
                pageSize: 50,
                myForm: {
                    id: 0,
                    name: '',
                    slug: ''
                },
                myRules: {
                    name: [
                        {required: true, type: "string", message: '请填写标签名称', trigger: 'blur'}
                    ],
                    slug: [
                        {required: true, type: "string", message: '请填写标签别名', trigger: 'blur'}
                    ]
                },
                editFormVisible: false,
                editFormLoading: false,
                listLoading: false,
                myFormTitle: '编辑',
                checkedAll: [],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
            }
        },
        methods: {
            formatterFlag: function (row, column) {
                if (row.slug === '') {
                    return '';
                }
                return decodeURI(row.slug);
            },
            getData: function () {
                let that = this;
                that.listLoading = true;
                util.ajax.get('/backend/tags', {
                    params: {
                        rows: this.pageSize,
                        page: this.currentPage
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.listData = res.list.data;
                        that.total = res.list.total;
                        that.currentPage = res.list.current_page;
                        that.listLoading = false;
                    } else {
                        that.$Notice.warning({
                            title: '数据获取失败',
                            desc: ''
                        });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            handleSizeChange(val) {
                //console.log(`每页 ${val} 条`);
                this.pageSize = val;
                this.getData();
            },
            handleCurrentChange(val) {
                this.currentPage = val;
                //console.log(`当前页: ${val}`);
                this.getData();
            },
            handleCreate: function () {
                let that = this;
                that.myFormTitle = '新增';
                that.myForm.id = 0;
                that.editFormVisible = true;
                that.setTopCategorys();
            },
            handleEdit: function (row) {
                let that = this;
                that.setTopCategorys();
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                that.editFormVisible = true;
                util.ajax.get('/backend/tags/' + row.id).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        res.slug = decodeURI(res.item.slug);
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
            handleDistory: function (type, row) {
                let that = this, idsParam = {};
                switch (type) {
                    case 'one':
                        if (parseInt(row.id) <= 0) {
                            that.$Notice.warning({
                                title: '请选择需要删除的数据',
                                desc: ''
                            });
                            return false;
                        }
                        idsParam = {ids: [row.id]};
                        break;
                    case 'multi':
                        let ids = that.util.getIdByArr(that.checkedAll);
                        if (ids.length <= 0) {
                            that.$Notice.warning({
                                title: '请选择需要删除的数据',
                                desc: ''
                            });
                            return false;
                        }
                        idsParam = {ids: ids};
                        break;
                    default:
                        break;
                }

                that.$Modal.confirm({
                    title: '确认删除选中的记录吗?',
                    content: '<p>您确认删除选中的记录吗?</p>',
                    onOk: () => {
                        that.listLoading = true;
                        util.ajax.post('/backend/tags/destroy', util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status == 200 ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type == 'one') {
                                that.util.removeByValue(that.listData, row.id);
                            } else {
                                for (let index in that.checkedAll) {
                                    that.util.removeByValue(that.listData, that.checkedAll[index].id);
                                }
                            }

                        }).catch(function (error) {
                            console.log(error);
                        });
                    },
                    onCancel: () => {
                        that.listLoading = false;
                    }
                });
            },
            submitMyForm: function () {
                let that = this;
                that.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }

                    if (that.myForm.id > 0) {
                        util.ajax.put('/backend/tags/update', that.myForm).then(function (response) {
                            let res = response.data;
                            that.$message({
                                message: res.status === 200 ? '编辑成功' : '编辑失败',
                                type: res.status
                            });
                            if (res.status === 200) {
                                that.closeForm('myForm');
                                that.getData();
                            }
                        }).catch(function (error) {
                            console.log(error);
                        });
                    } else {
                        util.ajax.post('/backend/tags/store', that.myForm).then(function (response) {
                            //console.log(response);
                            let res = response.data;
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
                                    for (let index in error.response.data) {
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
            closeForm() {
                this.editFormVisible = false;
                this.$refs['myForm'].resetFields();
                this.myForm = {
                    id: 0,
                    name: '',
                    slug: ''
                };
            },
            selectRow(row) {
                this.checkedAll = row;
            },
            selectAll(selection) {
                this.checkedAll = selection;
            },
            setTopCategorys: function () {
                let categorys = this.listData.concat();
                categorys.splice(0, 0, {id: 0, name: '顶级分类', hidden: true, pid: 0});
                this.categorys = categorys;
            },
            translateWords(event) {
                let that = this;
                let query = that.myForm.name;
                if (query.match(/\w+/g) != null) {
                    that.myForm.slug = query;
                }
                if (query == null || query == '') {
                    return false;
                }
                util.ajax.get('/backend/utils/fanyi/' + query).then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        that.myForm.slug = res.item
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        watch: {
            'myForm.name': {//监听路由改变
                handler(curVal, oldVal){
                    this.myForm.slug = curVal.replace(' ', '-');
                },
                deep: true
            }
        },
        mounted() {
            this.getData();
        }
    }
</script>
