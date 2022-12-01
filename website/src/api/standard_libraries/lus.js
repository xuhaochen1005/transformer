import request from '@/utils/request';
export function GetStdLusLibraries(query) {
    return request({
        url: '/std/lus/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLusLibraries(id) {
    return request({
        url: '/std/lus/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLusLibraries(lus) {
    return request({
        url: '/std/lus/' + lus.id,
        method: 'PATCH',
        data: lus
    });
}
export function CreateStdLusLibraries(lus) {
    return request({
        url: '/std/lus',
        method: 'POST',
        data: lus
    });
}
//# sourceMappingURL=lus.js.map