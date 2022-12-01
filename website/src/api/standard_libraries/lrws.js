import request from '@/utils/request';
export function GetStdLrwsLibraries(query) {
    return request({
        url: '/std/lrws/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLrwsLibraries(id) {
    return request({
        url: '/std/lrws/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLrwsLibraries(lrws) {
    return request({
        url: '/std/lrws/' + lrws.id,
        method: 'PATCH',
        data: lrws
    });
}
export function CreateStdLrwsLibraries(lrws) {
    return request({
        url: '/std/lrws',
        method: 'POST',
        data: lrws
    });
}
//# sourceMappingURL=lrws.js.map