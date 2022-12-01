import request from '@/utils/request';
export function GetStdLmdLibraries(query) {
    return request({
        url: '/std/lmd',
        method: 'GET',
        params: query
    });
}
export function CreateStdLmdLibraries(lmd) {
    return request({
        url: '/std/lmd',
        method: 'POST',
        data: lmd
    });
}
export function UpdateStdLmdLibraries(lmd) {
    return request({
        url: '/std/lmd/' + lmd.id,
        method: 'PATCH',
        data: lmd
    });
}
export function DeleteStdLmdLibraries(id) {
    return request({
        url: '/std/lmd/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lmd.js.map