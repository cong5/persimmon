import axios from 'axios';
import Cookies from 'js-cookie';
import Moment from 'moment/moment';

let util = {};

let csrfToken = document.querySelector('meta[name="csrf-token"]').getAttribute('content');
axios.defaults.headers.common = {
    'X-Requested-With': 'XMLHttpRequest',
    'Accept': 'application/json',
    'X-CSRFToken': csrfToken
};
util.ajax = axios.create({
    timeout: 30000
});

util.remenberMe = function (data) {
    let viewer = {
        name: data.name,
        email: data.email,
        url: data.url
    };
    Cookies.set('viewer', viewer);
};

util.getViewer = function () {
    let viewer = {
        name: '',
        email: '',
        url: ''
    };
    let viewerInfo = Cookies.get('viewer');
    if (viewerInfo !== undefined) {
        viewer = JSON.parse(viewerInfo);
    }
    return viewer;
};

util.checkUrl = function (str, msg) {
    let RegUrl = new RegExp();
    RegUrl.compile('^[A-Za-z]+://[A-Za-z0-9-_]+\\.[A-Za-z0-9-_%&\?\/.=]+$');
    if (!RegUrl.test(str)) {
        that.$layer.toast({
            icon: 'icon-check',
            content: msg,
            time: 2000
        });
        return false;
    }
    return true;
};

util.timeFormat = function (timestamp) {
    if(timestamp <= 0){
        return ""
    }
    let time = Moment.unix(timestamp);
    return time.format('ddd, DD MMM YYYY HH:mm:ss');
};

export default util;
