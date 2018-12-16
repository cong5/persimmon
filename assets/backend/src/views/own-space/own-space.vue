<style lang="less">
    @import './own-space.less';
</style>

<template>
    <div>
        <Card>
            <p slot="title">
                <Icon type="person"></Icon>
                个人信息
            </p>
            <div>
                <Form ref="myForm" :model="myForm" :label-width="100" label-position="right" :rules="myRules">
                    <FormItem label="用户名：">
                        <span>{{ userName }}</span>
                    </FormItem>
                    <FormItem label="旧密码：" prop="old_password">
                        <div style="display:inline-block;width:300px;">
                            <Input type="password" v-model="myForm.old_password">
                                <Icon type="ios-locked-outline" slot="prepend"></Icon>
                            </Input>
                        </div>
                    </FormItem>
                    <FormItem label="新密码：" prop="password">
                        <div style="display:inline-block;width:300px;">
                            <Input type="password" v-model="myForm.password">
                                <Icon type="ios-locked-outline" slot="prepend"></Icon>
                            </Input>
                        </div>
                    </FormItem>
                    <FormItem label="确认密码：" prop="password_confirmation">
                        <div style="display:inline-block;width:300px;">
                            <Input type="password" v-model="myForm.password_confirmation">
                                <Icon type="ios-locked-outline" slot="prepend"></Icon>
                            </Input>
                        </div>
                    </FormItem>
                    <div>
                        <Button type="text" style="width: 100px;" @click="closeForm">取消</Button>
                        <Button type="primary" style="width: 100px;" :loading="saveLoading" @click="submitMyForm">保存
                        </Button>
                    </div>
                </Form>
            </div>
        </Card>
    </div>
</template>

<script>
    import Util from '../../libs/util';
    import Cookies from 'js-cookie';

    export default {
        data() {
            return {
                userName: '',
                saveLoading: false,
                myForm: {
                    old_password: '',
                    password: '',
                    password_confirmation: ''
                },
                myRules: {
                    old_password: [
                        {required: true, type: "string", message: '请填写旧密码', trigger: 'blur'},
                    ],
                    password: [
                        {required: true, type: "string", message: '请填写新密码', trigger: 'blur'},
                        {min: 8, max: 64, message: '长度在 8 到 64 个字符', trigger: 'blur'}
                    ],
                    password_confirmation: [
                        {required: true, type: "string", message: '请再填写一次密码', trigger: 'blur'},
                        {min: 8, max: 64, message: '长度在 8 到 64 个字符', trigger: 'blur'}
                    ],
                }
            }
        },
        created() {
        },
        methods: {
            submitMyForm() {
                let that = this;
                that.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }
                    Util.ajax.post('/backend/user', that.myForm).then(function (response) {
                        let res = response.data;
                        that.$Notice.open({
                            title: res.status == 200 ? '信息更新成功' : '信息更新失败',
                            desc: ''
                        });
                        if (res.status == 200) {
                            that.closeForm('myForm');
                        }
                        setTimeout(function () {
                            Cookies.remove('user');
                            that.$router.push({path: '/login'});
                        }, 2 * 1000);
                    }).catch(function (error) {
                        if (error.response) {
                            if (error.response.status == 422) {
                                for (var index in error.response.data) {
                                    that.$Notice.open({
                                        title: '警告',
                                        message: error.response.data[index][0]
                                    });
                                }
                            }
                        } else {
                            console.log(error);
                        }
                    });
                });
            },
            closeForm() {
                this.saveLoading = false;
                this.$refs['myForm'].resetFields();
                this.myForm = {
                    old_password: '',
                    password: '',
                    password_confirmation: ''
                };
                console.log('closeForm');
            }
        },
        mounted() {
            let user = Cookies.get('user');
            let userInfo = JSON.parse(user);
            this.userName = userInfo.name;
        }
    }
</script>
