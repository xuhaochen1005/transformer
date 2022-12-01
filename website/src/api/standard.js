import request from '@/utils/request';
export function login(userInfo) {
    return request({
        url: '/user/token',
        method: 'POST',
        data: userInfo
    });
}
export function GetLibraries(param) {
    return request({
        url: '/std/libraries',
        method: 'GET',
        data: param
    });
}
export function GetFormulaParams(param) {
    return request({
        url: '/std/lfp',
        method: 'GET',
        data: param
    });
}
//# sourceMappingURL=standard.js.map