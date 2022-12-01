import request from '@/utils/request';
export function GetStdlflvmlLibraries(query) {
    return request({
        url: '/std/lflvml/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlflvmlLibraries(id) {
    return request({
        url: '/std/lflvml/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlflvmlLibraries(lflvml) {
    return request({
        url: '/std/lflvml/' + lflvml.id,
        method: 'PATCH',
        data: lflvml
    });
}
export function CreateStdlflvmlLibraries(lflvml) {
    return request({
        url: '/std/lflvml',
        method: 'POST',
        data: lflvml
    });
}
//# sourceMappingURL=lflvml.js.map