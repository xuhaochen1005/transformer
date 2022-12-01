import request from '@/utils/request';
export function GetStdlctLibraries(query) {
    return request({
        url: '/std/lct/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlctLibraries(id) {
    return request({
        url: '/std/lct/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlctLibraries(lct) {
    return request({
        url: '/std/lct/' + lct.id,
        method: 'PATCH',
        data: lct
    });
}
export function CreateStdlctLibraries(lct) {
    return request({
        url: '/std/lct',
        method: 'POST',
        data: lct
    });
}
//# sourceMappingURL=lct.js.map