import request from '@/utils/request';
export function GetStdLpnLibraries(query) {
    return request({
        url: '/std/lpn/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLpnLibraries(id) {
    return request({
        url: '/std/lpn/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLpnLibraries(lpn) {
    return request({
        url: '/std/lpn/' + lpn.id,
        method: 'PATCH',
        data: lpn
    });
}
export function CreateStdLpnLibraries(lpn) {
    return request({
        url: '/std/lpn',
        method: 'POST',
        data: lpn
    });
}
//# sourceMappingURL=lpn.js.map