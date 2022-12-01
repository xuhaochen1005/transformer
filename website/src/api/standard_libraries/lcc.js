import request from '@/utils/request';
export function GetStdlccLibraries(query) {
    return request({
        url: '/std/lcc/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlccLibraries(id) {
    return request({
        url: '/std/lcc/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlccLibraries(lcc) {
    return request({
        url: '/std/lcc/' + lcc.id,
        method: 'PATCH',
        data: lcc
    });
}
export function CreateStdlccLibraries(lcc) {
    return request({
        url: '/std/lcc',
        method: 'POST',
        data: lcc
    });
}
//# sourceMappingURL=lcc.js.map