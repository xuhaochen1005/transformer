import request from '@/utils/request';
export function GetStdlcsLibraries(query) {
    return request({
        url: '/std/lcs/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlcsLibraries(id) {
    return request({
        url: '/std/lcs/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlcsLibraries(lcs) {
    return request({
        url: '/std/lcs/' + lcs.id,
        method: 'PATCH',
        data: lcs
    });
}
export function CreateStdlcsLibraries(lcs) {
    return request({
        url: '/std/lcs',
        method: 'POST',
        data: lcs
    });
}
//# sourceMappingURL=lcs.js.map