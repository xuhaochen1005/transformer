import request from '@/utils/request';
export function GetStdlcfiyLibraries(query) {
    return request({
        url: '/std/lcfiy/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlcfiyLibraries(id) {
    return request({
        url: '/std/lcfiy/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlcfiyLibraries(lcfiy) {
    return request({
        url: '/std/lcfiy/' + lcfiy.id,
        method: 'PATCH',
        data: lcfiy
    });
}
export function CreateStdlcfiyLibraries(lcfiy) {
    return request({
        url: '/std/lcfiy',
        method: 'POST',
        data: lcfiy
    });
}
//# sourceMappingURL=lcfiy.js.map