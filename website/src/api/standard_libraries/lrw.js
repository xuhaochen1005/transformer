import request from '@/utils/request';
export function GetStdLrwLibraries(query) {
    return request({
        url: '/std/lrw/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLrwLibraries(id) {
    return request({
        url: '/std/lrw/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLrwLibraries(lrw) {
    return request({
        url: '/std/lrw/' + lrw.id,
        method: 'PATCH',
        data: lrw
    });
}
export function CreateStdLrwLibraries(lrw) {
    return request({
        url: '/std/lrw',
        method: 'POST',
        data: lrw
    });
}
//# sourceMappingURL=lrw.js.map