<style lang="less">
    @import '../../styles/common.css';
    @import './components/table.less';

    .navigation {
        padding: 2em;
    }

    .pit-nav-item {
        border-radius: 4px;
        min-height: 36px;
        background: #99a9bf;
        margin-bottom: 10px;
        padding-left: 10px;
        line-height: 2em;
    }

    .pit-header-title {
        line-height: 36px;
        color: #48576a;
        font-size: 14px;
    }

    .categorys li {
        list-style: none;
        position: relative;
        border-radius: 4px;
        background: rgba(153, 169, 191, 0.19);
        margin-bottom: 10px;
        padding: 2px 0 2px 10px;
        line-height: 2em;
        font-size: 14px;
    }

    .categorys li a {
        color: #585858;
        text-decoration: none;
    }

    .categorys li span {
        padding: 0 10px;
        position: absolute;
        right: 0;
        top: 2px;
    }

    .categorys li span:active {
        background: rgba(153, 169, 191, 0.5);
        border-radius: 5px;
    }

    .ivu-card-body {
        padding: 0;
    }
</style>
<template>
    <div class="navigation">
        <Row :gutter="20">
            <Col span="8">
            <div class="grid-content bg-purple">
                <Collapse v-model="activeMenuAddCollapse">
                    <Panel name="1">
                        分类添加
                        <p slot="content">
                        <ul class="categorys">
                            <li v-for="category in categorys">
                                <a href="javascript:void(0);">
                                    <Icon type="ios-arrow-right"></Icon>
                                    {{ category.name }}
                                </a>
                                <span @click="addToNav(category)">
                                    <Icon type="ios-plus-outline"></Icon>
                                </span>
                            </li>
                        </ul>
                        </p>
                    </Panel>
                    <Panel name="2">
                        添加导航
                        <p slot="content">
                        <Form ref="myForm" :rules="myRules" :model="myForm" :label-width="100">
                            <FormItem label="菜单名称" prop="name">
                                <Input size="small" v-model="myForm.name"></Input>
                            </FormItem>
                            <FormItem label="菜单地址" prop="url">
                                <Input size="small" v-model="myForm.url"></Input>
                            </FormItem>
                            <FormItem>
                                <Button type="primary" @click="onSubmit('myForm')">确定</Button>
                                <Button @click="onReset('myForm')">取消</Button>
                            </FormItem>
                        </Form>
                        </p>
                    </Panel>
                </Collapse>
            </div>
            </Col>
            <Col span="16">
            <div class="grid-content bg-purple-light">
                <Card class="box-card">
                    <div slot="header" class="clearfix">
                        <span class="pit-header-title">菜单列表</span>
                        <Button icon="edit" @click="updateMenu" size="small" type="primary"
                                style="float: right;margin-top: 5px;">保存
                        </Button>
                    </div>
                    <DragableTable refs="table" :columnsList="tableColumns" :tableData="navigations" :start="onStart" :end="onEnd"></DragableTable>
                </Card>
                <Spin size="large" fix v-if="listLoading"></Spin>
            </div>
            </Col>
        </Row>
    </div>
</template>
<script>
    import Util from '../../libs/util';
    import DragableTable from './components/dragableTable.vue';

    export default {
        name: 'dragable-table',
        components: {
            DragableTable
        },
        data() {
            return {
                myForm: {
                    name: '',
                    url: '',
                    sorting: 0
                },
                myRules: {
                    name: [
                        {required: true, type: "string", message: '请填写菜单名称', trigger: 'blur'},
                        {min: 1, message: '长度大于1个字符', trigger: 'blur'}
                    ],
                    url: [
                        {required: true, type: "string", message: '请填写url地址', trigger: 'blur'},
                        {min: 1, message: '长度大于1个字符', trigger: 'blur'}
                    ]
                },
                sortList: {},
                navigations: [],
                categorys: [],
                tableColumns: [
                    {
                        title: '拖拽',
                        key: 'drag',
                        width: 90,
                        align: 'center',
                        render: (h) => {
                            return h(
                                'Icon',
                                {
                                    props: {
                                        type: 'arrow-move',
                                        size: 24
                                    }
                                }
                            );
                        }
                    },
                    {
                        title: '名称',
                        key: 'name'
                    },
                    {
                        title: '地址',
                        key: 'url'
                    },
                    {
                        title: '删除',
                        type: '',
                        width: 80,
                        align: 'center',
                        render: (h, params) => {
                            return h(
                                'Button',
                                {
                                    props: {
                                        type: 'text',
                                        icon: 'ios-close-outline'
                                    },
                                    on: {
                                        click: () => {
                                            this.removeItem(params.row)
                                        }
                                    }
                                }
                            );
                        }
                    },
                ],
                table: {
                    hasDragged: false,
                    isDragging: false,
                    oldIndex: 0,
                    newIndex: 0,
                    draggingRecord: []
                },
                listLoading: false,
                editLoading: false,
                activeMenuAddCollapse: '2'
            }
        },
        methods: {
            onSubmit() {
                let that = this;
                that.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }
                    that.navigations.splice(-1, 0, this.myForm);
                    that.myForm = {
                        name: '',
                        url: '',
                        sorting: 0
                    };
                    //console.log(that.navigations);
                });
            },
            onReset() {
                this.myForm = {
                    name: '',
                    url: '',
                    sorting: 0
                };
            },
            getData() {
                let that = this;
                that.listLoading = true;
                Util.ajax.get('/backend/navigation').then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        /*
                        res.list.sort(function (x, y) {
                            return x.sorting > y.sorting ? 1 : -1;
                        });*/
                        that.navigations = res.list;
                        //获取排序值
                        if (that.navigations.length > 0) {
                            for (let index in that.navigations) {
                                that.sortList[that.navigations[index]['name']] = that.navigations[index]['sorting'];
                            }
                            //console.log( that.sortList);
                        }
                    } else {
                        console.log('数据获取失败或者数据为空。')
                    }
                    that.listLoading = false;
                }).catch(function (error) {
                    console.log(error);
                    that.listLoading = false;
                });
            },
            getCategory() {
                let that = this;
                that.listLoading = true;
                Util.ajax.get('/backend/categories', {
                    params: {
                        rows: 999,
                        page: 1
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        that.categorys = res.list.data;
                    } else {
                        that.$Notice.error({
                            title: '数据获取失败',
                            desc: ''
                        });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            removeItem(item) {
                let that = this;
                that.$Modal.confirm({
                    title: '确认删除选中的记录吗?',
                    content: '<p>您确认删除选中的记录吗?</p>',
                    onOk: () => {
                        for (let index in that.navigations) {
                            if (that.navigations[index].name == item.name) {
                                that.navigations.splice(index, 1);
                            }
                        }
                        that.updateMenu();
                    },
                    onCancel: () => {
                        that.listLoading = false;
                    }
                });
            },
            updateMenu() {
                let that = this;
                that.editLoading = true;
                let data = {
                    nav: JSON.stringify(that.navigations)
                }
                Util.ajax.post('/backend/navigation/update', Util.stringify(data)).then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        that.$Notice.success({
                            title: '更新成功',
                            desc: ''
                        });
                    } else {
                        that.$Notice.error({
                            title: '更新失败',
                            desc: ''
                        });
                    }
                    that.editLoading = false;
                }).catch(function (error) {
                    console.log(error);
                    that.editLoading = false;
                });
            },
            onStart(el) {
                this.table.oldIndex = el.oldIndex;
                this.table.hasDragged = true;
                this.table.isDragging = true;
            },
            onEnd(el) {
                this.table.newIndex = el.newIndex;
                this.table.isDragging = false;
                this.table.draggingRecord.unshift({
                    from: this.table.oldIndex + 1,
                    to: this.table.newIndex + 1
                });
                let newPos = this.navigations[this.table.oldIndex];
                let oldPos = this.navigations[this.table.newIndex];
                this.navigations[this.table.oldIndex] = oldPos;
                this.navigations[this.table.newIndex] = newPos;
                this.updateMenu();
            },
            addToNav: function (category) {
                let that = this;
                let navName = category.category_name;
                let navUrl = '/categories/' + category.category_flag;
                for (let i = 0, len = that.navigations.length; i < len; i++) {
                    if (navName == that.navigations[i].name) {
                        that.$Notice.warning({
                            title: '该菜单已经存在',
                            desc: ''
                        });
                        return false;
                        break;
                    }
                }
                let nav = {
                    name: navName,
                    url: navUrl
                };
                that.navigations.push(nav);
                that.updateMenu();
                that.getData();
            }
        },
        mounted() {
            this.getData();
            this.getCategory();
        }
    }
</script>
