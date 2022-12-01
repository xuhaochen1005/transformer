import request from '@/utils/request';
export function GetStdLlvomlLibraries(query) {
    return request({
        url: '/std/llvoml',
        method: 'GET',
        params: query
    });
}
export function CreateStdLlvomlLibraries(llvoml) {
    return request({
        url: '/std/llvoml',
        method: 'POST',
        data: llvoml
    });
}
export function UpdateStdLlvomlLibraries(llvoml) {
    return request({
        url: '/std/llvoml/' + llvoml.id,
        method: 'PATCH',
        data: llvoml
    });
}
export function DeleteStdLlvomlLibraries(id) {
    return request({
        url: '/std/llvoml/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=llvoml.js.map