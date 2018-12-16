import Vue from 'vue';
import layer from 'vue-layer-mobile'
import Prism from './libs/prism.min';
import './styles/app.css';
import 'vue-layer-mobile/need/layer.css'
import mycomment from './views/comment.vue';
Vue.use(layer);
Prism.highlightAll();

new Vue({
    el: '#comment',
    components: {mycomment},
    data () {
        return {};
    },
    methods: {},
    mounted () {
    },
    watch: {}
});
