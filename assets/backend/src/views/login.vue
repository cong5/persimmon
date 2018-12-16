<style lang="less">
    @import './login.less';
</style>

<template>
    <div class="login" @keydown.enter="handleSubmit">
        <div class="login-con">
            <Card :bordered="false">
                <p slot="title">
                    <Icon type="log-in"></Icon>
                    欢迎登录
                </p>
                <div class="form-con">
                    <Form ref="loginForm" :model="form" :rules="rules">
                        <FormItem prop="username">
                            <Input v-model="form.email" placeholder="请输入用户名">
                            <span slot="prepend">
                                    <Icon :size="16" type="person"></Icon>
                                </span>
                            </Input>
                        </FormItem>
                        <FormItem prop="password">
                            <Input type="password" v-model="form.password" placeholder="请输入密码">
                            <span slot="prepend">
                                    <Icon :size="14" type="locked"></Icon>
                                </span>
                            </Input>
                        </FormItem>
                        <FormItem>
                            <Button @click="handleSubmit" type="primary" long>登录</Button>
                        </FormItem>
                    </Form>
                    <p class="login-tip">输入邮箱和密码登录。</p>
                    <Spin size="large" fix v-if="spinShow" class="login-spin">登陆中...</Spin>
                </div>
            </Card>
        </div>
    </div>
</template>

<script>
    import Cookies from 'js-cookie';
    import Util from '../libs/util';
    export default {
        data() {
            return {
                spinShow:false,
                form: {
                    email: '',
                    password: ''
                },
                rules: {
                    email: [
                        {required: true, message: '账号不能为空', trigger: 'blur'},
                        {type: 'email', message: '邮箱格式不正确', trigger: 'blur'}
                    ],
                    password: [
                        {required: true, message: '密码不能为空', trigger: 'blur'},
                        { type: 'string', min: 6, message: '密码长度需要大于6位', trigger: 'blur' }
                    ]
                }
            };
        },
        methods: {
            handleSubmit() {
                this.$refs.loginForm.validate((valid) => {
                    if (valid) {
                        let that = this;
                        this.spinShow = true;
                        Util.ajax.post('/backend/login', Util.stringify(this.form))
                            .then(function (response) {
                                that.spinShow = false;
                                let data = response.data;
                                console.log(data)
                                if (data.status === 200) {
                                    Cookies.set('user', data.item);
                                    that.$store.commit('setAvator');
                                    that.$Notice.success({
                                        title: data.info,
                                        desc: ''
                                    });
                                    setTimeout(function () {
                                        that.$router.push({
                                            name: 'home_index'
                                        });
                                    }, 1000);
                                } else {
                                    that.$Notice.error({
                                        title: data.info,
                                        desc: ''
                                    });
                                }
                            })
                            .catch(function (error) {
                                that.spinShow = false;
                                console.log(error);
                            });
                    }
                });
            },
            handleSpinCustom() {
                this.$Spin.show({
                    render: (h) => {
                        return h('div', [
                            h('Icon', {
                                'class': 'login-spin-icon-load',
                                props: {
                                    type: 'load-c',
                                    size: 18
                                }
                            }),
                            h('div', '登陆中...')
                        ])
                    }
                });
            }
        }
    };
</script>
