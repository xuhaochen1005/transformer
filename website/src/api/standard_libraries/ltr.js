import request from '@/utils/request';
export function GetStdLtrLibraries(query) {
    return request({
        url: '/std/ltr/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLtrLibraries(id) {
    return request({
        url: '/std/ltr/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLtrLibraries(ltr) {
    return request({
        url: '/std/ltr/' + ltr.id,
        method: 'PATCH',
        data: ltr
    });
}
export function CreateStdLtrLibraries(ltr) {
    return request({
        url: '/std/ltr',
        method: 'POST',
        data: ltr
    });
}
//# sourceMappingURL=ltr.js.map