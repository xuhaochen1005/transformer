import request from '@/utils/request';
export function GetStdLhvomlLibraries(query) {
    return request({
        url: '/std/lhvoml',
        method: 'GET',
        params: query
    });
}
export function CreateStdLhvomlLibraries(lhvoml) {
    return request({
        url: '/std/lhvoml',
        method: 'POST',
        data: lhvoml
    });
}
export function UpdateStdLhvomlLibraries(lhvoml) {
    return request({
        url: '/std/lhvoml/' + lhvoml.id,
        method: 'PATCH',
        data: lhvoml
    });
}
export function DeleteStdLhvomlLibraries(id) {
    return request({
        url: '/std/lhvoml/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lhvoml.js.map