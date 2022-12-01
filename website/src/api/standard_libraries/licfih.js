import request from '@/utils/request';
export function GetStdlicfihLibraries(query) {
    return request({
        url: '/std/licfih/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlicfihLibraries(id) {
    return request({
        url: '/std/licfih/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlicfihLibraries(licfih) {
    return request({
        url: '/std/licfih/' + licfih.id,
        method: 'PATCH',
        data: licfih
    });
}
export function CreateStdlicfihLibraries(licfih) {
    return request({
        url: '/std/licfih',
        method: 'POST',
        data: licfih
    });
}
//# sourceMappingURL=licfih.js.map