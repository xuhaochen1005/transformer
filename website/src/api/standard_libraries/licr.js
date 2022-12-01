import request from '@/utils/request';
export function GetStdlicrLibraries(query) {
    return request({
        url: '/std/licr/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlicrLibraries(id) {
    return request({
        url: '/std/licr/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlicrLibraries(licr) {
    return request({
        url: '/std/licr/' + licr.id,
        method: 'PATCH',
        data: licr
    });
}
export function CreateStdlicrLibraries(licr) {
    return request({
        url: '/std/licr',
        method: 'POST',
        data: licr
    });
}
//# sourceMappingURL=licr.js.map