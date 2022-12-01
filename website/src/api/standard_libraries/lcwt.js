import request from '@/utils/request';
export function GetStdlcwtLibraries(query) {
    return request({
        url: '/std/lcwt/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlcwtLibraries(id) {
    return request({
        url: '/std/lcwt/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlcwtLibraries(lcwt) {
    return request({
        url: '/std/lcwt/' + lcwt.id,
        method: 'PATCH',
        data: lcwt
    });
}
export function CreateStdlcwtLibraries(lcwt) {
    return request({
        url: '/std/lcwt',
        method: 'POST',
        data: lcwt
    });
}
//# sourceMappingURL=lcwt.js.map