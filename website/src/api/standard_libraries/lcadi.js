import request from '@/utils/request';
export function GetStdLcadiLibraries(query) {
    return request({
        url: '/std/lcadi/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLcadiLibraries(id) {
    return request({
        url: '/std/lcadi/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLcadiLibraries(lcadi) {
    return request({
        url: '/std/lcadi/' + lcadi.id,
        method: 'PATCH',
        data: lcadi
    });
}
export function CreateStdLcadiLibraries(lcadi) {
    return request({
        url: '/std/lcadi',
        method: 'POST',
        data: lcadi
    });
}
//# sourceMappingURL=lcadi.js.map