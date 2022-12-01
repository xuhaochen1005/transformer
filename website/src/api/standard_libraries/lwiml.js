import request from '@/utils/request';
export function GetStdLwimlLibraries(query) {
    return request({
        url: '/std/lwiml',
        method: 'GET',
        params: query
    });
}
export function CreateStdLwimlLibraries(lwiml) {
    return request({
        url: '/std/lwiml',
        method: 'POST',
        data: lwiml
    });
}
export function UpdateStdLwimlLibraries(lwiml) {
    return request({
        url: '/std/lwiml/' + lwiml.id,
        method: 'PATCH',
        data: lwiml
    });
}
export function DeleteStdLwimlLibraries(id) {
    return request({
        url: '/std/lwiml/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lwiml.js.map