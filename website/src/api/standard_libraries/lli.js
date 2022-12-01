import request from '@/utils/request';
export function GetStdlliLibraries(query) {
    return request({
        url: '/std/lli/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlliLibraries(id) {
    return request({
        url: '/std/lli/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlliLibraries(lli) {
    return request({
        url: '/std/lli/' + lli.id,
        method: 'PATCH',
        data: lli
    });
}
export function CreateStdlliLibraries(lli) {
    return request({
        url: '/std/lli',
        method: 'POST',
        data: lli
    });
}
//# sourceMappingURL=lli.js.map