<template>
    <div class="pit-content">
        <div class="pit-action-btn">
            <Button @click="handleCreate" icon="plus"></Button>
            <Button @click="handleDistory('multi',{})" icon="trash-b"></Button>
        </div>

        <div class="data-list">
            <Table :loading="listLoading" :columns="tableColumns" :data="listData" stripe @on-select="selectRow" @on-select-all="selectAll"></Table>
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
                    <FormItem label="名称" prop="name">
                        <Input v-model="myForm.name" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="链接" prop="url">
                        <Input v-model="myForm.url" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="Logo">
                        <Input v-model="myForm.logo" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="分组">
                        <Input v-model="myForm.group" auto-complete="off"></Input>
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
<style type="text/css">
    .links {
        text-decoration: none;
        color: #000
    }
</style>
<script>
    import Util from '../../libs/util';

    export default {
        data() {
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
                    logo: '',
                    group: '',
                },
                myRules: {
                    name: [
                        {required: true, type: "string", message: '请填写链接名称', trigger: 'blur'}
                    ],
                    url: [
                        {required: true, type: "string", message: '请填写链接地址', trigger: 'blur'}
                    ]
                },
                editFormVisible: false,
                editFormLoading: false,
                listLoading: false,
                myFormTitle: '编辑',
                checkedAll: [],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],

                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '名称',
                        key: 'name'
                    },
                    {
                        title: '地址',
                        key: 'url',
                        render: (h, params) => {
                            return h('a', {
                                attrs: {
                                    'href': params.row.url,
                                    'target': '_blank'
                                }
                            }, params.row.url);
                        }
                    },
                    {
                        title: 'Logo',
                        key: 'logo',
                        render: (h, params) => {
                            return h('img', {
                                attrs: {
                                    'src': params.row.logo
                                },
                                style: {
                                    width: '100px'
                                }
                            });
                        }
                    },
                    {
                        title: '日期',
                        key: 'created_at',
                        render: (h, params) => {
                            let time = Util.timeFormat(params.row.created_at);
                            return h('span', time);
                        }
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        icon: 'edit',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.handleEdit(params.row)
                                        }
                                    }
                                }, ''),
                                h('Button', {
                                    props: {
                                        icon: 'trash-b',
                                        size: 'small'
                                    },
                                    on: {
                                        click: () => {
                                            this.handleDistory('one', params.row);
                                        }
                                    }
                                }, '')
                            ]);
                        }
                    }
                ],
            }
        },
        methods: {
            getData: function () {
                let that = this;
                that.listLoading = true;
                Util.ajax.get('/backend/links', {
                    params: {
                        rows: that.pageSize,
                        page: this.currentPage
                    }
                }).then(function (response) {
                    let res = response.data;
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
            handleSizeChange(val) {
                //console.log(`每页 ${val} 条`);
                this.pageSize = val;
                this.getData();
            },
            handleCurrentChange(val) {
                this.currentPage = val;
                this.getData();
                //console.log(`当前页: ${val}`);
            },
            handleCreate: function () {
                let that = this;
                that.myFormTitle = '新增';
                that.myForm.id = 0;
                that.editFormVisible = true;
            },
            handleEdit: function (row) {
                let that = this;
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                that.editFormVisible = true;
                Util.ajax.get('/backend/links/' + row.id).then(function (response) {
                    let res = response.data;
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
            handleDistory: function (type, row) {
                let that = this, idsParam = {};
                switch (type) {
                    case 'one':
                        if (parseInt(row.id) <= 0) {
                            that.$Notice.open({
                                title: '请选择需要删除的数据',
                                desc: ''
                            });
                            return false;
                        }
                        idsParam = {ids: [row.id]};
                        break;
                    case 'multi':
                        var ids = that.util.getIdByArr(that.checkedAll);
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
                        Util.ajax.post('/backend/links/destroy', Util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.warning({
                                title: res.status === 200 ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type === 'one') {
                                that.util.removeByValue(that.listData, row.id);
                            } else {
                                for (var index in that.checkedAll) {
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
                        Util.ajax.post('/backend/links/update', that.myForm).then(function (response) {
                            let res = response.data;
                            that.$Notice.warning({
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
                        Util.ajax.post('/backend/links', that.myForm).then(function (response) {
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
            closeForm: function () {
                this.editFormVisible = false;
                this.$refs['myForm'].resetFields();
                this.myForm = {
                    id: 0,
                    name: '',
                    url: '',
                    logo: '',
                    group: '',
                };
                console.log('closeForm');
            },
            selectRow(row) {
                this.checkedAll = row;
            },
            selectAll(selection) {
                this.checkedAll = selection;
            },
        },
        watch: {},
        mounted() {
            this.getData();
        }
    }
</script>
