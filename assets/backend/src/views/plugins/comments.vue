<style lang="less">
    @import './comments.less';
</style>

<template>
    <div class="pit-content">
        <div class="data-list">
            <div class="common-item">
                <div class="common-list-head">
                    <Row type="flex">
                        <Col span="4" order="1">
                        作者
                        </Col>
                        <Col span="12" order="2">
                        评论
                        </Col>
                        <Col span="4" order="3">
                        回复于
                        </Col>
                        <Col span="4" order="4">
                        日期
                        </Col>
                    </Row>
                </div>
                <div class="comment-list-body">
                    <Row class="comment-item" type="flex" v-for="row in listData" :key="row.id">
                        <Col span="4" order="1">
                        <div class="comment-item-author">
                            <div class="comment-item-avatar">
                                <strong>
                                    <img v-bind:src="'https://cn.gravatar.com/avatar/' + row.md5 + '?d=identicon&s=32'"
                                         alt=""/>{{ row.name }}
                                </strong>
                            </div>
                            <div class="comment-item-meta">
                                <a v-bind:href="row.email" target="_blank">{{ row.email }}</a>
                                <br/>
                                <a v-bind:href="row.url" target="_blank">{{ row.url }}</a>
                            </div>
                        </div>
                        </Col>
                        <Col span="12" order="2">
                        <p class="comment-item-content" v-html="row.content"></p>
                        <p>
                            <a href="javascript:void(0);" @click="handleEdit(row)">编辑</a> |
                            <template v-if="row.status == 1">
                                <a href="javascript:void(0);" @click="spam(row,3)">垃圾评论</a> |
                            </template>
                            <template v-else-if="row.status == 2">
                                <a href="javascript:void(0);" @click="spam(row,1)">审核</a> |
                            </template>
                            <template v-else-if="row.status == 3">
                                <a href="javascript:void(0);" @click="spam(row,2)">不是垃圾评论</a> |
                            </template>
                            <a href="javascript:void(0);" @click="handleDistory('one', row)">删除</a>
                        </p>
                        </Col>
                        <Col span="4" order="3">
                        <div class="comment-item-title">
                            <a class="myp-bold" href="javascript:void(0)" @click="postEditor(row.posts_id)"
                               target="_blank">{{ row.title }}</a>
                            <br/>
                            <a v-bind:href="'/post/' + row.slug" target="_blank">查看文章</a>
                        </div>
                        </Col>
                        <Col span="4" order="4">
                        <div class="comment-item-date">
                            {{ row.created_at * 1000 | moment("YYYY-MM-DD A h:mm") }}
                        </div>
                        </Col>
                    </Row>
                </div>
            </div>
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
                    <FormItem label="链接">
                        <Input v-model="myForm.url" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="E-Mail" prop="email">
                        <Input v-model="myForm.email" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="内容" prop="markdown">
                        <Input type="textarea" autosize v-model="myForm.markdown" auto-complete="off"></Input>
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
    import Util from '../../libs/util';

    export default {
        data () {
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
                    name: [
                        {required: true, type: 'string', message: '请填写链接名称', trigger: 'blur'}
                    ],
                    email: [
                        {required: true, type: 'string', message: '请填写邮箱', trigger: 'blur'}
                    ],
                    markdown: [
                        {required: true, type: 'string', message: '请填写评论内容', trigger: 'blur'}
                    ]
                },
                editFormVisible: false,
                editFormLoading: false,
                listLoading: false,
                myFormTitle: '编辑',
                checkedAll: [],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
            };
        },
        methods: {
            getData: function () {
                let that = this;
                that.listLoading = true;
                Util.ajax.get('/backend/comments', {
                    params: {
                        rows: that.pageSize,
                        page: this.currentPage
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.listData = Util.emailToMd5(res.list.data);
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
            handleSizeChange (val) {
                //console.log(`每页 ${val} 条`);
                this.pageSize = val;
                this.getData();
            },
            handleCurrentChange (val) {
                this.currentPage = val;
                this.getData();
                //console.log(`当前页: ${val}`);
            },
            spam (row, status) {
                let id = row.id;
                let that = this;
                Util.ajax.post('/backend/comments/spam', Util.stringify({
                    id: id,
                    status: status
                })).then(function (response) {
                    let res = response.data;
                    that.$Notice.warning({
                        title: res.status === 200 ? '操作成功' : '操作失败',
                        desc: ''
                    });
                    that.$nextTick(() => {
                        that.getData();
                    })
                }).catch(function (error) {
                    console.log(error);
                });
            },
            handleCreate () {
                let that = this;
                that.myFormTitle = '新增';
                that.myForm.id = 0;
                that.editFormVisible = true;
            },
            handleEdit (row) {
                let that = this;
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                that.editFormVisible = true;
                Util.ajax.get('/backend/comments/' + row.id).then(function (response) {
                    let res = response.data;
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
            handleDistory (type, row) {
                let that = this, idsParam = {};
                switch (type) {
                    case 'one':
                        if (parseInt(row.id) <= 0) {
                            that.$Notice.warning({
                                message: '请选择需要删除的数据',
                                type: 'warning'
                            });
                            return false;
                        }
                        idsParam = {ids: [row.id]};
                        break;
                    case 'multi':
                        let ids = Util.getIdByArr(that.checkedAll);
                        if (ids.length <= 0) {
                            that.$Notice.warning({
                                message: '请选择需要删除的数据',
                                type: 'warning'
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
                        Util.ajax.post('/backend/comments/destroy', Util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.warning({
                                title: res.status === 200 ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type === 'one') {
                                Util.removeByValue(that.listData, row.id);
                            } else {
                                for (let index in that.checkedAll) {
                                    Util.removeByValue(that.listData, that.checkedAll[index].id);
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
            postEditor (id) {
                let path = '/posts/edit/' + id;
                this.$router.push({path: path});
            },
            submitMyForm () {
                let that = this;
                that.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }

                    if (that.myForm.id > 0) {
                        Util.ajax.post('/backend/comments/update', Util.stringify(that.myForm)).then(function (response) {
                            let res = response.data;
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
                        Util.ajax.post('/backend/links', that.myForm).then(function (response) {
                            let res = response.data;
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
            closeForm () {
                this.editFormVisible = false;
                this.$refs['myForm'].resetFields();
                this.myForm = {
                    id: 0,
                    name: '',
                    url: '',
                    logo: '',
                    group: '',
                };
            },
            selectRow (row) {
                this.checkedAll = row;
            },
            selectAll (selection) {
                this.checkedAll = selection;
            },
        },
        watch: {},
        mounted () {
            this.getData();
        }
    };
</script>
