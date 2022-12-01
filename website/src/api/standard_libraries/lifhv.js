import request from '@/utils/request';
export function GetStdlifhvLibraries(query) {
    return request({
        url: '/std/lifhv/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlifhvLibraries(id) {
    return request({
        url: '/std/lifhv/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlifhvLibraries(lifhv) {
    return request({
        url: '/std/lifhv/' + lifhv.id,
        method: 'PATCH',
        data: lifhv
    });
}
export function CreateStdlifhvLibraries(lifhv) {
    return request({
        url: '/std/lifhv',
        method: 'POST',
        data: lifhv
    });
}
//# sourceMappingURL=lifhv.js.map