<template>
    <div class="pit-content">
        <div class="pit-action-btn">
            <Button @click="handleCreate" icon="plus"></Button>
            <Button @click="handleDistory('multi',{})" icon="trash-b"></Button>
        </div>
        <div class="data-list">
            <Table :loading="listLoading" :columns="tableColumns" :data="listData" stripe @on-select="selectRow" @on-select-all="selectAll"></Table>
        </div>

        <Modal :title="myFormTitle" v-model="editFormVisible" @on-ok="submitMyForm" @on-cancel="closeForm">
            <div class="pit-dialog-edit-form">
                <Form ref="myForm" :rules="myRules" class="myForm" :model="myForm" label-position="right"
                      :label-width="100" style="width: 90%;">
                    <FormItem label="分类名称" prop="name">
                        <Input v-model="myForm.name" auto-complete="off" @on-blur="translateWords"
                               @on-enter="translateWords"></Input>
                    </FormItem>
                    <FormItem label="分类别名" prop="slug">
                        <Input v-model="myForm.slug" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="分类描述">
                        <Input type="textarea" v-model="myForm.description"></Input>
                    </FormItem>
                    <FormItem label="父分类">
                        <Select v-model="myForm.pid" placeholder="请选择父分类">
                            <Option v-for="item in categorys" :value="item.id" :key="item.id">{{ item.name }}
                            </Option>
                        </Select>
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

    export default {
        data() {
            return {
                listData: [],
                categorys: [],
                myForm: {
                    id: 0,
                    name: '',
                    slug: '',
                    description: '',
                    pid: 0,
                },
                myRules: {
                    name: [
                        {required: true, type: "string", title: '请填写分类名称', trigger: 'blur'}
                    ],
                    slug: [
                        {required: true, type: "string", title: '请填写分类别名', trigger: 'blur'},
                        {pattern: /^[a-zA-Z0-9_-]+$/, title: '只允许英文或者拼音,横杠(-),下划线(_)', trigger: 'blur'}
                    ]
                },
                editFormVisible: false,
                editFormLoading: false,
                listLoading: false,
                myFormTitle: '编辑',
                checkedAll: [],
                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '分类名称',
                        key: 'name'
                    },
                    {
                        title: '分类别名',
                        key: 'slug'
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
                                        size: 'small',
                                        icon:'edit'
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
                                        size: 'small',
                                        icon:'trash-a'
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
                util.ajax.get('/backend/categories', {
                    params: {
                        page: 1,
                        rows: 20
                    }
                }).then(function (response) {
                    let res = response.data;
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
                util.ajax.get('/backend/categories/' + row.id).then(function (response) {
                    let res = response.data;
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
                        let ids = util.getIdByArr(that.checkedAll);
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
                        util.ajax.post('/backend/categories/destroy',util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.warning({
                                title: res.status === 200 ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type === 'one') {
                                util.removeByValue(that.listData, row.id);
                            } else {
                                for (var index in that.checkedAll) {
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
                        util.ajax.put('/backend/categories/update', that.myForm).then(function (response) {
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
                        util.ajax.post('/backend/categories/store', that.myForm).then(function (response) {
                            let res = response.data;
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
            closeForm: function () {
                this.editFormVisible = false;
                this.$refs['myForm'].resetFields();
                this.myForm = {
                    id: 0,
                    name: '',
                    slug: '',
                    description: '',
                    pid: 0,
                };
                console.log('closeForm');
            },
            setTopCategorys: function () {
                let categorys = this.listData.concat();
                categorys.splice(0, 0, {id: 0, name: '顶级分类', hidden: true, pid: 0});
                this.categorys = categorys;
            },
            selectRow(row) {
                this.checkedAll = row;
            },
            selectAll(selection) {
                this.checkedAll = selection;
            },
            translateWords(event) {
                let that = this;
                let query = that.myForm.name;
                if (query.match(/\w+/g) != null) {
                    that.myForm.slug = query;
                }
                if (query == null || query === '') {
                    return false;
                }
                util.ajax.get('/backend/utils/fanyi/' + query).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.myForm.slug = res.item
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        mounted() {
            this.getData();
        }
    }
</script>
