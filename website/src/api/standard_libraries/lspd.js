import request from '@/utils/request';
export function GetStdLspdLibraries(query) {
    return request({
        url: '/std/lspd/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLspdLibraries(id) {
    return request({
        url: '/std/lspd/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLspdLibraries(lspd) {
    return request({
        url: '/std/lspd/' + lspd.id,
        method: 'PATCH',
        data: lspd
    });
}
export function CreateStdLspdLibraries(lspd) {
    return request({
        url: '/std/lspd',
        method: 'POST',
        data: lspd
    });
}
//# sourceMappingURL=lspd.js.map