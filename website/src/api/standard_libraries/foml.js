import request from '@/utils/request';
export function GetStdlfomlLibraries(query) {
    return request({
        url: '/std/lfoml/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlfomlLibraries(id) {
    return request({
        url: '/std/lfoml/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlfomlLibraries(lfoml) {
    return request({
        url: '/std/lflvml/' + lfoml.id,
        method: 'PATCH',
        data: lfoml
    });
}
export function CreateStdlfomlLibraries(lfoml) {
    return request({
        url: '/std/lfoml',
        method: 'POST',
        data: lfoml
    });
}
//# sourceMappingURL=foml.js.map