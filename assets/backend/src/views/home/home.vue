<style lang="less">
    @import "./home.less";
    @import "../../styles/common.css";
</style>
<template>
    <div class="home-main">
        <Row>
            <Col span="20">
                <Row>
                    <Col span="6">
                        <infor-card
                            id-name="posts"
                            :end-val="count.posts"
                            iconType="document-text"
                            color="#2d8cf0"
                            :is-decimals="0"
                            intro-text="正常博文"
                        ></infor-card>
                    </Col>
                    <Col span="6" class-name="padding-left-5">
                    <infor-card
                            id-name="post_trash"
                            :end-val="count.post_trash"
                            iconType="trash-a"
                            color="#ffd572"
                            :is-decimals="0"
                            intro-text="回收站博文"
                    ></infor-card>
                    </Col>
                    <Col span="6" class-name="padding-left-5">
                        <infor-card
                            id-name="user_views"
                            :end-val="count.user_views"
                            iconType="ios-eye"
                            color="#64d572"
                            :is-decimals="2"
                            :iconSize="50"
                            intro-text="总浏览量"
                        ></infor-card>
                    </Col>
                    <Col span="6" class-name="padding-left-5">
                        <infor-card
                            id-name="comments"
                            :end-val="count.comments"
                            iconType="chatboxes"
                            color="#f25e43"
                            :is-decimals="0"
                            intro-text="评论"
                        ></infor-card>
                    </Col>
                </Row>
            </Col>
        </Row>
        <Spin size="large" fix v-if="loading"></Spin>
    </div>
</template>

<script>
import countUp from './components/countUp.vue';
import inforCard from './components/inforCard.vue';
import Util from '../../libs/util';

export default {
    name: 'home',
    components: {
        inforCard
    },
    data () {
        return {
            loading:false,
            count: {
                posts: 0,
                comments: 0,
                post_trash: 0,
                user_views: 1
            }
        };
    },
    created () {
        this.getMeta();
    },
    computed: {
        avatorPath () {
            return localStorage.avatorImgPath;
        }
    },
    methods: {
        getMeta: function () {
            let that = this;
            that.loading = true;
            Util.ajax.get('/backend/meta').then(function (response) {
                let data = response.data;
                if (data.status == 200) {
                    that.count = data.item;
                } else {
                    that.$Notice.warning({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
                that.loading = false;
            }).catch(function (error) {
                console.log(error);
                that.loading = false;
            });
        },
    }
};
</script>
