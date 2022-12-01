import request from '@/utils/request';
export function login(userInfo) {
    return request({
        url: '/user/token',
        method: 'POST',
        data: userInfo
    });
}
export function GetDepartments(param) {
    return request({
        url: '/department/',
        method: 'GET',
        data: param
    });
}
//# sourceMappingURL=department.js.map