import request from '@/utils/request';
export function GetStdlicoLibraries(query) {
    return request({
        url: '/std/lico/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlicoLibraries(id) {
    return request({
        url: '/std/lico/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlicoLibraries(lico) {
    return request({
        url: '/std/lico/' + lico.id,
        method: 'PATCH',
        data: lico
    });
}
export function CreateStdlicoLibraries(lico) {
    return request({
        url: '/std/lico',
        method: 'POST',
        data: lico
    });
}
//# sourceMappingURL=lico.js.map