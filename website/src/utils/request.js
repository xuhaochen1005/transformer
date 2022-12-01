import qs from 'qs';
import axios from 'axios';
import { ElMessage } from 'element-plus';
import { useStore } from '@/store';
import { UserMutation } from '@/store/user/mutation';
const notify = (msg) => {
    ElMessage.error({
        type: 'error',
        message: msg,
        showClose: true
    });
};
const request = axios.create({
    baseURL: process.env.VUE_APP_BASE_API
});
request.interceptors.request.use(config => {
    config.headers['token'] = useStore().state.user.token;
    if (config.timeout && config.timeout > 0) {
        config.headers['TIMEOUT'] = config.timeout / 1000;
    }
    else {
        config.timeout = 5000;
        config.headers['TIMEOUT'] = 5;
    }
    if (config.params) {
        const params = {};
        Object.keys(config.params).forEach(key => {
            if (config.params[key]) {
                params[key] = config.params[key];
            }
        });
        delete config.params;
        config.url += '?' + qs.stringify(params, { arrayFormat: 'repeat' });
    }
    return config;
}, error => {
    notify('请求出错：' + error.toString());
    return { data: { code: 0 } };
});
request.interceptors.response.use(response => {
    if (response.data.code !== 200) {
        notify('请求出错：' + response.data.error + '。错误编号：' + response.headers['x-request-id']);
        if (response.data.code === 401) {
            useStore().commit(UserMutation.LOGOUT);
        }
    }
    return response;
}, error => {
    notify('请求出错：' + error.toString());
    return { data: { code: 0 } };
});
export default request;
//# sourceMappingURL=request.js.map