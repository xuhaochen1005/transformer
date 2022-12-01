import request from '@/utils/request';
export function GetStdlckLibraries(query) {
    return request({
        url: '/std/lck/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlckLibraries(id) {
    return request({
        url: '/std/lck/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlckLibraries(lck) {
    return request({
        url: '/std/lck/' + lck.id,
        method: 'PATCH',
        data: lck
    });
}
export function CreateStdlckLibraries(lck) {
    return request({
        url: '/std/lck',
        method: 'POST',
        data: lck
    });
}
//# sourceMappingURL=lck.js.map