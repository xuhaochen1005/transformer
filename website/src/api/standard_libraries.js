import request from '@/utils/request';
export function GetStdAdblLibraries(param) {
    return request({
        url: '/std/adbl',
        method: 'GET',
        data: param
    });
}
//# sourceMappingURL=standard_libraries.js.map