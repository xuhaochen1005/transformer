import request from '@/utils/request';
export function GetStdLlLibraries(query) {
    return request({
        url: '/std/ll',
        method: 'GET',
        params: query
    });
}
export function CreateStdLlLibraries(ll) {
    return request({
        url: '/std/ll',
        method: 'POST',
        data: ll
    });
}
export function UpdateStdLlLibraries(ll) {
    return request({
        url: '/std/ll/' + ll.id,
        method: 'PATCH',
        data: ll
    });
}
export function DeleteStdLlLibraries(id) {
    return request({
        url: '/std/ll/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=ll.js.map