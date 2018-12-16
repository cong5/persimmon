<template>
    <div class="pit-post-form" style="width: 60%">
        <Form ref="myForm" :model="myForm" :label-width="150"
              class="pit-common">
            <FormItem v-for="option in options" :value="option.id" :key="option.id" :label="option.title">
                <Input :type="option.data_type" v-model="myForm[option.name]"></Input>
            </FormItem>
            <FormItem>
                <Button @click="closeForm('myForm')">取 消</Button>
                <Button type="primary" @click="submitMyForm('myForm')">确 定</Button>
            </FormItem>
        </Form>
        <Spin size="large" fix v-if="editFormLoading"></Spin>
    </div>
</template>
<style type="text/css">

</style>
<script>
    import util from '../../libs/util';

    export default {
        data () {
            return {
                editFormLoading: false,
                options: [],
                myForm: {}
            };
        },
        created () {
            this.getData();
        },
        methods: {
            getData: function () {
                let dataType = {
                    1: 'text',
                    2: 'textarea',
                };
                let that = this;
                that.editFormLoading = true;
                util.ajax.get('/backend/settings').then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        let data = res.list;
                        if (data.length > 0) {
                            for (let index in data) {
                                that.myForm[data[index].name] = data[index].value;
                            }

                            for (let j = 0; j < data.length; j++) {
                                data[j].data_type = dataType[data[j].data_type];
                            }

                            that.options = data;
                        }
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

            submitMyForm: function () {
                let that = this;
                util.ajax.post('/backend/settings/update', util.stringify(that.myForm)).then(function (response) {
                    let res = response.data;
                    that.$Notice.warning({
                        title: res.status === 200 ? '更新成功' : '更新失败',
                        desc: ''
                    });
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        watch: {},
        mounted () {
        }
    };
</script>
