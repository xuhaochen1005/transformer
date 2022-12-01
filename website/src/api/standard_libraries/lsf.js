import request from '@/utils/request';
export function GetStdLsfLibraries(query) {
    return request({
        url: '/std/lsf/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLsfLibraries(id) {
    return request({
        url: '/std/lsf/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLsfLibraries(lsf) {
    return request({
        url: '/std/lsf/' + lsf.id,
        method: 'PATCH',
        data: lsf
    });
}
export function CreateStdLsfLibraries(lsf) {
    return request({
        url: '/std/lsf',
        method: 'POST',
        data: lsf
    });
}
//# sourceMappingURL=lsf.js.map