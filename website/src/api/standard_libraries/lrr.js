import request from '@/utils/request';
export function GetStdLrrLibraries(query) {
    return request({
        url: '/std/lrr',
        method: 'GET',
        params: query
    });
}
export function CreateStdLrrLibraries(lrr) {
    return request({
        url: '/std/lrr',
        method: 'POST',
        data: lrr
    });
}
export function UpdateStdLrrLibraries(lrr) {
    return request({
        url: '/std/lrr/' + lrr.id,
        method: 'PATCH',
        data: lrr
    });
}
export function DeleteStdLrrLibraries(id) {
    return request({
        url: '/std/lrr/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lrr.js.map