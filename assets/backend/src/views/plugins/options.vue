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
                <Form ref="myForm" :rules="myRules" class="myForm" :label-width="100" :model="myForm" style="width: 80%">
                    <FormItem label="配置项说明" prop="title">
                        <Input v-model="myForm.title" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="配置项名称" prop="name">
                        <Input v-model="myForm.name" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="配置项值">
                        <Input :type="myForm.data_type" :autosize="myForm.data_type == 'textarea'" v-model="myForm.value" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="配置分组">
                        <Input v-model="myForm.group" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="配置项备注">
                        <Input type="textarea" :autosize="{ minRows: 3, maxRows: 5}" v-model="myForm.option_remark" auto-complete="off"></Input>
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

</style>
<script>
    import util from '../../libs/util';

    export default{
        data(){
            return {
                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '配置项说明',
                        key: 'title'
                    },
                    {
                        title: '配置项名称',
                        key: 'name'
                    },
                    {
                        title: '配置分组',
                        key: 'group'
                    },
                    {
                        title: '配置项备注',
                        key: 'remark'
                    },
                    {
                        title: '日期',
                        key: 'created_at',
                        render: (h, params) => {
                            let time = util.timeFormat(params.row.created_at);
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
                    title: [
                        {required: true, type: "string", message: '请填写配置项说明', trigger: 'blur'}
                    ],
                    name: [
                        {required: true, type: "string", message: '请填写配置项名称', trigger: 'blur'},
                        {pattern: /^[a-zA-Z0-9_]+$/, message: '只允许英文或者拼音,横杠(-)', trigger: 'blur'}
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
            getData: function () {
                let that = this;
                that.listLoading = true;
                util.ajax.get('/backend/options', {
                    params: {
                        rows: this.pageSize
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
                util.ajax.get('/backend/options/' + row.id).then(function (response) {
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
                        let ids = util.getIdByArr(that.checkedAll);
                        if (ids.length <= 0) {
                            that.$Notice.open({
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
                        util.ajax.post('/backend/options/destroy', util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status === 200 ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type === 'one') {
                                util.removeByValue(that.listData, row.id);
                            } else {
                                for (let index in that.checkedAll) {
                                    util.removeByValue(that.listData, that.checkedAll[index].id);
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
                        util.ajax.post('/backend/options/update', util.stringify(that.myForm)).then(function (response) {
                            let res = response.data;
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
                        util.ajax.post('/backend/options/store', util.stringify(that.myForm)).then(function (response) {
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
                                        that.$notify({
                                            title: '警告',
                                            message: error.response.data[index][0],
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
