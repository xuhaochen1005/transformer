import request from '@/utils/request';
export function GetStdLwimLibraries(query) {
    return request({
        url: '/std/lwim/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLwimLibraries(id) {
    return request({
        url: '/std/lwim/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLwimLibraries(lwim) {
    return request({
        url: '/std/lwim/' + lwim.id,
        method: 'PATCH',
        data: lwim
    });
}
export function CreateStdLwimLibraries(lwim) {
    return request({
        url: '/std/lwim',
        method: 'POST',
        data: lwim
    });
}
//# sourceMappingURL=lwim.js.map