import request from '@/utils/request';
export function login(userInfo) {
    return request({
        url: '/user/token',
        method: 'POST',
        data: userInfo
    });
}
export function GetDesignProjects(param) {
    return request({
        url: '/design/projects',
        method: 'GET',
        data: param
    });
}
//# sourceMappingURL=projects.js.map