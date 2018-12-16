<style lang="less">
    @import "./post-trash.less";
</style>

<template>
    <div class="pit-content">
        <div class="pit-action-btn">
            <Button @click="handleDistory('multi',{})" icon="trash-a" class="myp-search-item"></Button>
            <Select v-model="category_id" clearable @change="filterCategory" placeholder="请选择" class="myp-search-item"
                    style="width:200px">
                <Option v-for="item in categorys" :key="item.id" :value="item.id">{{ item.name }}</Option>
            </Select>
            <Input v-model="q" placeholder="请输入内容" icon="search" style="width: 200px" class="myp-search-item">
            <Button slot="append" icon="ios-search" @click="searchBtn"></Button>
            </Input>
        </div>

        <div class="data-list">
            <Table :loading="listLoading" :columns="tableColumns" :data="listData" stripe @on-select="selectRow"
                   @on-select-all="selectAll"></Table>
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

    </div>
</template>

<script>
    import Util from '../../libs/util';

    export default {
        data () {
            return {
                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '标题',
                        key: 'title'
                    },
                    {
                        title: '分类',
                        key: 'categories',
                        render: (h, params) => {
                            return h('span', params.row.categories.name);
                        }
                    },
                    {
                        title: '标签',
                        key: 'tags',
                        render: (h, params) => {
                            let tagsArr = [];
                            let tags = params.row.tags;
                            for (let i = 0, len = tags.length; i < len; i++) {
                                tagsArr[i] = h('Tag', {}, tags[i].name);
                            }
                            return h('div', tagsArr);
                        }
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
                                        icon: 'refresh'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.handleRestore(params.row);
                                        }
                                    }
                                }, ''),
                                h('Button', {
                                    props: {
                                        size: 'small',
                                        icon: 'trash-b'
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
                category_id: '',
                categorys: [],
                currentPage: 1,
                total: 0,
                pageSize: 20,
                listLoading: false,
                checkedAll: [],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
                q: ''
            };
        },
        methods: {
            filterCategory: function (value) {
                this.getData();
            },
            searchBtn: function (event) {
                this.getData();
            },
            getData: function () {
                let that = this;
                that.listLoading = true;
                let query = {
                    rows: that.pageSize,
                    categoryId: that.category_id,
                    q: that.q,
                    page: that.currentPage
                };

                Util.ajax.get('/backend/posts-trash', {params: query}).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.listData = res.list.data;
                        that.total = res.list.total;
                        that.currentPage = res.list.current_page;
                        that.listLoading = false;
                    } else {
                        that.$message({
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
                //console.log(`当前页: ${val}`);
            },
            handleRestore: function (row) {
                let that = this;

                that.$Modal.confirm({
                    title: '确认恢复选中的文章吗?',
                    content: '<p>确认恢复选中的文章吗?</p>',
                    onOk: () => {
                        that.listLoading = true;
                        Util.ajax.post('/backend/posts-trash/update', Util.stringify({'ids': [row.id]})).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status === 200 ? '恢复成功' : '恢复失败',
                                desc: ''
                            });
                            Util.removeByValue(that.listData, row.id);
                        }).catch(function (error) {
                            console.log(error);
                        });
                    },
                    onCancel: () => {
                        that.listLoading = false;
                    }
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
                        let ids = Util.getIdByArr(that.checkedAll);
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
                        Util.ajax.delete('/backend/posts-trash/destroy', util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.open({
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
            selectRow (row) {
                this.checkedAll = row;
            },
            selectAll (selection) {
                this.checkedAll = selection;
            },
            getCategorys: function () {
                let that = this;
                Util.ajax.get('/backend/categories', {
                    params: {
                        rows: 999,
                        page: 1
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        res.list.data.splice(0, 0, {id: 0, name: '全部', hidden: true, pid: 0});
                        that.categorys = res.list.data;
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
            setTopCategorys: function () {
                let categorys = this.listData.concat();
                categorys.splice(0, 0, {id: 0, name: '全部', hidden: true, pid: 0});
                this.categorys = categorys;
            }
        },
        mounted () {
            this.getCategorys();
            this.getData();
        }
    };
</script>
