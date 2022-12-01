import request from '@/utils/request';
export function GetStdlelfLibraries(query) {
    return request({
        url: '/std/lelf/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlelfLibraries(id) {
    return request({
        url: '/std/lelf/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlelfLibraries(lelf) {
    return request({
        url: '/std/lelf/' + lelf.id,
        method: 'PATCH',
        data: lelf
    });
}
export function CreateStdlelfLibraries(lelf) {
    return request({
        url: '/std/lelf',
        method: 'POST',
        data: lelf
    });
}
//# sourceMappingURL=lelf.js.map