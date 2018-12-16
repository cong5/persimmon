<style lang="less">
    @import './post-list.less';
</style>

<template>
    <div class="pit-content">
        <div class="pit-action-btn">
            <Button @click="addPost()" icon="android-add"  class="myp-search-item"></Button>
            <Button @click="handleDistory('multi',{})" icon="trash-b"  class="myp-search-item"></Button>
            <Select v-model="category_id" @on-change="filterCategory" style="width:200px" class="myp-search-item">
                <Option v-for="item in categorys" :value="item.id" :key="item.id">{{ item.name }}</Option>
            </Select>
            <Input v-model="q" placeholder="请输入内容" style="width: 200px" class="myp-search-item">
                <Button slot="append" icon="ios-search" @click="searchBtn"></Button>
            </Input>
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

    </div>
</template>

<script>
    import util from '../../libs/util';

    export default {
        data() {
            return {
                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '标题',
                        width: 400,
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
                                tagsArr[i] = h('Tag', {}, tags[i].name)
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
                                        icon:'edit'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.postEditor(params.row)
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
                listData: [],
                category_id: 0,
                categorys: [],
                currentPage: 1,
                total: 0,
                pageSize: 20,
                listLoading: false,
                checkedAll: [],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
                q: ''
            }
        },
        methods: {
            filterCategory(value) {
                this.getData();
            },
            searchBtn(event) {
                this.getData();
            },
            getData() {
                let that = this;
                that.listLoading = true;
                let query = {
                    rows: that.pageSize,
                    categoryId: that.category_id,
                    keywords: that.q,
                    page: that.currentPage
                };

                util.ajax.get('/backend/posts', {params: query}).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.listData = res.list.data;
                        that.total = res.list.total;
                        that.currentPage = res.list.current_page;
                        that.listLoading = false;
                    } else {
                        this.$Notice.warning({
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
            handleDistory(type, row) {
                let that = this, idsParam = {};
                switch (type) {
                    case 'one':
                        if (parseInt(row.id) <= 0) {
                            this.$Notice.warning({
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
                            this.$Notice.warning({
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
                        util.ajax.post('/backend/posts/destroy', util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status === 200 ? '删除成功' : '删除失败',
                                desc: '',
                                duration: 2
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
            selectRow(row) {
                this.checkedAll = row;
            },
            selectAll(selection) {
                this.checkedAll = selection;
            },
            getCategorys() {
                let that = this;
                util.ajax.get('/backend/categories', {
                    params: {
                        rows: 999,
                        page:1
                    }
                }).then(function (response) {
                    let data = response.data;
                    if (data.status === 200) {
                        data.list.data.splice(0, 0, {id: 0, name: '全部', hidden: true, pid: 0});
                        that.categorys = data.list.data;
                    } else {
                        this.$Notice.error({
                            title: '数据获取失败',
                            desc: ''
                        });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            postEditor(row) {
                let path = '/posts/edit/' + row.id;
                this.$router.push({path: path});
            },
            addPost() {
                let path = "/posts/add";
                this.$router.push({path: path});
            }
        },
        mounted() {
            this.getCategorys();
            this.getData();
        }
    }
</script>
