import request from '@/utils/request';
export function GetStdLrfLibraries(query) {
    return request({
        url: '/std/lrf',
        method: 'GET',
        params: query
    });
}
export function CreateStdLrfLibraries(lrf) {
    return request({
        url: '/std/lrf',
        method: 'POST',
        data: lrf
    });
}
export function UpdateStdLrfLibraries(lrf) {
    return request({
        url: '/std/lrf/' + lrf.id,
        method: 'PATCH',
        data: lrf
    });
}
export function DeleteStdLrfLibraries(id) {
    return request({
        url: '/std/lrf/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lrf.js.map