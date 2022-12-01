import request from '@/utils/request';
export function GetStdlfwsLibraries(query) {
    return request({
        url: '/std/lfws/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlfwsLibraries(id) {
    return request({
        url: '/std/lfws/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlfwsLibraries(lfws) {
    return request({
        url: '/std/lfws/' + lfws.id,
        method: 'PATCH',
        data: lfws
    });
}
export function CreateStdlfwsLibraries(lfws) {
    return request({
        url: '/std/lfws',
        method: 'POST',
        data: lfws
    });
}
//# sourceMappingURL=lfws.js.map